package i18n

type Provider interface {
	Set(lang, code, message string)
	Get(lang, code string) string
}
