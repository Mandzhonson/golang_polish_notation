package tokens

import (
	"bufio"
	"errors"
	"os"
	"unicode"
)

// определение типа токена и приоритет
const (
	Num = 1
	Var
	PlusMinus = 2
	MultDiv   = 3
	Func      = 4
	LParen    = 5
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

func Tokenize(str string, arr *[]Token) error {
	// разделение по токенам нашей строки для дальнейшей обработки
	rune_str := []rune(str)
	i := 0
	for i != len(rune_str) {
		if unicode.IsDigit(rune_str[i]) {
			start := i
			for unicode.IsDigit(rune_str[i]) && i < len(rune_str) {
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
			for unicode.IsLetter(rune_str[i]) && i < len(rune_str) {
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
		switch (*arr)[i].Str {
		case "(":
			if i+1 < len(*arr) && (*arr)[i+1].Str == ")" {
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

func CheckOper(arr *[]Token) error {
	// корректность ввода операторов (не учитывает унарные операции)
	for i := range *arr {
		if (*arr)[i].Tok == PlusMinus || (*arr)[i].Tok == MultDiv {
			if i+1 >= len(*arr) || i-1 < 0 {
				return errors.New("некорректная строка: неправильно стоят знаки")
			} else if (*arr)[i-1].Tok == PlusMinus || (*arr)[i-1].Tok == MultDiv || (*arr)[i+1].Tok == PlusMinus || (*arr)[i+1].Tok == MultDiv {
				return errors.New("некорректная строка: неправильно стоят знаки")
			} else if (*arr)[i-1].Str == "(" || (*arr)[i+1].Str == ")" {
				return errors.New("некорректная строка: неправильно стоят знаки")
			}
		}
	}
	for i := range *arr {
		if (*arr)[i].Tok == Num || (*arr)[i].Tok == Var {
			if i+1 >= len(*arr) && i-1 > 0 && (*arr)[i-1].Tok != PlusMinus && (*arr)[i-1].Tok != MultDiv {
				return errors.New("между числами должен быть знак")
			}
			if (i+1 < len(*arr) && ((*arr)[i+1].Tok == Num || (*arr)[i+1].Tok == Var)) || (i-1 > 0 && ((*arr)[i-1].Tok == Num || (*arr)[i-1].Tok == Var)) {
				return errors.New("между числами должен быть знак")
			}
			if i-1 >= 0 && i+1 < len(*arr) && (*arr)[i-1].Tok != PlusMinus && (*arr)[i-1].Tok != MultDiv && (*arr)[i-1].Str != "(" {
				return errors.New("между числами должен быть знак")
			}
		}
	}
	return nil
}

func CheckFunc(arr *[]Token) error {
	// проверка на то, что после функции обязательно идет скобка открывающаяся
	for i := range *arr {
		if (*arr)[i].Tok == Func {
			if i+1 >= len(*arr) {
				return errors.New("любая функция должна принимать аргумент в скобках")
			}
			if i+1 < len(*arr) && (*arr)[i+1].Str != "(" {
				return errors.New("любая функция должна принимать аргумент в скобках")
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
	err = CheckOper(arr)
	if err != nil {
		return err
	}
	err = CheckFunc(arr)
	if err != nil {
		return err
	}
	return nil
}
