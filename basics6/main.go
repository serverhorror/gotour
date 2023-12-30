package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("first", "second")
	fmt.Println(a, b)

	j, _ := swap("j param", "k param")
	fmt.Println(j)
}
