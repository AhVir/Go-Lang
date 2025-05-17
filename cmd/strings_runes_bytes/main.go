package main

import (
	"fmt"
	"strings"
)

func main() {
	var myStr string = "rèsume"
	fmt.Println(myStr)
	index_0 := myStr[1]
	fmt.Println(index_0) // it will print 114, GO usages utf-8 encoding to represent string, instead of ASCII

	fmt.Printf("Value: %v, Type: %T\n", index_0, index_0) // type: uint8

	for i, v := range myStr {
		fmt.Println(i, v)
	} // we will see that, index 2 is missing because range is actually decoding 2 bytes of char for that 'special e'

	// String in Go -> underline structure is, an array of bytes
	fmt.Printf("lenght(in bytes, not in characters) of the String myStr is %v\n", len(myStr))

	var myStr2 = []rune("rèsumè") // we're using Array of rune, not array of bytes
	myStr_0 := myStr2[1]
	fmt.Printf("Value: %v, Type: %T\n", myStr_0, myStr_0) // this time, it will print the whole 2 bytes for that e as we've used rune

	// String Building
	var strSlice = []string{"t", "a", "n"}
	concat := ""

	for i := range strSlice {
		concat += strSlice[i]
	}

	fmt.Println(concat)

	// concat[0] = 'a'   -> strings are immunitable in GO
	// so basically, concat += strSlice[i]  -> creating completely new string every time, and that's obviously inefficient

	// Solution for that:

	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	concat2 := strBuilder.String() // much faster
	fmt.Printf("concat2: %v\n", concat2)
}
