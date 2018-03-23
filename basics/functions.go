package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func multipleResults(x, y string) (string, string) {
	return y, x
}

func namedResults(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// Function basic
	fmt.Println(sum(42, 23))

	// Function multiple results
	a, b := multipleResults("world", "hello")
	fmt.Println(a, b)

	// Function named results
	fmt.Println(namedResults(27))
}