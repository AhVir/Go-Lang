package main

import (
	"fmt"
)

func main() {
	// Array
	fmt.Println("Array:")
	fmt.Println("------")

	var intArr [3]int32
	fmt.Println(intArr[0])
	intArr[1] = 99
	fmt.Println(intArr[1:3])
	fmt.Println(intArr)

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[2])

	// var intArr2 [3]int = [3]int{1, 2, 3}
	// intArr2 := [3]int{1, 2, 3} // -> little shortcut
	intArr2 := [...]int{1, 2, 3} // -> more little shortcut

	fmt.Println(intArr2)

	// Slices
	fmt.Println("\n\n\nSlices:")
	fmt.Println("-------")

	intSlice := []int{1, 2, 3, 4}
	fmt.Println(intSlice)
	fmt.Printf("Len of intSlice: %v, and Capacity of intSlice: %v\n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 9)
	fmt.Println(intSlice)
	fmt.Printf("Len of intSlice: %v, and Capacity of intSlice: %v\n", len(intSlice), cap(intSlice))
	// fmt.Println(intSlice[7])   ---> but this will still give index out of bound error

	intSlice2 := []int{11, 12, 13, 14}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Printf("intSlice -> Len: %v, and Cap: %v\n", len(intSlice), cap(intSlice))

	intSlice3 := []int{1}
	fmt.Println(intSlice3)
	intSlice3 = append(intSlice3, 2)
	fmt.Printf("intSlice -> Len: %v, and Cap: %v\n", len(intSlice3), cap(intSlice3))
	fmt.Println(intSlice3)
	intSlice3 = append(intSlice3, 2)
	fmt.Printf("intSlice -> Len: %v, and Cap: %v\n", len(intSlice3), cap(intSlice3))
	fmt.Println(intSlice3)

	// Another way to make a slice:
	intSlice4 := make([]int, 10, 13) // [data type, len, cap]
	fmt.Println(intSlice4)
	fmt.Printf("intSlice4 -> Len: %v, and Cap: %v\n", len(intSlice4), cap(intSlice4))

	for i, val := range intSlice4 {
		fmt.Printf("Index: %v -> %v\n", i, val)
	}

	// Map:
	fmt.Println("\n\n\nMap: ")
	fmt.Println("----")

	var myMap map[string]int = make(map[string]int)
	fmt.Println(myMap)

	var myMap2 = map[string]int{"tanvir": 22, "ahmed": 12}
	fmt.Println(myMap2)
	fmt.Println(myMap2["tanvir"])
	fmt.Println(myMap2["Sheikh"]) // although doesn't exist, shows default value of the return type

	var val, is_available = myMap2["Sheikh"]

	if !is_available {
		fmt.Printf("The Key %v, doesn't exist in myMap2 map.", "Sheikh")
	} else {
		fmt.Printf("Value: %v\n", val)
	}

	delete(myMap2, "Ahmed")
	fmt.Println(myMap2)

	// Loop
	fmt.Println("\n\n\nLoop:")
	fmt.Println("-----")

	// Using 'range'
	for name := range myMap2 {
		fmt.Println(name)
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Value: %v\n", name, age)
	}

	// And we can use for loop to replicate structures of traditional While & For Loop!
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 10; i >= 0; i-- {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	x := 11
	for {
		if x > 15 {
			break
		}

		fmt.Printf("%v ", x)
		// x = x + 1
		x++
	}
}
