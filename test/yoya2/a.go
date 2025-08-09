package main

import (
	"fmt"

	"github.com/llyb120/yoya2/y"
)

func main() {
	// var a = []int{1, 2, 3, 3}
	// b := y.Flex(a, func(v int, _ int) string {
	// 	if v <= 1 {
	// 		return ""
	// 	}
	// 	time.Sleep(1 * time.Second)
	// 	return fmt.Sprintf("%d", v)
	// }, y.NotEmpty, y.UseDistinct, y.UseAsync)

	// c := y.Filter(a, func(v int) bool {
	// 	return v != 1
	// })
	// c := y.Filter(a, y.Not, 1, 2)

	// fmt.Println(c)

	type Person struct {
		Name string
		Age  int

		Children []Person
	}

	var arr []Person
	for i := 0; i < 10; i++ {
		arr = append(arr, Person{
			Name: fmt.Sprintf("jian%d", i),
			Age:  i,
			Children: []Person{
				{
					Name: fmt.Sprintf("child%d", i),
					Age:  i,
				},
			},
		})
	}

	// 1 flex

	// 2 for append

	names := y.Pick[string](arr, "[Name*=child] Name")
	fmt.Println(names)

	// d := y.NewData[Person]()
	// d.Set(Person{Name: "jian", Age: 100})
	// d["abc"] = 1
	// d["def"] = 2

	// d.Set(1)
	// fmt.Println(d.Data())

	// str, _ := json.Marshal(d)
	// fmt.Println(string(str))
}
