package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func multipleResults(x, y string) (string, string) {
	return y, x
}

func main() {
	// Function basic
	fmt.Println(sum(42, 23))

	// Function multiple results
	a, b := multipleResults("world", "hello")
	fmt.Println(a, b)
}