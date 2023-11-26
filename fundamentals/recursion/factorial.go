/*
Given X, print its factorial
X! = X * (X-1) * (X-2) * .... * 1
*/
package main

import "fmt"

func factorial(input int) int {
	if input == 1 {
		return 1
	}

	oneLess := input - 1
	return input * factorial(oneLess)
}

func main() {
	var input int = 5
	output := factorial(input)

	fmt.Println(output)
}
