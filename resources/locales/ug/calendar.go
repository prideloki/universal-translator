package ug

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE، MMMM d، y", Long: "MMMM d، y", Medium: "MMM d، y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}، {0}", Short: "{1}، {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "ماي", 6: "ئىيۇن", 7: "ئىيۇل", 8: "ئاۋغۇست", 10: "ئۆكتەبىر", 11: "نويابىر", 1: "يانۋار", 2: "فېۋرال", 3: "مارت", 4: "ئاپرېل", 9: "سېنتەبىر", 12: "دېكابىر"}, Narrow: ut.CalendarMonthFormatNameValue{1: "1", 2: "2", 3: "3", 5: "5", 7: "7", 10: "10", 12: "12", 4: "4", 6: "6", 8: "8", 9: "9", 11: "11"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{4: "ئاپرېل", 6: "ئىيۇن", 7: "ئىيۇل", 12: "دېكابىر", 1: "يانۋار", 3: "مارت", 5: "ماي", 8: "ئاۋغۇست", 9: "سېنتەبىر", 10: "ئۆكتەبىر", 11: "بويابىر", 2: "فېۋرال"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "چا", 4: "پە", 5: "چۈ", 6: "شە", 0: "يە", 1: "دۈ", 2: "سە"}, Narrow: ut.CalendarDayFormatNameValue{6: "ش", 0: "ي", 1: "د", 2: "س", 3: "چ", 4: "پ", 5: "ج"}, Short: ut.CalendarDayFormatNameValue{3: "چ", 4: "پ", 5: "ج", 6: "ش", 0: "ي", 1: "د", 2: "س"}, Wide: ut.CalendarDayFormatNameValue{6: "شەنبە", 0: "يەكشەنبە", 1: "دۈشەنبە", 2: "سەيشەنبە", 3: "چارشەنبە", 4: "پەيشەنبە", 5: "جۈمە"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue{"am": "چۈشتىن بۇرۇن", "pm": "چۈشتىن كېيىن"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "چۈشتىن بۇرۇن", "pm": "چۈشتىن كېيىن"}}}}