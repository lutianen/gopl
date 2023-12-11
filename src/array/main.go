package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [...]int{15: -1}
	arr1 := [16]int{15: -1}

	fmt.Println("It should be true: ", arr == arr1)

	modifyArray(arr)
	fmt.Println("It should be true: ", arr == arr1)

	modifyArrayWithPtr(&arr)
	fmt.Println("It should be false: ", arr == arr1)
}

func modifyArray(arr [16]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = int(math.Sin(float64(i)))
	}
}

func modifyArrayWithPtr(arr *[16]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = int(math.Sin(float64(i)))
	}
}
