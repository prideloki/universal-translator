package ur

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: ".", Group: ",", Negative: "\u200e-\u200e", Percent: "%", PerMille: "‰"}
	formats = ut.NumberFormats{Decimal: "#,##0.###", Currency: "¤\u00a0#,##,##0.00", CurrencyAccounting: "¤\u00a0#,##,##0.00", Percent: "#,##,##0%"}
)
