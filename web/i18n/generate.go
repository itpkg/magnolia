package i18n

import (
	"fmt"
	"log"
	"os"
	"path"

	"golang.org/x/text/language"
)

//Generate generate translate file
func Generate(lang string) error {
	lng, err := language.Parse(lang)
	if err != nil {
		return err
	}
	if err = os.MkdirAll("locales", 0700); err != nil {
		return err
	}
	fn := path.Join("locales", fmt.Sprintf("%s.toml", lng.String()))
	log.Printf("generate locale file %s", fn)
	fd, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	defer fd.Close()
	return err
}
