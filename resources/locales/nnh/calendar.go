package nnh

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE , 'lyɛ'̌ʼ d 'na' MMMM, y", Long: "'lyɛ'̌ʼ d 'na' MMMM, y", Medium: "d MMM, y", Short: "dd/MM/yy"}, Time: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}, DateTime: ut.CalendarDateFormat{Full: "{1},{0}", Long: "{1}, {0}", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "saŋ mbʉ̀ŋ", 10: "saŋ tàŋa tsetsáʼ", 11: "saŋ mejwoŋó", 12: "saŋ lùm", 2: "saŋ kàg ngwóŋ", 3: "saŋ lepyè shúm", 6: "saŋ njÿoláʼ", 7: "saŋ tyɛ̀b tyɛ̀b mbʉ̀ŋ", 1: "saŋ tsetsɛ̀ɛ lùm", 4: "saŋ cÿó", 5: "saŋ tsɛ̀ɛ cÿó", 9: "saŋ ngwɔ̀ʼ mbÿɛ"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "saŋ tsetsɛ̀ɛ lùm", 2: "saŋ kàg ngwóŋ", 3: "saŋ lepyè shúm", 6: "saŋ njÿoláʼ", 7: "saŋ tyɛ̀b tyɛ̀b mbʉ̀ŋ", 8: "saŋ mbʉ̀ŋ", 10: "saŋ tàŋa tsetsáʼ", 4: "saŋ cÿó", 5: "saŋ tsɛ̀ɛ cÿó", 9: "saŋ ngwɔ̀ʼ mbÿɛ", 11: "saŋ mejwoŋó", 12: "saŋ lùm"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "lyɛʼɛ́ sẅíŋtè", 1: "mvfò lyɛ̌ʼ", 2: "mbɔ́ɔntè mvfò lyɛ̌ʼ", 3: "tsètsɛ̀ɛ lyɛ̌ʼ", 4: "mbɔ́ɔntè tsetsɛ̀ɛ lyɛ̌ʼ", 5: "mvfò màga lyɛ̌ʼ", 6: "màga lyɛ̌ʼ"}, Narrow: ut.CalendarDayFormatNameValue(nil), Short: ut.CalendarDayFormatNameValue{0: "lyɛʼɛ́ sẅíŋtè", 1: "mvfò lyɛ̌ʼ", 2: "mbɔ́ɔntè mvfò lyɛ̌ʼ", 3: "tsètsɛ̀ɛ lyɛ̌ʼ", 4: "mbɔ́ɔntè tsetsɛ̀ɛ lyɛ̌ʼ", 5: "mvfò màga lyɛ̌ʼ", 6: "màga lyɛ̌ʼ"}, Wide: ut.CalendarDayFormatNameValue{0: "lyɛʼɛ́ sẅíŋtè", 1: "mvfò lyɛ̌ʼ", 2: "mbɔ́ɔntè mvfò lyɛ̌ʼ", 3: "tsètsɛ̀ɛ lyɛ̌ʼ", 4: "mbɔ́ɔntè tsetsɛ̀ɛ lyɛ̌ʼ", 5: "mvfò màga lyɛ̌ʼ", 6: "màga lyɛ̌ʼ"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "mbaʼámbaʼ", "pm": "ncwònzém"}}}}