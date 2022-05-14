package main

//	Go's math/rand package provides pseudorandom number generation

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//	rand.Intn(100) returns a random int n -> 0 <= n < 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	//	rand.Float64 returns a float64 -> 0.0 <= f < 0
	fmt.Println(rand.Float64())

	//	This can be used to generate random floats in other ranges, e.g. 5.0 <= f < 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	//	The default number generator is deterministic, so it'll produce the same sequence of numbers
	//	each time by default. To produce varying sequence, give it a seed that changes. Note that this is Note
	//	safe to use for random numbers you intend to be secret, use crypto/rand for those
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1)

	//	Call the resulting rand.Rand just like function on the rand package
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	//	If you see a source with the same number, it produces the same sequence of random numbers
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
