/*
Right Angled Number Pyramid One
===============================
input = 5
1
2 2
3 3 3
4 4 4 4
5 5 5 5 5
===============================
*/
package main

import "fmt"

func rightAngledNumberPyramid(input int) {
	for i := 0; i < input; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(i+1, " ")
		}
		fmt.Println()
	}
}

func main() {
	var input int = 5
	rightAngledNumberPyramid(input)
}
