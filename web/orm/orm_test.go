package orm_test

import (
	"testing"
	"time"

	"github.com/itpkg/magnolia/web/orm"
	_ "github.com/lib/pq"
)

func TestOrm(t *testing.T) {
	if err := orm.Open("postgres", "user=postgres dbname=magnolia_test sslmode=disable"); err != nil {
		t.Fatal(err)
	}
	if err := orm.Generate(time.Now().Format("060102150405")); err != nil {
		t.Fatal(err)
	}
	if err := orm.Migrate(); err != nil {
		t.Fatal(err)
	}
	if err := orm.Rollback(); err != nil {
		t.Fatal(err)
	}
	if status, err := orm.Status(); err == nil {
		t.Logf("%+v", status)
	} else {
		t.Fatal(err)
	}
}
