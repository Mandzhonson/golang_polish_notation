package main

import (
	"fmt"
	"polish/tokens"
)

func main() {
	var list string
	if tokens.ReadString(&list) == 1 {
		var arr []tokens.Token
		tokens.Tokenize(list, &arr)
		for i := range arr {
			fmt.Printf("%s\n", arr[i].Str)
		}
	}
}
