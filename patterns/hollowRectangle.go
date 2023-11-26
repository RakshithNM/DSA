/*
Hollow Rectangle
=================
Input = 5
* * * * *
*       *
*       *
*       *
* * * * *
=================
*/

package main

import "fmt"

func hollowRectangle(input int) {
	for i := 1; i <= input; i++ {
		for j := 1; j <= input; j++ {
			if i == 1 || i == input || j == 1 || j == input {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	var input int = 5
	hollowRectangle(input)
}
