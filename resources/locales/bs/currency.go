package bs

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"ESA": ut.Currency{Currency: "ESA", DisplayName: "Španska pezeta (račun) ESA", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Malagasijski franak", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timorški eskudo", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "Brunejski dolar", Symbol: "BND"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Turska lira (1922–2005)", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Urugvajski pezo en unidades indeksades", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambijska kvača", Symbol: "ZMW"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "stara islandska kruna", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Latvijska rublja", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuvajtski dinar", Symbol: "KWD"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Makedonski denar (1992–1993)", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paragvajski gvarani", Symbol: "PYG"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoanska tala", Symbol: "WST"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Posebna prava", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kirgistanski som", Symbol: "KGS"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidadtobaški dolar", Symbol: "TTD"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Jugoslovenski konvertibilni dinar", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Južnoafrički rand (finansijski)", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesotski loti", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "SAD dolar (isti dan)", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Urugvajski pezo (1975–1993)", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Albanski lek (1946–1965)", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "argentinski pezo monedo nacional", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamajski dolar", Symbol: "JMD"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepalska rupija", Symbol: "NPR"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Nepoznata valuta", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Butanski ngultrum", Symbol: "BTN"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Alžirski dinar", Symbol: "DZD"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malavska kvača", Symbol: "MWK"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Srebro", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Čehoslovačka tvrda koruna", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Danska kruna", Symbol: "DKK"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Šrilankanska rupija", Symbol: "LKR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Srpski dinar", Symbol: "din."}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belgijski frank (konvertibilni)", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahreinski dinar", Symbol: "BHD"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ekvadorijski sukr", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Kenijski šiling", Symbol: "KES"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venecuelanski bolivar", Symbol: "VEF"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Australijski dolar", Symbol: "AUD"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Kineski juan", Symbol: "CNY"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ganijski cedi (1979–2007)", Symbol: ""}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Monegaskaški franak", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Sudanska funta (1957–1998)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brazilski real", Symbol: "BRL"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Čileanski pezos", Symbol: "CLP"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kazahstanski tenge", Symbol: "KZT"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Surinamski gilder", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Ugandijski šiling (1966–1987)", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Avganistanski avgani (1927–2002)", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Gulden Holandskih Antila", Symbol: "ANG"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Španska pezeta", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Mađarska forinta", Symbol: "HUF"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzanski šiling", Symbol: "TZS"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Bolivijski boliviano", Symbol: "BOB"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Kostarikanski kolon", Symbol: "CRC"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Irački dinar", Symbol: "IQD"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Marokanski franak", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portugalski eskudo", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra Sao Toma i Principa", Symbol: "STD"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Evropska novčana jedinica", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belizeanski dolar", Symbol: "BZD"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Britanska funta", Symbol: "GBP"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Meksijski unidad de inverzion", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistanska rupija", Symbol: "PKR"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zairski novi zair (1993–1998)", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabvejski dolar (1980–2008)", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afganistanski afgan", Symbol: "AFN"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fidži dolar", Symbol: "FJD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Letonski lats", Symbol: "LVL"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldavski lev", Symbol: "MDL"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turska lira", Symbol: "TRY"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Evropska jedinica računa (XBC)", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Istočnokaripski dolar", Symbol: "XCD"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbadoski dolar", Symbol: "BBD"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Brazilski kruzado (1986–1989)", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Hrvatska kuna", Symbol: "kn"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladeška taka", Symbol: "BDT"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Bugarski tvrdi lev", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Brazilski kruzado novo (1989–1990)", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Gruzijski lari", Symbol: "GEL"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somalski šiling", Symbol: "SOS"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Francuski UIC-frank", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirham Ujedinjenih Arapskih Emirata", Symbol: "AED"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Litvanski talonas", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Malteška lira", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jemenski rijal", Symbol: "YER"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angolijska kvanza (1977–1991)", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Azerbejdžanski manat", Symbol: "AZN"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Brazilski kruzeiro (1993–1994)", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Španska pezeta (konvertibilni račun)", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Portugalska Gvineja eskudo", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibijski dolar", Symbol: "NAD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Sejšelska rupija", Symbol: "SCR"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Jugoslovenski novi dinar", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libanska funta", Symbol: "LBP"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filipinski pezos", Symbol: "PHP"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Gvinejski sili", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Mijanmarski kjat", Symbol: "MMK"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudijski rijal", Symbol: "SAR"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Ukrajinski karbovaneti", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Evropska kompozitna jedinica", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Honduraska lempira", Symbol: "HNL"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Italijanska lira", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Luksemburški finansijski franak", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Evropska jedinica računa (XBD)", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angolska kvanza", Symbol: "AOA"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Beloruska nova rublja (1994–1999)", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libijski dinar", Symbol: "LYD"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigerijska naira", Symbol: "NGN"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongoanski franak", Symbol: "CDF"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Istočnoevropska marka", Symbol: ""}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Južno-korejski Von (1945–1953)", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Brazilski kruzeiro (1990–1993)", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Bocvanska pula", Symbol: "BWP"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR franak", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Izraelski novi šekel", Symbol: "ILS"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Meksički pezos", Symbol: "MXN"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Novozelandski dolar", Symbol: "NZD"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Tajlandski baht", Symbol: "฿"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Turkmenistanski manat (1993–2009)", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentinski pezos", Symbol: "ARS"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Bugarski socijalistički lev", Symbol: ""}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Dolar kineske narodne banke", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Gruzijski kupon larit", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltarska funta", Symbol: "GIP"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Meksijski srebrno pezo (1861–1992)", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "SAD dolar (sledeći dan)", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Austrijski šiling", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Bolivijski mvdol", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Srpski dinar (2002–2006)", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Estonska kruna", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Malijanski franak", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Slovačka kruna", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Svazilendski lilangeni", Symbol: "SZL"}, "COU": ut.Currency{Currency: "COU", DisplayName: "unidad de valor real", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Ekvadorski unidad de valor konstantin (UVC)", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Drahma", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordanski dinar", Symbol: "JOD"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Južno-korejski hvan (1953–1962)", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mozambijski metikal (1980–2006)", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Ukrajinska grivna", Symbol: "UAH"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosansko-Hercegovački dinar", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Gvajanski dolar", Symbol: "GYD"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tonganska panga", Symbol: "TOP"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Angolijska nova kvanza (1990–2000)", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Hrvatski dinar", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mozambijski metikal", Symbol: "MZN"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nikaragvanska kordoba", Symbol: "NIO"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Sudanska funta", Symbol: "SDG"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermudski dolar", Symbol: "BMD"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Luksemburški konvertibilni franak", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Džibutski franak", Symbol: "DJF"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indijska rupija", Symbol: "₹"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Luksemburški franak", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokanski dirham", Symbol: "MAD"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP franak", Symbol: "XPF"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Folklandska funta", Symbol: "FKP"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Francuski franak", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Iranski rijal", Symbol: "IRR"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panamska balboa", Symbol: "PAB"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Svetohelenska funta", Symbol: "SHP"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Salvadorski kolon", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina Papue Nove Gvineje", Symbol: "PGK"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Rumunski lev", Symbol: "RON"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Evropska valutna jedinica", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bugarski lev", Symbol: "BGN"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Gvineja bisao pezo", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Makaonska pataka", Symbol: "MOP"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Ruska rublja (1991–1998)", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Francuski zlatni frank", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Zimbabvejski dolar (2009)", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonežanska rupija", Symbol: "IDR"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Švedska kruna", Symbol: "SEK"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belgijski franak", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Nemačka marka", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrejska nakfa", Symbol: "ERN"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Sjevernokorejski von", Symbol: "KPW"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Poljski zloti (1950–1995)", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Rumunski leu (1952–2006)", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Solomonski dolar", Symbol: "SBD"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Jugoslovenski reformirani dinar", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambijska kvača (1968–2012)", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zairski zair (1971–1993)", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Irska funta", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tuniski dinar", Symbol: "TND"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongolski tugrik", Symbol: "MNT"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Peruvijski inti", Symbol: ""}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Vijetnamski dong (1978–1985)", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Jugoslovenski tvrdi dinar", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorska pezeta", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Češka kruna", Symbol: "CZK"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japanski jen", Symbol: "¥"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Južnokorejski von", Symbol: "₩"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laoski kip", Symbol: "LAK"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruanski novi sol", Symbol: "PEN"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapurski dolar", Symbol: "SGD"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Čileanski unidades de fomento", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Slovenski tolar", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tadžakistanska rublja", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Ugandski šiling", Symbol: "UGX"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "uzbekistanski som", Symbol: "UZS"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Bugarski lev (1879–1952)", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Kolumbijski pezos", Symbol: "COP"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Gvinejski franak", Symbol: "GNF"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Arubanski florin", Symbol: "AWG"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Burmanski kjat", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Kipratska funta", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Finska marka", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauritanijska ugvija", Symbol: "MRO"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Rodizijski dolar", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadžikistanski somoni", Symbol: "TJS"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Bolivijski boliviano (1863–1963)", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Brazilski kruzeiro (1942–1967)", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Kajmanski dolar", Symbol: "KYD"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omanski rijal", Symbol: "OMR"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Jemenski dinar", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahamski dolar", Symbol: "BSD"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Čileanski eskudo", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Islandska kruna", Symbol: "ISK"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "moldavski kupon", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Urugvajski pezos", Symbol: "UYU"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosansko-Hercegovačka konvertibilna marka", Symbol: "KM"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Evatorijalna gvineja ekvele", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldivska rufija", Symbol: "MVR"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mozambijski eskudo", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanadski dolar", Symbol: "CAD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kambodžanski rijel", Symbol: "KHR"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Južnosudanska funta", Symbol: "SSP"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vijetnamski dong", Symbol: "₫"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET fondovi", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Angolijska kvanza reajustado (1995–1999)", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Azerbejdžanski manat (1993–2006)", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Zelenortski eskudo", Symbol: "CVE"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komorski franak", Symbol: "KMF"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Norveška kruna", Symbol: "NOK"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Novi tajvanski dolar", Symbol: "NT$"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Bolivijski pezo", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominikanski pezos", Symbol: "DOP"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haićanski gurd", Symbol: "HTG"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Holandski gulden", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinamski dolar", Symbol: "SRD"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatski vatu", Symbol: "VUV"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Bosansko-hercegovački novi dinar", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egipatska funta", Symbol: "EGP"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Peruvijski sol (1863–1965)", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Paladijum", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Južnoafrički rand", Symbol: "ZAR"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belgijski frank (finansijski)", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberijski dolar", Symbol: "LRD"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Nikaragvanška kordoba (1988–1991)", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Američki dolar", Symbol: "USD"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Jermenski dram", Symbol: "AMD"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundski franak", Symbol: "BIF"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Zimbabvejski dolar (2008)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Bjeloruska rublja", Symbol: "BYR"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Etiopski bir", Symbol: "ETB"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litvanski litas", Symbol: "LTL"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Makedonski denar", Symbol: "MKD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauricijska rupija", Symbol: "MUR"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Argentinski pezo (1983–1985)", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Gvatemalski kecal", Symbol: "GTQ"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Katarski rijal", Symbol: "QAR"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Ruska rublja", Symbol: "RUB"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "zapadnoafrički franak CFA", Symbol: "CFA"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platina", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ganski cedi", Symbol: "GHS"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Poljski zlot", Symbol: "PLN"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Ruandski franak", Symbol: "RWF"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Sudanski dinar (1992–2007)", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sijeraleonski leone", Symbol: "SLL"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Venecuelanski bolivar (1871–2008)", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albanski lek", Symbol: "ALL"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Argentinski austral", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Kubanski konvertibilni pezos", Symbol: "CUC"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Kod testirane valute", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Honkonški dolar", Symbol: "HKD"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "stari izraelski šekeli", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmenistanski manat", Symbol: "TMT"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Brazilski kruzeiro novo (1967–1986)", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR Evro", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Malteška funta", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Sovjetska rublja", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Švicarski franak", Symbol: "CHF"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "kubanski pezos", Symbol: "CUP"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambijski dalasi", Symbol: "GMD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Malagaski ariari", Symbol: "MGA"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Zlato", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Izraelska funta", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "centralnoafrički franak CFA", Symbol: "FCFA"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "argentinski pezos lej", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malezijski ringit", Symbol: "MYR"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Sirijska funta", Symbol: "SYP"}}