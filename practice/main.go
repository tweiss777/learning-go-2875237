package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// an example program that will calculate the sum of 2 numbers
func main() {

	// step 1 is to take input from the user
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter the first number")
	var input1, _ = reader.ReadString('\n')
	fmt.Println("Enter the second number")
	var input2, _ = reader.ReadString('\n')

	// step 2 is to conver input of user from string to int
	// not working for some reason
	intVar1, err1 := strconv.Atoi(strings.TrimSpace(input1))
	intVar2, err2 := strconv.Atoi(strings.TrimSpace(input2))

	// step 3 is to sum the two numbers up

	if err1 == nil && err2 == nil {
		var sum int = intVar1 + intVar2
		fmt.Println("the sum of the two numbers is %s", sum)

	} else {
		fmt.Println("Something went wrong....")
		fmt.Println(err1)
		fmt.Println(err2)
	}

}
