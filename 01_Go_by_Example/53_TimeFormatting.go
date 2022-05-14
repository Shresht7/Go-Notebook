package main

//	Go supports time formatting and parsing via pattern based layouts

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	//	This is an example of formatting time according to RFC3339
	fmt.Println(t.Format(time.RFC3339))

	//	Time parsing uses the same layout values as Format
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",
	)
	if e != nil {
		panic(e)
	}
	fmt.Println(t1)

	//	Format and Parse use example-based layouts. Usually you'll use a constant from time for these layouts,
	//	but you can also supply custom layouts. Layouts must use the reference time
	//	Mon Jan 2 15:04:15 MST 2006 to show the pattern with which to format/parse
	// 	a given time/string.
	fmt.Println(t.Format("3:04PM"))
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	if e != nil {
		panic(e)
	}
	fmt.Println(t2)

	//	For purely numeric representations you can also use standard string formatting with the extracted components
	//	of time value
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	//	Parse will return an error on malformed input explaining the parsing problem
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	fmt.Println(e)
}
