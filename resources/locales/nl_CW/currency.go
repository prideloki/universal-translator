package nl_CW

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "ANG", DisplayName: "", Symbol: "NAf."},
	}
}
