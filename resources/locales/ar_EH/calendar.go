package ar_EH

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "يناير", 5: "مايو", 9: "سبتمبر", 4: "أبريل", 7: "يوليو", 6: "يونيو", 12: "ديسمبر", 10: "أكتوبر", 11: "نوفمبر", 2: "فبراير", 3: "مارس", 8: "أغسطس"}, Narrow: ut.CalendarMonthFormatNameValue{4: "أ", 5: "و", 7: "ل", 11: "ب", 12: "د", 3: "م", 6: "ن", 2: "ف", 8: "غ", 9: "س", 10: "ك", 1: "ي"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{9: "سبتمبر", 2: "فبراير", 7: "يوليو", 10: "أكتوبر", 5: "مايو", 4: "أبريل", 12: "ديسمبر", 3: "مارس", 6: "يونيو", 11: "نوفمبر", 1: "يناير", 8: "أغسطس"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء"}, Narrow: ut.CalendarDayFormatNameValue{1: "ن", 2: "ث", 3: "ر", 4: "خ", 5: "ج", 6: "س", 0: "ح"}, Short: ut.CalendarDayFormatNameValue{1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد"}, Wide: ut.CalendarDayFormatNameValue{0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "مساءً", "night2": "ليلاً", "afternoon1": "ظهرًا", "night1": "منتصف الليل", "morning1": "فجرا", "pm": "م", "afternoon2": "بعد الظهر", "am": "ص", "morning2": "صباحًا"}, Narrow: ut.CalendarPeriodFormatNameValue{"night2": "ليلاً", "morning1": "فجرا", "morning2": "صباحًا", "afternoon2": "بعد الظهر", "evening1": "مساءً", "pm": "م", "afternoon1": "ظهرًا", "night1": "منتصف الليل", "am": "ص"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night2": "ليلاً", "night1": "منتصف الليل", "morning1": "فجرا", "morning2": "صباحًا", "pm": "مساءً", "afternoon1": "ظهرًا", "am": "صباحًا", "afternoon2": "بعد الظهر", "evening1": "مساءً"}}}}