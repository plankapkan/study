package main

import "fmt"

type Point struct {
	X int
	Y int
	S string
}

func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

func draw() {
	p1 := Point{
		X: 2,
		Y: 2,
	}
	p2 := Point{
		X: 4,
	}
	fmt.Println(p1)
	fmt.Println(p1.X)
	p1.X = p1.Y + 1
	fmt.Println(p1.X)
	fmt.Println(p2)

	g := &p1        // var addr in mem
	fmt.Println(g)  //pointer
	fmt.Println(*g) //struct
	fmt.Println(&g) //adr in mem
}

func main() {
	a := 42 // variable
	b := &a // var addr in memory
	fmt.Println(b)
	c := *b / 2 // addr in memory -> value
	fmt.Println(c)
	draw()
	p1 := Point{
		X: 5,
		Y: 7,
		S: "string",
	}
	p1.method()
	p2 := &p1
	p2.method()

	// arrays
	var array [3]string //fixed length
	array[0] = "first"
	array[1] = "second"
	fmt.Println(array)

	numbers := [...]int{10, 20, 30, 40} //dynamic length
	fmt.Println(numbers)

	//slices no fixed length
	letters := []string{"a", "b", "c"}
	letters2 := []string{
		"d",
		"e",
		"f",
	}
	letters2 = append(letters2, "j")
	fmt.Println(letters, letters2)

	EmptySlice := make([]string, 3) //empty slice
	EmptySlice[0] = "f"
	EmptySlice[1] = "g"
	EmptySlice[2] = "h"
	fmt.Println(EmptySlice)
	fmt.Println(fmt.Sprintf("length   %d", len(EmptySlice))) //length
	fmt.Println(fmt.Sprintf("capacity %d", cap(EmptySlice))) //capacity
	EmptySlice = append(EmptySlice, "i", "j")
	fmt.Println(EmptySlice)
	fmt.Println(fmt.Sprintf("length   %d", len(EmptySlice))) //length
	fmt.Println(fmt.Sprintf("capacity %d", cap(EmptySlice))) //capacity
	EmptySlice = append(EmptySlice, "k", "l")
	fmt.Println(EmptySlice)
	fmt.Println(fmt.Sprintf("length   %d", len(EmptySlice))) //length
	fmt.Println(fmt.Sprintf("capacity %d", cap(EmptySlice))) //capacity

	animalsArray := [3]string{"dog", "cat", "pig"}
	var part1 []string = animalsArray[0:2]
	part2 := animalsArray[1:3]
	fmt.Println(part1)
	fmt.Println(part2)
	part2[0] = "frog"
	fmt.Println(animalsArray)
	fmt.Println(part2)
}
