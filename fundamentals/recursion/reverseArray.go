/*
Given an array, reverse it using iterative and recursive approach
*/
package main

import (
	"fmt"
)

func reverseIterative(input []int) []int {
	temp := input
	tempCopy := make([]int, len(input))
	copy(tempCopy, input)
	for i := len(input) - 1; i >= 0; i-- {
		temp[len(input)-1-i] = tempCopy[i]
	}
	return temp
}

func reverseIterativePointer(input []int) []int {
	left := 0
	right := len(input) - 1
	for left < len(input)/2 {
		input[left], input[right] = input[right], input[left]
		left++
		right--
	}
	return input
}

func reverseRecursive(input []int, start int, end int) []int {
	if start > end {
		return input
	}
	input[start], input[end] = input[end], input[start]
	fmt.Println(input)
	return reverseRecursive(input, start+1, end-1)
}

func main() {
	input1 := []int{1, 2, 3, 4, 5, 6}
	output1 := reverseIterative(input1)
	fmt.Println(output1)

	input2 := []int{3, 4, 5, 6, 7, 8, 9}
	output2 := reverseIterativePointer(input2)
	fmt.Println(output2)

	input3 := []int{5, 6, 7, 8, 9, 10}
	output3 := reverseRecursive(input3, 0, len(input3)-1)
	fmt.Println(output3)
}
