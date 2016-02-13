package sr

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd.MM.y.", Short: "d.M.yy."}, Time: ut.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "мај", 7: "јул", 9: "сеп", 11: "нов", 6: "јун", 8: "авг", 10: "окт", 12: "дец", 1: "јан", 2: "феб", 3: "мар", 4: "апр"}, Narrow: ut.CalendarMonthFormatNameValue{4: "а", 5: "м", 6: "ј", 12: "д", 1: "ј", 3: "м", 7: "ј", 8: "а", 9: "с", 10: "о", 11: "н", 2: "ф"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "март", 4: "април", 8: "август", 9: "септембар", 10: "октобар", 11: "новембар", 12: "децембар", 1: "јануар", 2: "фебруар", 5: "мај", 6: "јун", 7: "јул"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "суб", 0: "нед", 1: "пон", 2: "уто", 3: "сре", 4: "чет", 5: "пет"}, Narrow: ut.CalendarDayFormatNameValue{3: "с", 4: "ч", 5: "п", 6: "с", 0: "н", 1: "п", 2: "у"}, Short: ut.CalendarDayFormatNameValue{1: "по", 2: "ут", 3: "ср", 4: "че", 5: "пе", 6: "су", 0: "не"}, Wide: ut.CalendarDayFormatNameValue{5: "петак", 6: "субота", 0: "недеља", 1: "понедељак", 2: "уторак", 3: "среда", 4: "четвртак"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "поподне", "evening1": "вече", "night1": "ноћ", "midnight": "поноћ", "am": "пре подне", "noon": "подне", "pm": "по подне", "morning1": "јутро"}, Narrow: ut.CalendarPeriodFormatNameValue{"afternoon1": "поподне", "evening1": "вече", "night1": "ноћ", "midnight": "поноћ", "am": "пре подне", "noon": "подне", "pm": "по подне", "morning1": "јутро"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "поподне", "evening1": "вече", "night1": "ноћ", "midnight": "поноћ", "am": "пре подне", "noon": "подне", "pm": "по подне", "morning1": "јутро"}}}}