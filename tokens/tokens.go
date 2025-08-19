package tokens

import (
	"bufio"
	"os"
	"unicode"
)

const (
	Num  = 1
	Plus = 2
	Minus
	Mult = 3
	Div
	LParen = 4
	RParen = 5
)

func ReadString(str *string) int {
	flag := 1
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		flag = 0
	} else {
		*str = line
	}
	return flag
}

func Tokenize(str string, arr *[]Token) {
	rune_str := []rune(str)
	var tokens []Token
	i := 0
	for i != len(rune_str) {
		if unicode.IsDigit(rune_str[i]) {
			j := i
			for unicode.IsDigit(rune_str[i]) {
				i++
			}
			num := []rune(rune_str[j:i])
			tokens = append(tokens, Token{string(num), Num})
		}
		i++
	}
	*arr = tokens
}
