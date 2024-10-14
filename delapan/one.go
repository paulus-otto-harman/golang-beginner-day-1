package main

import (
	"fmt"
	"strconv"
)

func main() {
	hargaJual := 150000.0
	hargaBeli := 100000.0
	biayaOperasional := 1000.0
	diskon := 15.0
	jumlahTerjual := 100

	hargaJualSetelahDiskon := hargaJual - diskon/100*hargaJual
	fmt.Println("Harga jual setelah diskon :", hargaJualSetelahDiskon)

	totalPendapatan := float64(jumlahTerjual) * hargaJualSetelahDiskon
	fmt.Println("Total Pendapatan :", strconv.FormatFloat(totalPendapatan, 'f', 2, 64))

	totalBiaya := float64(jumlahTerjual) * biayaOperasional
	fmt.Println("Total Biaya :", totalBiaya)

	modal := float64(jumlahTerjual) * hargaBeli
	totalKeuntungan := totalPendapatan - modal - totalBiaya
	fmt.Println("Total Keuntungan :", strconv.FormatFloat(totalKeuntungan, 'f', 2, 64))
}
