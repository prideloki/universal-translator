package es_EC

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "USD", DisplayName: "", Symbol: "$"},
	}
}
