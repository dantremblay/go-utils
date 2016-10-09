package readinput

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(label string) string {
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s : ", label)

	input, err := consolereader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	count := len(input)

	return input[0 : count-1]
}
