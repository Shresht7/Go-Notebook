package main

import "fmt"

//	Variadic functions can be called with any number of trailing arguments
//	fmt.Println is a common variadic function

//	Sum can take a variable number of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	//	Variadic functions can be called the usual way
	sum(1, 2)
	sum(1, 2, 3)

	//	If you already have multiple args in a slice, apply them to a variadic function using func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)

}
