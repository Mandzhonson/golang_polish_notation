package main

import (
	"fmt"
	"polish/tokens"
)

func main() {
	var list string
	if tokens.ReadString(&list) == 1 {
		fmt.Printf("%s", list)
	}
}
