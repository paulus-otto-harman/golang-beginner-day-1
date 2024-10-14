package main

import (
	"fmt"
	"strconv"
)

//buat variabel slice
//jumlah data 10
//tampilkan index ke 5 sampai 9

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(n[5 : 9+1])

	fmt.Println(strconv.FormatFloat(float64(21), 'E', 2, 64))
}
