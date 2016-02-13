package fr_MU

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiyaa maldivien", Symbol: "MVR"}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunisien", Symbol: "TND"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portugais", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unité d’investissement chilienne", Symbol: "CLF"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "dông vietnamien (1978–1985)", Symbol: "VNN"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "dollar de la Banque populaire chinoise", Symbol: "CNX"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "peso lourd argentin (1970–1983)", Symbol: "ARL"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar bahreïni", Symbol: "BHD"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unité de compte européenne (ECU)", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar irakien", Symbol: "IQD"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dollar surinamais", Symbol: "$SR"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "couronne slovaque", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rouble letton", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "franc WIR", Symbol: ""}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "shekel israélien (1980–1985)", Symbol: "ILR"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dollar des îles Salomon", Symbol: "$SB"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sud-africain", Symbol: "ZAR"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afghani (1927–2002)", Symbol: "AFA"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "métical", Symbol: "MZM"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franc CFA (BCEAO)", Symbol: "CFA"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "pa’anga tongan", Symbol: "TOP"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari géorgien", Symbol: "GEL"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentin (1983–1985)", Symbol: "ARP"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nouveau rouble biélorusse (1994–1999)", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "nouveau dinar yougoslave", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franc CFA (BEAC)", Symbol: "FCFA"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "livre des îles Malouines", Symbol: "£FK"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zloty (1950–1995)", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar soudanais", Symbol: "SDD"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha zambien", Symbol: "ZMW"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "couronne islandaise", Symbol: "ISK"}, "COU": ut.Currency{Currency: "COU", DisplayName: "unité de valeur réelle colombienne", Symbol: "COU"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platine", Symbol: "XPT"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti lesothan", Symbol: "lLS"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unité de compte européenne (XBD)", Symbol: "XBD"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "dollar zimbabwéen (2008)", Symbol: "ZWR"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "roupie des Seychelles", Symbol: "SCR"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso cubain convertible", Symbol: "CUC"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigérian", Symbol: "NGN"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "mark convertible bosniaque", Symbol: "BAM"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saoudien", Symbol: "SAR"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "droit de tirage spécial", Symbol: "DTS"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rouble soviétique", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "nouveau cruzeiro brésilien (1967–1986)", Symbol: "BRB"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbovanetz", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "peso argentin (1881–1970)", Symbol: "ARM"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "sucre", Symbol: "XSU"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaïre zaïrois", Symbol: "ZRZ"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franc comorien", Symbol: "KMF"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolivar vénézuélien (1871–2008)", Symbol: "VEB"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croate", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cédi", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldave", Symbol: "MDL"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franc luxembourgeois", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "dollar des États-Unis", Symbol: "$US"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu roumain", Symbol: "RON"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat myanmarais", Symbol: "MMK"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "livre sterling", Symbol: "£GB"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "livre chypriote", Symbol: "£CY"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekwélé équatoguinéen", Symbol: "GQE"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexicain", Symbol: "$MX"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lire italienne", Symbol: "₤IT"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongol", Symbol: "MNT"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dollar de Singapour", Symbol: "$SG"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azéri (1993–2006)", Symbol: "AZM"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nouveau dollar taïwanais", Symbol: "TWD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letton", Symbol: "LVL"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dollar bermudien", Symbol: "$BM"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "nouveau manat turkmène", Symbol: "TMT"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dollar brunéien", Symbol: "$BN"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franc rwandais", Symbol: "RWF"}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra santoméen", Symbol: "STD"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kirghize", Symbol: "KGS"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "roupie népalaise", Symbol: "NPR"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dollar jamaïcain", Symbol: "JMD"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "florin néerlandais", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol péruvien", Symbol: "PES"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambien (1968–2012)", Symbol: "ZMK"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dollar canadien", Symbol: "$CA"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franc financier luxembourgeois", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florin arubais", Symbol: "AWG"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franc marocain", Symbol: "fMA"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dollar zimbabwéen", Symbol: "ZWD"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawite", Symbol: "MWK"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo capverdien", Symbol: "CVE"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angolais", Symbol: "AOA"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "denar macédonien", Symbol: "MKD"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial qatari", Symbol: "QAR"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "coupon de lari géorgien", Symbol: "GEK"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lev bulgare", Symbol: "BGN"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula botswanais", Symbol: "BWP"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franc convertible luxembourgeois", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso bolivien", Symbol: "BOP"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "livre sud-soudanaise", Symbol: "SSP"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dollar rhodésien", Symbol: "$RH"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dollar barbadien", Symbol: "BBD"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "euro WIR", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "devise inconnue ou non valide", Symbol: "XXX"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dollar libérien", Symbol: "LRD"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croate", Symbol: "HRK"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sud-africain (financier)", Symbol: "ZAL"}, "AON": ut.Currency{Currency: "AON", DisplayName: "nouveau kwanza angolais (1990–2000)", Symbol: "AON"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haïtienne", Symbol: "HTG"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franc burundais", Symbol: "BIF"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "livre israélienne", Symbol: "£IL"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franc suisse", Symbol: "CHF"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "schilling autrichien", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso d’argent mexicain (1861–1992)", Symbol: "MXP"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brésilien (1990–1993)", Symbol: "BRE"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unité de valeur constante équatoriale (UVC)", Symbol: "ECV"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dông vietnamien", Symbol: "₫"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo de Guinée portugaise", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit malais", Symbol: "MYR"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franc belge (financier)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rouble biélorusse", Symbol: "BYR"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "unité de compte ADB", Symbol: "XUA"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituanien", Symbol: "LTL"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre équatorien", Symbol: "ECS"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya mauritanien", Symbol: "MRO"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dollar fidjien", Symbol: "$FJ"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira hondurien", Symbol: "HNL"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dollar des Caraïbes orientales", Symbol: "XCD"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "or", Symbol: "XAU"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta espagnole (compte A)", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambodgien", Symbol: "KHR"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht thaïlandais", Symbol: "THB"}, "INR": ut.Currency{Currency: "INR", DisplayName: "roupie indienne", Symbol: "₹"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "roupie pakistanaise", Symbol: "PKR"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentin", Symbol: "$AR"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso bissau-guinéen", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "livre de Sainte-Hélène", Symbol: "SHP"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominicain", Symbol: "DOP"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentin", Symbol: "ARA"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franc guinéen", Symbol: "GNF"}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial yéménite", Symbol: "YER"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar du Yémen", Symbol: "YDD"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "dinar serbo-monténégrin", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dollar zimbabwéen (2009)", Symbol: "ZWL"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franc belge", Symbol: "FB"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rouble tadjik", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "lek albanais (1947–1961)", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar koweïtien", Symbol: "KWD"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franc belge (convertible)", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "roupie mauricienne", Symbol: "Rs"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "boliviano bolivien (1863–1963)", Symbol: "BOL"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek albanais", Symbol: "ALL"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram arménien", Symbol: "AMD"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "couronne estonienne", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan renminbi chinois", Symbol: "CNY"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmène", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba", Symbol: "NIC"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dollar néo-zélandais", Symbol: "$NZ"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dollar du Guyana", Symbol: "GYD"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angolais (1977–1990)", Symbol: "AOK"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro", Symbol: "BRR"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "hwan sud-coréen (1953–1962)", Symbol: "KRH"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franc CFP", Symbol: "FCFP"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nafka érythréen", Symbol: "ERN"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "couronne danoise", Symbol: "DKK"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "roupie maldivienne", Symbol: "MVP"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol bolivien", Symbol: "BOV"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano bolivien", Symbol: "BOB"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won sud-coréen", Symbol: "₩"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo timorais", Symbol: "TPE"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dollar de Hong Kong", Symbol: "HKD"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "escudo chilien", Symbol: "CLE"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "couronne norvégienne", Symbol: "NOK"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tadjik", Symbol: "TJS"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambien", Symbol: "GMD"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "couronne suédoise", Symbol: "SEK"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dollar des îles Caïmans", Symbol: "KYD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca macanaise", Symbol: "MOP"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unité européenne composée", Symbol: "XBA"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dollar trinidadien", Symbol: "$TT"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dollar australien", Symbol: "$AU"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "livre de Gibraltar", Symbol: "£GI"}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham des Émirats arabes unis", Symbol: "AED"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "livre soudanaise", Symbol: "SDG"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti péruvien", Symbol: "PEI"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolivar vénézuélien", Symbol: "VEF"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso uruguayen (1975–1993)", Symbol: "UYP"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unité de compte européenne (XBC)", Symbol: "XBC"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "shilling tanzanien", Symbol: "TZS"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rouble russe", Symbol: "RUB"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip loatien", Symbol: "LAK"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "nouveau cruzado", Symbol: "BRN"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghani afghan", Symbol: "AFN"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli guinéen", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dollar bélizéen", Symbol: "$BZ"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franc malien", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dollar bahaméen", Symbol: "$BS"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "roupie srilankaise", Symbol: "LKR"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unité de conversion mexicaine (UDI)", Symbol: "MXV"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoan", Symbol: "WS$"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "hryvnia ukrainienne", Symbol: "UAH"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chilien", Symbol: "$CL"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "livre soudanaise (1956–2007)", Symbol: "SDP"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical mozambicain", Symbol: "MZN"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka bangladeshi", Symbol: "BDT"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "zloty polonais", Symbol: "PLN"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbe", Symbol: "RSD"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatémaltèque", Symbol: "GTQ"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladium", Symbol: "XPD"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "dinar yougoslave Noviy", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colón salvadorien", Symbol: "SVC"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar jordanien", Symbol: "JOD"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "roupie indonésienne", Symbol: "IDR"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franc djiboutien", Symbol: "DJF"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "livre turque", Symbol: "TRY"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dollar namibien", Symbol: "$NA"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguayen", Symbol: "PYG"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "ancien leu roumain", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "(devise de test)", Symbol: "XTS"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unité monétaire européenne", Symbol: "XBB"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brésilien (1986–1989)", Symbol: "BRC"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "drachme grecque", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "couronne tchèque", Symbol: "CZK"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "shilling somalien", Symbol: "SOS"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iranien", Symbol: "IRR"}, "KES": ut.Currency{Currency: "KES", DisplayName: "shilling kényan", Symbol: "KES"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cédi ghanéen", Symbol: "GHS"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franc congolais", Symbol: "CDF"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azéri", Symbol: "AZN"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum bouthanais", Symbol: "BTN"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "livre syrienne", Symbol: "SYP"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso uruguayen (unités indexées)", Symbol: "UYI"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta espagnole (compte convertible)", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "dollar des Etats-Unis (jour suivant)", Symbol: "USN"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorrane", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguayen", Symbol: "$UY"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franc UIC", Symbol: "XFU"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "shilling ougandais (1966–1987)", Symbol: "UGS"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "mark est-allemand", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nouveau sol péruvien", Symbol: "PEN"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libyen", Symbol: "LYD"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen japonais", Symbol: "JPY"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni swazi", Symbol: "SZL"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "livre irlandaise", Symbol: "£IE"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "argent", Symbol: "XAG"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar bosniaque", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forint hongrois", Symbol: "HUF"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituanien", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone sierra-léonais", Symbol: "SLL"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panaméen", Symbol: "PAB"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colón costaricain", Symbol: "CRC"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta espagnole", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso philippin", Symbol: "PHP"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omanais", Symbol: "OMR"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lire maltaise", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rouble russe (1991–1998)", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "livre égyptienne", Symbol: "EGP"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "livre libanaise", Symbol: "£LB"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazakh", Symbol: "KZT"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "shilling ougandais", Symbol: "UGX"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu vanuatuan", Symbol: "VUV"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "livre turque (1844–2005)", Symbol: "TRL"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "réal brésilien", Symbol: "R$"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tolar slovène", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "type de fonds RINET", Symbol: "XRE"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr éthiopien", Symbol: "ETB"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "won sud-coréen (1945–1953)", Symbol: "KRO"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franc or", Symbol: "XFO"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "mark allemand", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgache", Symbol: "MGA"}, "USS": ut.Currency{Currency: "USS", DisplayName: "dollar des Etats-Unis (jour même)", Symbol: "USS"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marocain", Symbol: "MAD"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nouveau zaïre zaïrien", Symbol: "ZRN"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franc malgache", Symbol: "Fmg"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "mark finlandais", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "livre maltaise", Symbol: "£MT"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum ouzbek", Symbol: "UZS"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franc français", Symbol: "F"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nouveau shekel israélien", Symbol: "₪"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar algérien", Symbol: "DZD"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso colombien", Symbol: "$CO"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "cruzeiro brésilien (1942–1967)", Symbol: "BRZ"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "florin surinamais", Symbol: "SRG"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar yougoslave convertible", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina papouan-néo-guinéen", Symbol: "PGK"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "florin antillais", Symbol: "ANG"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birman", Symbol: "BUK"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubain", Symbol: "CUP"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev bulgare (1962–1999)", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "couronne forte tchécoslovaque", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "córdoba oro nicaraguayen", Symbol: "NIO"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won nord-coréen", Symbol: "KPW"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo mozambicain", Symbol: "MZE"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza angolais réajusté (1995–1999)", Symbol: "AOR"}}