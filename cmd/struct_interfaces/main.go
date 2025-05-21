package main

import (
	"fmt"
)

type house struct {
	rooms     int
	price     int
	ownerInfo owner
	// owner
	float32
}

type owner struct {
	name string
}

func (e house) helloFromOwner() string {
	return "Owner " + e.ownerInfo.name + " says hello to you!!"
}

type gasEngine struct {
	mpg     int
	gallons int
}

type electricEngine struct {
	mpkwh int
	kwh   int
}

func (e gasEngine) milesLeft() int {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() int {
	return e.kwh * e.mpkwh
}

// Interface
type engine interface {
	milesLeft() int
}

// func canMakeIt(e gasEngine, miles int) {
func canMakeIt(e engine, miles int) {
	if miles < e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	fmt.Println("Struct & Interface: ")

	// var myHouse house = house{price: 10, rooms: 20, ownerInfo: owner{"Tanvir"}}
	var myHouse house = house{price: 10, rooms: 20, ownerInfo: owner{"tanvir"}, float32: 11.12}
	// var myHouse house
	fmt.Println(myHouse)

	fmt.Println("myHouse.float32: ", myHouse.float32)

	fmt.Println(myHouse.helloFromOwner())

	var myEngine = struct {
		mph   int
		price int
	}{12, 10} // anonymous struct -> must have to declare and initialize at the same time/place and obviously not reuseable.

	fmt.Println(myEngine.mph, " -> ", myEngine.price)

	// -----------------------------
	var engGas gasEngine = gasEngine{25, 15}
	var engElec electricEngine = electricEngine{20, 10}

	canMakeIt(engGas, 5) // Generic function, works on both structure. Thanks to Interface!
	canMakeIt(engElec, 5)
	canMakeIt(engGas, 500)
	canMakeIt(engElec, 500)
}
