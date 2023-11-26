/*
Print 1 to N using recursion
*/
package main

import "fmt"

func print1toN(a int, b int) {
	if a > b {
		return
	}
	fmt.Print(a, " ")
	a = a + 1
	print1toN(a, b)
}

func main() {
	print1toN(1, 10)
}
