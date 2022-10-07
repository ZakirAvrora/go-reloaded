package main

import "strconv"

func HexToDec(hex string) string {
	num, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(int(num))
}

func BinToDec(bin string) string {
	num, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(int(num))
}

func Up(str string) string {
	runes := []rune(str)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = r + ('A' - 'a')
		}
	}

	return string(runes)
}

func Low(str string) string {
	runes := []rune(str)
	for i, r := range runes {
		if r >= 'A' && r <= 'Z' {
			runes[i] = r + ('a' - 'A')
		}
	}

	return string(runes)
}

func Cap(str string) string {
	runes := []rune(str)
	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] = runes[0] + ('A' - 'a')
	}
	return string(runes)
}

func StrArrToStr(words []string) string {
	sen := words[0]
	for i := 1; i < len(words); i++ {
		sen += " "
		sen += words[i]
	}
	return sen
}

func IsSpace(r rune) bool {
	if r == ' ' || r == '\t' {
		return true
	}
	return false
}

func IsVomel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' ||
		r == 'h' {
		return true
	}

	return false
}

func IsPunctuation(r rune) bool {
	if r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';' {
		return true
	}
	return false
}
