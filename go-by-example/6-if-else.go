package main

import (
	f "fmt"
)

func main(){
	f.Println("Hi from 6!")

	if 7%2 == 0 {
		f.Println("7 is even!")
	} else {
		f.Println("7 is odd.")
	}

	if 8%2 == 0 || 7%2 == 0 {
		f.Println("Either 7 or 8 is even.")
	}

	if x:=9; x<9 {
		f.Println("x is less than 9")
	} else {
		f.Println("x is >= 9")
	}

	whatAmI := func(i interface{}){
		switch x := i.(type) {
		case bool:
			f.Println("I'm a bool")
		case int:
			f.Println("I'm an int")
		default:
			f.Printf("the type is %T\n", x)
		}
	}

	whatAmI(true)
	whatAmI(12)
	whatAmI(1.2)
}
