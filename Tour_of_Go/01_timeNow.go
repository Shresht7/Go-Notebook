package main

/*
 *	Packages can aslo be imported individually
 *	import "fmt"
 *	import "math/rand"
 *	However, it is good practice to use factored import like below
 */

import (
	"fmt"
	"time"
)

func main() {
	//	Returns the current time
	fmt.Println("The time is", time.Now())
}
