package main

import (
	"math/rand"
)

/*
switch is a conditional statement that evaluates an expression and compares,
it against a list of possible matches and executes the corresponding block of code.
*/

/*
It is an idiomatic way of replacing complex if else clauses.
Cases are evaluated from top to bottom.
*/

func main() {
	// finger := 4
	// A switch can include an optional statement that is executed before the expression is evaluated.
	// The scope of finger in this case is limited to the switch block.
	switch finger := 4; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Invalid")
	}

	// Multiple Expressions in case, it is possible to include multiple expressions in a case by separating them with comma.
	switch vowel := a; vowel {
	case "a", "e", "i", "o", "u":
		fmt.Println("it's a vowel")
	default:
		fmt.Println("it's not a vowel")
	}

	// Expression less switch, the expression in a switch is optional and it can be omitted.
	// If the expression is omitted, the switch is considered to be switch true.
	num := 57
	switch {
	case num > 50:
		fmt.Printf("%v is greater than %v\n",num, 50)
	}

	// Fallthrough
	// The control comes out of the switch statement immediately after a case is executed.
	// fallthrough should be the last statement in a case.
	// If it is present somewhere in the middle, the compiler will complain that fallthrough statement out of place.

	switch num := 75; { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}

	// When fallthrough is encountered the control moves to the first statement of the next case

	switch x := 10; {
	case x < 100:
		fmt.Println("x is less than 100")
		fallthrough
	case x > 100: // Fallthrough happens even when the case evaluates to false
		fmt.Println("x is not greater than 100")
	}

	// fallthrough cannot be used in the last case of a switch since there are no more cases to fallthrough.
	// If fallthrough is present in the last case, it will result in the following compilation error.
	// cannot fallthrough final case in switch

	// Breaking switch
	switch num := 20; {
	case num > 30:
		break
	}

	// Lables in switch
	someLabel:
		for {
			switch i := rand.Intn(100); {
			case i%2 == 0:
				fmt.Printf("Generated even number %d", i)
				break someLabel
			}
		}
	// Please note that if the break statement is used without the label, the switch statement will only be broken and the loop will continue running. 
	// So labeling the loop and using it in the break statement inside the switch is necessary to break the outer for loop.
}
