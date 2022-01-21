package main

import "fmt"

func main() {

	//	For is Go's only looping construct.

	//	Single-Condition. Essentially a while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	//	Standard for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//	Infinite loop
	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
