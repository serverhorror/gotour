package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	fmt.Printf("x: %v(%T), y: %v(%T), f: %v(%T), z: %v(%T)", x, x, y, y, f, f, z, z)
	fmt.Printf("x: %v(%T), y: %v(%T), z: %v(%T)", x, x, y, y, z, z)
}
