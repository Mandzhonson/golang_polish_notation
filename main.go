package main

import (
	"fmt"
	"polish/tokens"
)

func main() {
	var list string
	err := tokens.ReadString(&list)
	if err == nil {
		var arr []tokens.Token
		err = tokens.Tokenize(list, &arr)
		if err == nil {
			for i := range arr {
				fmt.Printf("%s\n", arr[i].Str)
			}
		} else {
			fmt.Printf("Error: %s", err)
		}
	} else {
		fmt.Printf("Error: %s", err)
	}
}
