package algorithm

import (
	"errors"
	"fmt"
	"math"
	"polish/stack"
	"polish/tokens"
	"strconv"
)

const (
	Width  = 80
	Height = 25
	YMin   = -1
	YMax   = 1
)

func PolishNotation() {
	// PolishNotation входная точка программы
	var list string
	if err := tokens.ReadString(&list); err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	arr, err := tokens.Tokenize(list)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(arr)
	if err = tokens.CheckToken(arr); err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	res := Algorithm(&arr)
	if err = DrawGraphic(&res); err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}

func Algorithm(arr *[]tokens.Token) []tokens.Token {
	// Algorithm получает постфиксную строку алгоритмом Дейкстры
	var st_op stack.Stack
	var str []tokens.Token
	for i := range *arr {
		if (*arr)[i].Tok == tokens.Var || (*arr)[i].Tok == tokens.Num {
			str = append(str, (*arr)[i])
		} else if (*arr)[i].Str == "(" {
			st_op.Push((*arr)[i])
		} else if (*arr)[i].Str == ")" {
			for st_op.Peek().Str != "(" && !st_op.IsEmpty() {
				str = append(str, st_op.Peek())
				st_op.Pop()
			}
			st_op.Pop()
		} else {
			if (*arr)[i].Tok > st_op.Peek().Tok && st_op.Peek().Str != "(" {
				st_op.Push((*arr)[i])
			} else {
				for (*arr)[i].Tok <= st_op.Peek().Tok && st_op.Peek().Str != "(" {
					str = append(str, st_op.Peek())
					st_op.Pop()
				}
				st_op.Push((*arr)[i])
			}
		}
	}
	for !st_op.IsEmpty() {
		str = append(str, st_op.Peek())
		st_op.Pop()
	}
	return str
}

func Calculate(arr []tokens.Token, num float64) (float64, error) {
	// Calculate вычисляет значение выражения
	var st_num stack.Stack
	workArr := make([]tokens.Token, len(arr))
	copy(workArr, arr)
	for i := range workArr {
		if workArr[i].Str == "x" {
			strNum := strconv.FormatFloat(num, 'f', -1, 64)
			workArr[i].Str = strNum
			workArr[i].Tok = tokens.Num
		}
		if workArr[i].Tok == tokens.Num {
			st_num.Push(workArr[i])
		} else if workArr[i].Tok != tokens.Func {
			num2, err := strconv.ParseFloat(st_num.Peek().Str, 64)
			if err != nil {
				return 0, err
			}
			st_num.Pop()
			num1, err := strconv.ParseFloat(st_num.Peek().Str, 64)
			if err != nil {
				return 0, err
			}
			st_num.Pop()
			switch workArr[i].Str {
			case "+":
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(num1+num2, 'f', -1, 64), Tok: tokens.Num})
			case "-":
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(num1-num2, 'f', -1, 64), Tok: tokens.Num})
			case "*":
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(num1*num2, 'f', -1, 64), Tok: tokens.Num})
			case "/":
				if num2 == 0 {
					return 0, errors.New("деление на 0 запрещено")
				}
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(num1/num2, 'f', -1, 64), Tok: tokens.Num})
			}
		} else {
			num1, err := strconv.ParseFloat(st_num.Peek().Str, 64)
			if err != nil {
				return 0, err
			}
			st_num.Pop()
			switch workArr[i].Str {
			case "sin":
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(math.Sin(num1), 'f', -1, 64)})
			case "cos":
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(math.Cos(num1), 'f', -1, 64)})
			case "ln":
				if num1 <= 0 {
					return 0, errors.New("операция ln: число не соответствует допустимым значениям")
				}
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(math.Log(num1), 'f', -1, 64)})
			case "sqrt":
				if num1 < 0 {
					return 0, errors.New("операция sqrt: число не соответствует допустимым значениям")
				}
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(math.Sqrt(num1), 'f', -1, 64)})
			case "ctg":
				cos := math.Cos(num1)
				sin := math.Sin(num1)
				if sin == 0 {
					return 0, errors.New("операция ctg: sin не должен быть равен 0")
				}
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(cos/sin, 'f', -1, 64)})
			case "tan":
				cos := math.Cos(num1)
				sin := math.Sin(num1)
				if cos == 0 {
					return 0, errors.New("операция tan: cos не должен быть равен 0")
				}
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(sin/cos, 'f', -1, 64)})
			}
		}
	}
	res, err := strconv.ParseFloat(st_num.Peek().Str, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func DrawGraphic(arr *[]tokens.Token) error {
	//DrawGraphic вычисляет значения функции и выводит на экран готовый результат
	var res [Height][Width]int
	step := (4 * math.Pi) / 79
	for i := 0; i < Width; i++ {
		val, err := Calculate(*arr, (step * float64(i)))
		if err != nil {
			return err
		}
		yNormal := 24 - int(math.Round(24*(val-YMin)/(YMax-YMin)))
		if yNormal >= 0 && yNormal < Height {
			res[yNormal][i] = 1
		}
	}
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			if res[i][j] == 0 && j+1 != Width {
				fmt.Print(". ")
			} else if res[i][j] == 1 && j+1 != Width {
				fmt.Print("* ")
			} else if res[i][j] == 1 && j+1 == Width {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		if i+1 != Height {
			fmt.Print("\n")
		}
	}
	return nil
}
