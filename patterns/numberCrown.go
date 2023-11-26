/*
Number Crown
===================
input = 5
1        1
12      21
123    321
1234  4321
1234554321
===================
*/
package main

import "fmt"

func numberCrown1(input int) {
	for i := 1; i <= input; i++ {
		for j := 1; j <= i; j++ {
			var toPrint = j
			fmt.Print(toPrint)
			toPrint = toPrint + 1
		}

		for j := 1; j <= 2*input-2*i; j++ {
			fmt.Print(" ")
		}

		var toPrint = i
		for j := 1; j <= i; j++ {
			fmt.Print(toPrint)
			toPrint = toPrint - 1
		}
		fmt.Println()
	}
}

func numberCrown2(input int) {
	for i := 1; i <= input; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}

		for j := 1; j <= 2*input-2*i; j++ {
			fmt.Print(" ")
		}

		for j := i; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func main() {
	var input int = 5
	numberCrown1(input)
	numberCrown2(input)
}
