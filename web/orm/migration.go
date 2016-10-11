package orm

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

type Migration struct {
	Name string   `toml:"-"`
	Up   []string `toml:"up"`
	Down []string `toml:"down"`
}

func (p *Migration) File() string {
	return path.Join("db", driver, "migrations", p.Name)
}

func Generate(name string) error {
	mig := Migration{
		Name: fmt.Sprintf("%s_%s.toml", time.Now().Format("20060102150405"), name),
		Up:   []string{fmt.Sprintf("CREATE TABLE t%s(id int)", name)},
		Down: []string{fmt.Sprintf("DROP TABLE t%s", name)},
	}
	log.Printf("generate file %s", mig.File())

	fd, err := os.OpenFile(mig.File(), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	defer fd.Close()
	return toml.NewEncoder(fd).Encode(mig)
}

func Migrate() error {
	items, err := load()
	if err != nil {
		return err
	}
	for _, mig := range items {
		var count int
		switch driver {
		case "postgres":
			row := db.QueryRow("SELECT COUNT(*) FROM migration_schemes WHERE name = $1", mig.Name)
			err = row.Scan(&count)
		default:
			err = unknownDriverError
		}
		if err != nil {
			return err
		}
		if count == 0 {
			switch driver {
			case "postgres":
				for _, up := range mig.Up {
					if _, err := db.Exec(up); err != nil {
						return err
					}
				}

			default:
				return unknownDriverError
			}

		}
		switch driver {
		case "postgres":
			_, err = db.Exec("INSERT INTO migration_schemes(name) VALUES($1)", mig.Name)
		default:
			err = unknownDriverError
		}
	}
	return err
}

func Rollback() error {
	var name string
	var err error
	switch driver {
	case "postgres":
		row := db.QueryRow("SELECT name FROM migration_schemes ORDER BY id DESC LIMIT 1")
		err = row.Scan(&name)
	default:
		err = unknownDriverError
	}
	if err != nil {
		return err
	}
	var mig Migration
	_, err = toml.DecodeFile(path.Join("db", driver, "migrations", fmt.Sprintf("%s.toml", name)), &mig)
	if err != nil {
		return err
	}
	for _, down := range mig.Down {
		if _, err := db.Exec(down); err != nil {
			return err
		}
	}
	switch driver {
	case "postgres":
		_, err = db.Exec("DELETE FROM migration_schemes WHERE name = $1", name)
	default:
		err = unknownDriverError
	}

	return err
}

func Status() (map[string]time.Time, error) {
	rows, err := db.Query("SELECT name, created_at FROM migration_schemes ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	ret := make(map[string]time.Time)
	for rows.Next() {
		var created time.Time
		var name string
		if err := rows.Scan(&name, &created); err != nil {
			return nil, err
		}
		ret[name] = created
	}
	return ret, nil
}

func load() ([]*Migration, error) {
	const ext = ".toml"
	root := path.Join("db", driver, "migrations")
	var items []*Migration
	err := filepath.Walk(root, func(p string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			name := f.Name()
			if path.Ext(name) == ext {
				var mig Migration
				_, err := toml.DecodeFile(p, &mig)
				if err != nil {
					return err
				}
				mig.Name = name[0 : len(name)-len(ext)]
				items = append(items, &mig)
			}
		}
		return nil
	})
	return items, err
}
