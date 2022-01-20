package main

/*
	Packages can also be imported individually
	import "fmt"
	import "math/rand"
	However, it is good practice to use factored import like below
*/

import (
	"fmt"
	"time"
)

/*
	By convention, the package name is the same as the last element of the import path
	For instance, the `math/rand` package comprises files that begin with the statement `package rand`
*/

func main() {
	//	Returns the current time
	fmt.Println("The time is", time.Now())
}

/*
	If a variable/function name is capitalized then it is exported from the package.
	`Pizza` and `Pi` are exported names while `pineapple` and `cake` are not.
	`Println` function is exported from the fmt package

	When importing a package, you can only refer to the publically available exported names.
	An un-exported name is not accessible from outside the package and hence is private.
*/
