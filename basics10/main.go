package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)

	var n float32 = 3.0

	fmt.Printf("n: %#0.00f, n: %T", n, n)
}
