package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}
func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

type Product2[T uint | string] struct {
	id          T
	description string
	price       float32
}

func main() {
	product1 := Product2[uint]{
		1,
		"Buso",
		33.200,
	}
	product2 := Product2[string]{
		"HJZ-OEU92-MDJ",
		"Buso",
		33.200,
	}
	fmt.Println(product1, product2)
}
