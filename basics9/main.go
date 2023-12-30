package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	fmt.Printf("i: %T, j: %T, c: %T, python: %T, java: %T\n",
		i, j, c, python, java)
}
