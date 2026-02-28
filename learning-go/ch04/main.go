package main

import (
	"fmt"
	"math/rand"
)

func exercise1() []int {
	var radSlice []int
	for i := 1; i <= 100; i++ {
		radSlice = append(radSlice, rand.Intn(i))
	}
	return radSlice
}
func exercise2(radSlice []int) {
	for _, val := range radSlice {
		switch {
		case val%2 == 0 && val%3 == 0:
			fmt.Println("Six!")
		case val%2 == 0:
			fmt.Println("Two!")
		case val%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
func exercise3() {
	var total int
	for i := 0; i < 10; i++ {
		total = total + i
		fmt.Println(total)
	}
	fmt.Println(total, "outside")

}
func main() {
	// result := exercise1()
	// exercise2(result)
	exercise3()
}
