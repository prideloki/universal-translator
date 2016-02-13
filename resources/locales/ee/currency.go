package ee

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BMD": ut.Currency{Currency: "BMD", DisplayName: "bermudaga dollar", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "omanga rial", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "europa gadzidzenu xbb", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "ɣetoɖofe afrikaga CFA franc BEAC", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "europa kɔnta dzidzenu xbd", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP ga franc", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "burundiga franc", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "turkmenistanga manat", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "uruguayga peso (1975–1993)", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "laosga kip", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "serbiaga dinar", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ugandaga shilling", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "andorraga peseta", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "kubaga peso", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "saipriɔtga pound", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "lithuaniaga litas", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "mongoliaga tugrik", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "sudanga pound (1957–1998)", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "guatemalaga quetzal", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "italiaga lira", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "kenyaga shilling", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "peruga nuevo sol (1863–1965)", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Taiwanga dollar", Symbol: "NT$"}, "USD": ut.Currency{Currency: "USD", DisplayName: "US ga dollar", Symbol: "US$"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "bosnia kple herzegovinaga dinar yeyètɔ (1994–1997)", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "tsɛk repɔblikga koruna", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "israelga yeyetɔ sheqel", Symbol: "₪"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "pɔtugaltɔwo ƒe giniga escudo", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "nepalga rupee", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghanistanga afghani", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "angolaga kwanza (1990–2000)", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "ghana siɖi (1979–2007)", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "spaniaga peseta (Convertible)", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "são tomé kple príncipega dobra", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "yugoslaviaga convertible dinar (1990–1992)", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "anyiehe afrikaga rand (gadzikpɔtɔ)", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "armeniaga dram", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "angolaga kwanza", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Canadaga dollar", Symbol: "CA$"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "yugoslaviaga hard dinar (1966–1990)", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "bosnia-herzegovinaga convertible mark", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "trinidad kple tobagoga dollar", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "ukrainega karbovanet", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "mɔritiusga rupee", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "uzbekistanga som", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "beldziumga franc (convertible)", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "kubaga convertible peso", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Denmarkga krone", Symbol: "DKK"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "saint helenaga pound", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "syriaga pound", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "fransemega sika franc", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "gaɖuɖu dodokpɔ dzesi xts", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "aldzeriaga dinar", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "libyaga dinar", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "sɛtselsga rupee", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "zimbabwega dollar (2008)", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "tsɛkoslovakiaga hard koruna", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "haitiga gourde", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "komoroga franc", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "germaniaga mark", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "frentsiga franc", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "venezuelaga bolívar (1871–2008)", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "angolaga kwanza (1977–1991)", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "bulgariaga lev", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "bhutanga ngultrum", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "serbiaga dinar (2002–2006)", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "ekuadɔ dzidzenu matrɔmatrɔ", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "giniga franc", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "irakga dinar", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "surinamga dollar", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "argentinaga peso (1983–1985)", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "argentinaga peso", Symbol: ""}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "bulgariaga lev (1879–1952)", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "klosalo", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "europa dzidzenu xba", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "bahamiaga dollar", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "maltaga lira", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "South Koreaga won", Symbol: "₩"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "sierra leonega leone", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "us ga dollar (ŋkeke si gbɔna tɔ)", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "barbadiaga dollar", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "dziboutiga franc", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "estoniaga kroon", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "venezuelaga bolívar", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "sika", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "zambiaga kwacha (1968–2012)", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR euro CHE", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "falklanɖ ƒudomekpo dukɔwo ƒe ga pound", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "sovietga rouble", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "slovaniaga tolar", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "samaoga tala", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "azerbaidzanga manat", Symbol: ""}, "COU": ut.Currency{Currency: "COU", DisplayName: "kolombiaga vavãtɔ", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "gambiaga dalasi", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "latviaga ruble", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "hollandga guilder", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "tunisiaga dinar", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "ɣedzeƒe caribbeaga dollar", Symbol: "EC$"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "yemeniga dinar", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afghanistanga afghani (1927–2002)", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Swissga franc", Symbol: "CHF"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hong Kongga dollar", Symbol: "HK$"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zairega zaire (1971–1993)", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "spaniaga peseta (A)", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "gini-bisau peso", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "panamaga balboa", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "tajikistanga somoni", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "ugandaga shilling (1966–1987)", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "zimbabwega dollar (1980–2008)", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "zimbabwega dollar (2009)", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "fidziga dollar", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "dzɔdziaga lari", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "slovakga koruna", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "albaniaga lek", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "ɣetoɖofe afrikaga CFA franc BCEAO", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "yugoslaviaga yeyetɔ dinar (1994–2002)", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "egyptega pound", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "grisiga drachma", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "hungariaga forint", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "makanesega pataca", Symbol: ""}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "yugoslaviaga dinar (1992–1993)", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "tsilega peso", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "swaziga lilangeni", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "beldziumga franc", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "maldiviaga rufiyaa", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudi Arabiaga riyal", Symbol: "SAR"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "sri lankaga rupee", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Norwayga krone", Symbol: "NOK"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "nikaraguaga córdoba (1988–1991)", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "ɣedzeƒe germaniaga mark", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "yɔdanga dinar", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "lebanonga pound", Symbol: ""}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "anyiehe koreaga hwan (1953–1962)", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "katarga rial", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "romaniaga leu (1952–2006)", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "brazilia cruzado (1986–1989)", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Chinesega yuan", Symbol: "CN¥"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "new zealanɖga dollar", Symbol: "NZ$"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "botswanaga pula", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "mɔritaniaga ouguiya", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "angolaga kwanza xoxotɔ (1995–1999)", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "belarusiaga ruble yeytɔ (1994–1999)", Symbol: ""}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "tsainatɔwo ƒe gadzraɖoƒe dollar", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "kolombiaga peso", Symbol: ""}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "makedoniaga denar (1992–1993)", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "us ga dollar (ŋkeke ma ke tɔ)", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "morokoga franc", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "malagasega ariary", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "malaysiaga ringit", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "nikaraguaga córdoba", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vanuatuga vatu", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "arubaga lorin", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "dominicaga peso", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "liberiaga dollar", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "myanmaga kyat", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkishga lira", Symbol: "TRY"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "vietnamga dong (1978–1985)", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "pakistaniga rupee", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "belizega dollar", Symbol: ""}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "aiselandga króna (1918–1981)", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "filipiniga peso", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "latviaga lats", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "rhodesiaga dollar", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "tajikistanga ruble", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "uruguayga peso UYI", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "uruguayga peso", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR euro CHW", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "guyanaga dollar", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "kuwaitga dinar", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "vietnamga dong", Symbol: "₫"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "moldovaga leu", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "surinamega guilder", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "timɔga escudo", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "azerbaidzanga manat (1993–2006)", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "burmaga kyat", Symbol: ""}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "monegaskga franc", Symbol: ""}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "israelga sheqel (1980–1985)", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "papua new guineaga kina", Symbol: ""}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "tsilega escudo", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "kosta rikaga kolón", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "kape verdega escudo", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "braziliaga cruzeiro xoxotɔ gbãtɔ (1990–1993)", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "peruga nuevo sol", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "turkmenistanga manat (1993–2009)", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rɔtsiaga ruble (1991–1998)", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "sudanga dinar (1992–2007)", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "bangladeshga taka", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "gilbrataga pound", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesiaga rupiah", Symbol: "IDR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "polanɖga zloty", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzaniatɔwofɛgadudu", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "kambodiaga riel", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Mexicoga peso", Symbol: "MX$"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "mozambikga metikal", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "kazakhstanga tenge", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "lazembɔgga franc", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "makedoniaga denar", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "maliga franc", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "europa gaɖuɖu", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "albaniaga lek (1946–1965)", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Australiaga dollar", Symbol: "AU$"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "dzɔdziaga kupon larit", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "bahrainga dinar", Symbol: ""}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "boliviaga boliviano (1863–1963)", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "lazembɔgga gadzikpɔ franc", Symbol: ""}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "moldovaga cupon", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "salvadɔga colón", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladiumga", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET gadodo XRE", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "boliviaga mvdol", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "giniga syli", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "lithuaniaga talonas", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "kirgistanga som", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "dziehe koreaga won", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "maltaga pound", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "argentinaga austral", Symbol: ""}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "bɔlgariaga socialist lev", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "belarusiaga ruble", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "namibiaga dollar", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "ethiopiaga birr", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Britishga pound", Symbol: "£"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "iranga rial", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "solomon ƒudomekpo dukɔwo ƒe ga dollar", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "argentinaga peso (1881–1970)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Braziliaga real", Symbol: "R$"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japanesega yen", Symbol: "JP¥"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "ekuadɔga sucre", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "kroatiaga dinar", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "peruga inti", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "sudanga pound", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "singapɔga dollar", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "tɛkiiga lira (1922–2005)", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "europa kɔnta dzidzenu xbc", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "eritreaga nakfa", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indiaga rupee", Symbol: "₹"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "lazembɔgga convertible franc", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "gaɖuɖu manya", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "gaɖuɖu ɖoɖo tɔxɛ", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "yemeniga rial", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "ireland pound", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naidzeriaga naira", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russiaga ruble", Symbol: "RUB"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "bɔlgariaga hard lev", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "israelga pound", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "polanɖga zloty (1950–1995)", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "somaliaga shilling", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "anyiehe sudanga pound", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "frentsi UIC-franc", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "nedalands antilleaga guilder", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "beldziumga franc (financial)", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "bruneiga dollar", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kroatiaga kuna", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "tonagaga pa’anga", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "argentinaga peso ley (1970–1983)", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "brazilia cruzado xoxotɔ (1989–1990)", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "EUR", Symbol: "€"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "South Africaga rand", Symbol: "ZAR"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "anyiehe koreaga won (1945–1953)", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "romaniaga leu", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Swedishga krone", Symbol: "SEK"}, "AED": ut.Currency{Currency: "AED", DisplayName: "united arab emiratesga dirham", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "ɔstriaga schilling", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platinum", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "aiselandga króna", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dzamaikaga dollar", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "rwandaga franc", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "Thailandga baht", Symbol: "฿"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "braziliaga cruzeiro (1993–1994)", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "braziliaga cruzeiro (1942–1967)", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "kongoga franc", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "boliviaga peso", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "finlandga markka", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "pɔtugalga escudo", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "ukrainega hryvnia", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "spaniaga peseta", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ghana siɖi", Symbol: "GH₵"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "malagasega franc", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "paraguayga guarani", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "tsilegakɔnta dzidzenu UF", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "honduraga lempira", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "lesotoga loti", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "morokoga dirham", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "zambiaga kwacha", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "bosnia-herzegovinaga dinar (1992–1994)", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviaga boliviano", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "braziliaga cruzeiro xoxotɔ (1967–1986)", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "zairega yeyetɔ zaire", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekuatorial giniga ekwele", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "kayman ƒudomekpoga dollar", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "malawiga kwacha", Symbol: ""}}