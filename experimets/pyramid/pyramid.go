package main

import "fmt"

func main() {
	rows := 5
	for y := 1; y <= rows; y++ {
		for x := 1; x <= y; x++ {
			fmt.Print(x)
		}
		fmt.Println()
	}
}
