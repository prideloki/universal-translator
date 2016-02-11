package ne

import "github.com/go-playground/universal-translator"

func newCurrencies() []ut.Currency {
	return []ut.Currency{
		{Currency: "AED", DisplayName: "संयुक्त अरब एमिराट्स डिर्हाम", Symbol: "AED"},
		{Currency: "AFA", DisplayName: "अफ्गानी(१९२७–२००२)", Symbol: ""},
		{Currency: "AFN", DisplayName: "अफ्गान अफ्गानी", Symbol: "AFN"},
		{Currency: "ALL", DisplayName: "अल्बानियन लेक", Symbol: "ALL"},
		{Currency: "AMD", DisplayName: "आर्मेनियाली ड्राम", Symbol: "AMD"},
		{Currency: "ANG", DisplayName: "नेदरल्याण्ड्स एन्टिलियन गिल्डर", Symbol: "ANG"},
		{Currency: "AOA", DisplayName: "एङ्गोलान क्वान्जा", Symbol: "AOA"},
		{Currency: "ARS", DisplayName: "अर्जेन्टिनी पेसो", Symbol: "ARS"},
		{Currency: "AUD", DisplayName: "अष्ट्रेलियन डलर", Symbol: "A$"},
		{Currency: "AWG", DisplayName: "आरूबन फ्लोरिन", Symbol: "AWG"},
		{Currency: "AZN", DisplayName: "अजरबैजानी मानात", Symbol: "AZN"},
		{Currency: "BAM", DisplayName: "बोस्निया-हर्जगोभिनिया रूपान्तरयोग्य मार्क", Symbol: "BAM"},
		{Currency: "BBD", DisplayName: "बर्बाडियन डलर", Symbol: "BBD"},
		{Currency: "BDT", DisplayName: "बङ्गलादेशी टाका", Symbol: "BDT"},
		{Currency: "BGN", DisplayName: "बुल्गारियाली लेभ", Symbol: "BGN"},
		{Currency: "BHD", DisplayName: "बाहारैनी डिनार", Symbol: "BHD"},
		{Currency: "BIF", DisplayName: "बुरूण्डियाली फ्रान्क", Symbol: "BIF"},
		{Currency: "BMD", DisplayName: "बर्मुडन डलर", Symbol: "BMD"},
		{Currency: "BND", DisplayName: "ब्रुनाई डलर", Symbol: "BND"},
		{Currency: "BOB", DisplayName: "बोलिभियन बोलिभियानो", Symbol: "BOB"},
		{Currency: "BRL", DisplayName: "ब्राजिलियन रियल", Symbol: "R$"},
		{Currency: "BSD", DisplayName: "बहामियाली डलर", Symbol: "BSD"},
		{Currency: "BTN", DisplayName: "भुटानी एन्\u200cगुल्ट्रुम", Symbol: "BTN"},
		{Currency: "BWP", DisplayName: "बोट्सवानान पुला", Symbol: "BWP"},
		{Currency: "BYR", DisplayName: "बेलारूसी रूबल", Symbol: "BYR"},
		{Currency: "BZD", DisplayName: "वेलिज डलर", Symbol: "BZD"},
		{Currency: "CAD", DisplayName: "क्यानाडियाली डलर", Symbol: "CA$"},
		{Currency: "CDF", DisplayName: "कङ्गोली फ्रान्क", Symbol: "CDF"},
		{Currency: "CHF", DisplayName: "स्विस् फ्रैङ्क", Symbol: "CHF"},
		{Currency: "CLP", DisplayName: "चिलियन पेसो", Symbol: "CLP"},
		{Currency: "CNY", DisplayName: "चिनिँया युआन", Symbol: "CN¥"},
		{Currency: "COP", DisplayName: "कोलम्वियन पेसो", Symbol: "COP"},
		{Currency: "CRC", DisplayName: "कोष्टारिकन कोलोन", Symbol: "CRC"},
		{Currency: "CUC", DisplayName: "क्यूवाली रूपान्तरणयोग्य पेसो", Symbol: "CUC"},
		{Currency: "CUP", DisplayName: "क्यूवाली पेसो", Symbol: "CUP"},
		{Currency: "CVE", DisplayName: "केप भर्डियन एस्कुडो", Symbol: "CVE"},
		{Currency: "CZK", DisplayName: "चेख गणतञ्त्र कोरूना", Symbol: "CZK"},
		{Currency: "DJF", DisplayName: "जिबौंटियाली फ्रान्क", Symbol: "DJF"},
		{Currency: "DKK", DisplayName: "ड्यानिश क्रोन", Symbol: "DKK"},
		{Currency: "DOP", DisplayName: "डोमिनिकन पेसो", Symbol: "DOP"},
		{Currency: "DZD", DisplayName: "अल्जेरियाली डिनार", Symbol: "DZD"},
		{Currency: "EGP", DisplayName: "इजिप्सियन पाउन्ड", Symbol: "EGP"},
		{Currency: "ERN", DisplayName: "एरिट्रियन नाक्फा", Symbol: "ERN"},
		{Currency: "ETB", DisplayName: "इथियोपियाली बिर", Symbol: "ETB"},
		{Currency: "EUR", DisplayName: "युरो", Symbol: "€"},
		{Currency: "FJD", DisplayName: "फिजीयाली डलर", Symbol: "FJD"},
		{Currency: "FKP", DisplayName: "फक्\u200cल्याण्ड टापुहरूका पाउन्ड", Symbol: "FKP"},
		{Currency: "GBP", DisplayName: "बेलायती पाउण्ड स्टर्लिङ", Symbol: "£"},
		{Currency: "GEL", DisplayName: "जर्जियाली लारी", Symbol: "GEL"},
		{Currency: "GHS", DisplayName: "घानाली सेडी", Symbol: "GHS"},
		{Currency: "GIP", DisplayName: "जिब्राल्टर पाउण्ड", Symbol: "GIP"},
		{Currency: "GMD", DisplayName: "गाम्वियाली डालासी", Symbol: "GMD"},
		{Currency: "GNF", DisplayName: "गिनियाली फ्रान्क", Symbol: "GNF"},
		{Currency: "GTQ", DisplayName: "ग्वाटेमाला क्वेट्जाल", Symbol: "GTQ"},
		{Currency: "GYD", DisplayName: "गाइनिज डलर", Symbol: "GYD"},
		{Currency: "HKD", DisplayName: "हङकङ डलर", Symbol: "HK$"},
		{Currency: "HNL", DisplayName: "होन्डुरान लेम्पिरा", Symbol: "HNL"},
		{Currency: "HRK", DisplayName: "क्रोएशियाली कुना", Symbol: "HRK"},
		{Currency: "HTG", DisplayName: "हैटियाली गुर्ड", Symbol: "HTG"},
		{Currency: "HUF", DisplayName: "हङ्गेरियन फोरिन्ट", Symbol: "HUF"},
		{Currency: "IDR", DisplayName: "इण्डोनेशियाली रूपियाँ", Symbol: "IDR"},
		{Currency: "ILS", DisplayName: "इजरायली नयाँ शेकेल", Symbol: "₪"},
		{Currency: "INR", DisplayName: "भारतीय रूपिँया", Symbol: "₹"},
		{Currency: "IQD", DisplayName: "इराकी डिनार", Symbol: "IQD"},
		{Currency: "IRR", DisplayName: "इरानियाली रियाल", Symbol: "IRR"},
		{Currency: "ISK", DisplayName: "आइसल्याण्डिक क्रोना", Symbol: "ISK"},
		{Currency: "JMD", DisplayName: "जमाइकाली डलर", Symbol: "JMD"},
		{Currency: "JOD", DisplayName: "जोर्डानियाली डलर", Symbol: "JOD"},
		{Currency: "JPY", DisplayName: "जापानी येन", Symbol: "JP¥"},
		{Currency: "KES", DisplayName: "केन्याली शिलिङ", Symbol: "KES"},
		{Currency: "KGS", DisplayName: "किर्गिस्तानी सोम", Symbol: "KGS"},
		{Currency: "KHR", DisplayName: "कम्बोडिनेयाली रियल", Symbol: "KHR"},
		{Currency: "KMF", DisplayName: "कोमोरियन फ्रान्क", Symbol: "KMF"},
		{Currency: "KPW", DisplayName: "उत्तर कोरियाली वन", Symbol: "KPW"},
		{Currency: "KRW", DisplayName: "दक्षिण कोरियाली वन", Symbol: "₩"},
		{Currency: "KWD", DisplayName: "कुवेती डिनार", Symbol: "KWD"},
		{Currency: "KYD", DisplayName: "केम्यान टापुहरूका डलर", Symbol: "KYD"},
		{Currency: "KZT", DisplayName: "काजाखस्तानी टेन्ज", Symbol: "KZT"},
		{Currency: "LAK", DisplayName: "लाओशियन किप", Symbol: "LAK"},
		{Currency: "LBP", DisplayName: "लेबनाली पाउन्ड", Symbol: "LBP"},
		{Currency: "LKR", DisplayName: "श्रीलङ्काली रूपिया", Symbol: "LKR"},
		{Currency: "LRD", DisplayName: "लिबेरियाली डलर", Symbol: "LRD"},
		{Currency: "LTL", DisplayName: "लिथुनियाली लिटास", Symbol: "LTL"},
		{Currency: "LVL", DisplayName: "लाट्भियाली लाट्स", Symbol: "LVL"},
		{Currency: "LYD", DisplayName: "लिवियाली डिनार", Symbol: "LYD"},
		{Currency: "MAD", DisplayName: "मोरोक्काली डिर्\u200cहाम", Symbol: "MAD"},
		{Currency: "MDL", DisplayName: "माल्डोभन लेउ", Symbol: "MDL"},
		{Currency: "MGA", DisplayName: "मालागासी एरिआरी", Symbol: "MGA"},
		{Currency: "MKD", DisplayName: "म्यासेडोनियाली डेनार", Symbol: "MKD"},
		{Currency: "MMK", DisplayName: "म्यान्मा क्याट", Symbol: "MMK"},
		{Currency: "MNT", DisplayName: "मङ्गोलियाली टुग्रिक", Symbol: "MNT"},
		{Currency: "MOP", DisplayName: "माकानिज पटाका", Symbol: "MOP"},
		{Currency: "MRO", DisplayName: "माउरिटानियानली औगुइया", Symbol: "MRO"},
		{Currency: "MUR", DisplayName: "माउरिटियन रूपी", Symbol: "MUR"},
		{Currency: "MVR", DisplayName: "मालडिभियाली रूफियाँ", Symbol: "MVR"},
		{Currency: "MWK", DisplayName: "मलाविअन क्वाचा", Symbol: "MWK"},
		{Currency: "MXN", DisplayName: "मेक्सिकन पेसो", Symbol: "MX$"},
		{Currency: "MYR", DisplayName: "मलेशियाली रिङ्गेट", Symbol: "MYR"},
		{Currency: "MZN", DisplayName: "मोजाम्विकन मेटिकल", Symbol: "MZN"},
		{Currency: "NAD", DisplayName: "नामिबियन डलर", Symbol: "NAD"},
		{Currency: "NGN", DisplayName: "नाइजेरियन नाइरा", Symbol: "NGN"},
		{Currency: "NIO", DisplayName: "निकारागुवान कोर्डोवा", Symbol: "NIO"},
		{Currency: "NOK", DisplayName: "नर्वेजियाली क्रोन", Symbol: "NOK"},
		{Currency: "NPR", DisplayName: "नेपाली रूपैयाँ", Symbol: "नेरू"},
		{Currency: "NZD", DisplayName: "न्यूजिल्याण्ड डलर", Symbol: "NZ$"},
		{Currency: "OMR", DisplayName: "ओमनी रियल", Symbol: "OMR"},
		{Currency: "PAB", DisplayName: "पानामानियाली बाल्बोआ", Symbol: "PAB"},
		{Currency: "PEN", DisplayName: "पेरूभियाली न्यूभो सोल", Symbol: "PEN"},
		{Currency: "PGK", DisplayName: "पपुआ न्यू गिनियाली किना", Symbol: "PGK"},
		{Currency: "PHP", DisplayName: "फिलिपिनी पेसो", Symbol: "PHP"},
		{Currency: "PKR", DisplayName: "पाकिस्तानी रूपियाँ", Symbol: "PKR"},
		{Currency: "PLN", DisplayName: "पोलिश ज्लोटाई", Symbol: "PLN"},
		{Currency: "PYG", DisplayName: "पारागुयाली गुरानी", Symbol: "PYG"},
		{Currency: "QAR", DisplayName: "कतारी रियल", Symbol: "QAR"},
		{Currency: "RON", DisplayName: "रोमानियाली लेऊ", Symbol: "RON"},
		{Currency: "RSD", DisplayName: "सर्बियाली डिनार", Symbol: "RSD"},
		{Currency: "RUB", DisplayName: "रूसी रूबल", Symbol: "RUB"},
		{Currency: "RWF", DisplayName: "र्\u200cवाण्डाली फ्रान्क", Symbol: "RWF"},
		{Currency: "SAR", DisplayName: "साउदी रियालहरू", Symbol: "SAR"},
		{Currency: "SBD", DisplayName: "सोलोमन टापुहरूका डलर", Symbol: "SBD"},
		{Currency: "SCR", DisplayName: "सेचेलोइस रूपी", Symbol: "SCR"},
		{Currency: "SDG", DisplayName: "सुडानी पाउन्ड", Symbol: "SDG"},
		{Currency: "SEK", DisplayName: "स्विडिश क्रोना", Symbol: "SEK"},
		{Currency: "SGD", DisplayName: "सिङ्गापुर डलर", Symbol: "SGD"},
		{Currency: "SHP", DisplayName: "सेन्ट हेलेना पाउन्ड", Symbol: "SHP"},
		{Currency: "SLL", DisplayName: "सियरा लियोनेन लियोन", Symbol: "SLL"},
		{Currency: "SOS", DisplayName: "सोमाली शिलिङ", Symbol: "SOS"},
		{Currency: "SRD", DisplayName: "सुरिनामिज डलर", Symbol: "SRD"},
		{Currency: "SSP", DisplayName: "दक्षिण सुडानी पाउन्ड", Symbol: "SSP"},
		{Currency: "STD", DisplayName: "साओ टोम र प्रिन्सिप डोब्रा", Symbol: "STD"},
		{Currency: "SYP", DisplayName: "सिरियाली पाउन्ड", Symbol: "SYP"},
		{Currency: "SZL", DisplayName: "स्वाजी लिलान्गेनी", Symbol: "SZL"},
		{Currency: "THB", DisplayName: "थाई भाट", Symbol: "฿"},
		{Currency: "TJS", DisplayName: "ताजिक्स्तानी सोमोनी", Symbol: "TJS"},
		{Currency: "TMT", DisplayName: "टुर्क्मेनिस्तानी मानात", Symbol: "TMT"},
		{Currency: "TND", DisplayName: "टुनिसियाली डिनार", Symbol: "TND"},
		{Currency: "TOP", DisplayName: "टङ्गन पाङ्गा", Symbol: "TOP"},
		{Currency: "TRY", DisplayName: "टर्किश लिरा", Symbol: "TRY"},
		{Currency: "TTD", DisplayName: "त्रिनिडाड र टोबागो डलर", Symbol: "TTD"},
		{Currency: "TWD", DisplayName: "नयाँ ताइवान डलर", Symbol: "NT$"},
		{Currency: "TZS", DisplayName: "ताञ्जानियाली शिलिङ", Symbol: "TZS"},
		{Currency: "UAH", DisplayName: "युक्रेनी हिर्भिनिया", Symbol: "UAH"},
		{Currency: "UGX", DisplayName: "युगाण्डाली शिलिङ", Symbol: "UGX"},
		{Currency: "USD", DisplayName: "अमेरिकी डलर", Symbol: "US$"},
		{Currency: "UYU", DisplayName: "उरूगुवायाली पेसो", Symbol: "UYU"},
		{Currency: "UZS", DisplayName: "उज्बेकिस्तान सोम", Symbol: "UZS"},
		{Currency: "VEF", DisplayName: "भेनेजुएलन बोलिभर", Symbol: "VEF"},
		{Currency: "VND", DisplayName: "भियतनामी डङ्", Symbol: "₫"},
		{Currency: "VUV", DisplayName: "भानुआतू भातु", Symbol: "VUV"},
		{Currency: "WST", DisplayName: "सामोआन ताला", Symbol: "WST"},
		{Currency: "XAF", DisplayName: "सीएफ्\u200cए फ्रान्क बीइएसी", Symbol: "FCFA"},
		{Currency: "XCD", DisplayName: "पूर्वी क्यारिबियन डलर", Symbol: "EC$"},
		{Currency: "XOF", DisplayName: "सीएफ्\u200cए फ्रान्क बीसीइएओ", Symbol: "CFA"},
		{Currency: "XPF", DisplayName: "सीएफ्\u200cपी फ्रान्क", Symbol: "CFPF"},
		{Currency: "XXX", DisplayName: "अज्ञात मुद्रा", Symbol: ""},
		{Currency: "YER", DisplayName: "येमेनी रियाल", Symbol: "YER"},
		{Currency: "ZAR", DisplayName: "दक्षिण अफ्रिकी र्\u200dयान्ड", Symbol: "ZAR"},
		{Currency: "ZMK", DisplayName: "जाम्बियाली क्वाचा (१९६८–२०१२)", Symbol: ""},
		{Currency: "ZMW", DisplayName: "जाम्बियाली क्वाचा", Symbol: "ZMW"},
	}
}
