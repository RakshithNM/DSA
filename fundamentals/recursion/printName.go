/*
Print name N times using recursive function call
*/

package main

import "fmt"

func printName(input int) {
	if input == 0 {
		return
	}
	fmt.Println("Rakshith")
	input = input - 1
	printName(input)
}

func main() {
	printName(5)
}
