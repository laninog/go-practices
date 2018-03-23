package main

import (
	"fmt"
	"math/cmplx"
)

var (
	MaxInt uint64     = 1<<64 - 1
	i      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	MyConst   = 100
	YourConst = "You"
)

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "now"

func main() {
	// Implicit
	cPlusPlus, javaScript := true, false

	fmt.Println(x, y, z, c, python, java)

	fmt.Println(cPlusPlus, javaScript)

	const format = "%T(%v)\n"
	fmt.Printf(format, MaxInt, MaxInt)
	fmt.Printf(format, z, z)

	fmt.Println(MyConst, YourConst)
}
