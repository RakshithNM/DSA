package main

// import fmt package
import "fmt"

func main() {

	fmt.Print("Hello, \n")
	fmt.Print("World!")

	if 12%2 == 0 {
		fmt.Print("even")
	}

	i := 3
	switch i {
	case 3:
		fmt.Print("three")
	case 4:
		fmt.Print("four")
	}

}
