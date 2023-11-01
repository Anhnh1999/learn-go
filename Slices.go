package main

import (
	"fmt"
)

func main() {
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	//tao mot slice s
	s = make([]string, 3)
	//hoac don gian hon s := []string{1,2,3}
	//de tao mang thi can chi ra kich thuoc mang s := [3]string{}

	fmt.Println("emp", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("apd", s)

	c := make([]string, len(s))

	//c = s
	copy(c, s)
	fmt.Println("c:", c)

	hehe := [5]string{"1", "2", "3", "4", "5"}

	//hehe = append(hehe, "6", "7")
	l := hehe[2:5]
	fmt.Println(l)

	s1 := []int{}
	fmt.Println(s1)

}
