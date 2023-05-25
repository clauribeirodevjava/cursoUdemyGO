package main

import (
	"fmt"
)

func main() {
	slice := []int{5, 1, 2}
	newInts := reverse(slice)
	slice2 := []string{"a", "e", "f", "b"}
	newString := reverse(slice2)
	fmt.Print(newInts)
	fmt.Print(newString)
}

type constraintCustom interface {
	int | string
}

func reverse[T constraintCustom](slice []T) []T {
	newints := make([]T, len(slice))
	newIntsLen := len(slice) - 1
	for i := 0; i < len(slice); i++ {
		newints[newIntsLen] = slice[i]
		newIntsLen--
	}

	return newints
}
