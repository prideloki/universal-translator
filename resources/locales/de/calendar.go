package de

import "github.com/go-playground/universal-translator"

func newCalendar() ut.Calendar {
	return ut.Calendar{
		Formats: ut.CalendarFormats{
			Date:     ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"},
			Time:     ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"},
			DateTime: ut.CalendarDateFormat{Full: "{1} 'um' {0}", Long: "{1} 'um' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"},
		},
		FormatNames: ut.CalendarFormatNames{
			Months: ut.CalendarMonthFormatNames{
				Abbreviated: ut.CalendarMonthFormatNameValue{Jan: "Jan", Feb: "Feb", Mar: "Mär", Apr: "Apr", May: "Mai", Jun: "Jun", Jul: "Jul", Aug: "Aug", Sep: "Sep", Oct: "Okt", Nov: "Nov", Dec: "Dez"},
				Narrow:      ut.CalendarMonthFormatNameValue{Jan: "J", Feb: "F", Mar: "M", Apr: "A", May: "M", Jun: "J", Jul: "J", Aug: "A", Sep: "S", Oct: "O", Nov: "N", Dec: "D"},
				Short:       ut.CalendarMonthFormatNameValue{},
				Wide:        ut.CalendarMonthFormatNameValue{Jan: "Januar", Feb: "Februar", Mar: "März", Apr: "April", May: "Mai", Jun: "Juni", Jul: "Juli", Aug: "August", Sep: "September", Oct: "Oktober", Nov: "November", Dec: "Dezember"},
			},
			Days: ut.CalendarDayFormatNames{
				Abbreviated: ut.CalendarDayFormatNameValue{Sun: "So", Mon: "Mo", Tue: "Di", Wed: "Mi", Thu: "Do", Fri: "Fr", Sat: "Sa"},
				Narrow:      ut.CalendarDayFormatNameValue{Sun: "S", Mon: "M", Tue: "D", Wed: "M", Thu: "D", Fri: "F", Sat: "S"},
				Short:       ut.CalendarDayFormatNameValue{Sun: "So.", Mon: "Mo.", Tue: "Di.", Wed: "Mi.", Thu: "Do.", Fri: "Fr.", Sat: "Sa."},
				Wide:        ut.CalendarDayFormatNameValue{Sun: "Sonntag", Mon: "Montag", Tue: "Dienstag", Wed: "Mittwoch", Thu: "Donnerstag", Fri: "Freitag", Sat: "Samstag"},
			},
			Periods: ut.CalendarPeriodFormatNames{
				Abbreviated: ut.CalendarPeriodFormatNameValue{AM: "vorm.", PM: "nachm."},
				Narrow:      ut.CalendarPeriodFormatNameValue{AM: "vorm.", PM: "nachm."},
				Short:       ut.CalendarPeriodFormatNameValue{},
				Wide:        ut.CalendarPeriodFormatNameValue{AM: "vorm.", PM: "nachm."},
			},
		},
	}
}
