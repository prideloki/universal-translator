package ha

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan", 3: "Mar", 6: "Yun", 9: "Sat", 10: "Okt", 12: "Dis", 2: "Fab", 4: "Afi", 5: "May", 7: "Yul", 8: "Agu", 11: "Nuw"}, Narrow: ut.CalendarMonthFormatNameValue{7: "Y", 9: "S", 10: "O", 1: "J", 2: "F", 3: "M", 8: "A", 11: "N", 12: "D", 4: "A", 5: "M", 6: "Y"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{5: "Mayu", 6: "Yuni", 10: "Oktoba", 1: "Janairu", 2: "Faburairu", 3: "Maris", 4: "Afirilu", 7: "Yuli", 8: "Agusta", 9: "Satumba", 11: "Nuwamba", 12: "Disamba"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Ta", 3: "Lr", 4: "Al", 5: "Ju", 6: "As", 0: "Lh", 1: "Li"}, Narrow: ut.CalendarDayFormatNameValue{4: "A", 5: "J", 6: "A", 0: "L", 1: "L", 2: "T", 3: "L"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "Lahadi", 1: "Litinin", 2: "Talata", 3: "Laraba", 4: "Alhamis", 5: "Jummaʼa", 6: "Asabar"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}}}