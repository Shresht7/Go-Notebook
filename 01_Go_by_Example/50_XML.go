package main

//	Go offers built-in support for XML-like formats

import (
	"encoding/xml"
	"fmt"
)

//	Plant will be mapped to XML. Similar to the JSON example, field tags contain directives for the encoder
//	and decoder.
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=$v, name=$v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	//	Emit XML representing our plant; using MarshalIndent to produce a more human-readable output
	out, _ := xml.MarshalIndent(coffee, " ", " ")
	fmt.Println(string(out))

	//	To add a generic XML header to output append it explicitly
	fmt.Println(xml.Header + string(out))

	// Use unmarshal to parse a stream of bytes with XML in a data structure
	//	If the XML is malformed or cannot be mapped onto Plant, a descriptive error will be returned
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// the parent>child>plant field tells the encoder to nest all plants under <parent><child>
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent("nesting", " ", " ")
	fmt.Println(string(out))
}
