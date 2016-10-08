package controllers

import (
	"fmt"
	"path"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/itpkg/magnolia/models"
	"golang.org/x/text/language"
)

//Controller base controller
type Controller struct {
	beego.Controller

	Locale string
}

//Prepare prepare
func (p *Controller) Prepare() {
	p.setLangVer()
}

func (p *Controller) setLangVer() {
	const key = "locale"

	// 1. Check URL arguments.
	lang := p.Input().Get(key)

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = p.Ctx.GetCookie(key)
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := p.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. Default language is English.
	lng, _, _ := matcher.Match(language.Make(lang))
	lang = lng.String()

	// Save language information in cookies.
	p.Ctx.SetCookie(key, lang, 1<<31-1, "/")

	// Set language properties.
	p.Locale = lang
	p.Data["L"] = lang
	p.Data["Locales"] = locales

}

// ----------------------------------------------------------------------------
var matcher language.Matcher
var locales = []language.Tag{
	language.AmericanEnglish,
	language.SimplifiedChinese,
	language.TraditionalChinese,
}

func init() {
	matcher = language.NewMatcher(locales)

	for _, lang := range locales {
		lng := lang.String()
		beego.Info("Loading language: " + lng)
		if err := i18n.SetMessage(lng, path.Join("conf", "locales", fmt.Sprintf("%s.ini", lng))); err != nil {
			beego.Error(err)
		}
	}

	beego.AddFuncMap("T", models.T)
}
