package ps

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "AFN", DisplayName: "افغانۍ", Symbol: "؋"},
	}
}
