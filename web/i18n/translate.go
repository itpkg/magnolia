package i18n

import(
	"log"
	"fmt"
	"os"
	"path/filepath"
	"path"

		"golang.org/x/text/language"
		"github.com/BurntSushi/toml"
)


//Load load from locales path
func Load() error{
	const ext = ".toml"
	err := filepath.Walk("locales", func(p string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			name := f.Name()
			if path.Ext(name) == ext {
				lang, err:=language.Parse(name[0 : len(name)-len(ext)])
				if err!=nil{
					return err
				}
				log.Printf("find locale %s", lang)
	items := make(map[string]string)
				_, err = toml.DecodeFile(p, &items)
				if err != nil {
					return err
				}
				locales[lang.String()] = items
			}
		}
		return nil
	})
	return err

}

//T translate 
func T(lang, code string, args...interface{}) string{
	var msg string
	if provider != nil{
		msg = provider.Get(lang, code)
	}
	if msg == ""{
		if items, ok := locales[lang]; ok{
			if item,ok:=items[code]; ok{
				msg = item
			}
		}
	}
	if msg == ""{
		return code
	}
	return fmt.Sprintf(msg, args...)
}


var locales  map[string]map[string]string
var provider Provider

func init(){
	locales = make(map[string]map[string]string)
}
