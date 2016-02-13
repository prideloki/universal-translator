package rm

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BEF": ut.Currency{Currency: "BEF", DisplayName: "franc beltg", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum butanais", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colon da la Costa Rica", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu rumen", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "rubel russ (vegl)", Symbol: "RUR"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso da l’Uruguay", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladi", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nov rubel bieloruss (1994–1999)", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar iracais", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iranais", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kirghis", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza angolan reajustado", Symbol: ""}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "veglia cruna islandaisa", Symbol: "ISJ"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical dal mozambican", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serb", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni dal Tadschikistan", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platin", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haitian", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "rubel letton", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht tailandais", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "aur", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brasilian", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya da la Mauretania", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "rubel russ (nov)", Symbol: "RUB"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "rubel dal Tadschikistan", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unitad monetara europeica", Symbol: "XEU"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dollar rodesian", Symbol: ""}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "vegl dong vietnamais", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franc CFP", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nov zaire dal Zaire", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "cruna islandaisa", Symbol: "ISK"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "franc monegas", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nov sol peruan", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "schilling tansanian", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "fonds RINET", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "peseta andorrana", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franc convertibel luxemburgais", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "rupia pakistana", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escudo da Timor", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso da la Guinea-Bissau", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituan", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit da la Malaisia", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu dal Vanuatu", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "code per verifitgar la valuta", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "riyal da Katar", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone da la Sierra Leone", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso da l’Uruguay (unidades indexadas)", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "riyal saudit", Symbol: "SAR"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franc d’aur franzos", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unidades de fomento chilenas", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cuban", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr etiopic", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha dal Malawi", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa dal Panama", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaire dal Zaire", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "dirham dals Emirats Arabs Unids", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marc convertibel bosniac", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "nov lev bulgar", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "glivra da Gibraltar", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "schilling ucranais", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "dollar dals Stadis Unids da l’America (proxim di)", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "dollar dal Brunei", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franc congolais", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "peseta spagnola", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekwele da la Guinea Equatoriala", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dollar da la Nova Zelanda", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "glivra siriana", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "boliviano", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brasilian", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "kupon larit georgian", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal da la Guatemala", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "dinar da la Macedonia", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escudo portugais", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat aserbaidschanic", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka bangladais", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dollar dal Belize", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lira taliana", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libic", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franc dal Mali", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angolan", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra da São Tomé e Principe", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghani", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "franc WIR", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "glivra cipriota", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupia indonaisa", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "nova lira tirca", Symbol: "TRY"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dollar dals Stadis Unids da l’America", Symbol: "$"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentin", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franc beltg (finanzial)", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "cruzeiro novo brasilian (1967–1986)", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi ghanais", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franc da la Guinea", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "glivra maltaisa", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigeriana", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentin", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franc dal Dschibuti", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dollar da Hongkong", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti dal Lesotho", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "glivra da Sontg’Elena", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "sheqel", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dollar surinam", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong vietnamais", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "schilling austriac", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat dal Myanmar", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dollar namibian", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "nov kwanza angolan", Symbol: ""}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "nov dinar da la Bosnia ed Erzegovina", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari georgian", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escudo da la Guinea Portugaisa", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afghani (1927–2002)", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dollar da Barbados", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "peseta spagnola (conto convertibel)", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambodschan", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni dal Swaziland", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "dretgs da prelevaziun spezials", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rufiyaa da las Maledivas", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso d’argient mexican (1861–1992)", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "lira tirca", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franc dal Burundi", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "rupia da la Sri Lanka", Symbol: ""}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "cupon moldav", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexican", Symbol: "MX$"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "schilling ugandais", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "cruna tschecoslovaca", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "glivra indonaisa", Symbol: "IEP"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "glivra sudanaisa", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "mvdol bolivian", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franc finanzial luxemburgais", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol peruan", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dollar dal Singapur", Symbol: ""}, "COU": ut.Currency{Currency: "COU", DisplayName: "unidad de valor real", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franc franzos", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "drachma greca", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dollar giamaican", Symbol: "JMD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "rupia da las Seychellas", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dollar da la Caribica Orientala", Symbol: "EC$"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar dal Bahrain", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "glivra sudanaisa (1956–2007)", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha da la sambia (1968–2012)", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dollar dal Simbabwe", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won da la Corea dal Sid", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar dal Kuwait", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franc marocan", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca dal Macao", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "pa’anga da Tonga", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nov dollar taiwanais", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unitad da scuntrada da l’Ecuador", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "glivra libanaisa", Symbol: ""}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "rupia da las Maledivas", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "vegl leu rumen", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "dollar dals Stadis Unids da l’America (medem di)", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escudo dal Cap Verd", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escudo dal mozambican", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar algerian", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha da la sambia", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev bulgar", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "euro WIR", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso filippin", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angolan (1977–1990)", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "peso argentin ley", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franc beltg (convertibel)", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dollar da la Guyana", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "schilling kenian", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "zloty polac", Symbol: "PLN"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar da la Bosnia ed Erzegovina", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dollar da las Bermudas", Symbol: ""}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "vegl boliviano", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "rupia dal Mauritius", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "rupia nepalaisa", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guarani paraguaian", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "flurin da las Antillas Olandaisas", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dollar australian", Symbol: "A$"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "lev bulgar socialistic", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "glivra sterlina", Symbol: "£"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldav", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti peruan", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "yuan renminbi chinais", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "forint ungarais", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "schilling somalian", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turkmen", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "peseta spagnola (conto A)", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "glivra dal Falkland", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial da l’Oman", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "cruna tscheca", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cedi ghanais (1979–2007)", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franc CFA BEAC", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unitad dal quint europeica (XBC)", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat aserbaidschanic (1993–2006)", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franc luxemburgais", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbovanetz ucranais", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franc UIC franzos", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar dal Jemen", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "peso argentin moneda nacional", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso chilen", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso columbian", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar jordanic", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franc svizzer", Symbol: "CHF"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "cordoba nicaraguan", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "argient", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unitad europeica cumponida", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "valuta nunenconuschenta", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro brasilian", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dollar da las Bahamas", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa eritreic", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won da la Corea dal Nord", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lira maltaisa", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala da la Samoa", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sidafrican (finanzial)", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "glivra egipziana", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "cruna svedaisa", Symbol: "SEK"}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial dal Jemen", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "nov dinar jugoslav", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentin (1983–1985)", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "flurin surinam", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sidafrican", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "lek albanais (1947–1961)", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "marc da la Germania da l’Ost", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "vegl cruzeiro brasilian", Symbol: ""}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "escudo chilen", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominican", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli da la Guinea", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira hondurian", Symbol: "HNL"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dirham marocan", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "flurin ollandais", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "sum usbec", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "rupia indica", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zloty polac (1950–1995)", Symbol: ""}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "vegl lev bulgar", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "marc tudestg", Symbol: ""}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "vegl sheqel israelian", Symbol: ""}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "vegl won da la Corea dal Sid", Symbol: ""}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "vegl dinar macedon", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "cruna norvegiaisa", Symbol: "NOK"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franc ruandais", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colon da l’El Salvador", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "cruzado novo brasilian", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dollar canadais", Symbol: "CA$"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "cruna danaisa", Symbol: "DKK"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "cruna estona", Symbol: "EEK"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franc madagasc", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tugrik mongolic", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar sudanais", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "dinar jugoslav (1966–1990)", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dollar dal Fidschi", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tolar sloven", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franc CFA BCEAO", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso bolivian", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambic", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franc comorian", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge casac", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dollar da las Salomonas", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brasilian (1990–1993)", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croat", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "glivra israeliana", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary madagasc", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unidad de inversion mexicana (UDI)", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "vegl metical mozambican", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "cruna slovaca", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "flurin da l’Aruba", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituan", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "hryvnia ucranais", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar jugoslav convertibel", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dollar da las Inslas Cayman", Symbol: "KYD"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "rubel sovietic", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolivar venezuelan (1871–2008)", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolivar venezuelan", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unitad dal quint europeica (XBD)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "rubel bieloruss", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre equadorian", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croata", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dollar liberian", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "cordoba oro nicaraguan", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunesian", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula da la Botswana", Symbol: ""}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "hwan da la Corea dal Sid", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letton", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina da la Papua Nova Guinea", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Kyat burmais", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "vegl dinar serb", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "marc finlandais", Symbol: "FIM"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dollar da Trinidad e Tobago", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram armen", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "yen giapunais", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip laot", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "nov peso da l’Uruguay (1975–1993)", Symbol: ""}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "dinar jugoslav refurmà", Symbol: ""}}