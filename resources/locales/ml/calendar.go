package ml

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "y, MMMM d, EEEE", Long: "y, MMMM d", Medium: "y, MMM d", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "ഓഗ", 10: "ഒക്ടോ", 11: "നവം", 12: "ഡിസം", 1: "ജനു", 4: "ഏപ്രി", 5: "മേയ്", 7: "ജൂലൈ", 9: "സെപ്റ്റം", 2: "ഫെബ്രു", 3: "മാർ", 6: "ജൂൺ"}, Narrow: ut.CalendarMonthFormatNameValue{1: "ജ", 3: "മാ", 6: "ജൂ", 10: "ഒ", 11: "ന", 2: "ഫെ", 4: "ഏ", 5: "മെ", 7: "ജൂ", 8: "ഓ", 9: "സെ", 12: "ഡി"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "മാർച്ച്", 5: "മേയ്", 6: "ജൂൺ", 7: "ജൂലൈ", 11: "നവംബർ", 12: "ഡിസംബർ", 1: "ജനുവരി", 2: "ഫെബ്രുവരി", 4: "ഏപ്രിൽ", 8: "ഓഗസ്റ്റ്", 9: "സെപ്റ്റംബർ", 10: "ഒക്\u200cടോബർ"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "ഞായർ", 1: "തിങ്കൾ", 2: "ചൊവ്വ", 3: "ബുധൻ", 4: "വ്യാഴം", 5: "വെള്ളി", 6: "ശനി"}, Narrow: ut.CalendarDayFormatNameValue{0: "ഞാ", 1: "തി", 2: "ചൊ", 3: "ബു", 4: "വ്യാ", 5: "വെ", 6: "ശ"}, Short: ut.CalendarDayFormatNameValue{0: "ഞാ", 1: "തി", 2: "ചൊ", 3: "ബു", 4: "വ്യാ", 5: "വെ", 6: "ശ"}, Wide: ut.CalendarDayFormatNameValue{0: "ഞായറാഴ്\u200cച", 1: "തിങ്കളാഴ്\u200cച", 2: "ചൊവ്വാഴ്\u200cച", 3: "ബുധനാഴ്\u200cച", 4: "വ്യാഴാഴ്\u200cച", 5: "വെള്ളിയാഴ്\u200cച", 6: "ശനിയാഴ്\u200cച"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "morning1": "പുലർച്ചെ", "morning2": "രാവിലെ", "afternoon2": "ഉച്ചതിരിഞ്ഞ്", "night1": "രാത്രി", "midnight": "അർദ്ധരാത്രി", "noon": "ഉച്ച", "pm": "PM", "afternoon1": "ഉച്ചയ്ക്ക്", "evening1": "വൈകുന്നേരം", "evening2": "സന്ധ്യ"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "അർദ്ധരാത്രി", "am": "AM", "morning2": "രാവിലെ", "evening2": "സന്ധ്യ", "night1": "രാത്രി", "noon": "ഉച്ച", "pm": "PM", "morning1": "പുലർച്ചെ", "afternoon1": "ഉച്ചയ്ക്ക്", "afternoon2": "ഉച്ചതിരിഞ്ഞ്", "evening1": "വൈകുന്നേരം"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"afternoon2": "ഉച്ചതിരിഞ്ഞ്", "evening1": "വൈകുന്നേരം", "midnight": "അർദ്ധരാത്രി", "noon": "ഉച്ച", "pm": "PM", "afternoon1": "ഉച്ചയ്ക്ക്", "night1": "രാത്രി", "am": "AM", "morning1": "പുലർച്ചെ", "morning2": "രാവിലെ", "evening2": "സന്ധ്യ"}}}}