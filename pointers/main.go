package main

import (
	"fmt"
)

/* 
	Pointers yakni berupa variable yang dapat menyimpan memory address.
	jadi ketika ada variable yang usdah kita buat, kita coba kirimkan dengan poointer, maka variable tersebut akan berubah
*/
// Contoh pointer 1: 
func pointer1(satu1 *int) {
	*satu1 = *satu1 + 1
}


func main(){
	satu1 := 10
	// fmt.Printf("value nya %d\n satu1 :", satu1)

	pointer1(&satu1)
	fmt.Printf("value nya %X, satu1 : %d\n", &satu1, satu1)


	fmt.Println("Pointers")


	// kita bisa juga membuat pointer kosong
	var kosong *int 
	fmt.Printf("value nya %X, kosong : %v\n", &kosong, kosong)

	// Mengisi pointer kosong dengan nilai dari input
	var input *int
	fmt.Scan("Masukkan nilai untuk pointer kosong: ", &input)
	
}