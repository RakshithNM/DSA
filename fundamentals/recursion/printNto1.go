/*
Print N to 1 using recursion
*/
package main

import "fmt"

func printNto1(a int, b int) {
	if a < b {
		return
	}
	fmt.Print(a, " ")
	a = a - 1
	printNto1(a, 1)
}

func main() {
	printNto1(10, 1)
}
