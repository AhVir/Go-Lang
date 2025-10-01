package main

import (
	f "fmt"
	"math"
)

const s string = "constant"

func main(){
	p := f.Println

	p(s)

	const n = 5000000000000

	const d = 3e20 / n
	p(d)

	p("int64(d): ", int64(d))
	p("math.Sin: ", math.Sin(d))
}
