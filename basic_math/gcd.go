/*
Find GCD of two numbers
=========================
Input1 = 4, Input2 = 8
Output = 4
=========================
*/
package main

import "fmt"

func gcd(input1 int, input2 int) {
	larger := input1
	smaller := input2
	if input1 < input2 {
		larger = input2
		smaller = input1
	}
	remainder := smaller
	for remainder != 0 {
		remainder = larger % smaller
		larger = smaller
		smaller = remainder
	}
	fmt.Println(larger)
}

func main() {
	gcd(4, 8)
}
