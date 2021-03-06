package i18n

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

//LocaleHandler detect locale from http header
func LocaleHandler(c *gin.Context) {
	const key = "locale"
	// 1. Check URL arguments.
	lng := c.Request.URL.Query().Get(key)

	// 2. Get language information from cookies.
	if len(lng) == 0 {
		if ck, er := c.Request.Cookie(key); er == nil {
			lng = ck.Value
		}
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lng) == 0 {
		al := c.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			lng = al[:5]
		}
	}
	tag, _, _ := matcher.Match(language.Make(lng))
	// 4. Write to cookie.
	c.SetCookie(key, tag.String(), 1<<31-1, "/", "", false, false)
	c.Set(key, &tag)
}

var matcher language.Matcher

func init() {
	matcher = language.NewMatcher([]language.Tag{
		language.AmericanEnglish,
		language.SimplifiedChinese,
		language.TraditionalChinese,
	})
}
