package yi

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, dטן MMMM y", Long: "dטן MMMM y", Medium: "dטן MMM y", Short: "dd/MM/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}, {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "אקט", 3: "מערץ", 4: "אַפּר", 7: "יולי", 9: "סעפּ", 8: "אויג", 11: "נאוו", 12: "דעצ", 1: "יאַנ", 2: "פֿעב", 5: "מיי", 6: "יוני"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "מערץ", 5: "מיי", 7: "יולי", 8: "אויגוסט", 11: "נאוועמבער", 12: "דעצעמבער", 1: "יאַנואַר", 2: "פֿעברואַר", 4: "אַפּריל", 6: "יוני", 9: "סעפּטעמבער", 10: "אקטאבער"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "פֿרײַטיק", 6: "שבת", 0: "זונטיק", 1: "מאָנטיק", 2: "דינסטיק", 3: "מיטוואך", 4: "דאנערשטיק"}, Narrow: ut.CalendarDayFormatNameValue(nil), Short: ut.CalendarDayFormatNameValue{5: "פֿרײַטיק", 6: "שבת", 0: "זונטיק", 1: "מאָנטיק", 2: "דינסטיק", 3: "מיטוואך", 4: "דאנערשטיק"}, Wide: ut.CalendarDayFormatNameValue{2: "דינסטיק", 3: "מיטוואך", 4: "דאנערשטיק", 5: "פֿרײַטיק", 6: "שבת", 0: "זונטיק", 1: "מאָנטיק"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "פֿאַרמיטאָג", "pm": "נאָכמיטאָג"}}}}