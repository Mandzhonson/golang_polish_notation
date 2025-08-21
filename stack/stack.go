package stack

import (
	"polish/tokens"
)

type Stack struct {
	Arr []tokens.Token
	Top int
}

func (st *Stack) Init() {
	st.Top = 0
	st.Arr = []tokens.Token{}
}

func (st *Stack) Push(el tokens.Token) {
	st.Arr = append(st.Arr, el)
	st.Top++
}

func (st *Stack) Pop() {
	if st.Top-1 >= 0 {
		st.Arr = st.Arr[:st.Top-1]
		st.Top--
	}
}

func (st *Stack) Peek() tokens.Token {
	if st.Top-1 >= 0 {
		return st.Arr[st.Top-1]
	}
	return tokens.Token{Str: "", Tok: 0}
}
