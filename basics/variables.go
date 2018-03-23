package main

import "fmt"

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "now"

func main() {
	// Implicit
	cPlusPlus, javaScript := true, false

	fmt.Println(x, y, z, c, python, java)

	fmt.Println(cPlusPlus, javaScript)
}