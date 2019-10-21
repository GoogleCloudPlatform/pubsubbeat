package cp

var cp874 *charsetMap = &charsetMap{
	sb: [256]rune{
		0x0000, //NULL
		0x0001, //START OF HEADING
		0x0002, //START OF TEXT
		0x0003, //END OF TEXT
		0x0004, //END OF TRANSMISSION
		0x0005, //ENQUIRY
		0x0006, //ACKNOWLEDGE
		0x0007, //BELL
		0x0008, //BACKSPACE
		0x0009, //HORIZONTAL TABULATION
		0x000A, //LINE FEED
		0x000B, //VERTICAL TABULATION
		0x000C, //FORM FEED
		0x000D, //CARRIAGE RETURN
		0x000E, //SHIFT OUT
		0x000F, //SHIFT IN
		0x0010, //DATA LINK ESCAPE
		0x0011, //DEVICE CONTROL ONE
		0x0012, //DEVICE CONTROL TWO
		0x0013, //DEVICE CONTROL THREE
		0x0014, //DEVICE CONTROL FOUR
		0x0015, //NEGATIVE ACKNOWLEDGE
		0x0016, //SYNCHRONOUS IDLE
		0x0017, //END OF TRANSMISSION BLOCK
		0x0018, //CANCEL
		0x0019, //END OF MEDIUM
		0x001A, //SUBSTITUTE
		0x001B, //ESCAPE
		0x001C, //FILE SEPARATOR
		0x001D, //GROUP SEPARATOR
		0x001E, //RECORD SEPARATOR
		0x001F, //UNIT SEPARATOR
		0x0020, //SPACE
		0x0021, //EXCLAMATION MARK
		0x0022, //QUOTATION MARK
		0x0023, //NUMBER SIGN
		0x0024, //DOLLAR SIGN
		0x0025, //PERCENT SIGN
		0x0026, //AMPERSAND
		0x0027, //APOSTROPHE
		0x0028, //LEFT PARENTHESIS
		0x0029, //RIGHT PARENTHESIS
		0x002A, //ASTERISK
		0x002B, //PLUS SIGN
		0x002C, //COMMA
		0x002D, //HYPHEN-MINUS
		0x002E, //FULL STOP
		0x002F, //SOLIDUS
		0x0030, //DIGIT ZERO
		0x0031, //DIGIT ONE
		0x0032, //DIGIT TWO
		0x0033, //DIGIT THREE
		0x0034, //DIGIT FOUR
		0x0035, //DIGIT FIVE
		0x0036, //DIGIT SIX
		0x0037, //DIGIT SEVEN
		0x0038, //DIGIT EIGHT
		0x0039, //DIGIT NINE
		0x003A, //COLON
		0x003B, //SEMICOLON
		0x003C, //LESS-THAN SIGN
		0x003D, //EQUALS SIGN
		0x003E, //GREATER-THAN SIGN
		0x003F, //QUESTION MARK
		0x0040, //COMMERCIAL AT
		0x0041, //LATIN CAPITAL LETTER A
		0x0042, //LATIN CAPITAL LETTER B
		0x0043, //LATIN CAPITAL LETTER C
		0x0044, //LATIN CAPITAL LETTER D
		0x0045, //LATIN CAPITAL LETTER E
		0x0046, //LATIN CAPITAL LETTER F
		0x0047, //LATIN CAPITAL LETTER G
		0x0048, //LATIN CAPITAL LETTER H
		0x0049, //LATIN CAPITAL LETTER I
		0x004A, //LATIN CAPITAL LETTER J
		0x004B, //LATIN CAPITAL LETTER K
		0x004C, //LATIN CAPITAL LETTER L
		0x004D, //LATIN CAPITAL LETTER M
		0x004E, //LATIN CAPITAL LETTER N
		0x004F, //LATIN CAPITAL LETTER O
		0x0050, //LATIN CAPITAL LETTER P
		0x0051, //LATIN CAPITAL LETTER Q
		0x0052, //LATIN CAPITAL LETTER R
		0x0053, //LATIN CAPITAL LETTER S
		0x0054, //LATIN CAPITAL LETTER T
		0x0055, //LATIN CAPITAL LETTER U
		0x0056, //LATIN CAPITAL LETTER V
		0x0057, //LATIN CAPITAL LETTER W
		0x0058, //LATIN CAPITAL LETTER X
		0x0059, //LATIN CAPITAL LETTER Y
		0x005A, //LATIN CAPITAL LETTER Z
		0x005B, //LEFT SQUARE BRACKET
		0x005C, //REVERSE SOLIDUS
		0x005D, //RIGHT SQUARE BRACKET
		0x005E, //CIRCUMFLEX ACCENT
		0x005F, //LOW LINE
		0x0060, //GRAVE ACCENT
		0x0061, //LATIN SMALL LETTER A
		0x0062, //LATIN SMALL LETTER B
		0x0063, //LATIN SMALL LETTER C
		0x0064, //LATIN SMALL LETTER D
		0x0065, //LATIN SMALL LETTER E
		0x0066, //LATIN SMALL LETTER F
		0x0067, //LATIN SMALL LETTER G
		0x0068, //LATIN SMALL LETTER H
		0x0069, //LATIN SMALL LETTER I
		0x006A, //LATIN SMALL LETTER J
		0x006B, //LATIN SMALL LETTER K
		0x006C, //LATIN SMALL LETTER L
		0x006D, //LATIN SMALL LETTER M
		0x006E, //LATIN SMALL LETTER N
		0x006F, //LATIN SMALL LETTER O
		0x0070, //LATIN SMALL LETTER P
		0x0071, //LATIN SMALL LETTER Q
		0x0072, //LATIN SMALL LETTER R
		0x0073, //LATIN SMALL LETTER S
		0x0074, //LATIN SMALL LETTER T
		0x0075, //LATIN SMALL LETTER U
		0x0076, //LATIN SMALL LETTER V
		0x0077, //LATIN SMALL LETTER W
		0x0078, //LATIN SMALL LETTER X
		0x0079, //LATIN SMALL LETTER Y
		0x007A, //LATIN SMALL LETTER Z
		0x007B, //LEFT CURLY BRACKET
		0x007C, //VERTICAL LINE
		0x007D, //RIGHT CURLY BRACKET
		0x007E, //TILDE
		0x007F, //DELETE
		0x20AC, //EURO SIGN
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0x2026, //HORIZONTAL ELLIPSIS
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0x2018, //LEFT SINGLE QUOTATION MARK
		0x2019, //RIGHT SINGLE QUOTATION MARK
		0x201C, //LEFT DOUBLE QUOTATION MARK
		0x201D, //RIGHT DOUBLE QUOTATION MARK
		0x2022, //BULLET
		0x2013, //EN DASH
		0x2014, //EM DASH
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0x00A0, //NO-BREAK SPACE
		0x0E01, //THAI CHARACTER KO KAI
		0x0E02, //THAI CHARACTER KHO KHAI
		0x0E03, //THAI CHARACTER KHO KHUAT
		0x0E04, //THAI CHARACTER KHO KHWAI
		0x0E05, //THAI CHARACTER KHO KHON
		0x0E06, //THAI CHARACTER KHO RAKHANG
		0x0E07, //THAI CHARACTER NGO NGU
		0x0E08, //THAI CHARACTER CHO CHAN
		0x0E09, //THAI CHARACTER CHO CHING
		0x0E0A, //THAI CHARACTER CHO CHANG
		0x0E0B, //THAI CHARACTER SO SO
		0x0E0C, //THAI CHARACTER CHO CHOE
		0x0E0D, //THAI CHARACTER YO YING
		0x0E0E, //THAI CHARACTER DO CHADA
		0x0E0F, //THAI CHARACTER TO PATAK
		0x0E10, //THAI CHARACTER THO THAN
		0x0E11, //THAI CHARACTER THO NANGMONTHO
		0x0E12, //THAI CHARACTER THO PHUTHAO
		0x0E13, //THAI CHARACTER NO NEN
		0x0E14, //THAI CHARACTER DO DEK
		0x0E15, //THAI CHARACTER TO TAO
		0x0E16, //THAI CHARACTER THO THUNG
		0x0E17, //THAI CHARACTER THO THAHAN
		0x0E18, //THAI CHARACTER THO THONG
		0x0E19, //THAI CHARACTER NO NU
		0x0E1A, //THAI CHARACTER BO BAIMAI
		0x0E1B, //THAI CHARACTER PO PLA
		0x0E1C, //THAI CHARACTER PHO PHUNG
		0x0E1D, //THAI CHARACTER FO FA
		0x0E1E, //THAI CHARACTER PHO PHAN
		0x0E1F, //THAI CHARACTER FO FAN
		0x0E20, //THAI CHARACTER PHO SAMPHAO
		0x0E21, //THAI CHARACTER MO MA
		0x0E22, //THAI CHARACTER YO YAK
		0x0E23, //THAI CHARACTER RO RUA
		0x0E24, //THAI CHARACTER RU
		0x0E25, //THAI CHARACTER LO LING
		0x0E26, //THAI CHARACTER LU
		0x0E27, //THAI CHARACTER WO WAEN
		0x0E28, //THAI CHARACTER SO SALA
		0x0E29, //THAI CHARACTER SO RUSI
		0x0E2A, //THAI CHARACTER SO SUA
		0x0E2B, //THAI CHARACTER HO HIP
		0x0E2C, //THAI CHARACTER LO CHULA
		0x0E2D, //THAI CHARACTER O ANG
		0x0E2E, //THAI CHARACTER HO NOKHUK
		0x0E2F, //THAI CHARACTER PAIYANNOI
		0x0E30, //THAI CHARACTER SARA A
		0x0E31, //THAI CHARACTER MAI HAN-AKAT
		0x0E32, //THAI CHARACTER SARA AA
		0x0E33, //THAI CHARACTER SARA AM
		0x0E34, //THAI CHARACTER SARA I
		0x0E35, //THAI CHARACTER SARA II
		0x0E36, //THAI CHARACTER SARA UE
		0x0E37, //THAI CHARACTER SARA UEE
		0x0E38, //THAI CHARACTER SARA U
		0x0E39, //THAI CHARACTER SARA UU
		0x0E3A, //THAI CHARACTER PHINTHU
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0x0E3F, //THAI CURRENCY SYMBOL BAHT
		0x0E40, //THAI CHARACTER SARA E
		0x0E41, //THAI CHARACTER SARA AE
		0x0E42, //THAI CHARACTER SARA O
		0x0E43, //THAI CHARACTER SARA AI MAIMUAN
		0x0E44, //THAI CHARACTER SARA AI MAIMALAI
		0x0E45, //THAI CHARACTER LAKKHANGYAO
		0x0E46, //THAI CHARACTER MAIYAMOK
		0x0E47, //THAI CHARACTER MAITAIKHU
		0x0E48, //THAI CHARACTER MAI EK
		0x0E49, //THAI CHARACTER MAI THO
		0x0E4A, //THAI CHARACTER MAI TRI
		0x0E4B, //THAI CHARACTER MAI CHATTAWA
		0x0E4C, //THAI CHARACTER THANTHAKHAT
		0x0E4D, //THAI CHARACTER NIKHAHIT
		0x0E4E, //THAI CHARACTER YAMAKKAN
		0x0E4F, //THAI CHARACTER FONGMAN
		0x0E50, //THAI DIGIT ZERO
		0x0E51, //THAI DIGIT ONE
		0x0E52, //THAI DIGIT TWO
		0x0E53, //THAI DIGIT THREE
		0x0E54, //THAI DIGIT FOUR
		0x0E55, //THAI DIGIT FIVE
		0x0E56, //THAI DIGIT SIX
		0x0E57, //THAI DIGIT SEVEN
		0x0E58, //THAI DIGIT EIGHT
		0x0E59, //THAI DIGIT NINE
		0x0E5A, //THAI CHARACTER ANGKHANKHU
		0x0E5B, //THAI CHARACTER KHOMUT
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
		0xFFFD, //UNDEFINED
	},
}
