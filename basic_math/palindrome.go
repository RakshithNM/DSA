/*
Given an integer, check if it is a palindrome
=============================================
Input = 121
Output = true

Input = 123
Output = false
=============================================
*/

package main

import "fmt"

func reverse(input int) int {
	reverse := 0
	for input != 0 {
		digit := input % 10
		reverse = reverse*10 + digit
		input = input / 10
	}
	return reverse
}

func main() {
	var input1 int = 121
	reverse1 := reverse(input1)

	fmt.Println(input1 == reverse1)

	var input2 int = 223
	reverse2 := reverse(input2)

	fmt.Println(input2 == reverse2)
}
