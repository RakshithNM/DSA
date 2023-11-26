/*
Right Angled Triangle
======================
input = 5
*
* *
* * *
* * * *
* * * * *
======================
*/
package main

import "fmt"

func rightAngledTriangle(input int) {
	for i := 0; i < input; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("5")
	var input1 int = 5
	rightAngledTriangle(input1)

	fmt.Println("10")
	var input2 int = 10
	rightAngledTriangle(input2)
}
