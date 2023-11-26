/*
Binary Number Triangle
======================
input = 5
1
0 1
1 0 1
0 1 0 1
1 0 1 0 1
=====================
*/
package main

import "fmt"

func binaryNumberTriangle1(input int) {
	for i := 1; i <= input; i++ {
		for j := 1; j <= i; j++ {
			if i%2 == 0 && j%2 == 0 {
				fmt.Print("1")
			} else if i%2 != 0 && j%2 == 0 {
				fmt.Print("0")
			} else if i%2 == 0 && j%2 != 0 {
				fmt.Print("0")
			} else {
				fmt.Print("1")
			}
		}
		fmt.Println()
	}
}

func binaryNumberTriangle2(input int) {
	var start int = 1
	for i := 1; i <= input; i++ {
		if i%2 == 0 {
			start = 0
		}

		for j := 1; j <= i; j++ {
			fmt.Print(start)
			start = 1 - start
		}
		fmt.Println()
	}
}

func main() {
	var input int = 6
	binaryNumberTriangle1(input)
	binaryNumberTriangle2(input)
}
