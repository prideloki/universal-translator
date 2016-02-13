package nb

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BYB": ut.Currency{Currency: "BYB", DisplayName: "hviterussiske nye rubler (1994–1999)", Symbol: "BYB"}, "COU": ut.Currency{Currency: "COU", DisplayName: "colombianske unidad de valor real", Symbol: "COU"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "spanske peseta (A–konto)", Symbol: "ESA"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "georgiske lari", Symbol: "GEL"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "guatemalanske quetzal", Symbol: "GTQ"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "peruanske nuevo sol", Symbol: "PEN"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "testvalutakode", Symbol: "XTS"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "brasilianske cruzeiro novo (1967–1986)", Symbol: "BRB"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "kirgisiske som", Symbol: "KGS"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "kambodsjanske riel", Symbol: "KHR"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "filippinske pesos", Symbol: "PHP"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "saudiarabiske riyaler", Symbol: "SAR"}, "USN": ut.Currency{Currency: "USN", DisplayName: "amerikanske dollar (neste dag)", Symbol: "USN"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "estiske kroon", Symbol: "EEK"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "finske mark", Symbol: "FIM"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "laotiske kip", Symbol: "LAK"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "madagassiske ariary", Symbol: "MGA"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "ukrainske karbovanetz", Symbol: "UAK"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "kubanske konvertible pesos", Symbol: "CUC"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "nicaraguanske córdoba", Symbol: "NIO"}, "AON": ut.Currency{Currency: "AON", DisplayName: "angolanske nye kwanza (1990–2000)", Symbol: "AON"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "sørkoreanske hwan (1953–1962)", Symbol: "KRH"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "sørkoreanske won (1945–1953)", Symbol: "KRO"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "nederlandske gylden", Symbol: "NLG"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "bolivianske boliviano (1863–1963)", Symbol: "BOL"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "meksikanske unidad de inversion (UDI)", Symbol: "MXV"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "zambiske kwacha", Symbol: "ZMW"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "spanske peseta (konvertibel konto)", Symbol: "ESB"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "ecuadorianske unidad de valor constante (UVC)", Symbol: "ECV"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hongkong-dollar", Symbol: "HKD"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "israelske pund", Symbol: "ILP"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "polske zloty", Symbol: "PLN"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "europeisk valutaenhet", Symbol: "XEU"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP-franc", Symbol: "XPF"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "argentinske pesos", Symbol: "ARS"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "meksikanske pesos", Symbol: "MXN"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "chilenske pesos", Symbol: "CLP"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "vietnamesiske dong (1978–1985)", Symbol: "VNN"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "costaricanske colón", Symbol: "CRC"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "bermudiske dollar", Symbol: "BMD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "libanesiske pund", Symbol: "LBP"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "australske dollar", Symbol: "AUD"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "arubiske floriner", Symbol: "AWG"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "bhutanske ngultrum", Symbol: "BTN"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "egyptiske pund", Symbol: "EGP"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "ungarske forinter", Symbol: "HUF"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "salvadoranske colon", Symbol: "SVC"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "trinidadiske dollar", Symbol: "TTD"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nye taiwanske dollar", Symbol: "TWD"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghansk afghani", Symbol: "AFN"}, "VND": ut.Currency{Currency: "VND", DisplayName: "vietnamesiske dong", Symbol: "VND"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR euro", Symbol: "CHE"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "moldovske leu", Symbol: "MDL"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "turkmenske manat", Symbol: "TMT"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "bolivianske mvdol", Symbol: "BOV"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "zimbabwiske dollar (1980–2008)", Symbol: "ZWD"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "slovenske tolar", Symbol: "SIT"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "latviske lats", Symbol: "LVL"}, "AED": ut.Currency{Currency: "AED", DisplayName: "emiratarabiske dirham", Symbol: "AED"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "djiboutiske franc", Symbol: "DJF"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "fijianske dollar", Symbol: "FJD"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "georgiske kupon larit", Symbol: "GEK"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "honduranske lempira", Symbol: "HNL"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "salomonske dollar", Symbol: "SBD"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "tyrkiske lire (1922–2005)", Symbol: "TRL"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vanuatiske vatu", Symbol: "VUV"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "bulgarske lev", Symbol: "BGN"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Kinas folkebank dollar", Symbol: "CNX"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "zambiske kwacha (1968–2012)", Symbol: "ZMK"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "andorranske pesetas", Symbol: "ADP"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "liberiske dollar", Symbol: "LRD"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "jugoslaviske noviy-dinarer", Symbol: "YUM"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR franc", Symbol: "CHW"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "greske drakmer", Symbol: "GRD"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "sudanesiske dinarer (1992–2007)", Symbol: "SDD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "uruguayanske pesos", Symbol: "UYU"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "argentinske pesos (1983–1985)", Symbol: "ARP"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "maldiviske rufiyaa", Symbol: "MVR"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "surinamske dollar", Symbol: "SRD"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "nye bosnisk-hercegovinske dinarer (1994–1997)", Symbol: "BAN"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "kappverdiske escudos", Symbol: "CVE"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "haitiske gourde", Symbol: "HTG"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "srilankiske rupier", Symbol: "LKR"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "latviske rubler", Symbol: "LVR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "sankthelenske pund", Symbol: "SHP"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "aserbajdsjanske manat", Symbol: "AZN"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "jordanske dinarer", Symbol: "JOD"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "komoriske franc", Symbol: "KMF"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "brasilianske real", Symbol: "BRL"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "dominikanske pesos", Symbol: "DOP"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "guineanske franc", Symbol: "GNF"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "venezuelanske bolivar", Symbol: "VEF"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franske UIC-franc", Symbol: "XFU"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "jugoslaviske dinarer (hard)", Symbol: "YUD"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "belgiske franc (finansielle)", Symbol: "BEL"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "brasilianske cruzados (1986–1989)", Symbol: "BRC"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "etiopiske birr", Symbol: "ETB"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "iranske rialer", Symbol: "IRR"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litauiske litas", Symbol: "LTL"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "libyske dinarer", Symbol: "LYD"}, "PES": ut.Currency{Currency: "PES", DisplayName: "peruvianske sol (1863–1965)", Symbol: "PES"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "portugisiske escudo", Symbol: "PTE"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "argentinske australer", Symbol: "ARA"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "surinamske gylden", Symbol: "SRG"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "sovjetiske rubler", Symbol: "SUR"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "gull", Symbol: "XAU"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "østkaribiske dollar", Symbol: "XCD"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "sudanske pund", Symbol: "SDG"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "caymanske dollar", Symbol: "KYD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "makedonske denarer", Symbol: "MKD"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "malaysiske ringgit", Symbol: "MYR"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "jemenittiske dinarer", Symbol: "YDD"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "albanske lek (1946–1965)", Symbol: "ALK"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "russiske rubler (1991–1998)", Symbol: "RUR"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "somaliske shilling", Symbol: "SOS"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "europeisk sammensatt enhet", Symbol: "XBA"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "zimbabwisk dollar (2009)", Symbol: "ZWL"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "albanske lek", Symbol: "ALL"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kroatiske kuna", Symbol: "HRK"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "MCF", Symbol: "MCF"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "syriske pund", Symbol: "SYP"}, "WST": ut.Currency{Currency: "WST", DisplayName: "samoanske tala", Symbol: "WST"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "ecuadorianske sucre", Symbol: "ECS"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "chilenske escudo", Symbol: "CLE"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinea-Bissau-pesos", Symbol: "GWP"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "omanske rialer", Symbol: "OMR"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "singaporske dollar", Symbol: "SGD"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "turkmenske manat (1993–2009)", Symbol: "TMM"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "sucre", Symbol: "XSU"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "nederlandske antillegylden", Symbol: "ANG"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "brasilianske cruzeiro (1993–1994)", Symbol: "BRR"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "irske pund", Symbol: "IEP"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "sørkoreanske won", Symbol: "KRW"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "kuwaitiske dinarer", Symbol: "KWD"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "makedonske denarer (1992–1993)", Symbol: "MKN"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET-fond", Symbol: "XRE"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "argentinsk pesos (1881–1970)", Symbol: "ARM"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "guyanske dollar", Symbol: "GYD"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "islandske kroner", Symbol: "ISK"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "seychelliske rupier", Symbol: "SCR"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "sierraleonske leone", Symbol: "SLL"}, "THB": ut.Currency{Currency: "THB", DisplayName: "thailandske baht", Symbol: "THB"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "ukrainske hryvnia", Symbol: "UAH"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franske gullfranc", Symbol: "XFO"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "sveitsiske franc", Symbol: "CHF"}, "YER": ut.Currency{Currency: "YER", DisplayName: "jemenittiske rialer", Symbol: "YER"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "burmesiske kyat", Symbol: "BUK"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "bolivianske pesos", Symbol: "BOP"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "maltesiske lira", Symbol: "MTL"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "mosambikiske escudo", Symbol: "MZE"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladium", Symbol: "XPD"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "belgiske franc", Symbol: "BEF"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "bulgarske lev (hard)", Symbol: "BGL"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "guineanske syli", Symbol: "GNS"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "sørafrikanske rand (finansielle)", Symbol: "ZAL"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "zairiske nye zaire", Symbol: "ZRN"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "bangladeshiske taka", Symbol: "BDT"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "bulgarske lev (1879–1952)", Symbol: "BGO"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "kroatiske dinarer", Symbol: "HRD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "mauritiske rupier", Symbol: "MUR"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "angolanske omjusterte kwanza (1995–1999)", Symbol: "AOR"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "ghanesisk cedi (1979–2007)", Symbol: "GHC"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "gambiske dalasi", Symbol: "GMD"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "maltesiske pund", Symbol: "MTP"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "qatarske rialer", Symbol: "QAR"}, "RON": ut.Currency{Currency: "RON", DisplayName: "rumenske leu", Symbol: "RON"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "tsjekkoslovakiske koruna (hard)", Symbol: "CSK"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "bahamanske dollar", Symbol: "BSD"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "østtyske mark", Symbol: "DDM"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franske franc", Symbol: "FRF"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "norske kroner", Symbol: "kr"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ugandiske shilling", Symbol: "UGX"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "usbekiske som", Symbol: "UZS"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "bahrainske dinarer", Symbol: "BHD"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "kubanske pesos", Symbol: "CUP"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "islandske kroner (1918–1981)", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "mauritanske ouguiya", Symbol: "MRO"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "uruguayanske pesos (1975–1993)", Symbol: "UYP"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "belgiske franc (konvertible)", Symbol: "BEC"}, "KES": ut.Currency{Currency: "KES", DisplayName: "kenyanske shilling", Symbol: "KES"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "tadsjikiske rubler", Symbol: "TJR"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "europeisk kontoenhet (XBC)", Symbol: "XBC"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "østerrikske shilling", Symbol: "ATS"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "serbiske dinarer (2002–2006)", Symbol: "CSD"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "marokkanske dirham", Symbol: "MAD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "mongolske tugrik", Symbol: "MNT"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "malawiske kwacha", Symbol: "MWK"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "sørsudanske pund", Symbol: "SSP"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "tonganske paʻanga", Symbol: "TOP"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "vestafrikanske CFA-franc", Symbol: "CFA"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "aserbajdsjanske manat (1993–2006)", Symbol: "AZM"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "japanske yen", Symbol: "JPY"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "argentinske peso ley", Symbol: "ARL"}, "BND": ut.Currency{Currency: "BND", DisplayName: "bruneiske dollar", Symbol: "BND"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "bolivianske boliviano", Symbol: "BOB"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "hviterussiske rubler", Symbol: "BYR"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "jamaikanske dollar", Symbol: "JMD"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "luxemburgske franc", Symbol: "LUF"}, "TND": ut.Currency{Currency: "TND", DisplayName: "tunisiske dinarer", Symbol: "TND"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "armenske dram", Symbol: "AMD"}, "USS": ut.Currency{Currency: "USS", DisplayName: "amerikanske dollar (samme dag)", Symbol: "USS"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "lesothiske loti", Symbol: "LSL"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "newzealandske dollar", Symbol: "NZD"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "italienske lire", Symbol: "ITL"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "gibraltarske pund", Symbol: "GIP"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekvatorialguineanske ekwele guineana", Symbol: "GQE"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "marokkanske franc", Symbol: "MAF"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "rhodesiske dollar", Symbol: "RHD"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "russiske rubler", Symbol: "RUB"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "kinesiske yuan", Symbol: "CNY"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "mosambikiske metical", Symbol: "MZN"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "nigerianske naira", Symbol: "NGN"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "papuanske kina", Symbol: "PGK"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "rwandiske franc", Symbol: "RWF"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "timoresiske escudo", Symbol: "TPE"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "meksikanske sølvpesos (1861–1992)", Symbol: "MXP"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "sudanesiske pund", Symbol: "SDP"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "svenske kroner", Symbol: "SEK"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "namibiske dollar", Symbol: "NAD"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "litauiske talonas", Symbol: "LTT"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "slovakiske koruna", Symbol: "SKK"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "zimbabwisk dollar (2008)", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "eritreiske nakfa", Symbol: "ERN"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "myanmarske kyat", Symbol: "MMK"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "angolanske kwanza", Symbol: "AOA"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "luxemburgske konvertible franc", Symbol: "LUC"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "maliske franc", Symbol: "MLF"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "maldiviske rupier", Symbol: "MVP"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "rumenske leu (1952–2006)", Symbol: "ROL"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "uruguyanske pesos (indekserte enheter)", Symbol: "UYI"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "venezuelanske bolivar (1871–2008)", Symbol: "VEB"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "sentralafrikanske CFA-franc", Symbol: "XAF"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "tsjekkiske koruna", Symbol: "CZK"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "sørafrikanske rand", Symbol: "ZAR"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "portugisiske guinea escudo", Symbol: "GWE"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "luxemburgske finansielle franc", Symbol: "LUL"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "panamanske balboa", Symbol: "PAB"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "serbiske dinarer", Symbol: "RSD"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "botswanske pula", Symbol: "BWP"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "madagassiske franc", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "europeisk monetær enhet", Symbol: "XBB"}, "COP": ut.Currency{Currency: "COP", DisplayName: "colombianske pesos", Symbol: "COP"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zairiske zaire", Symbol: "ZRZ"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "bosnisk-hercegovinske konvertible mark", Symbol: "BAM"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "brasilianske cruzeiro (1990–1993)", Symbol: "BRE"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "kongolesiske franc", Symbol: "CDF"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "gamle mosambikiske metical", Symbol: "MZM"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "ugandiske shilling (1966–1987)", Symbol: "UGS"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "jugoslaviske konvertible dinarer", Symbol: "YUN"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "angolanske kwanza (1977–1990)", Symbol: "AOK"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "britiske pund", Symbol: "£"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "israelske shekler (1980–1985)", Symbol: "ILR"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "nicaraguanske cordoba (1988–1991)", Symbol: "NIC"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "danske kroner", Symbol: "DKK"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "swazilandske lilangeni", Symbol: "SZL"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "paraguayanske guarani", Symbol: "PYG"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "nepalske rupier", Symbol: "NPR"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "polske zloty (1950–1995)", Symbol: "PLZ"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "sølv", Symbol: "XAG"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "algeriske dinarer", Symbol: "DZD"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "kypriotiske pund", Symbol: "CYP"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "tyske mark", Symbol: "DEM"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ghanesiske cedi", Symbol: "GHS"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "nordkoreanske won", Symbol: "KPW"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "pakistanske rupier", Symbol: "PKR"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "tadsjikiske somoni", Symbol: "TJS"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "tyrkiske lire", Symbol: "TRY"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "brasilianske cruzeiro (1942–1967)", Symbol: "BRZ"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "europeisk kontoenhet (XBD)", Symbol: "XBD"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "ADB-kontoenhet", Symbol: "XUA"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "tanzanianske shilling", Symbol: "TZS"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "barbadiske dollar", Symbol: "BBD"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nye israelske shekler", Symbol: "ILS"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "irakske dinarer", Symbol: "IQD"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "peruvianske inti", Symbol: "PEI"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "jugoslaviske reformerte dinarer (1992–1993)", Symbol: "YUR"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "bosnisk-hercegovinske dinarer (1992–1994)", Symbol: "BAD"}, "INR": ut.Currency{Currency: "INR", DisplayName: "indiske rupier", Symbol: "INR"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "kasakhstanske tenge", Symbol: "KZT"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "spesielle trekkrettigheter", Symbol: "XDR"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platina", Symbol: "XPT"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "chilenske unidades de fomento", Symbol: "CLF"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "burundiske franc", Symbol: "BIF"}, "USD": ut.Currency{Currency: "USD", DisplayName: "amerikanske dollar", Symbol: "USD"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "bulgarske lev (sosialist)", Symbol: "BGM"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "ukjent valuta", Symbol: "XXX"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afgansk afghani (1927–2002)", Symbol: "AFA"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "spanske peseta", Symbol: "ESP"}, "STD": ut.Currency{Currency: "STD", DisplayName: "São Tomé og Príncipe-dobra", Symbol: "STD"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "brasilianske cruzado novo (1989–1990)", Symbol: "BRN"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "kanadiske dollar", Symbol: "CAD"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "falklandspund", Symbol: "FKP"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "beliziske dollar", Symbol: "BZD"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "moldovske cupon", Symbol: "MDC"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "makaoiske pataca", Symbol: "MOP"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "indonesiske rupier", Symbol: "IDR"}}