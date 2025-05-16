package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello, Dude!")

	var intNum int = 23
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var str string = "Hey Tanvir"
	fmt.Println(str)

	var str2 string = `
	hi
	my 
	friend
	dear!`

	fmt.Println(str2)

	// Counting number of variables
	fmt.Println(len("len of str2: " + str2))
	fmt.Println("len of str2: ", len(str2))
	fmt.Println("actual len of str2: ", utf8.RuneCountInString(str2))

	var myBool bool = true
	fmt.Println(myBool)
	var myBool2 bool = false
	fmt.Println(myBool2)

	// rune
	var myRune rune = '😊'
	fmt.Printf("Emoji: %c\n", myRune)

	myFunc := hello() // -> bad practice
	// var myFunc = hello()  -> little bad practice
	// var myFunc string = hello() -> good practice to use if the value is returning from a function
	fmt.Println(myFunc)

	const pi float32 = 3.1416
	// const pi = 3.1416 -> both works with const, mentioning type or not
	fmt.Println(pi)

	test()
}

func hello() string {
	return "hello, DUDEEEEEEEEE!"
}
