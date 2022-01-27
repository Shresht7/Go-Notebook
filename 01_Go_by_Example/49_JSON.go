package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//	We'll use these two structs to demonstrate encoding and decoding of custom types
type response1 struct {
	Page   int
	Fruits []string
}

//	Only exported fields will be encoded/decoded in JSON.
//	Fields must start with a capital letter to be exported
type response2 struct {
	Page   int      `json: "page"`
	Fruits []string `json: "fruits"`
}

func main() {

	//	First we'll look at encoding basic data types to JSON strings
	bolB, _ := json.Marshal(true)
	fmt.Println(bolB)

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//	Slices and maps
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	//	The JSON package can automatically encode your custom data types
	//	It will only include exported fields in the encoded output and will by default
	//	use those names as the JSON keys
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	//	You can use tags on struct field declration to customize the encoded JSON jey names
	//	Check the definition of response2 above to see an example of such tags
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//	Decoding JSON data
	byt := []byte(`{"num": 6.13, "str": ["a", "b"]}`)

	//	We need to provide a variable where the JSON package can put the decoded data
	//	This map[string]interface{} will hold a map of strings to arbitrary data types
	var dat map[string]interface{}

	//	Here's the actual decoding and a check for associated errors
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	//	In order to use the values in the decoded map, we'll need to convert them to their appropriate type.
	num := dat["num"].(float64)
	fmt.Println(num)

	//	Accessing nested data requires a series of conversions
	arr := dat["str"].([]interface{})
	str1 := arr[0].(string)
	fmt.Println(str1)

	//	We can also decode JSON into custom data types.
	//	This has the advantage of adding additional type-safety to our programs
	//	and eliminating the need for type assertions when accessing the decoded data
	x := `{"page": 1, "fruits": ["applie", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(x), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//	In the examples above, we always use bytes and strings as intermediates between
	//	data and JSON representation on standard out. We can also stream JSON encodings directly
	//	to os.Writers like os.Stdout or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
