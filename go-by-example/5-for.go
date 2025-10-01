package main

import (
	f "fmt"
)

func main(){
	f.Println("hi, from 5-For")

	i := 1
	for i <= 3 {
		f.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		f.Println("range", j)
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
	}

	for x:=1; x<=5; x++ {
		if x==4 {
			continue
		}

		f.Println("x =", x)
	}
}


