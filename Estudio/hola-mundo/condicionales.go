package main

import (
	"fmt"
	"runtime"
)

func main() {
	// var v int64 = 3
	// if v == 1 {
	// 	fmt.Println("True")
	// } else if v == 2 {
	// 	fmt.Println("False")
	// } else {
	// 	fmt.Println("Invalid")
	// }
	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("El OS es Windows")
	case "linux":
		fmt.Println("El OS Linux")
	case "mac":
		fmt.Println("El OS Mac")
	}

}
