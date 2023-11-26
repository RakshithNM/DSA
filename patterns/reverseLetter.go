/*
Reverse Letter
================
A B C D E
A B C D
A B C
A B
A
================
*/
package main

import "fmt"

func reverseLetter(input int) {
	var startingASCII = 65
	for i := 0; i < input; i++ {
		for j := 0; j < input-i; j++ {
			fmt.Print(string(rune(startingASCII+j)), " ")
		}
		fmt.Println()
	}
}

func main() {
	var input int = 5
	reverseLetter(input)
}
