package tokens

import (
	"bufio"
	"errors"
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

func ReadString(str *string) error {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return err
	} else {
		*str = line
	}
	return nil
}

func Tokenize(str string, arr *[]Token) error {
	rune_str := []rune(str)
	i := 0
	for i != len(rune_str) {
		if unicode.IsDigit(rune_str[i]) {
			start := i
			for unicode.IsDigit(rune_str[i]) {
				i++
			}
			num := []rune(rune_str[start:i])
			*arr = append(*arr, Token{Str: string(num), Tok: Num})
		}
		if rune_str[i] == '+' || rune_str[i] == '-' {
			*arr = append(*arr, Token{Str: string(rune_str[i]), Tok: PlusMinus})
		}
		if rune_str[i] == '*' || rune_str[i] == '/' {
			*arr = append(*arr, Token{Str: string(rune_str[i]), Tok: MultDiv})
		}
		if rune_str[i] == '(' {
			*arr = append(*arr, Token{Str: string(rune_str[i]), Tok: LParen})
		}
		if rune_str[i] == ')' {
			*arr = append(*arr, Token{Str: string(rune_str[i]), Tok: RParen})
		}
		if unicode.IsLetter(rune_str[i]) {
			start := i
			for unicode.IsLetter(rune_str[i]) {
				i++
			}
			res := string(rune_str[start:i])
			switch res {
			case "x":
				*arr = append(*arr, Token{Str: res, Tok: Var})
				i--
			case "sin", "cos", "ctg", "sqrt", "ln":
				*arr = append(*arr, Token{Str: res, Tok: Func})
				i--
			default:
				return errors.New("недопустимый символ")
			}
		}
		i++
	}
	return nil
}
