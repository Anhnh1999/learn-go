package main

import "fmt"

func main() {
	nums := []int{3, 56, 2}
	sum_index := 0
	sum_value := 0

	//neu chi co 1 gia tri  for i := range nums thi se la index
	for index, value := range nums {
		sum_index += index
		sum_value += value
	}
	fmt.Println("sum index", sum_index)
	fmt.Println("sum value", sum_value)

	fmt.Println("-------------------------------------------------------------")

	//voi maps thi se la key/value chu khong phai la index/value

	map1 := map[string]string{"a": "apple", "b": "ball"}

	for k, v := range map1 {
		fmt.Print("gia tri key la: ", k)
		fmt.Print("  ")
		fmt.Print("gia tri value la: ", v)
		fmt.Print("\n")
	}

	fmt.Println("-------------------------------------------------------------")

	//voi strings thi se la index/chr(index)
	for i, v := range "abcdef" {
		fmt.Print("index la ", i)
		fmt.Print("  ")
		fmt.Print("gia tri rune(ascii) la ", v)
		fmt.Print("\n")

	}
}
