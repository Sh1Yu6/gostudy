package main

import (
	"fmt"
	"strings"
)

func Addupper() func(int) int {
	n := 10
	return func(x int) int {
		n += x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}
func main() {
	f1 := Addupper()
	fmt.Println(f1(10))
	fmt.Println(f1(2))
	fmt.Println(f1(3))

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("asd"))
	fmt.Println(f2("asd.jpg"))
}
package main

import (
	"fmt"
	"strings"
)

func Addupper() func(int) int {
	n := 10
	return func(x int) int {
		n += x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}
func main() {
	f1 := Addupper()
	fmt.Println(f1(10))
	fmt.Println(f1(2))
	fmt.Println(f1(3))

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("asd"))
	fmt.Println(f2("asd.jpg"))
}
package main

import (
	"fmt"
	"strings"
)

func Addupper() func(int) int {
	n := 10
	return func(x int) int {
		n += x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}
func main() {
	f1 := Addupper()
	fmt.Println(f1(10))
	fmt.Println(f1(2))
	fmt.Println(f1(3))

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("asd"))
	fmt.Println(f2("asd.jpg"))
}
package main

import (
	"fmt"
	"strings"
)

func Addupper() func(int) int {
	n := 10
	return func(x int) int {
		n += x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}
func main() {
	f1 := Addupper()
	fmt.Println(f1(10))
	fmt.Println(f1(2))
	fmt.Println(f1(3))

	f2 := makeSuffix(".jpg")
	fmt.Println(f2("asd"))
	fmt.Println(f2("asd.jpg"))
}
