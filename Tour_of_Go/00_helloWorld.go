//	A package clause starts every source file
/*
 *	Every go program is made up of packages.
 *	Main is a special name declaring this source file
 *	an executable rather than a library.
 *	Executables start running in the main package.
 * 	This was a multi-line comment btw
 */
package main

//	Import declarations declare library packages used in this source file
import (
	"fmt" //	Formatting package in the Go standard library
)

//	Functions are defined using the `func` keyword
//	Main is a special function determining the entry-point of the application
func main() {
	//	Prinntln from the fmt package is used to output to stdout
	fmt.Println("Hello World!")
}
