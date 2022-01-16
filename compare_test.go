package go_struct_comparator

import (
	"fmt"
	"testing"
)

type Person struct {
	Age    int     `compare_key:"年龄"`
	Name   string  `compare_key:"姓名"`
	Child  Child   `compare_key:"孩子"`
	Phones []Phone `compare_key:"电话"`
}

// the comparator can be used across different type
type PersonB struct {
	Age    int     `compare_key:"年龄"`
	Name   string  `compare_key:"姓名"`
	Child  Child   `compare_key:"孩子"`
	Phones []Phone `compare_key:"电话"`
}

type Child struct {
	Age  int    `compare_key:"年龄"`
	Name string `compare_key:"姓名"`
}

type Phone struct {
	Num     int    `compare_key:"号码"`
	Address string `compare_key:"归属地"`
}

func TestComparator(t *testing.T) {
	a := Person{
		Age:  10,
		Name: "Wxl",
		Phones: []Phone{
			{Num: 123, Address: "1"},
			{Num: 456, Address: "2"},
		},
		Child: Child{
			Age:  5,
			Name: "Libaisi",
		},
	}
	b := PersonB{
		Age:  11,
		Name: "Zzw",
		Phones: []Phone{
			{Num: 123, Address: "2"},
			{Num: 457, Address: "2"},
		},
		Child: Child{
			Age:  6,
			Name: "Huangyou",
		},
	}
	result := Compare(a, b)
	fmt.Println(result)
}
