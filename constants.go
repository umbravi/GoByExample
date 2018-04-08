package main

import "fmt"
import "math"

const constantString string = "constant"

func constants() {
	fmt.Println(constantString)

	const constantNumber = 500000000

	const anotherConstantNumber = 3e20 / constantNumber
	fmt.Println(anotherConstantNumber)

	fmt.Println(int64(anotherConstantNumber))

	fmt.Println(math.Sin(constantNumber))
}
