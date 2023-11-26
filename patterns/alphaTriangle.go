/*
Alpha Triangle
==================
Input = 5
E
D E
C D E
B C D E
A B C D E
===================
*/

package main

import "fmt"

func alphaTrianlge(input int) {
	var startingASCII int = 70
	for i := 1; i <= input; i++ {
		startingASCII = startingASCII - i
		for j := 1; j <= i; j++ {
			fmt.Print(string(rune(startingASCII)))
			startingASCII = startingASCII + 1
		}
		fmt.Println()
	}

}

func main() {
	var input int = 5
	alphaTrianlge(input)
}
