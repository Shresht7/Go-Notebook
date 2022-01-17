package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	loops()
}

//	LOOPS
//	=====

func loops() {
	sum := 0

	//	Go's only looping construct is for
	//	It can be used in many ways however
	//	This is the standard initializer, condition, iterator construct
	//	No enclosing parantheses for whatever reason. Not a fan
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += 1
	}

	fmt.Println("The sum is", sum)

	sum = 1 //	Reset sum to 1. (0 wont work for the following example)

	//	The init and post conditions are optional
	for sum < 1000 {
		sum += sum
	}
	//	If this looks like a while loop, then you're right. Go has no special while keyword

	//	If no condition is given then it loops forever...
	for {
		break //	...least until it breaks
		// continue //	Also exists
	}
	fmt.Println("The sum is", sum)

	//	Move to if conditions
	ifConditions()
}

//	CONDITIONAL STATEMENTS
//	======================

func ifConditions() {
	sqrt(4)
	sqrt(-3)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	switchStatements()
}

func sqrt(x float64) string {

	//	like for loops, if conditions do not have parantheses
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {

	//	Like for loops, if statements can start with a initialized variable
	//	It is only valid inside the if block
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		//	Also valid in its else block. Oh, this is an else clause btw
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim

}

func switchStatements() {
	//	A switch statement is a shorter way to write a sequence of if-else statements
	//	It runs the first case whose value matches the conditional expression
	//	Unlike other languages, Go only runs the selected case, not all the cases that follow
	//	There is no need to break after every case in Go!
	//	Go's switch cases need not be constants and the values involved need not be integers

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)
	}

	//	Switch statements evaluate from top to bottom and return when a case succeeds

	//	A switch without a condition is the same as switch true
	//	Can be substituted for a long if-else chain
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	deferStatements()
}

func deferStatements() {
	//	A defer statements defers the execution of a function until the surrounding function returns
	//	The deferred call's arguments are evaluated immediately, but the function call is deferred until the surrounding function completes

	defer fmt.Println("I am not the first to execute")

	fmt.Println("YOLO")

	//	Deferred function calls are pushed onto a stack and executed in LIFO (last-in first-out) order
	defer fmt.Println("Am I the first?")
}
