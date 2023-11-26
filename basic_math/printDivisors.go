/*
Print all divisors of a number
Input = 36
Output = 1 2 3 4 6 9 12 18 36
*/

package main

import (
	"fmt"
	"math"
)

func printDivisors1(input int) {
	for i := 1; i <= input; i++ {
		if input%i == 0 {
			fmt.Print(i, " ")
		}
	}
}

func printDivisors2(input int) {
	sqrt := int(math.Sqrt(float64(input)))
	for i := 1; i <= sqrt; i++ {
		if input%i == 0 {
			fmt.Print(i, " ")
			if i != sqrt {
				fmt.Print(input/i, " ")
			}
		}
	}
}

func main() {
	var input int = 36
	printDivisors1(input)
	fmt.Println()
	printDivisors2(input)
}
