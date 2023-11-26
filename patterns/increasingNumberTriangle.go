/*
Increasing Number Triangle
==========================
Input = 5
1
2 3
4 5 6
7 8 9 10
11 12 13 14 15
=========================
*/
package main

import "fmt"

func increasingNumberTriangle(input int) {
	var toPrint int = 1
	for i := 1; i <= input; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(toPrint, " ")
			toPrint = toPrint + 1
		}
		fmt.Println()
	}
}

func main() {
	var input int = 5
	increasingNumberTriangle(input)
}
