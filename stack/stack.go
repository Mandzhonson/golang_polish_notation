package stack

import "polish/tokens"

type Stack struct {
	Arr []tokens.Token
	Top int
}

func (st *Stack) Init() {
	st.Top = -1
}

func (st *Stack) Push(el *tokens.Token) {
	st.Top++
	st.Arr = append(st.Arr, *el)
}

func (st *Stack) Pop() {
	if st.Top-1 >= 0 {
		newSlice := make([]tokens.Token, st.Top-1)
		copy(newSlice, st.Arr[:st.Top-1])
		copy(st.Arr, newSlice)
		st.Top--
	} else {
		st.Arr = []tokens.Token{}
		st.Top = -1
	}
}

func (st *Stack) Peek() tokens.Token {
	if st.Top != -1 {
		return st.Arr[st.Top]
	}
	return tokens.Token{}
}
