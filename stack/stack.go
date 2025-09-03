package stack

import (
	"polish/tokens"
)

type Stack struct {
	Arr []tokens.Token
}

func (st *Stack) Push(el tokens.Token) {
	st.Arr = append(st.Arr, el)
}

func (st *Stack) Pop() {
	if len(st.Arr)-1 >= 0 {
		st.Arr = st.Arr[:len(st.Arr)-1]
	}
}

func (st *Stack) Peek() tokens.Token {
	if len(st.Arr)-1 >= 0 {
		return st.Arr[len(st.Arr)-1]
	}
	return tokens.Token{Str: "", Tok: 0}
}

func (st *Stack) IsEmpty() bool {
	return len(st.Arr) == 0
}
