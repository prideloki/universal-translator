package fr_SY

import "github.com/go-playground/universal-translator"

// New returns a new instance of the locale
func New() *ut.Locale {
	return &ut.Locale{
		Locale: "fr_SY",
		Number: ut.Number{
			Symbols:    newSymbols(),
			Formats:    newFormats(),
			Currencies: newCurrencies(),
		},
		Calendar:   newCalendar(),
		PluralRule: pluralRule,
	}
}
