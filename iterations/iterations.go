package iterations

import "fmt"

func MapIterate () {
	sample := map[string]int {
		"age": 12,
		"weight": 123,
	}

	for k, v := range sample {
		fmt.Println(k, v)
	}
}


func ArrIterate () {
	arr := [...]string{"a", "b", "c", "d"}

	for index, value := range arr {
		fmt.Println(index, value)
	}
}