/*
Half Diamond Star
=================
input = 5
*
* *
* * *
* * * *
* * * * *
* * * *
* * *
* *
*
=================
*/
package main

import "fmt"

func halfDiamondStar(input int) {
	for i := 1; i <= (2*input)-1; i++ {
		var stars int = i
		if i > input {
			stars = (2 * input) - i
		}
		for j := 1; j <= stars; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	var input int = 5
	halfDiamondStar(input)
}
