package main

import "fmt"

func maps() {
	dictionary := make(map[string]int)

	dictionary["key1"] = 7
	dictionary["key2"] = 13

	fmt.Println("map one:", dictionary)

	valueOne := dictionary["key1"]
	fmt.Println("value one:", valueOne)

	fmt.Println("length of dictionary", len(dictionary))

	delete(dictionary, "key2")
	fmt.Println("map:", dictionary)

	_, keyTwoPresent := dictionary["key2"]
	fmt.Println("key two present:", keyTwoPresent)

	declaredMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("declared map:", declaredMap)
}
