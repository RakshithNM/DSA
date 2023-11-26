/*
Increasing Letter
==================
Input = 5
A
A B
A B C
A B C D
A B C D E
*/

package main

import "fmt"

func increasingLetter(input int) {
	var startingASCII int = 64
	for i := 1; i <= input; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(string(rune(startingASCII+j)), " ")
		}
		fmt.Println()
	}
}

func main() {
	var input int = 5
	increasingLetter(input)
}
