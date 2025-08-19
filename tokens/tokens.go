package tokens

import (
	"bufio"
	"os"
)

func ReadString(str *string) int {
	flag := 1
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		flag = 0
	} else {
		*str = line
	}
	return flag
}
