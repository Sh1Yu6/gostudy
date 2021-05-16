package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type A struct {
}

type B struct {
}

func (a A) Start() {
	fmt.Println("A start")
}

func (a A) Stop() {
	fmt.Println("A stop")

}

func (a B) Start() {
	fmt.Println("B start")

}

func (a B) Stop() {
	fmt.Println("B stop")

}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	a := A{}
	b := B{}
	fmt.Println(a, b)

	c := Computer{}
	c.Working(a)
	c.Working(b)
}
