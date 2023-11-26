/*
Number
===============
Input = 4
4 4 4 4 4 4 4
4 3 3 3 3 3 4
4 3 2 2 2 3 4
4 3 2 1 2 3 4
4 3 2 2 2 3 4
4 3 3 3 3 3 4
4 4 4 4 4 4 4
===============
*/
package main

import "fmt"

func number(input int) {
	for i := 1; i < 2*input-1; i++ {
		toPrint := input
		for j := 1; j < 2*input-1; j++ {
			fmt.Print(toPrint)
		}
		toPrint = toPrint - 1
		fmt.Println()
	}
}

func main() {
	fmt.Println("vim-go")
}
