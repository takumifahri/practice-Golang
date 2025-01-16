package main

import (
	"fmt"
)

// Input pada Golang adalah data yang diberikan kepada program untuk diproses. 
// Input bisa berasal dari berbagai sumber seperti keyboard, file, atau jaringan.
// Output pada Golang adalah data yang dihasilkan oleh program setelah memproses input. 
// Output bisa ditampilkan ke layar, disimpan ke file, atau dikirim melalui jaringan.

func greet(nama string, NIM string) {
	// fmt.Println("Hello", nama, "dengan NIM : ",NIM)
	fmt.Printf("Hello %s dengan NIM : %s\n", nama, NIM) // ]
	fmt.Println("Hello", nama, "dengan NIM :", NIM)
}

func test2 (name string, alamat string) (hasil string) {
	// kita bisa langsung return maupun kita bisa juga dengan hasil
	hasil = "Hello " + name + " dengan alamat " + alamat
	return
}

func test3 (name, alamat string) (hasil string) {
	// kita bisa langsung return maupun kita bisa juga dengan hasil
	return "Hello " + name + " dengan alamat " + alamat
}


/* 
	NOTE : 
	1. Untuk function sendiri di golang digunakan kapitalisasi pada title function. jadi tidak ada public ataupun private
	2. Misalkan kita membuat function dengan nama test() maka kita bisa memanggilnya dengan cara test() pada main function
	yang dimana jika kapitalisasinya itu huruf kecil maka hanya bisa dipanggil di package main saja. Jika kapitalisasinya itu huruf besar
	maka bisa dipanggil di package lain

*/
func main() {
	greet("Rafi", "J04032313044" )
	fmt.Println(test2("Rafi", "Jl. Kebon Jeruk"))
	fmt.Println(test3("Rafi", "Jl. Kebon Jeruk"))
	// Memanggil fungsi Test() dari package lain

}