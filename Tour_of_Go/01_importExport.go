package main

/*
	Packages can aslo be imported individually
	import "fmt"
	import "math/rand"
	However, it is good practice to use factored import like below
*/

import (
	"fmt"
	"time"
)

func main() {
	//	Returns the current time
	// If a variable/function name is capitalized then it is exported from the package.
	//	`Pizza` and `Pi` are exported names while `pineapple` and `cake` are not.
	//	`Println` function is exported from the fmt package
	fmt.Println("The time is", time.Now())
}

//	When importing a package, you can only refer to the publically available exported names.
//	An unexported name is not accessible from outside the package and hence is private.
