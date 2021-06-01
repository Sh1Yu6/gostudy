package db

import (
	"fmt"
	"testing"
)

func TestF123(t *testing.T) {
	rss := UserSignup("asde", "123")
	fmt.Println(rss)
}
