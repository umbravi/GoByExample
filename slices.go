package main

import "fmt"

func slices() {
	slice := make([]string, 3)
	fmt.Println("emptySlice:", slice)

	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Println("set:", slice)
	fmt.Println("get:", slice[2])
	fmt.Println("length:", len(slice))

	slice = append(slice, "d")
	slice = append(slice, "e", "f")
	fmt.Println("appended:", slice)

	copyOfSlice := make([]string, len(slice))
	copy(copyOfSlice, slice)
	fmt.Println("copy of slice:", copyOfSlice)

	cut := slice[2:5]
	fmt.Println("cut one of slice:", cut)

	cut2 := slice[:5]
	fmt.Println("cut two of slice", cut2)

	cut3 := slice[2:]
	fmt.Println("cut three of slice", cut3)

	declaredSlice := []string{"g", "h", "i"}
	fmt.Println("declaredSlice:", declaredSlice)

	twodimensionalSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twodimensionalSlice[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twodimensionalSlice[i][j] = i + j
		}
	}
	fmt.Println("two dimensional slice:", twodimensionalSlice)
}
