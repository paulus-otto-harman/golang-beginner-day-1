package main

import "fmt"

var satu int

func main() {
	m := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(m[3])

	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n[4])
}
