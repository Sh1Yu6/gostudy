package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name  string
	Age   int
	Sal   float64
	Skill string
}

func testStruct() []uint8 {
	monster := A{
		"请问",
		20,
		2.343,
		"卡三等奖",
	}

	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("Marshal error %v\n", err)
	}
	fmt.Println(string(data))
	return data
}

func testMap() {
	m := make(map[string]interface{})
	m["name"] = "几十块"
	m["age"] = 123
	m["addr"] = 123123.123123
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal error %v\n", err)
	}
	fmt.Println(string(data))

}

func testSlice() []uint8 {
	var s []map[string]interface{}
	m := make(map[string]interface{})
	m["name"] = "几十块"
	m["age"] = 123
	m["addr"] = 123123.123123
	s = append(s, m)
	m1 := make(map[string]interface{})
	m1["name"] = "几十123块"
	m1["age"] = 13423
	m1["addr"] = 155523123.123123
	s = append(s, m1)
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("Marshal error %v\n", err)
	}
	fmt.Println(string(data))
	return data
}

func testUnStruct(data []uint8) {

	str := `{"Name":"请问","Age":20,"Sal":2.343,"Skill":"卡三等奖"}`

	var a A
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("Unmarshal error %v\n", err)
	}
	fmt.Println(a)
}

func testUnMap(data []uint8) {

	str := `{"Name":"请问","Age":20,"Sal":2.343,"Skill":"卡三等奖"}`
	a := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("Unmarshal error %v\n", err)
	}
	fmt.Println(a)
}

func testUnSlice(data []uint8) {

	//str := `{"Name":"请问","Age":20,"Sal":2.343,"Skill":"卡三等奖"}`
	var a []map[string]interface{}
	err := json.Unmarshal([]byte(data), &a)
	if err != nil {
		fmt.Printf("Unmarshal error %v\n", err)
	}
	fmt.Println(a)
}

func main() {
	testStruct()
	testMap()
	data := testSlice()
	//testUnStruct(data)
	//testUnMap(data)
	testUnSlice(data)

}
