package main

import "fmt"

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

func main() {
	greet("Rafi", "J04032313044" )
	fmt.Println(test2("Rafi", "Jl. Kebon Jeruk"))
	fmt.Println(test3("Rafi", "Jl. Kebon Jeruk"))
}