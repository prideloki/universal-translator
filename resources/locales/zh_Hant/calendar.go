package zh_Hant

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "y年M月d日 EEEE", Long: "y年M月d日", Medium: "y年M月d日", Short: "y/M/d"},
			Time:     ut.CalendarDateFormat{Full: "ah:mm:ss [zzzz]", Long: "ah:mm:ss [z]", Medium: "ah:mm:ss", Short: "ah:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "1月", Feb: "2月", Mar: "3月", Apr: "4月", May: "5月", Jun: "6月", Jul: "7月", Aug: "8月", Sep: "9月", Oct: "10月", Nov: "11月", Dec: "12月"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "1", Feb: "2", Mar: "3", Apr: "4", May: "5", Jun: "6", Jul: "7", Aug: "8", Sep: "9", Oct: "10", Nov: "11", Dec: "12"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "1月", Feb: "2月", Mar: "3月", Apr: "4月", May: "5月", Jun: "6月", Jul: "7月", Aug: "8月", Sep: "9月", Oct: "10月", Nov: "11月", Dec: "12月"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "週日", Mon: "週一", Tue: "週二", Wed: "週三", Thu: "週四", Fri: "週五", Sat: "週六"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "日", Mon: "一", Tue: "二", Wed: "三", Thu: "四", Fri: "五", Sat: "六"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "日", Mon: "一", Tue: "二", Wed: "三", Thu: "四", Fri: "五", Sat: "六"},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "星期日", Mon: "星期一", Tue: "星期二", Wed: "星期三", Thu: "星期四", Fri: "星期五", Sat: "星期六"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "上午", PM: "下午"},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{},
			},
		},
	}
}
