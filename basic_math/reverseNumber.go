/*
Given an integer, reverse it
Input = 12345
Output = 54321
*/
package main

import "fmt"

func reverseNumber(input int) {
	reverse := 0
	for input != 0 {
		digit := input % 10
		reverse = reverse*10 + digit
		input = input / 10
	}
	fmt.Print(reverse)
}

func main() {
	var input int = 123
	reverseNumber(input)
}
