/*
Inverted right pyramid
======================
input = 5
* * * * *
* * * *
* * *
* *
*
======================
*/
package main

import "fmt"

func invertedRightPyramid(input int) {
	for i := 0; i < input; i++ {
		for j := 0; j < input-i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	var input1 int = 5
	invertedRightPyramid(input1)
}
