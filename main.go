package main

import (
	"fmt"
	"learn-utils/utils/list"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	persons := []Person{
		{
			Name: "kocak",
			Age:  1,
		},

		{
			Name: "test",
			Age:  1,
		},

		{
			Name: "Budi",
			Age:  12,
		},

		{
			Name: "Suso",
			Age:  12,
		},

		{
			Name: "Sus2",
			Age:  12,
		},

		{
			Name: "Susoi 123",
			Age:  12,
		},
	}
	var l = new(list.List)
	for _, v := range persons {
		l.Add(v)
	}

	var dic = l.GroupBy(func(s interface{}) interface{} {
		d, _ := s.(Person)
		return d.Age
	})

	fmt.Println(dic)
}
