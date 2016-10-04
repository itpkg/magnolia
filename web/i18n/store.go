package i18n

import "golang.org/x/text/language"

//Store i18n store
type Store interface {
	Set(lang *language.Tag, code, message string)
	Get(lang *language.Tag, code string) string
	Del(lang *language.Tag, code string)
	Keys(lang *language.Tag) ([]string, error)
}
