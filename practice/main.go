// this tells the compiler this is our main entry point
package main

import (
	"bufio"
	"fmt"
	"os"
)

func sumTwo(x int, y int) int {
	var sum int = x + y
	return sum

}

func main() {
	fmt.Printf("Welcome to the calculator\n")
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	input2, _ := reader.ReadString('\n')

}
