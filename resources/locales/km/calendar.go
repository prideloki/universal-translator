package km

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} នៅ {0}", Long: "{1} នៅ {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "សីហា", 9: "កញ្ញា", 10: "តុលា", 11: "វិច្ឆិកា", 2: "កុម្ភៈ", 3: "មីនា", 5: "ឧសភា", 7: "កក្កដា", 12: "ធ្នូ", 1: "មករា", 4: "មេសា", 6: "មិថុនា"}, Narrow: ut.CalendarMonthFormatNameValue{1: "1", 3: "3", 5: "5", 6: "6", 7: "7", 11: "11", 12: "12", 2: "2", 4: "4", 8: "8", 9: "9", 10: "10"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "មីនា", 4: "មេសា", 7: "កក្កដា", 12: "ធ្នូ", 2: "កុម្ភៈ", 5: "ឧសភា", 6: "មិថុនា", 8: "សីហា", 9: "កញ្ញា", 10: "តុលា", 11: "វិច្ឆិកា", 1: "មករា"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "ច័ន្ទ", 2: "អង្គារ", 3: "ពុធ", 4: "ព្រហស្បតិ៍", 5: "សុក្រ", 6: "សៅរ៍", 0: "អាទិត្យ"}, Narrow: ut.CalendarDayFormatNameValue{5: "សុ", 6: "ស", 0: "អា", 1: "ច", 2: "អ", 3: "ពុ", 4: "ព្រ"}, Short: ut.CalendarDayFormatNameValue{5: "សុក្រ", 6: "សៅរ៍", 0: "អាទិត្យ", 1: "ច័ន្ទ", 2: "អង្គារ", 3: "ពុធ", 4: "ព្រហស្បតិ៍"}, Wide: ut.CalendarDayFormatNameValue{5: "សុក្រ", 6: "សៅរ៍", 0: "អាទិត្យ", 1: "ច័ន្ទ", 2: "អង្គារ", 3: "ពុធ", 4: "ព្រហស្បតិ៍"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "រសៀល", "evening1": "ល្ងាច", "night1": "យប់", "midnight": "កណ្តាលអាធ្រាត្រ", "am": "ព្រឹក", "noon": "ថ្ងៃ\u200bត្រង់", "pm": "ល្ងាច", "morning1": "ព្រឹក"}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "ថ្ងៃ\u200bត្រង់", "pm": "ល្ងាច", "morning1": "ព្រឹក", "afternoon1": "រសៀល", "evening1": "ល្ងាច", "night1": "យប់", "midnight": "កណ្តាលអាធ្រាត្រ", "am": "ព្រឹក"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"morning1": "ព្រឹក", "afternoon1": "រសៀល", "evening1": "ល្ងាច", "night1": "យប់", "midnight": "កណ្តាលអាធ្រាត្រ", "am": "ព្រឹក", "noon": "ថ្ងៃ\u200bត្រង់", "pm": "ល្ងាច"}}}}