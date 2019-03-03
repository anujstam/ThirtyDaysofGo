package main

import (
	"strconv"
	"fmt"
	"bufio"
)

func makeint(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Bad String")
		return 0
	} else {
		return n
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(makeint(text))
}
