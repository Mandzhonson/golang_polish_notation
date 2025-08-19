package tokens

import (
	"bufio"
	"errors"
	"os"
	"unicode"
)

// определение типа токена
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
	// считывание строки
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
	// разделение по токенам нашей строки для дальнейшей обработки
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

func CheckParen(arr *[]Token) error {
	// проверка на корректный ввод скобок, сначала количество открытых(-1) и закрытых(+1)
	// в сумме должны давать 0
	// далее под каждую открывающуюся скобку ищем закрывающаюся
	count := 0
	for i := range *arr {
		if (*arr)[i].Tok == LParen {
			count += -1

		}
		if (*arr)[i].Tok == RParen {
			count += 1
		}
	}
	if count != 0 {
		return errors.New("нарушен баланс скобок")
	} else {
		for i := range *arr {
			if (*arr)[i].Tok == RParen && count == 0 {
				return errors.New("неверный порядок расстановки скобок")
			} else {
				if (*arr)[i].Tok == LParen {
					if i+1 < len(*arr) {
						j := i + 1
						count -= 1
						for (*arr)[j].Tok != RParen {
							j++
						}
						if (*arr)[j].Tok == RParen && j-i == 1 {
							return errors.New("пустые скобки")
						}
					} else {
						return errors.New("неверный порядок расстановки скобок")
					}
				}
			}
		}
	}
	return nil
}

func CheckToken(arr *[]Token) error {
	// проверка токенов на корректность для дальнейшей работы
	err := CheckParen(arr)
	if err != nil {
		return err
	}
	// next work
	return nil
}
