package main

import "fmt"

func split_without_named_returns(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y

}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
	fmt.Println(split_without_named_returns(17))
}
