package main

import "fmt"

func main() {
	// Correct variable declaration
	var myAge = 17
	
	// Conditional check without unnecessary parentheses
	if myAge == 5 {
		fmt.Println("You're too young")
	}

	if myAge == 17 {
		fmt.Println("So sweet")
	}
	
	// Conditional check for a range of values
	if myAge >= 17 && myAge <= 30 {
		fmt.Println("My age is between 17 and 30")
	}
	
	// Proper usage of scoped variable declaration within an if statement
	if dadAge := 9; dadAge < myAge {
		fmt.Println(dadAge)
	}
}
var myAge = 17
if myAge == 5 {
    fmt.Println("You're too young")
}

if myAge == 17 {
    fmt.Println("So sweet")
}

if myAge >= 17 && myAge <= 30 {
    fmt.Println("My age is between 17 and 30")
}
if dadAge := 9; dadAge < myAge {
    fmt.Println(dadAge)
}
