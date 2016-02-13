package se_SE

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}, Time: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "borg", 12: "juov", 4: "cuo", 6: "geas", 7: "suoi", 9: "čakč", 2: "guov", 5: "mies", 11: "skáb", 3: "njuk", 1: "ođđj", 10: "golg"}, Narrow: ut.CalendarMonthFormatNameValue{9: "Č", 2: "G", 8: "B", 3: "N", 4: "C", 5: "M", 6: "G", 1: "O", 10: "G", 11: "S", 12: "J", 7: "S"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{9: "čakčamánnu", 10: "golggotmánnu", 1: "ođđajagemánnu", 3: "njukčamánnu", 8: "borgemánnu", 12: "juovlamánnu", 6: "geassemánnu", 7: "suoidnemánnu", 11: "skábmamánnu", 2: "guovvamánnu", 4: "cuoŋománnu", 5: "miessemánnu"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "duor", 5: "bear", 6: "láv", 0: "sotn", 1: "vuos", 2: "maŋ", 3: "gask"}, Narrow: ut.CalendarDayFormatNameValue{6: "L", 0: "S", 1: "V", 2: "M", 3: "G", 4: "D", 5: "B"}, Short: ut.CalendarDayFormatNameValue{}, Wide: ut.CalendarDayFormatNameValue{4: "duorasdat", 5: "bearjadat", 6: "lávvardat", 0: "sotnabeaivi", 1: "vuossárga", 2: "maŋŋebárga", 3: "gaskavahkku"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "i.b.", "pm": "e.b."}, Narrow: ut.CalendarPeriodFormatNameValue{}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"pm": "eahketbeaivi", "am": "iđitbeaivi"}}}}