/*
Alpha Ramp
===============
Input = 5
A
B B
C C C
D D D D
E E E E E
===============
*/
package main

import "fmt"

func alphaRamp(input int) {
	startingASCII := 65
	for i := 0; i < input; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(string(rune(startingASCII)), " ")
		}
		startingASCII = startingASCII + 1
		fmt.Println()
	}
}

func main() {
	var input int = 5
	alphaRamp(input)
}
