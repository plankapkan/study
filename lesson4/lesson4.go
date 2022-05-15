package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Point struct {
	X int //`mapstructure:"xx"`
	Y int `mapstructure:"yy"`
}

func (p Point) method() {
	fmt.Println("coordinates are", "x", p.X, "y", p.Y)
}

func main() {
	pointMap := map[string]int{
		"x":  11,
		"yy": 14,
	}
	pointStruct := Point{}
	mapstructure.Decode(pointMap, &pointStruct)
	fmt.Println(pointStruct)

	for k, v := range pointMap {
		fmt.Println(k)
		fmt.Println(v)
	}
}
