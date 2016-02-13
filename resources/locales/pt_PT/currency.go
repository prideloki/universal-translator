package pt_PT

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"LBP": ut.Currency{Currency: "LBP", DisplayName: "Libra libanesa", Symbol: "LBP"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Franco francês", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Peso uruguaio en unidades indexadas", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya da Mauritânia", Symbol: "MRO"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariari de Madagáscar", Symbol: "MGA"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Franco do Mali", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Zloti polaco (1950–1995)", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Dinar jordaniano", Symbol: "JOD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupia mauriciana", Symbol: "MUR"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Inti peruano", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu moldavo", Symbol: "MDL"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Escudo português", Symbol: "\u200b"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vatu de Vanuatu", Symbol: "VUV"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Franco CFA (BEAC)", Symbol: "FCFA"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Dong vietnamita", Symbol: "₫"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Libra maltesa", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Ostmark da Alemanha Oriental", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Cruzado brasileiro (1986–1989)", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirham marroquino", Symbol: "MAD"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Peso argentino (1983–1985)", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Bolívar venezuelano", Symbol: "VEF"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Rublo soviético", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dólar australiano", Symbol: "AU$"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Florim das Antilhas Holandesas", Symbol: "ANG"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni da Suazilândia", Symbol: "SZL"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Peseta de Andorra", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Peso cubano", Symbol: "CUP"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Rial saudita", Symbol: "SAR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Libra de Santa Helena", Symbol: "SHP"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Dólar de Singapura", Symbol: "SGD"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Unidad de Valor Constante (UVC) do Equador", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Franco comoriano", Symbol: "KMF"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Novo dinar da Bósnia-Herzegovina (1994–1997)", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Lev forte búlgaro", Symbol: ""}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Escudo chileno", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colon costa-riquenho", Symbol: "CRC"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Moeda desconhecida", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial de Omã", Symbol: "OMR"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Unidad de Inversion (UDI) mexicana", Symbol: ""}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Cupon moldávio", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Rublo bielorrusso", Symbol: "BYR"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Dólar rodesiano", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dinar líbio", Symbol: "LYD"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Cruzeiro brasileiro (1990–1993)", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Peso da Guiné-Bissau", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Dólar das Ilhas Caimão", Symbol: "KYD"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Libra egípcia", Symbol: "EGP"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinar tunisino", Symbol: "TND"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Rand sul-africano (financeiro)", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Coroa estoniana", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Libra de Chipre", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Novo sol peruano", Symbol: "PEN"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Forint húngaro", Symbol: "HUF"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Lev socialista búlgaro", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira das Honduras", Symbol: "HNL"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha zambiano (1968–2012)", Symbol: "ZMK"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Dólar do Zimbábue (2008)", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Talonas lituano", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Libra de Gibraltar", Symbol: "GIP"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Dólar do Zimbábue (2009)", Symbol: ""}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Franco monegasco", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum do Butão", Symbol: "BTN"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Prata", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Dinar croata", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Xelim ugandense (1966–1987)", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kuna croata", Symbol: "HRK"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Dólar neozelandês", Symbol: "NZ$"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Super Dinar jugoslavo", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Franco UIC francês", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florim de Aruba", Symbol: "AWG"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Sheqel antigo israelita", Symbol: ""}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Lev búlgaro (1879–1952)", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Ouro", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Peso Plata mexicano (1861–1992)", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Sheqel novo israelita", Symbol: "₪"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Cuanza angolano reajustado (1995–1999)", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Sucre equatoriano", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Rupia nepalesa", Symbol: "NPR"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "Franco WIR", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi da Gâmbia", Symbol: "GMD"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Leu romeno (1952–2006)", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Dinar kuwaitiano", Symbol: "KWD"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira nigeriana", Symbol: "NGN"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dólar liberiano", Symbol: "LRD"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Coroa sueca", Symbol: "SEK"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde haitiano", Symbol: "HTG"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Franco ruandês", Symbol: "RWF"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupia seichelense", Symbol: "SCR"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupia do Sri Lanka", Symbol: "LKR"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Marco alemão", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial iemenita", Symbol: "YER"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Dólar da Guiana", Symbol: "GYD"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Franco-ouro francês", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Franco marroquino", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Cruzeiro novo brasileiro (1967–1986)", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Xelim tanzaniano", Symbol: "TZS"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat do Azerbaijão", Symbol: "AZN"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Kyat birmanês", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Florim holandês", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Franco conversível de Luxemburgo", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Rand sul-africano", Symbol: "ZAR"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Cedi do Gana", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Unidade de Conta Europeia (XBC)", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "Dólar norte-americano (Mesmo dia)", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Franco de Madagascar", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha zambiano", Symbol: "ZMW"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "Euro WIR", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Franco belga", Symbol: ""}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Dólar do Banco Popular da China", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Dinar macedónio", Symbol: "MKD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Peso cubano conversível", Symbol: "CUC"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Bolívar venezuelano (1871–2008)", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza angolano", Symbol: "AOA"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidade de Valor Real", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Won norte-coreano", Symbol: "KPW"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Lari georgiano", Symbol: "GEL"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Marco bósnio-herzegóvino conversível", Symbol: "BAM"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Sol peruano (1863–1965)", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Córdoba nicaraguano (1988–1991)", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Coroa checa", Symbol: "CZK"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Franco congolês", Symbol: "CDF"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra de São Tomé e Príncipe", Symbol: "STD"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Guarani paraguaio", Symbol: "PYG"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Dinar da Bósnia-Herzegóvina", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metical de Moçambique", Symbol: "MZN"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Dinar forte jugoslavo", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Libra sudanesa (1957–1998)", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Coroa islandesa", Symbol: "ISK"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Kip de Laos", Symbol: "LAK"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Lira turca", Symbol: "TRY"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Lek albanês", Symbol: "ALL"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloti polaco", Symbol: "PLN"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Unidade Monetária Europeia", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "Fundos RINET", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Dólar das Caraíbas Orientais", Symbol: "EC$"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Córdoba nicaraguano", Symbol: "NIO"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Libra israelita", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Peso colombiano", Symbol: "COP"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Mvdol boliviano", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zaire Novo zairense (1993–1998)", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirham dos Emirados Árabes Unidos", Symbol: "AED"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats da Letónia", Symbol: "LVL"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha do Malawi", Symbol: "MWK"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Dólar do Suriname", Symbol: "SRD"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Baht da Tailândia", Symbol: "฿"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Coroa antiga islandesa", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Franco belga (financeiro)", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dólar das Bahamas", Symbol: "BSD"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rupia das Ilhas Maldivas", Symbol: "MVR"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dólar dos Estados Unidos", Symbol: "US$"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Cruzeiro brasileiro (1942–1967)", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dinar argelino", Symbol: "DZD"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula de Botswana", Symbol: "BWP"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal da Guatemala", Symbol: "GTQ"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Peseta espanhola (conta conversível)", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Unidade de Conta Europeia (XBD)", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Cedi de Gana", Symbol: "GHS"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Hryvnia da Ucrânia", Symbol: "UAH"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Novo cuanza angolano (1990–2000)", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Escudo da Guiné Portuguesa", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Manat azerbaijano (1993–2006)", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Franco financeiro de Luxemburgo", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Franco jibutiano", Symbol: "DJF"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas da Lituânia", Symbol: "LTL"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Iene japonês", Symbol: "JP¥"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Franco suíço", Symbol: "CHF"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Rublo russo", Symbol: "RUB"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Cruzeiro brasileiro (1993–1994)", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Franco belga (convertível)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Real brasileiro", Symbol: "R$"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Peso mexicano", Symbol: "MX$"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Franco burundiano", Symbol: "BIF"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Escudo cabo-verdiano", Symbol: "CVE"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dólar de Hong Kong", Symbol: "HK$"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Libra sul-sudanesa", Symbol: "SSP"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Paʻanga de Tonga", Symbol: "TOP"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Dinar conversível jugoslavo", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Dinar sérvio (2002–2006)", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa do Panamá", Symbol: "PAB"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dólar do Zimbabwe", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Xelim queniano", Symbol: "KES"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Rupia indonésia", Symbol: "IDR"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Franco CFA (BCEAO)", Symbol: "CFA"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviano", Symbol: "BOB"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat do Turquemenistão", Symbol: "TMT"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Peseta espanhola", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Dinar iemenita", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Rublo do Tadjiquistão", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Rublo russo (1991–1998)", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metical de Moçambique (1980–2006)", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Libra esterlina britânica", Symbol: "£"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni do Tajaquistão", Symbol: "TJS"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Tala samoano", Symbol: "WST"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Libra sudanesa", Symbol: "SDG"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Xelim austríaco", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leone de Serra Leoa", Symbol: "SLL"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik da Mongólia", Symbol: "MNT"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Código de Moeda de Teste", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Coroa Forte checoslovaca", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "direito especial de saque", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Novo dólar taiwanês", Symbol: "NT$"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Franco luxemburguês", Symbol: ""}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Dong vietnamita (1978–1985)", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Xelim somali", Symbol: "SOS"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Tolar Bons esloveno", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Unidade Composta Europeia", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Franco guineense", Symbol: "GNF"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu romeno", Symbol: "RON"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Peso dominicano", Symbol: "DOP"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dólar canadiano", Symbol: "CA$"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won sul-coreano", Symbol: "₩"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaire zairense (1971–1993)", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Peso uruguaio", Symbol: "UYU"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Peso boliviano", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Escudo timorense", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuan chinês", Symbol: "CN¥"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Peso argentino (1881–1970)", Symbol: ""}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Dinar macedônio (1992–1993)", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "Dólar bruneíno", Symbol: "BND"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti do Lesoto", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa da Eritreia", Symbol: "ERN"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Lira turca (1922–2005)", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Austral argentino", Symbol: ""}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Won da Coreia do Sul (1945–1953)", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Unidade da Moeda Europeia", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Karbovanetz ucraniano", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Peso chileno", Symbol: "CLP"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Manat do Turcomenistão (1993–2009)", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca de Macau", Symbol: "MOP"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Marca finlandesa", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Peso filipino", Symbol: "PHP"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Dinar sudanês (1992–2007)", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afeghani (1927–2002)", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Rublo novo bielorusso (1994–1999)", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dólar da Namíbia", Symbol: "NAD"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupia indiana", Symbol: "₹"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Syli da Guiné", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Libra irlandesa", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka de Bangladesh", Symbol: "BDT"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Dólar de Fiji", Symbol: "FJD"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial iraniano", Symbol: "IRR"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platina", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Ringgit malaio", Symbol: "MYR"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Rupia paquistanesa", Symbol: "PKR"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Ekwele da Guiné Equatorial", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Dracma grego", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram arménio", Symbol: "AMD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Riel cambojano", Symbol: "KHR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Dinar sérvio", Symbol: "RSD"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Florim do Suriname", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Escudo de Moçambique", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Lira maltesa", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Peso argentino", Symbol: "ARS"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som do Uzbequistão", Symbol: "UZS"}, "USN": ut.Currency{Currency: "USN", DisplayName: "Dólar norte-americano (Dia seguinte)", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Xelim ugandense", Symbol: "UGX"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial do Catar", Symbol: "QAR"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Libra síria", Symbol: "SYP"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Lek Albanês (1946–1965)", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Coroa eslovaca", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Colom salvadorenho", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Unidades de Fomento chilenas", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Rublo letão", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Libra das Ilhas Falkland", Symbol: "FKP"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Dólar jamaicano", Symbol: "JMD"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Dólar das Ilhas Salomão", Symbol: "SBD"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Peso uruguaio (1975–1993)", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Peseta espanhola (conta A)", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Dólar barbadense", Symbol: "BBD"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Dólar bermudense", Symbol: "BMD"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat de Mianmar", Symbol: "MMK"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge do Cazaquistão", Symbol: "KZT"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Dólar de Trindade e Tobago", Symbol: "TTD"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Cuanza angolano (1977–1990)", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Coroa norueguesa", Symbol: "NOK"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Lira italiana", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Peso lei argentino (1970–1983)", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Birr etíope", Symbol: "ETB"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar baremita", Symbol: "BHD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina da Papua-Nova Guiné", Symbol: "PGK"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Coroa dinamarquesa", Symbol: "DKK"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Cruzado novo brasileiro (1989–1990)", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Dinar iraquiano", Symbol: "IQD"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afegani do Afeganistão", Symbol: "AFN"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som do Quirguistão", Symbol: "KGS"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Paládio", Symbol: ""}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Hwan da Coreia do Sul (1953–1962)", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "Franco CFP", Symbol: "CFPF"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Cupom Lari georgiano", Symbol: ""}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Dinar reformado iugoslavo (1992–1993)", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dólar belizense", Symbol: "BZD"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Boliviano (1863–1963)", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Lev búlgaro", Symbol: "BGN"}}