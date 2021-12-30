package go_struct_comparator

import (
	"fmt"
	"testing"
)

type A struct {
	A int    `compare_key:"test"`
	B string `compare_key:"test2"`
}

type B struct {
	A int    `compare_key:"test"`
	B string `compare:"test2"`
}

func TestComparator(t *testing.T) {
	a := A{
		A: 1,
		B: "2",
	}
	b := B{
		A: 2,
		B: "2",
	}
	fmt.Println(structCompare(a, b))
}
