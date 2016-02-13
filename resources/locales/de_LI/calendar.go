package de_LI

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'um' {0}", Long: "{1} 'um' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "Aug", 9: "Sep", 11: "Nov", 12: "Dez", 10: "Okt", 2: "Feb", 3: "Mär", 7: "Jul", 1: "Jan", 5: "Mai", 4: "Apr", 6: "Jun"}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 12: "D", 1: "J", 2: "F", 5: "M", 8: "A", 6: "J", 3: "M", 9: "S", 10: "O", 11: "N", 7: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{11: "November", 4: "April", 5: "Mai", 7: "Juli", 8: "August", 9: "September", 12: "Dezember", 10: "Oktober", 1: "Januar", 2: "Februar", 6: "Juni", 3: "März"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Mo", 2: "Di", 3: "Mi", 4: "Do", 5: "Fr", 6: "Sa", 0: "So"}, Narrow: ut.CalendarDayFormatNameValue{4: "D", 5: "F", 6: "S", 0: "S", 1: "M", 2: "D", 3: "M"}, Short: ut.CalendarDayFormatNameValue{6: "Sa.", 0: "So.", 1: "Mo.", 2: "Di.", 3: "Mi.", 4: "Do.", 5: "Fr."}, Wide: ut.CalendarDayFormatNameValue{4: "Donnerstag", 5: "Freitag", 6: "Samstag", 0: "Sonntag", 1: "Montag", 2: "Dienstag", 3: "Mittwoch"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"night1": "Nacht", "am": "vorm.", "morning1": "Morgen", "afternoon1": "Mittag", "midnight": "Mitternacht", "pm": "nachm.", "evening1": "Abend", "afternoon2": "Nachmittag", "morning2": "Vormittag"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "vm.", "midnight": "Mitternacht", "afternoon2": "Nachmittag", "pm": "nm.", "morning1": "Morgen", "night1": "Nacht", "evening1": "Abend", "morning2": "Vormittag", "afternoon1": "Mittag"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"afternoon2": "Nachmittag", "am": "vorm.", "afternoon1": "Mittag", "pm": "nachm.", "morning1": "Morgen", "night1": "Nacht", "midnight": "Mitternacht", "evening1": "Abend", "morning2": "Vormittag"}}}}