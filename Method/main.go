package main

import (
	"fmt"
)

// Method Interface di Golang
// 1. Type Alias = Tipe data baru yang kita ambil dari tipe data lain,
// Misalkan kita ingin membuat type yakni TinggiBadan, merupakan angka integer
type TinggiBadan int

// 2, Struct, disini struct berfungsi membaut object yang berisikan beberapa type alias
type Segitiga struct {
	Alas int
	Tinggi int
}

// 3. interface method adalah method yang dideklarasikan di luar struct, dan bisa diakses oleh struct
type Bentuk interface {
	Luas() int
	Keliling() int
}

type PersegiPanjang struct {
	Panjang int
	Lebar int
}

func (per PersegiPanjang) Luas() int {
	return per.Panjang * per.Lebar
}


func (tb TinggiBadan) TerlaluTinggi() int {
	if tb > 200 {
		fmt.Printf("Tinggi badan %d cm, terlalu tinggi\n", tb)
		return 1
	} else{
		fmt.Printf("Tinggi badan %d cm, sudah ideal\n", tb)
		return 0
	}
}

func main(){
	fahri := TinggiBadan(170)
	fmt.Printf("Tinggi badan Fahri adalah %d cm\n", fahri)

	fahri.TerlaluTinggi()

	// kita coba struct segitiga
	MySegitiga := Segitiga{
		Alas:10,
		Tinggi:5,
	}
	fmt.Printf("type: %T, value : %+v \n", MySegitiga, MySegitiga )

	PersegiKu := PersegiPanjang{
		Panjang: 10,
		Lebar: 5,
	}

	fmt.Printf("Luas Persegi Panjang: %d cm\n", PersegiKu.Luas())
	
}