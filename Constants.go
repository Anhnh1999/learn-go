package main

import (
	"fmt"
)

const s string = "hang so"

func main() {
	fmt.Println(s)

	const n = 6

	const d = 6.0 / 2

	fmt.Println(d)

	fmt.Println(int64(d))
}
