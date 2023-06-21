package main

import (
	"fmt"
	"testing"
)

func TestFunWithSlices(t *testing.T) {

	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mySlice := myArray[1:5]

	if len(mySlice) != 4 {
		t.Error("it is not 5")
		t.Log(mySlice)
	}
}

func TestVal(t *testing.T) {

	if 1 == 1 {
		t.Error("olmadıı")
	}
}

type SomeType struct {
	i *int
	y *string
}

func changeMe(i *int) {
	*i = 27
}

func changeMeWithoutRef(i int) {
	i = 27
}

func TestMap(t *testing.T) {

	xi := 6

	x := 5

	xp := &x

	fmt.Printf("%p \n", xp)

	fmt.Printf("%p \n", &xi)

	changeMe(&x)

	fmt.Println(x)

	myMap := make(map[string]interface{})

	myMap["emre"] = "Kantar"
	myMap["emre2"] = true
	myMap["emre3"] = 12.4
	myMap["test"] = nil

	for key, i := range myMap {
		fmt.Printf("index is: %d key is: %s value is: %s \n", i, key, myMap[key])

		switch myMap[key].(type) {
		case string:
			fmt.Println("type is string")
		case int:
			fmt.Println("type is int")
		case float32:
			fmt.Println("type is float32")
		case float64:
			fmt.Println("type is float64")
		case bool:
			fmt.Println("type is boolean")
		default:
			fmt.Println("type is default")
		}
	}

	t.Log("success")
}
