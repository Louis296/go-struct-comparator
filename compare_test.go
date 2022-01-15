package go_struct_comparator

import (
	"fmt"
	"testing"
)

type Person struct {
	Age   int    `compare_key:"年龄"`
	Name  string `compare_key:"姓名"`
	Child Child  `compare_key:"孩子"`
}

// the comparator can be used across different type
type PersonB struct {
	Age   int    `compare_key:"年龄"`
	Name  string `compare_key:"姓名"`
	Child Child  `compare_key:"孩子"`
}

type Child struct {
	Age  int    `compare_key:"年龄"`
	Name string `compare_key:"姓名"`
}

func TestComparator(t *testing.T) {
	a := Person{
		Age:  10,
		Name: "Wxl",
		Child: Child{
			Age:  5,
			Name: "Libaisi",
		},
	}
	b := PersonB{
		Age:  11,
		Name: "Zzw",
		Child: Child{
			Age:  6,
			Name: "Huangyou",
		},
	}
	result := Compare(a, b)
	fmt.Println(result)
}
