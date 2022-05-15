package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	// arr := []string{"a", "b", "c"}
	// for index, element := range arr { // index + element
	// 	fmt.Println(index)
	// 	fmt.Println(element)
	// }
	// for _, element := range arr { // _ to skip index
	// 	fmt.Println(element)
	// }
	pointMap := map[string]Point{
		"b": {
			X: 11,
			Y: 13,
		},
	}
	otherMap := make(map[int]Point)
	// var oneMoreOtherMap map[string]Point
	pointMap["a"] = Point{
		X: 0,
		Y: 1,
	}
	// fmt.Println(pointMap)
	// fmt.Println(pointMap["a"])
	otherMap[1] = Point{3, 4}
	// fmt.Println(otherMap)
	// fmt.Println(otherMap[1])

	key := "c"
	value, isok := pointMap[key]
	if isok {
		fmt.Printf("key %s exists\n", key)
		fmt.Printf("value is %d", value)
	} else {
		fmt.Printf("key %s not exist", key)
	}

}
