package main

import (
	"fmt"
	"math"
)

func isPrime(input int) bool {
	sqrt := int(math.Sqrt(float64(input)))
	remainder := input
	for i := 2; i < sqrt; i++ {
		remainder = input % i
		if remainder == 0 {
			return false
		}
	}
	return true
}

func main() {
	var input int = 5
	ans := isPrime(input)

	if input != 1 && ans == true {
		fmt.Print(input, " is a prime number")
	} else {
		fmt.Print(input, " is not a prime number")
	}
}
