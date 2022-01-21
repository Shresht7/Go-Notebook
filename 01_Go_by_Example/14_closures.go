package main

import "fmt"

//	Go supports anonymous functions which can form closures.

//	This function intSeq returns another function, which we define anonymously in the body of intSeq
//	The returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	//	We call intSeq, assigning the result (a function) to nextInt
	//	This function value captures its own it value, which will be updated each time we call nextInt
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//	To confirm that the state is unique to that particular function, create and test a new one
	newInts := intSeq()
	fmt.Println(newInts())
}
