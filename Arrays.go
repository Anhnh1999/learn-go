package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("arr:", arr)

	//gan gia tri y nhu python
	arr[4] = 100

	fmt.Println("phan tu thu 5", arr[4])
	fmt.Println("arr sau khi gan gia tr", arr)

	//lay do dai cua arr, y nhu python
	fmt.Println("do dai cua arr", len(arr))

	//Khai bao mang int y nhu c

	// /var arr1[5] int =  {1,2,3,4,5}
	arr2 := [5]int{5, 4, 3, 2, 1}

	fmt.Println(arr2)

	//mang 2 chieu
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
