package main

import "fmt"

// callback
func dosomething(callback func(int, int) int, s string) int {
	result := callback(5, 10)
	fmt.Println(s)
	return result
}

// zamikanie
func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(i int) int {
		sum += i
		return sum
	}
}

// structs + pointers
type Point struct {
	X, Y int
}

func (p Point) movePoint(x, y int) Point { // usless
	p.X += x
	p.Y += y
	return p
}
func (p *Point) movePointPtr(x, y int) { // good example
	p.X += x
	p.Y += y
}

func main() {

	// callback
	sumCallback := func(n1, n2 int) int {
		return n1 + n2
	}
	result := dosomething(sumCallback, "plus")
	fmt.Println(result)

	multiplyCallback := func(n1, n2 int) int {
		return n1 * n2
	}
	result = dosomething(multiplyCallback, "multiply")
	fmt.Println(result)

	// zamikanie
	newPrice := totalPrice(4)
	fmt.Println(newPrice(1))  // sum 4 +1 = 5
	fmt.Println(newPrice(1))  // 5+1
	fmt.Println(newPrice(10)) // 6+10

	// structs + pointers
	p := Point{1, 2}
	// fmt.Println(p)
	// fmt.Println(p.movePoint(1, 1))
	// fmt.Println(p)
	// p.movePointPtr(5, 5)
	// fmt.Println(p)
	v := &p
	fmt.Println(p)
	v.movePoint(1, 1)
	fmt.Println(p)
	v.movePointPtr(3, 5)
	fmt.Println(p)

}
