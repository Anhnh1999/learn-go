package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 2; j <= 8; j++ {
		fmt.Print(j)
		fmt.Print(" ")
	}

}
