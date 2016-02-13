package en_TK

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "Oct", 12: "Dec", 1: "Jan", 8: "Aug", 9: "Sep", 11: "Nov", 5: "May", 7: "Jul", 2: "Feb", 3: "Mar", 4: "Apr", 6: "Jun"}, Narrow: ut.CalendarMonthFormatNameValue{10: "O", 1: "J", 12: "D", 2: "F", 6: "J", 7: "J", 8: "A", 3: "M", 4: "A", 5: "M", 9: "S", 11: "N"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "May", 6: "June", 8: "August", 10: "October", 2: "February", 3: "March", 12: "December", 1: "January", 7: "July", 9: "September", 11: "November", 4: "April"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri"}, Narrow: ut.CalendarDayFormatNameValue{2: "T", 3: "W", 4: "T", 5: "F", 6: "S", 0: "S", 1: "M"}, Short: ut.CalendarDayFormatNameValue{0: "Su", 1: "Mo", 2: "Tu", 3: "We", 4: "Th", 5: "Fr", 6: "Sa"}, Wide: ut.CalendarDayFormatNameValue{5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight"}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "n", "pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight", "am": "AM", "noon": "noon"}}}}