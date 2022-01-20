package main

import "fmt"

//	One of the most ubiquitous interfaces is `Stringer` defined by the `fmt` package
//	```
//	type Stringer interface {
//		String() string
//	}
//	```

//	A `Stringer` is a type that can describe itself as a string.
//	The `fmt` package (and many others) look for this interface to print values

type Person struct {
	Name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	b := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, b)

	exercise()
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func exercise() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

}
