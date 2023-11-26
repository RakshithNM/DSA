/*
Given a string, check if it is a palindrom
Input = "ABCBA"
Ouput " "Palindrom"

Input = "INDIA"
Output = "Not palindrome"
*/
package main

import "fmt"

func isPalindromeIterative(input string) bool {
	start := 0
	end := len(input) - 1

	for start < end {
		if input[start] != input[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func isPalindromeRecursive(input string, start int, end int) bool {
	if start > end {
		return input[start] == input[end]
	}

	if input[start] != input[end] {
		return false
	}
	start = start + 1
	end = end - 1
	return isPalindromeRecursive(input, start, end)
}

func main() {
	var input1 string = "ABCBA"
	output1 := isPalindromeIterative(input1)

	if output1 == true {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}

	var input2 string = "ABCB##"
	output2 := isPalindromeRecursive(input2, 0, len(input2)-1)

	if output2 == true {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}
}
