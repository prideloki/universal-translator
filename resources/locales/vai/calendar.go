package vai

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue(nil), Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "ꖨꕪꖃ ꔞꕮ", 6: "6", 7: "7", 11: "ꔞꘋꕔꕿ ꕸꖃꗏ", 9: "ꕢꕌ", 10: "ꕭꖃ", 12: "ꖨꕪꕱ ꗏꕮ", 2: "ꕒꕡꖝꖕ", 3: "ꕾꖺ", 4: "ꖢꖕ", 5: "ꖑꕱ", 8: "ꗛꔕ"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue(nil), Narrow: ut.CalendarDayFormatNameValue(nil), Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{4: "ꕉꔤꕆꕢ", 5: "ꕉꔤꕀꕮ", 6: "ꔻꔬꔳ", 0: "ꕞꕌꔵ", 1: "ꗳꗡꘉ", 2: "ꕚꕞꕚ", 3: "ꕉꕞꕒ"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}}}