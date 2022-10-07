package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func IsolateParathes(str string) string {
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		if runes[i] == ')' && i < len(runes)-1 && !IsSpace(runes[i+1]) {
			temp := make([]rune, len(runes[i+1:]))
			copy(temp, runes[i+1:])
			runes[i+1] = ' '
			runes = append(runes[:i+2], temp...)
		} else if runes[i] == '(' && i > 0 && !IsSpace(runes[i-1]) {
			temp := make([]rune, len(runes[i:]))
			copy(temp, runes[i:])
			runes[i] = ' '
			runes = append(runes[:i+1], temp...)
		}
	}
	return string(runes)
}

func PunctCorr(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if IsPunctuation(runes[i]) {
			for j := i - 1; j > 0; j-- {
				if runes[j] == ' ' || runes[j] == '\t' {
					runes = append(runes[:j], runes[j+1:]...)
					i--
				} else {
					break
				}
			}
			if i+1 < len(runes) && runes[i+1] != ' ' {
				temp := make([]rune, len(runes[i+1:]))
				copy(temp, runes[i+1:])
				runes[i+1] = ' '
				runes = append(runes[:i+2], temp...)

			}
		}
		if IsSpace(runes[i]) && i > 0 {
			if IsSpace(runes[i-1]) {
				runes = append(runes[:i], runes[i+1:]...)
				i--
			}
		}
	}
	return string(runes)
}

func QuotesCorr(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if runes[i] == '\'' {
			for j := i + 1; j < len(runes); j++ {
				if IsSpace(runes[j]) && runes[j-1] == '\'' {
					runes = append(runes[:j], runes[j+1:]...)
					j--
				} else if runes[j] == '\'' && IsSpace(runes[j-1]) {
					runes = append(runes[:j-1], runes[j:]...)
					j = j - 2
				} else if runes[j] == '\'' && !IsSpace(runes[j-1]) {
					i = j
					break
				}
			}
		}
	}
	return string(runes)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid number of arguments")
		return
	}

	readFile := os.Args[1]
	resultFile := os.Args[2]

	b, err := ioutil.ReadFile(readFile)
	if err != nil {
		log.Fatal(err)
	}

	sen := string(b)
	words := strings.Split(IsolateParathes(sen), " ")

	for i := 0; i < len(words); i++ {

		if words[i] == "" {
			words = append(words[:i], words[i+1:]...)
			i--
			continue
		}

		if rune(words[i][0]) == '(' && rune(words[i][len(words[i])-1]) == ')' {
			if words[i-1] != "" && words[i-1] != " " {
				switch words[i] {
				case "(hex)":
					words[i-1] = HexToDec(words[i-1])
				case "(bin)":
					words[i-1] = BinToDec(words[i-1])
				case "(up)":
					words[i-1] = Up(words[i-1])
				case "(low)":
					words[i-1] = Low(words[i-1])
				case "(cap)":
					words[i-1] = Cap(words[i-1])
				}
				words = append(words[:i], words[i+1:]...)
				i--
			}
		} else if rune(words[i][0]) == '(' && rune(words[i][len(words[i])-1]) == ',' {
			runes := []rune(words[i+1])
			if runes[len(runes)-1] == ')' {
				num, err := strconv.Atoi(string(runes[:len(runes)-1]))
				if err != nil {
					log.Fatal(err)
				}
				for j := num; j > 0; j-- {
					if i-j >= 0 {
						switch words[i] {
						case "(up,":
							words[i-j] = Up(words[i-j])
						case "(low,":
							words[i-j] = Low(words[i-j])
						case "(cap,":
							words[i-j] = Cap(words[i-j])
						}
					}
				}
				words = append(words[:i+1], words[i+2:]...)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		} else if words[i] == "a" {
			if IsVomel(rune(words[i+1][0])) {
				words[i] = "an"
			}
		}
	}

	///////////////////////////////////////////
	resStr := StrArrToStr(words)

	f, err := os.Create(resultFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err2 := f.WriteString(QuotesCorr(PunctCorr(resStr)))
	if err2 != nil {
		log.Fatal(err2)
	}
}
