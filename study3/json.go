package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name   string `json:"name"`
	Age    int    `josn:"age"`
	School string `json:"school"`
}

type B struct {
	Name    string `json:"name"`
	Age     int    `josn:"age"`
	School  string
	Address string
}

func main() {
	a := A{
		Name: "wang",
		Age:  12,
	}

	var b B

	data, _ := json.Marshal(a)

	fmt.Println(string(data))

	fmt.Println(b)
	json.Unmarshal(data, &b)
	fmt.Println(b)
}
