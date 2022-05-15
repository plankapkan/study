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
}
