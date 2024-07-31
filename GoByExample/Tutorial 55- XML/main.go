package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`   // XMLName is the main tag
	Id      int      `xml:"id,attr"` // x,attr is an attribute and not a tag
	Name    string   `xml:"name"`    // Inside tag
	Origin  []string `xml:"origin"`  // Inside tag
}

func (p Plant) String() string { // Constructor for a plant struct
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}       // Create tag
	coffee.Origin = []string{"Ethiopia", "Brazil"} // Add tags

	out, _ := xml.MarshalIndent(coffee, " ", "  ") // MarshalIndent indents tags
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out)) // Add XML header

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil { // Decode xml and check for errors
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"` // Nesting, Plants will be entered last
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
