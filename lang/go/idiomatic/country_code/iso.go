package country_code

import (
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// https://en.wikipedia.org/wiki/List_of_country_calling_codes
// https://en.wikipedia.org/wiki/ISO_3166-1
// https://www.cia.gov/the-world-factbook/references/country-data-codes/
// https://gist.github.com/radcliff/f09c0f88344a7fcef373#file-wikipedia-iso-country-codes-csv

type IsoCountryCode struct {
	Name    string `json:"name,omitempty"`
	Alpha2  string `json:"alpha-2,omitempty"`
	Alpha3  string `json:"alpha-3,omitempty"`
	Numeric int    `json:"numeric,omitempty"`
}

func CountryCodeLookup(s string) (IsoCountryCode, bool) {
	lower := stringer.ToLower(s)

	if code, ok := countryCode_alpha22code[lower]; ok {
		return code, true
	}

	if code, ok := countryCode_alpha32code[lower]; ok {
		return code, true
	}

	if code, ok := countryCode_name2code[lower]; ok {
		return code, true
	}

	if numeric, err := strconv.Atoi(lower); err == nil {
		if code, ok := countryCode_numeric2code[numeric]; ok {
			return code, true
		}
	}

	return IsoCountryCode{}, false
}

func init() {
	for _, line := range stringer.Split(isoDataCsv, "\n") {
		line = stringer.TrimSpace(line)

		if stringer.IsEmpty(line) || stringer.ContainsAny(line, "Comment") {
			continue
		}

		fields := slicer.Map(func(s string) string {
			return stringer.Trim(s, `"`)
		}, stringer.Split(line, ";")...)

		code := IsoCountryCode{}

		code.Name = fields[0]
		isoAll := fields[2]

		isoComps := slicer.Map(func(s string) string {
			return stringer.TrimSpace(s)
		}, stringer.Split(isoAll, "|")...)

		code.Alpha2 = isoComps[0]
		code.Alpha3 = isoComps[1]

		numeric, _ := strconv.Atoi(isoComps[2])
		code.Numeric = numeric

		if code.Alpha2 == "-" && code.Alpha3 == "-" && code.Numeric == 0 {
			continue
		}

		IsoCountryCodes = append(IsoCountryCodes, code)
		countryCode_name2code[stringer.ToLower(code.Name)] = code

		if code.Alpha2 != "-" {
			countryCode_alpha22code[stringer.ToLower(code.Alpha2)] = code
		}

		if code.Alpha3 != "-" {
			countryCode_alpha32code[stringer.ToLower(code.Alpha3)] = code
		}

		if code.Numeric != 0 {
			countryCode_numeric2code[code.Numeric] = code
		}
	}
}

var (
	IsoCountryCodes          = []IsoCountryCode{}
	countryCode_alpha22code  = map[string]IsoCountryCode{}
	countryCode_alpha32code  = map[string]IsoCountryCode{}
	countryCode_numeric2code = map[int]IsoCountryCode{}
	countryCode_name2code    = map[string]IsoCountryCode{}

	isoDataCsv = `
    Name;GENC;ISO 3166;Stanag;Internet;Comment
    "Afghanistan";"AFG";"AF | AFG | 004";"AFG";".af";""
    "Akrotiri";"XQZ";"- | - | -";"-";"-";""
    "Albania";"ALB";"AL | ALB | 008";"ALB";".al";""
    "Algeria";"DZA";"DZ | DZA | 012";"DZA";".dz";""
    "American Samoa";"ASM";"AS | ASM | 016";"ASM";".as";""
    "Andorra";"AND";"AD | AND | 020";"AND";".ad";""
    "Angola";"AGO";"AO | AGO | 024";"AGO";".ao";""
    "Anguilla";"AIA";"AI | AIA | 660";"AIA";".ai";""
    "Antarctica";"ATA";"AQ | ATA | 010";"ATA";".aq";"ISO defines as the territory south of 60 degrees south latitude"
    "Antigua and Barbuda";"ATG";"AG | ATG | 028";"ATG";".ag";""
    "Argentina";"ARG";"AR | ARG | 032";"ARG";".ar";""
    "Armenia";"ARM";"AM | ARM | 051";"ARM";".am";""
    "Aruba";"ABW";"AW | ABW | 533";"ABW";".aw";""
    "Ashmore and Cartier Islands";"XAC";"- | - | -";"AUS";"-";"ISO includes with Australia"
    "Australia";"AUS";"AU | AUS | 036";"AUS";".au";"ISO includes Ashmore and Cartier Islands, Coral Sea Islands"
    "Austria";"AUT";"AT | AUT | 040";"AUT";".at";""
    "Azerbaijan";"AZE";"AZ | AZE | 031";"AZE";".az";""
    "Bahamas, The";"BHS";"BS | BHS | 044";"BHS";".bs";""
    "Bahrain";"BHR";"BH | BHR | 048";"BHR";".bh";""
    "Baker Island";"XBK";"- | - | -";"UMI";"-";"ISO includes with the US Minor Outlying Islands"
    "Bangladesh";"BGD";"BD | BGD | 050";"BGD";".bd";""
    "Barbados";"BRB";"BB | BRB | 052";"BRB";".bb";""
    "Bassas da India";"XBI";"- | - | -";"-";"-";"administered as part of French Southern and Antarctic Lands; no ISO codes assigned"
    "Belarus";"BLR";"BY | BLR | 112";"BLR";".by";""
    "Belgium";"BEL";"BE | BEL | 056";"BEL";".be";""
    "Belize";"BLZ";"BZ | BLZ | 084";"BLZ";".bz";""
    "Benin";"BEN";"BJ | BEN | 204";"BEN";".bj";""
    "Bermuda";"BMU";"BM | BMU | 060";"BMU";".bm";""
    "Bhutan";"BTN";"BT | BTN | 064";"BTN";".bt";""
    "Bolivia";"BOL";"BO | BOL | 068";"BOL";".bo";""
    "Bosnia and Herzegovina";"BIH";"BA | BIH | 070";"BIH";".ba";""
    "Botswana";"BWA";"BW | BWA | 072";"BWA";".bw";""
    "Bouvet Island";"BVT";"BV | BVT | 074";"BVT";".bv";""
    "Brazil";"BRA";"BR | BRA | 076";"BRA";".br";""
    "British Indian Ocean Territory";"IOT";"IO | IOT | 086";"IOT";".io";""
    "British Virgin Islands";"VGB";"VG | VGB | 092";"VGB";".vg";""
    "Brunei";"BRN";"BN | BRN | 096";"BRN";".bn";""
    "Bulgaria";"BGR";"BG | BGR | 100";"BGR";".bg";""
    "Burkina Faso";"BFA";"BF | BFA | 854";"BFA";".bf";""
    "Burma";"MMR";"MM | MMR | 104";"MMR";".mm";"ISO uses the name Myanmar"
    "Burundi";"BDI";"BI | BDI | 108";"BDI";".bi";""
    "Cabo Verde";"CPV";"CV | CPV | 132";"CPV";".cv";""
    "Cambodia";"KHM";"KH | KHM | 116";"KHM";".kh";""
    "Cameroon";"CMR";"CM | CMR | 120";"CMR";".cm";""
    "Canada";"CAN";"CA | CAN | 124";"CAN";".ca";""
    "Cayman Islands";"CYM";"KY | CYM | 136";"CYM";".ky";""
    "Central African Republic";"CAF";"CF | CAF | 140";"CAF";".cf";""
    "Chad";"TCD";"TD | TCD | 148";"TCD";".td";""
    "Chile";"CHL";"CL | CHL | 152";"CHL";".cl";""
    "China";"CHN";"CN | CHN | 156";"CHN";".cn";"see also Taiwan"
    "Christmas Island";"CXR";"CX | CXR | 162";"CXR";".cx";""
    "Clipperton Island";"CPT";"- | - | -";"FYP";"-";"ISO includes with France"
    "Cocos (Keeling) Islands";"CCK";"CC | CCK | 166";"AUS";".cc";""
    "Colombia";"COL";"CO | COL | 170";"COL";".co";""
    "Comoros";"COM";"KM | COM | 174";"COM";".km";""
    "Congo, Democratic Republic of the";"COD";"CD | COD | 180";"COD";".cd";"formerly Zaire"
    "Congo, Republic of the";"COG";"CG | COG | 178";"COG";".cg";""
    "Cook Islands";"COK";"CK | COK | 184";"COK";".ck";""
    "Coral Sea Islands";"XCS";"- | - | -";"AUS";"-";"ISO includes with Australia"
    "Costa Rica";"CRI";"CR | CRI | 188";"CRI";".cr";""
    "Cote d'Ivoire";"CIV";"CI | CIV | 384";"CIV";".ci";""
    "Croatia";"HRV";"HR | HRV | 191";"HRV";".hr";""
    "Cuba";"CUB";"CU | CUB | 192";"CUB";".cu";""
    "Curacao";"CUW";"CW | CUW | 531";"-";".cw";""
    "Cyprus";"CYP";"CY | CYP | 196";"CYP";".cy";""
    "Czechia";"CZE";"CZ | CZE | 203";"CZE";".cz";""
    "Denmark";"DNK";"DK | DNK | 208";"DNK";".dk";""
    "Dhekelia";"XXD";"- | - | -";"-";"-";""
    "Djibouti";"DJI";"DJ | DJI | 262";"DJI";".dj";""
    "Dominica";"DMA";"DM | DMA | 212";"DMA";".dm";""
    "Dominican Republic";"DOM";"DO | DOM | 214";"DOM";".do";""
    "Ecuador";"ECU";"EC | ECU | 218";"ECU";".ec";""
    "Egypt";"EGY";"EG | EGY | 818";"EGY";".eg";""
    "El Salvador";"SLV";"SV | SLV | 222";"SLV";".sv";""
    "Equatorial Guinea";"GNQ";"GQ | GNQ | 226";"GNQ";".gq";""
    "Eritrea";"ERI";"ER | ERI | 232";"ERI";".er";""
    "Estonia";"EST";"EE | EST | 233";"EST";".ee";""
    "Eswatini";"SWZ";"SZ | SWZ | 748";"SWZ";".sz";""
    "Ethiopia";"ETH";"ET | ETH | 231";"ETH";".et";""
    "Europa Island";"XEU";"- | - | -";"-";"-";"administered as part of French Southern and Antarctic Lands; no ISO codes assigned"
    "Falkland Islands (Islas Malvinas)";"FLK";"FK | FLK | 238";"FLK";".fk";""
    "Faroe Islands";"FRO";"FO | FRO | 234";"FRO";".fo";""
    "Fiji";"FJI";"FJ | FJI | 242";"FJI";".fj";""
    "Finland";"FIN";"FI | FIN | 246";"FIN";".fi";""
    "France";"FRA";"FR | FRA | 250";"FRA";".fr";"ISO includes metropolitan France along with the dependencies of Clipperton Island, French Guiana, French Polynesia, French Southern and Antarctic Lands, Guadeloupe, Martinique, Mayotte, New Caledonia, Reunion, Saint Pierre and Miquelon, Wallis and Futuna"
    "France, Metropolitan";"-";"FX | FXX | 249";"-";".fx";"ISO limits to the European part of France"
    "French Guiana";"GUF";"GF | GUF | 254";"GUF";".gf";""
    "French Polynesia";"PYF";"PF | PYF | 258";"PYF";".pf";""
    "French Southern and Antarctic Lands";"ATF";"TF | ATF | 260";"ATF";".tf";"GENC does not include the French-claimed portion of Antarctica (Terre Adelie)"
    "Gabon";"GAB";"GA | GAB | 266";"GAB";".ga";""
    "Gambia, The";"GMB";"GM | GMB | 270";"GMB";".gm";""
    "Gaza Strip";"XGZ";"PS | PSE | 275";"PSE";".ps";"ISO identifies as Occupied Palestinian Territory"
    "Georgia";"GEO";"GE | GEO | 268";"GEO";".ge";""
    "Germany";"DEU";"DE | DEU | 276";"DEU";".de";""
    "Ghana";"GHA";"GH | GHA | 288";"GHA";".gh";""
    "Gibraltar";"GIB";"GI | GIB | 292";"GIB";".gi";""
    "Glorioso Islands";"XGL";"- | - | -";"-";"-";"administered as part of French Southern and Antarctic Lands; no ISO codes assigned"
    "Greece";"GRC";"GR | GRC | 300";"GRC";".gr";"For its internal communications, the  European Union recommends the use of the code EL in lieu of the ISO 3166-2 code of GR"
    "Greenland";"GRL";"GL | GRL | 304";"GRL";".gl";""
    "Grenada";"GRD";"GD | GRD | 308";"GRD";".gd";""
    "Guadeloupe";"GLP";"GP | GLP | 312";"GLP";".gp";""
    "Guam";"GUM";"GU | GUM | 316";"GUM";".gu";""
    "Guatemala";"GTM";"GT | GTM | 320";"GTM";".gt";""
    "Guernsey";"GGY";"GG | GGY | 831";"UK";".gg";""
    "Guinea";"GIN";"GN | GIN | 324";"GIN";".gn";""
    "Guinea-Bissau";"GNB";"GW | GNB | 624";"GNB";".gw";""
    "Guyana";"GUY";"GY | GUY | 328";"GUY";".gy";""
    "Haiti";"HTI";"HT | HTI | 332";"HTI";".ht";""
    "Heard Island and McDonald Islands";"HMD";"HM | HMD | 334";"HMD";".hm";""
    "Holy See (Vatican City)";"VAT";"VA | VAT | 336";"VAT";".va";""
    "Honduras";"HND";"HN | HND | 340";"HND";".hn";""
    "Hong Kong";"HKG";"HK | HKG | 344";"HKG";".hk";""
    "Howland Island";"XHO";"- | - | -";"UMI";"-";"ISO includes with the US Minor Outlying Islands"
    "Hungary";"HUN";"HU | HUN | 348";"HUN";".hu";""
    "Iceland";"ISL";"IS | ISL | 352";"ISL";".is";""
    "India";"IND";"IN | IND | 356";"IND";".in";""
    "Indonesia";"IDN";"ID | IDN | 360";"IDN";".id";""
    "Iran";"IRN";"IR | IRN | 364";"IRN";".ir";""
    "Iraq";"IRQ";"IQ | IRQ | 368";"IRQ";".iq";""
    "Ireland";"IRL";"IE | IRL | 372";"IRL";".ie";""
    "Isle of Man";"IMN";"IM | IMN | 833";"UK";".im";""
    "Israel";"ISR";"IL | ISR | 376";"ISR";".il";""
    "Italy";"ITA";"IT | ITA | 380";"ITA";".it";""
    "Jamaica";"JAM";"JM | JAM | 388";"JAM";".jm";""
    "Jan Mayen";"XJM";"- | - | -";"SJM";"-";"ISO includes with Svalbard"
    "Japan";"JPN";"JP | JPN | 392";"JPN";".jp";""
    "Jarvis Island";"XJV";"- | - | -";"UMI";"-";"ISO includes with the US Minor Outlying Islands"
    "Jersey";"JEY";"JE | JEY | 832";"UK";".je";""
    "Johnston Atoll";"XJA";"- | - | -";"UMI";"-";"ISO includes with the US Minor Outlying Islands"
    "Jordan";"JOR";"JO | JOR | 400";"JOR";".jo";""
    "Juan de Nova Island";"XJN";"- | - | -";"-";"-";"administered as part of French Southern and Antarctic Lands; no ISO codes assigned"
    "Kazakhstan";"KAZ";"KZ | KAZ | 398";"KAZ";".kz";""
    "Kenya";"KEN";"KE | KEN | 404";"KEN";".ke";""
    "Kingman Reef";"XKR";"- | - | -";"UMI";"-";"ISO includes with the US Minor Outlying Islands"
    "Kiribati";"KIR";"KI | KIR | 296";"KIR";".ki";""
    "Korea, North";"PRK";"KP | PRK | 408";"PRK";".kp";""
    "Korea, South";"KOR";"KR | KOR | 410";"KOR";".kr";""
    "Kosovo";"XKS";"XK | XKS | -";"-";"-";"XK and XKS are  ISO 3166 user assigned codes; ISO 3166 Maintenace Authority has not assigned codes"
    "Kuwait";"KWT";"KW | KWT | 414";"KWT";".kw";""
    "Kyrgyzstan";"KGZ";"KG | KGZ | 417";"KGZ";".kg";""
    "Laos";"LAO";"LA | LAO | 418";"LAO";".la";""
    "Latvia";"LVA";"LV | LVA | 428";"LVA";".lv";""
    "Lebanon";"LBN";"LB | LBN | 422";"LBN";".lb";""
    "Lesotho";"LSO";"LS | LSO | 426";"LSO";".ls";""
    "Liberia";"LBR";"LR | LBR | 430";"LBR";".lr";""
    "Libya";"LBY";"LY | LBY | 434";"LBY";".ly";""
    "Liechtenstein";"LIE";"LI | LIE | 438";"LIE";".li";""
    "Lithuania";"LTU";"LT | LTU | 440";"LTU";".lt";""
    "Luxembourg";"LUX";"LU | LUX | 442";"LUX";".lu";""
    "Macau";"MAC";"MO | MAC | 446";"MAC";".mo";""
    "Madagascar";"MDG";"MG | MDG | 450";"MDG";".mg";""
    "Malawi";"MWI";"MW | MWI | 454";"MWI";".mw";""
    "Malaysia";"MYS";"MY | MYS | 458";"MYS";".my";""
    "Maldives";"MDV";"MV | MDV | 462";"MDV";".mv";""
    "Mali";"MLI";"ML | MLI | 466";"MLI";".ml";""
    "Malta";"MLT";"MT | MLT | 470";"MLT";".mt";""
    "Marshall Islands";"MHL";"MH | MHL | 584";"MHL";".mh";""
    "Martinique";"MTQ";"MQ | MTQ | 474";"MTQ";".mq";""
    "Mauritania";"MRT";"MR | MRT | 478";"MRT";".mr";""
    "Mauritius";"MUS";"MU | MUS | 480";"MUS";".mu";""
    "Mayotte";"MYT";"YT | MYT | 175";"FRA";".yt";""
    "Mexico";"MEX";"MX | MEX | 484";"MEX";".mx";""
    "Micronesia, Federated States of";"FSM";"FM | FSM | 583";"FSM";".fm";""
    "Midway Islands";"XMW";"- | - | -";"UMI";"-";"ISO includes with the US Minor Outlying Islands"
    "Moldova";"MDA";"MD | MDA | 498";"MDA";".md";""
    "Monaco";"MCO";"MC | MCO | 492";"MCO";".mc";""
    "Mongolia";"MNG";"MN | MNG | 496";"MNG";".mn";""
    "Montenegro";"MNE";"ME | MNE | 499";"MNE";".me";""
    "Montserrat";"MSR";"MS | MSR | 500";"MSR";".ms";""
    "Morocco";"MAR";"MA | MAR | 504";"MAR";".ma";""
    "Mozambique";"MOZ";"MZ | MOZ | 508";"MOZ";".mz";""
    "Myanmar";"-";"- | - | -";"-";"-";"see Burma"
    "Namibia";"NAM";"NA | NAM | 516";"NAM";".na";""
    "Nauru";"NRU";"NR | NRU | 520";"NRU";".nr";""
    "Navassa Island";"XNV";"- | - | -";"UMI";"-";"ISO includes with the US Minor Outlying Islands"
    "Nepal";"NPL";"NP | NPL | 524";"NPL";".np";""
    "Netherlands";"NLD";"NL | NLD | 528";"NLD";".nl";""
    "New Caledonia";"NCL";"NC | NCL | 540";"NCL";".nc";""
    "New Zealand";"NZL";"NZ | NZL | 554";"NZL";".nz";""
    "Nicaragua";"NIC";"NI | NIC | 558";"NIC";".ni";""
    "Niger";"NER";"NE | NER | 562";"NER";".ne";""
    "Nigeria";"NGA";"NG | NGA | 566";"NGA";".ng";""
    "Niue";"NIU";"NU | NIU | 570";"NIU";".nu";""
    "Norfolk Island";"NFK";"NF | NFK | 574";"NFK";".nf";""
    "North Macedonia";"MKD";"MK | MKD | 807";"FYR";".mk";""
    "Northern Mariana Islands";"MNP";"MP | MNP | 580";"MNP";".mp";""
    "Norway";"NOR";"NO | NOR | 578";"NOR";".no";""
    "Oman";"OMN";"OM | OMN | 512";"OMN";".om";""
    "Pakistan";"PAK";"PK | PAK | 586";"PAK";".pk";""
    "Palau";"PLW";"PW | PLW | 585";"PLW";".pw";""
    "Palmyra Atoll";"XPL";"- | - | -";"UMI";"-";"ISO includes with the US Minor Outlying Islands"
    "Panama";"PAN";"PA | PAN | 591";"PAN";".pa";""
    "Papua New Guinea";"PNG";"PG | PNG | 598";"PNG";".pg";""
    "Paracel Islands";"XPR";"- | - | -";"-";"-";""
    "Paraguay";"PRY";"PY | PRY | 600";"PRY";".py";""
    "Peru";"PER";"PE | PER | 604";"PER";".pe";""
    "Philippines";"PHL";"PH | PHL | 608";"PHL";".ph";""
    "Pitcairn Islands";"PCN";"PN | PCN | 612";"PCN";".pn";""
    "Poland";"POL";"PL | POL | 616";"POL";".pl";""
    "Portugal";"PRT";"PT | PRT | 620";"PRT";".pt";""
    "Puerto Rico";"PRI";"PR | PRI | 630";"PRI";".pr";""
    "Qatar";"QAT";"QA | QAT | 634";"QAT";".qa";""
    "Reunion";"REU";"RE | REU | 638";"REU";".re";""
    "Romania";"ROU";"RO | ROU | 642";"ROU";".ro";""
    "Russia";"RUS";"RU | RUS | 643";"RUS";".ru";""
    "Rwanda";"RWA";"RW | RWA | 646";"RWA";".rw";""
    "Saint Barthelemy";"BLM";"BL | BLM | 652";"-";".bl";"ccTLD  .fr and .gp may also be used"
    "Saint Helena, Ascension, and Tristan da Cunha";"SHN";"SH | SHN | 654";"SHN";".sh";"includes Saint Helena Island, Ascension Island, and the Tristan da Cunha archipelago"
    "Saint Kitts and Nevis";"KNA";"KN | KNA | 659";"KNA";".kn";""
    "Saint Lucia";"LCA";"LC | LCA | 662";"LCA";".lc";""
    "Saint Martin";"MAF";"MF | MAF | 663";"-";".mf";"ccTLD  .fr and .gp may also be used"
    "Saint Pierre and Miquelon";"SPM";"PM | SPM | 666";"SPM";".pm";""
    "Saint Vincent and the Grenadines";"VCT";"VC | VCT | 670";"VCT";".vc";""
    "Samoa";"WSM";"WS | WSM | 882";"WSM";".ws";""
    "San Marino";"SMR";"SM | SMR | 674";"SMR";".sm";""
    "Sao Tome and Principe";"STP";"ST | STP | 678";"STP";".st";""
    "Saudi Arabia";"SAU";"SA | SAU | 682";"SAU";".sa";""
    "Senegal";"SEN";"SN | SEN | 686";"SEN";".sn";""
    "Serbia";"SRB";"RS | SRB | 688";"-";".rs";""
    "Seychelles";"SYC";"SC | SYC | 690";"SYC";".sc";""
    "Sierra Leone";"SLE";"SL | SLE | 694";"SLE";".sl";""
    "Singapore";"SGP";"SG | SGP | 702";"SGP";".sg";""
    "Sint Maarten";"SXM";"SX | SXM | 534";"-";".sx";""
    "Slovakia";"SVK";"SK | SVK | 703";"SVK";".sk";""
    "Slovenia";"SVN";"SI | SVN | 705";"SVN";".si";""
    "Solomon Islands";"SLB";"SB | SLB | 090";"SLB";".sb";""
    "Somalia";"SOM";"SO | SOM | 706";"SOM";".so";""
    "South Africa";"ZAF";"ZA | ZAF | 710";"ZAF";".za";""
    "South Georgia and the Islands";"SGS";"GS | SGS | 239";"SGS";".gs";""
    "South Sudan";"SSD";"SS | SSD | 728";"-";"-";"IANA has designated .ss as the ccTLD for South Sudan, however it has not been activated in DNS root zone"
    "Spain";"ESP";"ES | ESP | 724";"ESP";".es";""
    "Spratly Islands";"XSP";"- | - | -";"-";"-";""
    "Sri Lanka";"LKA";"LK | LKA | 144";"LKA";".lk";""
    "Sudan";"SDN";"SD | SDN | 729";"SDN";".sd";""
    "Suriname";"SUR";"SR | SUR | 740";"SUR";".sr";""
    "Svalbard";"XSV";"SJ | SJM | 744";"SJM";".sj";"ISO includes Jan Mayen"
    "Sweden";"SWE";"SE | SWE | 752";"SWE";".se";""
    "Switzerland";"CHE";"CH | CHE | 756";"CHE";".ch";""
    "Syria";"SYR";"SY | SYR | 760";"SYR";".sy";""
    "Taiwan";"TWN";"TW | TWN | 158";"TWN";".tw";""
    "Tajikistan";"TJK";"TJ | TJK | 762";"TJK";".tj";""
    "Tanzania";"TZA";"TZ | TZA | 834";"TZA";".tz";""
    "Thailand";"THA";"TH | THA | 764";"THA";".th";""
    "Timor-Leste";"TLS";"TL | TLS | 626";"TLS";".tl";""
    "Togo";"TGO";"TG | TGO | 768";"TGO";".tg";""
    "Tokelau";"TKL";"TK | TKL | 772";"TKL";".tk";""
    "Tonga";"TON";"TO | TON | 776";"TON";".to";""
    "Trinidad and Tobago";"TTO";"TT | TTO | 780";"TTO";".tt";""
    "Tromelin Island";"XTR";"- | - | -";"-";"-";"administered as part of French Southern and Antarctic Lands; no ISO codes assigned"
    "Tunisia";"TUN";"TN | TUN | 788";"TUN";".tn";""
    "Turkey (Turkiye)";"TUR";"TR | TUR | 792";"TUR";".tr";""
    "Turkmenistan";"TKM";"TM | TKM | 795";"TKM";".tm";""
    "Turks and Caicos Islands";"TCA";"TC | TCA | 796";"TCA";".tc";""
    "Tuvalu";"TUV";"TV | TUV | 798";"TUV";".tv";""
    "Uganda";"UGA";"UG | UGA | 800";"UGA";".ug";""
    "Ukraine";"UKR";"UA | UKR | 804";"UKR";".ua";""
    "United Arab Emirates";"ARE";"AE | ARE | 784";"ARE";".ae";""
    "United Kingdom";"GBR";"GB | GBR | 826";"GBR";".uk";"for its internal communications, the  European Union recommends the use of the code UK in lieu of the ISO 3166-2 code of GB"
    "United States";"USA";"US | USA | 840";"USA";".us";""
    "United States Minor Outlying Islands";"-";"UM | UMI | 581";"-";".um";"ISO includes Baker Island, Howland Island, Jarvis Island, Johnston Atoll, Kingman Reef, Midway Islands, Navassa Island, Palmyra Atoll, Wake Island"
    "Uruguay";"URY";"UY | URY | 858";"URY";".uy";""
    "Uzbekistan";"UZB";"UZ | UZB | 860";"UZB";".uz";""
    "Vanuatu";"VUT";"VU | VUT | 548";"VUT";".vu";""
    "Venezuela";"VEN";"VE | VEN | 862";"VEN";".ve";""
    "Vietnam";"VNM";"VN | VNM | 704";"VNM";".vn";""
    "Virgin Islands";"VIR";"VI | VIR | 850";"VIR";".vi";""
    "Virgin Islands (UK)";"-";"- | - | -";"-";".vg";"see British Virgin Islands"
    "Virgin Islands (US)";"-";"- | - | -";"-";".vi";"see Virgin Islands"
    "Wake Island";"XWK";"- | - | -";"UMI";"-";"ISO includes with the US Minor Outlying Islands"
    "Wallis and Futuna";"WLF";"WF | WLF | 876";"WLF";".wf";""
    "West Bank";"XWB";"PS | PSE | 275";"PSE";".ps";"ISO identifies as Occupied Palestinian Territory"
    "Western Sahara";"WI";"EH | ESH | 732";"ESH";".eh";""
    "Western Samoa";"-";"- | - | -";"-";".ws";"see Samoa"
    "World";"-";"- | - | -";"-";"-";"the Factbook uses the W data code from DIAM 65-18 Geopolitical Data Elements and Related Features, Data Standard No. 3, December 1994, published by the Defense Intelligence Agency"
    "Yemen";"YEM";"YE | YEM | 887";"YEM";".ye";""
    "Zaire";"-";"- | - | -";"-";"-";"see Democratic Republic of the Congo"
    "Zambia";"ZMB";"ZM | ZMB | 894";"ZMB";".zm";""
    "Zimbabwe";"ZWE";"ZW | ZWE | 716";"ZWE";".zw";""`
)
