package tokens

import (
	"bufio"
	"os"
	"unicode"
)

const (
	Num       = 1
	PlusMinus = 2
	MultDiv   = 3
	LParen    = 4
	RParen    = 5
	Var       = 6
	Func      = 7
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

func Tokenize(str string, arr *[]Token) int {
	flag := 1
	rune_str := []rune(str)
	i := 0
	for i != len(rune_str) {
		if unicode.IsDigit(rune_str[i]) && flag == 1 {
			start := i
			for unicode.IsDigit(rune_str[i]) {
				i++
			}
			num := []rune(rune_str[start:i])
			*arr = append(*arr, Token{Str: string(num), Tok: Num})
		}
		if (rune_str[i] == '+' || rune_str[i] == '-') && flag == 1 {
			*arr = append(*arr, Token{Str: string(rune_str[i]), Tok: PlusMinus})
		}
		if (rune_str[i] == '*' || rune_str[i] == '/') && flag == 1 {
			*arr = append(*arr, Token{Str: string(rune_str[i]), Tok: MultDiv})
		}
		if rune_str[i] == '(' && flag == 1 {
			*arr = append(*arr, Token{Str: string(rune_str[i]), Tok: LParen})
		}
		if rune_str[i] == ')' && flag == 1 {
			*arr = append(*arr, Token{Str: string(rune_str[i]), Tok: RParen})
		}
		if unicode.IsLetter(rune_str[i]) && flag == 1 {
			start := i
			for unicode.IsLetter(rune_str[i]) {
				i++
			}
			res := string(rune_str[start:i])
			switch res {
			case "x":
				*arr = append(*arr, Token{Str: res, Tok: Var})
			case "sin", "cos", "ctg", "sqrt", "ln":
				*arr = append(*arr, Token{Str: res, Tok: Func})
			default:
				flag = 0
			}
		}
		i++
	}
	return flag
}
