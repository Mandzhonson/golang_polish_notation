package algorithm

import (
	"fmt"
	"polish/stack"
	"polish/tokens"
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

