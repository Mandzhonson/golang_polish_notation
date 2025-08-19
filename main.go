package main

import (
	"fmt"
	"polish/tokens"
)

func main() {
	var list string
	if tokens.ReadString(&list) == 1 {
		var arr []tokens.Token
		if tokens.Tokenize(list, &arr) == 1 {
			for i := range arr {
				fmt.Printf("%s\n", arr[i].Str)
			}
		} else {
			// error added
		}
	} else {
		// error add
	}
}
