package algorithm

import (
	"errors"
	"fmt"
	"math"
	"polish/stack"
	"polish/tokens"
	"strconv"
)

func PolishNotation() {
	var list string
	err := tokens.ReadString(&list)
	if err == nil {
		var arr []tokens.Token
		err = tokens.Tokenize(list, &arr)
		if err == nil {
			err = tokens.CheckToken(&arr)
			if err != nil {
				fmt.Printf("Error: %s", err)
			} else {
				res := Algorithm(&arr)
				err = Calculate(&res)
				if err != nil {
					fmt.Printf("Error: %s", err)
				}
			}
		} else {
			fmt.Printf("Error: %s", err)
		}
	} else {
		fmt.Printf("Error: %s", err)
	}
}

func Algorithm(arr *[]tokens.Token) []tokens.Token {
	var st_op stack.Stack
	st_op.Init()
	var str []tokens.Token
	for i := range *arr {
		if (*arr)[i].Tok == tokens.Var || (*arr)[i].Tok == tokens.Num {
			str = append(str, (*arr)[i])
		} else if (*arr)[i].Str == "(" {
			st_op.Push((*arr)[i])
		} else if (*arr)[i].Str == ")" {
			for st_op.Peek().Str != "(" && st_op.Top != -1 {
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
	for st_op.Top != 0 {
		str = append(str, st_op.Peek())
		st_op.Pop()
	}
	return str
}

func Calculate(arr *[]tokens.Token) error {
	var st_num stack.Stack
	st_num.Init()
	for i := range *arr {
		if (*arr)[i].Tok == tokens.Var || (*arr)[i].Tok == tokens.Num {
			st_num.Push((*arr)[i])
		} else if (*arr)[i].Tok != tokens.Func {
			num2, err := strconv.ParseFloat(st_num.Peek().Str, 64)
			if err != nil {
				return err
			}
			st_num.Pop()
			num1, err := strconv.ParseFloat(st_num.Peek().Str, 64)
			if err != nil {
				return err
			}
			st_num.Pop()
			switch (*arr)[i].Str {
			case "+":
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(num1+num2, 'f', -1, 64), Tok: tokens.Num})
			case "-":
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(num1-num2, 'f', -1, 64), Tok: tokens.Num})
			case "*":
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(num1*num2, 'f', -1, 64), Tok: tokens.Num})
			case "/":
				if num2 == 0 {
					return errors.New("деление на 0 запрещено")
				}
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(num1/num2, 'f', -1, 64), Tok: tokens.Num})
			}
		} else {
			num1, err := strconv.ParseFloat(st_num.Peek().Str, 64)
			if err != nil {
				return err
			}
			st_num.Pop()
			switch (*arr)[i].Str {
			case "sin":
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(math.Sin(num1), 'f', -1, 64)})
			case "cos":
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(math.Cos(num1), 'f', -1, 64)})
			case "ln":
				if num1 <= 0 {
					return errors.New("операция ln: число не соответствует допустимым значениям")
				}
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(math.Log(num1), 'f', -1, 64)})
			case "sqrt":
				if num1 < 0 {
					return errors.New("операция ln: число не соответствует допустимым значениям")
				}
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(math.Sqrt(num1), 'f', -1, 64)})
			case "ctg":
				// проверка на хз что?
				cos := math.Cos(num1)
				sin := math.Sin(num1)
				if sin == 0 {
					return errors.New("операция ctg: sin не должен быть равен 0")
				}
				st_num.Push(tokens.Token{Str: strconv.FormatFloat(cos/sin, 'f', -1, 64)})
			}
		}
	}
	fmt.Println(st_num)
	return nil
}
