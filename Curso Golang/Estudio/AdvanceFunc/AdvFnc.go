package main

import "fmt"

func suma(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	fmt.Printf("%s, %d\n", name, total)
	return total
}
func main() {
	fmt.Println(suma("Hola", 12, 1, 24, 3))
	fmt.Println(suma("Sandra", 12, 1, 24, 13))
	fmt.Println(suma("Sebas", 12, 1, 24, 10))
}
