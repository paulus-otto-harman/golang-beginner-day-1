package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 1
	j := "-1x"
	k := 3.14
	var j2, _ = strconv.Atoi(j)

	var k2 = float32(k)
	fmt.Println(i + j2)
	fmt.Println(k2)

	l := "3,14"
	var l2, _ = strconv.ParseFloat(l, 64)

	fmt.Println("l2 = ", l2)

	m := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(m)

	n := []int{1, 2, 3}
	fmt.Println(n)
}
