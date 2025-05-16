package main

import (
	"errors"
	"fmt"
)

func main() {
	var div, reminder, err = intDivision(10, 0)

	// if err != nil {
	// 	fmt.Println("ERRORRR: ", err.Error())
	// } else if reminder == 0 {
	// 	fmt.Printf("Div: %v\n", div)
	// } else {
	// 	fmt.Printf("div is %v, and reminder is %v\n", div, reminder)
	// }

	// If Else alternative Switch statements
	switch {
	case err != nil:
		fmt.Println(err.Error())
		// fmt.Println(err)  -> this would also output the same!
	case reminder == 0:
		fmt.Printf("Div: %v\n", div)
	default:
		fmt.Printf("Div: %v, and Reminder is %v\n", div, reminder)
	}

}

func intDivision(nume int, denume int) (int, int, error) {
	var err error

	if denume == 0 {
		err = errors.New("demume is 0, can't divide by 0")
		return 0, 0, err
	}
	division := nume / denume
	reminder := nume % denume

	return division, reminder, err
}
