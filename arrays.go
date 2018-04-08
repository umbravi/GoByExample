package main

import "fmt"

func arrays() {
	emptyArray := [5]int{}
	fmt.Println("empty:", emptyArray)

	emptyArray[4] = 100
	fmt.Println("set:", emptyArray)
	fmt.Println("get:", emptyArray[4])
	fmt.Println("length:", len(emptyArray))

	declaredArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declared:", declaredArray)

	twoDimensional := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			twoDimensional[i][j] = i + j
		}
	}
	fmt.Println("twoDimensional:", twoDimensional)
}
