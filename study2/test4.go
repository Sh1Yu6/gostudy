package main

import "fmt"

type demoI interface {
	funDemo1(string) int
	funDemo2(int64) string
}
type demoS1 struct{}

func (d *demoS1) funDemo1(string) int {
	fmt.Println("ppp")
	return 0
}
func (d demoS1) funDemo2(int64) string {
	fmt.Println("vvv")
	return ""
}
func main() {
	s := demoS1{}
	demo3(s)
}
func demo3(d demoI) {
	d.funDemo1("asdf")
	d.funDemo2(123)
}
