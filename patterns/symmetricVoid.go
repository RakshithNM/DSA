/*
Symmetric void
================
Input = 5
**********
****  ****
***    ***
**      **
*        *
*        *
**      **
***    ***
****  ****
**********
=================
*/
package main

import "fmt"

func symmetricVoid1(input int) {
	for i := 1; i <= 2*input; i++ {
		var starLimit = input - i + 1
		if i > input {
			starLimit = i % input
		}
		if i == 2*input {
			starLimit = input
		}
		for j := 1; j <= starLimit; j++ {
			fmt.Print("*")
		}

		var spaceLimit = i - 1
		if i > input {
			spaceLimit = input - (i % input)
		}
		if i == 2*input {
			spaceLimit = 0
		}

		for j := 1; j <= spaceLimit; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= spaceLimit; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= starLimit; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*================
Input = 5
**********
****  ****
***    ***
**      **
*        *
*        *
**      **
***    ***
****  ****
**********
=================
*/

func symmetricVoid2(input int) {
	initS := 0
	for i := 1; i <= input; i++ {
		for j := 1; j <= input-i+1; j++ {
			fmt.Print("*")
		}

		for j := 1; j <= initS; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= input-i+1; j++ {
			fmt.Print("*")
		}

		initS += 2
		fmt.Println()
	}

	initS = 8
	for i := 1; i <= input; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		for j := 1; j <= initS; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		initS -= 2
		fmt.Println()
	}
}

func main() {
	var input int = 5
	symmetricVoid1(input)
	fmt.Println()
	symmetricVoid2(input)
}
