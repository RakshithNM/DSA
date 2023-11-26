/*
Rectangle
======================
input = 5
* * * * *
* * * * *
* * * * *
* * * * *
* * * * *
======================
*/
package main

import (
	"fmt"
)

func rectangular(input int) {
	for i := 0; i < input; i++ {
		for j := 0; j < input; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("3 Rectangular")
	var input1 int = 3
	rectangular(input1)

	fmt.Println("6 Rectangular")
	var input2 int = 6
	rectangular(input2)
}
