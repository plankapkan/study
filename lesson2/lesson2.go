package main

import "fmt"

func main() {
	defer println("deferred print")
	return2, i := withReturn2()
	fmt.Println(return2, i)
	fmt.Println(withReturn())

	num := i
	if isEven(num) == 0 { // test if
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	switch isEven(num) { // test switch
	case 0:
		fmt.Printf("%d is even\n", num)
	case 1:
		fmt.Printf("%d is odd\n", num)
	}

}

func isEven(num int) int {
	return num % 2
}

func withReturn() (str string, num int) {
	str = "length"
	num = 42
	return str, num
}
func withReturn2() (string, int) {
	str := "width"
	num := 24
	return str, num
}
