package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	//	Test whether a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	//	Above we used a string pattern directly, but for other regexp tasks
	//	you'll need to Compile an optimized regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")
	//	Many methods are available on these structs.
	fmt.Println(r.MatchString("peach"))

	//	This finds the match for the regex
	fmt.Println(r.FindString("peach punch"))

	//	This also finds the first match but returns the start and end indexes
	//	for the match instead of matching text
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	//	The submatch variant includes information about both the whole pattern
	//	matches and submatches within those matches. For example this will return information
	//	for both p([a-z]+)ch and ([a-z]+)
	fmt.Println(r.FindStringSubmatch("peach punch"))

	//	The All variants are available for other functions we saw above
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	//	Providing a non-negative integer as the second argument to these functions will limit the
	//	number of matches
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	//	Our examples above had string arguments and used names like MatchString.
	//	We can also provide []byte arguments and drop String from the function name
	fmt.Println(r.Match([]byte("peach")))

	//	When creating global variables with regular expressions you can use the MustCompile
	//	variation of Compile. MustCompile panics instead of returning an error, which makes it
	//	safe to use for global variables.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	//	The regexp package can also be used to replace subset of strings with other values
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	//	The Func variant allows you to transform matched text with a given function
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
