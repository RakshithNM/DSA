/*
Inverted Star Pyramid
============

	 *********
		*******
		 *****
			***
			 *

============
*/
package main

import "fmt"

func invertedStarPyramid(input int) {
	for i := 0; i < input; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*input-(2*i)-1; j++ {
			fmt.Print("*")
		}
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func main() {
	var input int = 10
	invertedStarPyramid(input)
}
