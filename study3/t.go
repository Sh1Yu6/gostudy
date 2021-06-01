package main

import (
	"fmt"
	"math"
	"math/big"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main() {

	i := big.NewInt(math.MaxInt64)
	i.Mul(i, big.NewInt(123123123))
	fmt.Printf("%v", i)
}
