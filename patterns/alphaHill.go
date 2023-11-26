/*
Alpha Hill
================
Input = 4

	      A
			 ABA
			ABCBA
		 ABCDCBA

================
*/
package main

import "fmt"

func alphaHill(input int) {
	for i := 1; i <= input; i++ {
		var startingASCII int = 64
		for j := 1; j <= input-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= (2*i)-1; j++ {
			if j > (2*i)/2 {
				startingASCII = startingASCII - 1
			} else {
				startingASCII = startingASCII + 1
			}
			fmt.Print(string(rune(startingASCII)))
		}
		for j := 1; j <= input-i; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func main() {
	var input int = 10
	alphaHill(input)
}
