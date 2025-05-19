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
}
