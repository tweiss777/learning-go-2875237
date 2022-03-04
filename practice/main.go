// this tells the compiler this is our main entry point
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumTwo(x int, y int) int {
	var sum int = x + y
	return sum

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the first number")
	number1, _ := reader.ReadString('\n')
	fmt.Println("Enter the second number")
	number2, _ := reader.ReadString('\n')

	conversion1, err1 := strconv.ParseInt(strings.TrimSpace(number1), 0, 64)
	conversion2, err2 := strconv.ParseInt(strings.TrimSpace(number2), 0, 64)

	if err1 != nil && err2 != nil {
		sum = sumTwo(number1, number2)
	}

}

