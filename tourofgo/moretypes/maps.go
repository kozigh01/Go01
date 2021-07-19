package moretypes

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
) 

type VertexMaps struct {
	Lat, Long float64
}

func maps1() {
	var m map[string]VertexMaps
	fmt.Printf("m as declared: %v (m = nil: %v)\n", m, m == nil)

	m = make(map[string]VertexMaps)
	fmt.Printf("m after make(map[string]Vertex: %v (m = nil: %v)\n", m, m == nil)

	m["Bell Labs"] = VertexMaps{40.68433, -74.39967}	
	m["Google"] = VertexMaps{37.42202, -122.08408 }
	fmt.Printf("Vertex for Bell Labs: %v\n", m)

	// map literal
	m = map[string]VertexMaps {
		"Bell Labs": {40.68433, -74.39967},
		"Google": {37.42202, -122.08408 },
	}

	fmt.Printf("Vertex for Bell Labs: %v\n", m)
}

func maps2() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Printf("Adding - m[Answer]: %v\n", m["Answer"])
	
	m["Answer"] = 46
	fmt.Printf("Updating - m[Answer]: %v\n", m["Answer"])

	delete(m, "Answer")
	fmt.Printf("Deleting - m[Answer]: %v\n", m["Answer"])

	v, ok := m["Answer"]
	fmt.Printf("Deleting (with existance check) - m[Answer]: %v, key exists: %v\n", v, ok)
}

func maps3() {
	wc.Test(WordCount)
}
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	
	for _, w := range strings.Fields(s) {
		fmt.Println(w)
		m[string(w)]++
	}
	return m
}