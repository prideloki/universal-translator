package ar

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MDL": ut.Currency{Currency: "MDL", DisplayName: "ليو مولدوفي", Symbol: "MDL"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "بيزو فضي مكسيكي - 1861-1992", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "بولا بتسواني", Symbol: "BWP"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "بيزو كوبي قابل للتحويل", Symbol: "CUC"}, "USS": ut.Currency{Currency: "USS", DisplayName: "دولار أمريكي (نفس اليوم)\u200f", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "الوحدة المالية الأوروبية", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "كود اختبار العملة", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "كينا بابوا غينيا الجديدة", Symbol: "PGK"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "دولار روديسي", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "زائير زائيري جديد", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "لوتي ليسوتو", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "فرنك لوكسمبرج", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "دينار صربي", Symbol: "RSD"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "فضة", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "بيزو أرجنتيني", Symbol: "ARS"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "كوتزال جواتيمالا", Symbol: "GTQ"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "دولار غيانا", Symbol: "GYD"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "روبل سوفيتي", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "دولار برمودي", Symbol: "BMD"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "جنيه جنوب السودان", Symbol: "ج.ج.س."}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "أفغاني", Symbol: "AFN"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "دولار باهامي", Symbol: "BSD"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "بيزو مكسيكي", Symbol: "MX$"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "روبل طاجيكستاني", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "شلن أوغندي", Symbol: "UGX"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "عملة غير معروفة", Symbol: "***"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "زائير زائيري", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "بيزو كولومبي", Symbol: "COP"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "ين ياباني", Symbol: "JP¥"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "ليتا ليتوانية", Symbol: "LTL"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "توغروغ منغولي", Symbol: "MNT"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "ريال قطري", Symbol: "ر.ق.\u200f"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "مارك البوسنة والهرسك قابل للتحويل", Symbol: "BAM"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "لارى جورجي", Symbol: "GEL"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "دينار كويتي", Symbol: "د.ك.\u200f"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "أوقية موريتانية", Symbol: "أ.م.\u200f"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "وحدة النقد الأوروبية", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "روفيه جزر المالديف", Symbol: "MVR"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "سوم أوزبكستاني", Symbol: "UZS"}, "VND": ut.Currency{Currency: "VND", DisplayName: "دونج فيتنامي", Symbol: "₫"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "فرنك سي إف بي", Symbol: "CFPF"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "فلورن أروبي", Symbol: "AWG"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "دينار عراقي", Symbol: "د.ع.\u200f"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "دينار ليبي", Symbol: "د.ل.\u200f"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "سول جديد البيرو", Symbol: "PEN"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "ليو روماني قديم", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "فاتو فانواتو", Symbol: "VUV"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "كيات بورمي", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "تالوناس ليتواني", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "جنيه إسرائيلي", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "شلن صومالي", Symbol: "SOS"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "ليرة تركي", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "دينار يوغسلافي", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "بيستا أندوري", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "سوم قيرغستاني", Symbol: "KGS"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "دولار سورينامي", Symbol: "SRD"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "استرال أرجنتيني", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "رينغيت ماليزي", Symbol: "MYR"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "زلوتي بولندي - 1950-1995", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "مانات أذربيجان", Symbol: "AZN"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "دينار صربي قديم", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "أوستمارك ألماني شرقي", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "كرونة دانماركي", Symbol: "DKK"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "بيزو غينيا بيساو", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "روبية إندونيسية", Symbol: "ر.إن."}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "نايرا نيجيري", Symbol: "NGN"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "روبل روسي", Symbol: "RUB"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "دينار يمني", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "دولار بربادوسي", Symbol: "BBD"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "بيزو بوليفي", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "فرنك كونغولي", Symbol: "CDF"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "جلدر هولندي", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "بالبوا بنمي", Symbol: "PAB"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "مانات تركمانستان", Symbol: "TMT"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "جوردى هايتي", Symbol: "HTG"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "مانات تركمنستاني", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "ذهب", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "البلاتين", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "دينار البوسنة والهرسك", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "بيزيتا إسباني", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "كرونه أيسلندي", Symbol: "ISK"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "قرطبة نيكاراغوا", Symbol: "NIO"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "روبية نيبالي", Symbol: "NPR"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "اسكود تيموري", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "فرنك فرنسي ذهبي", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "كروزايرو برازيلي - 1990-1993", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "جنيه إسترليني", Symbol: "£"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "بيزو أوروجواي - 1975-1993", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "الوحدة الأوروبية المركبة", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "راند جنوب أفريقيا", Symbol: "ZAR"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "دولار زمبابوي", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "يوان صيني", Symbol: "CN¥"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "جواراني باراجواي", Symbol: "PYG"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "دينار سوداني", Symbol: "د.س.\u200f"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "درام أرميني", Symbol: "AMD"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "فرنك سويسري", Symbol: "CHF"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "يورو", Symbol: "€"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "دولار هونغ كونغ", Symbol: "HK$"}, "USN": ut.Currency{Currency: "USN", DisplayName: "دولار أمريكي (اليوم التالي)\u200f", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "ليف بلغاري", Symbol: "BGN"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "كروزادو برازيلي", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "دينار تونسي", Symbol: "د.ت.\u200f"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "كوانزا أنجولي معدلة - 1995 - 1999", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "تاكا بنجلاديشي", Symbol: "BDT"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "فرنك غينيا", Symbol: "GNF"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "روبية باكستاني", Symbol: "ر.ب."}, "THB": ut.Currency{Currency: "THB", DisplayName: "باخت تايلاندي", Symbol: "฿"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "شلن تنزاني", Symbol: "TZS"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "حقوق السحب الخاصة", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "درهم إماراتي", Symbol: "د.إ.\u200f"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "وون كوريا الجنوبية", Symbol: "₩"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "روبل لاتفيا", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "كولون سلفادوري", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "فرنك مدغشقر", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "دولار ناميبي", Symbol: "NAD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "بيزو اوروغواي", Symbol: "UYU"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "اسكود برتغالي غينيا", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "فرنك لوكسمبرج المالي", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "فرنك غرب أفريقي", Symbol: "CFA"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "دينار يوغسلافي قابل للتحويل", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "كواشا زامبي - 1968-2012", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "بيزو أرجنتيني - 1983-1985", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "بيزو الدومنيكان", Symbol: "DOP"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "سيدي غانا", Symbol: "GHS"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "وون كوريا الشمالية", Symbol: "KPW"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "كواشا مالاوي", Symbol: "MWK"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "دولار سنغافوري", Symbol: "SGD"}, "INR": ut.Currency{Currency: "INR", DisplayName: "روبيه هندي", Symbol: "₹"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "كوانزا أنجولي", Symbol: "AOA"}, "AON": ut.Currency{Currency: "AON", DisplayName: "كوانزا أنجولي جديدة - 1990-2000", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "دلاسي جامبي", Symbol: "GMD"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "شلن أوغندي - 1966-1987", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "روبل بيلاروسي", Symbol: "BYR"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "بير أثيوبي", Symbol: "ETB"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "كيب لاوسي", Symbol: "LAK"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "روبية موريشيوسية", Symbol: "MUR"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "فرنك رواندي", Symbol: "RWF"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "مفدول بوليفي", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "فورينت مجري", Symbol: "HUF"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "ريال سعودي", Symbol: "ر.س.\u200f"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "نولتوم بوتاني", Symbol: "BTN"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "متكال موزمبيقي", Symbol: "MZN"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "دولار ترينداد وتوباجو", Symbol: "TTD"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "(UIC)فرنك فرنسي", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "ريال يمني", Symbol: "ر.ي.\u200f"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "كونا كرواتي", Symbol: "HRK"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "ريال إيراني", Symbol: "ر.إ."}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "دولار جامايكي", Symbol: "JMD"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "فرنك جزر القمر", Symbol: "ف.ج.ق.\u200f"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "جلدر سورينامي", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "جنيه إيرلندي", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "دولار جزر كيمن", Symbol: "KYD"}, "STD": ut.Currency{Currency: "STD", DisplayName: "دوبرا ساو تومي وبرينسيبي", Symbol: "STD"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "ليك ألباني", Symbol: "ALL"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "سيلي غينيا", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "رييال كمبودي", Symbol: "KHR"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "جنيه مالطي", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "دولار بليزي", Symbol: "BZD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "شلن كينيي", Symbol: "KES"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "اسكود موزمبيقي", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ريال عماني", Symbol: "ر.ع.\u200f"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "جنيه جزر فوكلاند", Symbol: "FKP"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "دينار أردني", Symbol: "د.أ.\u200f"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "لاتس لاتفيا", Symbol: "LVL"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "زلوتي بولندي", Symbol: "PLN"}, "USD": ut.Currency{Currency: "USD", DisplayName: "دولار أمريكي", Symbol: "US$"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "بوليفار فنزويلي", Symbol: "VEF"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "(XBD)وحدة الحساب الأوروبية", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "دولار كندي", Symbol: "CA$"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "كرونة استونية", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "جنيه سوداني قديم", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "فرنك بروندي", Symbol: "BIF"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "ليمبيرا هنداروس", Symbol: "HNL"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "فرنك مغربي", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "كيات ميانمار", Symbol: "MMK"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "باتاكا ماكاوي", Symbol: "MOP"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "بيزو فلبيني", Symbol: "PHP"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "ليلانجيني سوازيلندي", Symbol: "SZL"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "بانغا تونغا", Symbol: "TOP"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "راند جنوب أفريقيا -مالي", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "نوفو كروزايرو برازيلي - 1967-1986", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "جنيه مصري", Symbol: "ج.م.\u200f"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "دولار نيوزيلندي", Symbol: "NZ$"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "ليرة سورية", Symbol: "ل.س.\u200f"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "روبل بيلاروسي جديد - 1994-1999", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "أرياري مدغشقر", Symbol: "MGA"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "كرونة سلوفاكية", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "فرنك بلجيكي قابل للتحويل", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "ريال برازيلي", Symbol: "R$"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "اكويل جونينا غينيا الاستوائيّة", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "ليرة إيطالية", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "كرونة نرويجية", Symbol: "NOK"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "دولار أسترالي", Symbol: "AU$"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "بيزو شيلي", Symbol: "CLP"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "كرونة تشيكية", Symbol: "CZK"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "دينار كرواتي", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "شيكل إسرائيلي جديد", Symbol: "₪"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "تينغ كازاخستاني", Symbol: "KZT"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "دولار جزر سليمان", Symbol: "SBD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "ليون سيراليوني", Symbol: "SLL"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "دراخما يوناني", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "مانات أذريبجاني", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "فرنك بلجيكي مالي", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "فرنك لوكسمبرج قابل للتحويل", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "اسكودو الرأس الخضراء", Symbol: "CVE"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "كرونة سويدية", Symbol: "SEK"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "هريفنيا أوكراني", Symbol: "UAH"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "بوليفار فنزويلي - 1871-2008", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "فرنك وسط أفريقي", Symbol: "FCFA"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "جنيه قبرصي", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "ناكفا أريتري", Symbol: "ERN"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "دولار فيجي", Symbol: "FJD"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "روبل روسي - 1991-1998", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "دولار تايواني", Symbol: "NT$"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "بيزو كوبي", Symbol: "CUP"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "فرنك جيبوتي", Symbol: "DJF"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "ماركا فنلندي", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "فرنك فرنسي", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "دينار مقدوني", Symbol: "MKD"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "دولار شرق الكاريبي", Symbol: "EC$"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "دينار بحريني", Symbol: "د.ب.\u200f"}, "BND": ut.Currency{Currency: "BND", DisplayName: "دولار بروناي", Symbol: "BND"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "أفغاني - 1927-2002", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "فرنك بلجيكي", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "مارك ألماني", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "درهم مغربي", Symbol: "د.م.\u200f"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "فرنك مالي", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "روبية سيشيلية", Symbol: "SCR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "جنيه سانت هيلين", Symbol: "SHP"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "ليرة تركية", Symbol: "ل.ت."}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "دولار زمبابوي 2009", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "جنيه جبل طارق", Symbol: "GIP"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "روبية سريلانكية", Symbol: "LKR"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "غيلدر أنتيلي هولندي", Symbol: "ANG"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "دينار جزائري", Symbol: "د.ج.\u200f"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "دولار ليبيري", Symbol: "LRD"}, "RON": ut.Currency{Currency: "RON", DisplayName: "ليو روماني", Symbol: "RON"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "جنيه سوداني", Symbol: "ج.س."}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "تولار سلوفيني", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "تالا ساموا", Symbol: "WST"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "كوانزا أنجولي - 1977-1990", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "بوليفيانو بوليفي", Symbol: "BOB"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "سيدي غاني", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "جنيه لبناني", Symbol: "ل.ل.\u200f"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "ليرة مالطية", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "شلن نمساوي", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "كولن كوستا ريكي", Symbol: "CRC"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "كرونة تشيكوسلوفاكيا", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "كوردوبة نيكاراجوا", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "اسكود برتغالي", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "سوموني طاجيكستاني", Symbol: "TJS"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "الوحدة الحسابية الأوروبية", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "بالاديوم", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "كواشا زامبي", Symbol: "ZMW"}}