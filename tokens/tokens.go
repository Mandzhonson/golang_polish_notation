package tokens

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode"
)

// определение типа токена и приоритет
const (
	Num = 1
	Var
	PlusMinus  = 2
	MultDiv    = 3
	UnaryMinus = 4
	Func       = 5
	LParen     = 6
	RParen
)

func ReadString(str *string) error {
	// считывание строки
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	*str = line
	return nil
}

func Tokenize(str string) ([]Token, error) {
	// разделение по токенам нашей строки для дальнейшей обработки
	var arr []Token
	rune_str := []rune(str)
	i := 0
	for i != len(rune_str) {
		if unicode.IsDigit(rune_str[i]) {
			start := i
			for unicode.IsDigit(rune_str[i]) && i < len(rune_str) {
				i++
			}
			num := []rune(rune_str[start:i])
			arr = append(arr, Token{Str: string(num), Tok: Num})
		}
		if rune_str[i] == '+' || rune_str[i] == '-' {
			if i-1 >= 0 && (arr)[i-1].Tok == Num {
				arr = append(arr, Token{Str: string(rune_str[i]), Tok: PlusMinus})
			} else if i-1 < 0 || (i-1 >= 0 && (arr[i-1].Str == "(" || arr[i-1].Tok == PlusMinus || arr[i-1].Tok == MultDiv)) {
				arr = append(arr, Token{Str: string(rune_str[i]), Tok: UnaryMinus})
			} else {
				return []Token{}, errors.New("ошибка распознавания знака или неправильная расстановка")
			}
		}
		if rune_str[i] == '*' || rune_str[i] == '/' {
			arr = append(arr, Token{Str: string(rune_str[i]), Tok: MultDiv})
		}
		if rune_str[i] == '(' {
			arr = append(arr, Token{Str: string(rune_str[i]), Tok: LParen})
		}
		if rune_str[i] == ')' {
			arr = append(arr, Token{Str: string(rune_str[i]), Tok: RParen})
		}
		if unicode.IsLetter(rune_str[i]) {
			start := i
			for unicode.IsLetter(rune_str[i]) && i < len(rune_str) {
				i++
			}
			res := string(rune_str[start:i])
			switch res {
			case "x":
				arr = append(arr, Token{Str: res, Tok: Var})
				i--
			case "sin", "cos", "ctg", "sqrt", "ln", "tan":
				arr = append(arr, Token{Str: res, Tok: Func})
				i--
			default:
				return []Token{}, errors.New("недопустимый символ")
			}
		}
		i++
	}
	return arr, nil
}

func CheckParen(arr []Token) error {
	// проверка на корректный ввод скобок, сначала количество открытых(-1) и закрытых(+1)
	// в сумме должны давать 0
	// далее под каждую открывающуюся скобку ищем закрывающаюся
	count := 0
	for i := range arr {
		switch arr[i].Str {
		case "(":
			if i+1 < len(arr) && (arr)[i+1].Str == ")" {
				return errors.New("пустые скобки")
			}
			count++
		case ")":
			count--
			if count < 0 {
				return errors.New("неверный порядок расстановки скобок")
			}
		}
	}
	if count != 0 {
		return errors.New("несбалансированные скобки")
	}
	return nil
}

func CheckOper(arr []Token) error {
	// корректность ввода операторов
	var prevToken Token
	for i := range arr {
		if arr[i].Tok == Num {
			if prevToken.Tok == Num || prevToken.Str == ")" {
				fmt.Println(arr[i], prevToken)
				return errors.New("некорректная строка: перед числом может стоять знак или открывающаяся скобка")
			}
		}
		if arr[i].Tok == PlusMinus || arr[i].Tok == MultDiv {
			if prevToken.Tok != Num {
				return errors.New("некорректная строка: перед бинарным знаком должно быть число")
			}
			if i+1 >= len(arr) || (i+1 < len(arr) && arr[i+1].Tok != Num && arr[i+1].Tok != Func) {
				return errors.New("некорректная строка: после бинарного знака может находиться число или математическая функция")
			}
		}
		if arr[i].Str == "(" {
			if prevToken.Str != "(" && prevToken.Tok != PlusMinus && prevToken.Tok != MultDiv && prevToken.Tok != Func {
				fmt.Println(arr[i], prevToken)
				return errors.New("некорректная строка: перед открывающейся скобкой может быть операция или открывающаяся скобка")
			}
		}
		if arr[i].Str == ")" {
			if prevToken.Tok == PlusMinus || prevToken.Tok == MultDiv || prevToken.Tok == UnaryMinus || prevToken.Tok == 0 || prevToken.Str == "(" {
				return errors.New("некорректная строка: перед закрывающейся скобкой не могут находиться операции и ()")
			}
		}
		if arr[i].Tok == Func {
			if i+1 >= len(arr) || (i+1 < len(arr) && arr[i+1].Str != "(") {
				return errors.New("некорректная строка: аргументы функции должны быть заключены в скобки")
			}
		}
		prevToken = arr[i]
	}
	return nil
}

func CheckFunc(arr []Token) error {
	// проверка на то, что после функции обязательно идет скобка открывающаяся
	for i := range arr {
		if arr[i].Tok == Func {
			if i+1 >= len(arr) {
				return errors.New("любая функция должна принимать аргумент в скобках")
			}
			if i+1 < len(arr) && (arr)[i+1].Str != "(" {
				return errors.New("любая функция должна принимать аргумент в скобках")
			}
		}
	}
	return nil
}

func CheckToken(arr []Token) error {
	// проверка токенов на корректность для дальнейшей работы
	if err := CheckParen(arr); err != nil {
		return err
	}
	if err := CheckOper(arr); err != nil {
		return err
	}
	if err := CheckFunc(arr); err != nil {
		return err
	}
	return nil
}
