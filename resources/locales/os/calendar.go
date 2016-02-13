package os

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM, y 'аз'", Long: "d MMMM, y 'аз'", Medium: "dd MMM y 'аз'", Short: "dd.MM.yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "Нояб.", 2: "Февр.", 3: "Март.", 4: "Апр.", 5: "Май", 6: "Июнь", 8: "Авг.", 1: "Янв.", 7: "Июль", 9: "Сент.", 10: "Окт.", 12: "Дек."}, Narrow: ut.CalendarMonthFormatNameValue{2: "Ф", 3: "М", 6: "И", 8: "А", 10: "О", 11: "Н", 12: "Д", 1: "Я", 4: "А", 5: "М", 7: "И", 9: "С"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{7: "Июль", 9: "Сентябрь", 10: "Октябрь", 2: "Февраль", 3: "Мартъи", 4: "Апрель", 5: "Май", 6: "Июнь", 11: "Ноябрь", 1: "Январь", 8: "Август", 12: "Декабрь"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Крс", 2: "Дцг", 3: "Ӕрт", 4: "Цпр", 5: "Мрб", 6: "Сбт", 0: "Хцб"}, Narrow: ut.CalendarDayFormatNameValue{2: "Д", 3: "Ӕ", 4: "Ц", 5: "М", 6: "С", 0: "Х", 1: "К"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{4: "Цыппӕрӕм", 5: "Майрӕмбон", 6: "Сабат", 0: "Хуыцаубон", 1: "Къуырисӕр", 2: "Дыццӕг", 3: "Ӕртыццӕг"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "ӕмбисбоны размӕ", "pm": "ӕмбисбоны фӕстӕ"}}}}