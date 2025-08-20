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
				Algorithm(&arr)
			}
		} else {
			fmt.Printf("Error: %s", err)
		}
	} else {
		fmt.Printf("Error: %s", err)
	}
}

func Algorithm(arr *[]tokens.Token) {
	var st_op stack.Stack
	st_op.Init()
	str := ""
	for i := range *arr {
		if (*arr)[i].Tok == tokens.Var || (*arr)[i].Tok == tokens.Num {
			str += (*arr)[i].Str + " " // not optimize just test
		} else if st_op.Top == -1 && (*arr)[i].Tok != tokens.Var && (*arr)[i].Tok != tokens.Num {
			st_op.Push(&(*arr)[i])
		} else {
			if (*arr)[i].Tok > st_op.Peek().Tok {
				st_op.Push(&(*arr)[i])
			} else {
				for (*arr)[i].Tok <= st_op.Peek().Tok {
					str += st_op.Peek().Str + " "
					st_op.Pop()
				}
				st_op.Push(&(*arr)[i])
			}
		}
	}
	for st_op.Top != -1 {
		str += st_op.Peek().Str + " "
		st_op.Pop()
	}
	fmt.Printf("%s", str)
}
