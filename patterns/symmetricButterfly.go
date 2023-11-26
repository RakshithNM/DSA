/*
Symmetric Butterfly
====================
Input = 5
*        *
**      **
***    ***
****  ****
**********
****  ****
***    ***
**      **
*        *
====================
*/
package main

import "fmt"

func symmetricButterfly2(input int) {
	spaces := 2*input - 2
	for i := 1; i <= 2*input-1; i++ {
		stars := i
		if i > input {
			stars = 2*input - i
		}

		for j := 1; j <= stars; j++ {
			fmt.Print("*")
		}

		for j := 1; j <= spaces; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= stars; j++ {
			fmt.Print("*")
		}

		fmt.Println()
		if i >= input {
			spaces = spaces + 2
		} else {
			spaces = spaces - 2
		}
	}
}

func symmetricButterfly1(input int) {
	for i := 1; i < 2*input; i++ {
		limitJ := i
		limitI := input - i
		if i > input {
			limitJ = input - (i % input)
			limitI = input - limitJ
		}

		for j := 1; j <= limitJ; j++ {
			fmt.Print("*")
		}

		for j := 1; j <= 2*limitI; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= limitJ; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var input int = 5
	symmetricButterfly1(input)
	fmt.Println()
	symmetricButterfly2(input)
}
