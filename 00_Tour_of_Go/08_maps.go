package main

import "fmt"

func main() {
	/* 1. */ showcaseMaps()
}

type Vertex struct {
	Lat, Long float64
}

func showcaseMaps() {
	//	A map maps keys to values
	//	The zero value of a map is nil. A nil map has no keys, nor can keys be added
	//	The make function returns a map of the given type, initialized and ready for use
	m := make(map[string]Vertex)

	//	Assign's the struct to Bell Labs
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	//	Map literals are like struct literals but keys are required
	var n = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(n)

	//	Mutating Maps
	//	-------------

	//	Insert or update an element in map m
	m["Google"] = Vertex{37.42202, -122.08408}

	//	Retrieve an element
	ret := m["Google"]
	fmt.Println(ret)

	//	Delete an element
	delete(m, "Google")

	//	Test that a key is present with a two value assignment
	var elem, ok = m["Bell Labs"]
	fmt.Println(elem, ok)
	//	If key is in m, ok is true, otherwise ok is false
}
