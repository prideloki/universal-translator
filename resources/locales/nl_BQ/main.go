package nl_BQ

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "nl_BQ",
	Number: ut.Number{
		Symbols:    symbols,
		Formats:    formats,
		Currencies: currencies,
	},
	Calendar:   calendar,
	PluralRule: pluralRule,
}

func init() {
	ut.RegisterLocale(locale)
}