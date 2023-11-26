/*
Given an integer N, count the number of digits
Input = 12345
Output = 5
*/
package main

import "fmt"

func count(input int) {
	count := 0
	for input != 0 {
		input = input / 10
		count++
	}
	fmt.Print(count)
}

func main() {
	var input int = 123556666
	count(input)
}
