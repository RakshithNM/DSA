/*
Find the sum of first N numbers
*/
package main

import "fmt"

var sum int = 0

func sumOfN(input int) int {
	if input == 0 {
		return 0
	}

	oneLess := input - 1
	return input + sumOfN(oneLess)
}

func main() {
	var input int = 10
	output := sumOfN(input)
	fmt.Println(output)
}
