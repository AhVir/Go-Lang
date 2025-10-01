package main

import (
	"fmt"
)

func main(){
	var a[5] int
	fmt.Println("a[] initially:", a)

	a[4] = 12
	fmt.Println("a[]:", a)

	b := [5] int{1, 3, 2, 9, 0}
	fmt.Println("b[]:", b)

	c := [...] int{1, 3, 2}
	fmt.Println("c[]:", c)

	twoD := [2][3] int{
		{1, 2, 3},
		{0, 1, 0},
	}
	fmt.Println("twoD[][]:", twoD)
	
	fmt.Println("Array: ")
	for x := range 2 {
		for y := range 3 {
			// fmt.Println(twoD[x][y])
			fmt.Printf("%d ", twoD[x][y])
		}
	}
	fmt.Println()
}
