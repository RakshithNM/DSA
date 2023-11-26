/*
Example 1:
Input:153
Output: Yes, it is an Armstrong Number
Explanation: 1^3 + 5^3 + 3^3 = 153

Example 2:
Input:170
Output: No, it is not an Armstrong Number
Explanation: 1^3 + 7^3 + 0^3 != 170
*/
package main

import (
	"fmt"
	"math"
)

func digits(input int) float64 {
	count := float64(0)
	for input != 0 {
		input = input / 10
		count++
	}
	return count
}

func armstrong(input int) {
	var numDigits float64 = digits(input)
	var temp float64 = float64(input)
	var remainder int = input
	sum := float64(0)
	for temp != 0 {
		remainder = int(temp) % 10
		sum = sum + math.Pow(float64(remainder), numDigits)
		temp = temp / 10
	}
	if input == int(sum) {
		fmt.Print("It is an armstrong num")
	} else {
		fmt.Print("It is not an armstrong num")
	}
}

func main() {
	var input1 int = 153
	armstrong(input1)
	var input2 int = 170
	armstrong(input2)
}
