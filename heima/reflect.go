package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 123
	fmt.Println(reflect.TypeOf(a).Name())

	fmt.Println(reflect.ValueOf(a))
	v := reflect.ValueOf(a)
	c := v.Interface().(int)
	fmt.Println(c)

	x := 123
	y := 1235
	le(x, y)
}

func le(x, y interface{}) interface{} {
	xType := reflect.TypeOf(x)
	yType := reflect.TypeOf(y)
	if xType != yType {
		fmt.Println("!= err")
		return nil
	}
	xV := reflect.ValueOf(x).Interface()
	yV := reflect.ValueOf(y).Interface()
	aaa := xV.(int)
	//if xV > yV {
	//return xV
	//}
	return yV

}
