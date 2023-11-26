/*
Inverted numbered right pyramid
===============================
input = 5
1 2 3 4 5
1 2 3 4
1 2 3
1 2
1
===============================
*/
package main

import "fmt"

func invertedNumberedRightPyramid(input int) {
	for i := 0; i < input; i++ {
		for j := 0; j < input-i; j++ {
			fmt.Print(j+1, " ")
		}
		fmt.Println()
	}
}

func main() {
	var input int = 5
	invertedNumberedRightPyramid(input)
}
