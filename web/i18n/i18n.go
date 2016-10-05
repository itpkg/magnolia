package i18n

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
	"github.com/op/go-logging"
	"golang.org/x/text/language"
)

//I18n i18n helper
type I18n struct {
	Store  Store           `inject:""`
	Logger *logging.Logger `inject:""`

	Locales map[string]map[string]string
}

//Exist is lang exist?
func (p *I18n) Exist(lang string) bool {
	_, ok := p.Locales[lang]
	return ok
}

//Items list all items
// func (p *I18n) Items(lng string) map[string]interface{} {
// 	rt := make(map[string]interface{})
// 	if items, ok := p.Locales[lng]; ok {
// 		for k, v := range items {
// 			if strings.HasPrefix(k, "web.") {
// 				k = k[4:]
// 				codes := strings.Split(k, ".")
// 				tmp := rt
// 				for i, c := range codes {
// 					if i+1 == len(codes) {
// 						tmp[c] = v
// 					} else {
// 						if tmp[c] == nil {
// 							tmp[c] = make(map[string]interface{})
// 						}
// 						tmp = tmp[c].(map[string]interface{})
// 					}
// 				}
//
// 			}
// 		}
// 	}
// 	return rt
// }

//Load load locales
func (p *I18n) Load(dir string) error {
	const ext = ".ini"
	// Load locale from filesystem
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			name := info.Name()
			if filepath.Ext(name) != ext {
				return fmt.Errorf("Ingnore locale file %s", name)
			}
			p.Logger.Debugf("Find locale file %s", path)
			lang := language.Make(name[0 : len(name)-len(ext)])
			cfg, err := ini.Load(path)
			if err != nil {
				return err
			}
			items := cfg.Section("DEFAULT").KeysHash()
			// p.Logger.Debugf("Find %d items for locale %s", len(items), lang)
			for k, v := range items {
				p.set(&lang, k, v)
			}

		}
		return nil
	}); err != nil {
		return err
	}

	// Load locale from database
	for lang := range p.Locales {
		lng := language.Make(lang)
		ks, err := p.Store.Keys(&lng)
		if err != nil {
			return err
		}
		for _, k := range ks {
			p.Locales[lang][k] = p.Store.Get(&lng, k)
		}
		p.Logger.Debugf("Find locale %s, %d items.", lang, len(p.Locales[lang]))
	}
	return nil
}

func (p *I18n) set(lng *language.Tag, code, message string) {
	lang := lng.String()
	if _, ok := p.Locales[lang]; !ok {
		p.Locales[lang] = make(map[string]string)
	}
	p.Locales[lang][code] = message
}

//Ts translate by lang
func (p *I18n) Ts(args ...interface{}) string {
	if len(args) < 2 {
		return "null"
	}
	l := language.Make(args[0].(string))
	return p.T(&l, args[1].(string), args[2:]...)
}

//T translate by lang tag
func (p *I18n) T(lng *language.Tag, code string, args ...interface{}) string {
	lang := lng.String()
	msg := p.Store.Get(lng, code)
	if len(msg) == 0 {
		if items, ok := p.Locales[lang]; ok {
			msg = items[code]
		}
	}
	return fmt.Sprintf(msg, args...)
}
