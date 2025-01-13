package main

import (
	"errors"
	"fmt"
)

// main adalah fungsi utama yang akan dijalankan pertama kali ketika program dieksekusi.
// Fungsi ini mendemonstrasikan deklarasi variabel panjang dan pendek, serta penggunaan tipe data primitif.
// Perbedaan antara int, int32, dan int64 adalah:
// - int: Tipe data integer yang ukurannya bergantung pada arsitektur mesin (32-bit atau 64-bit).
// - int32: Tipe data integer dengan ukuran tetap 32-bit.
// - int64: Tipe data integer dengan ukuran tetap 64-bit.

// Tipe data byte adalah tipe data yang hanya bisa diisi oleh bilangan bulat positif yang lebih kecil.
// Byte merupakan alias dari uint8, yang berarti byte dapat menyimpan nilai dari 0 hingga 255.
// une adalah tipe data di Go yang merepresentasikan sebuah karakter Unicode. Rune pada dasarnya adalah alias untuk tipe data int32. Ini memungkinkan kita untuk menyimpan karakter Unicode yang lebih besar dari 8-bit, yang tidak bisa ditampung oleh tipe data byte (alias untuk uint8).
func main(){
	// Long Declaration
	var a int = 10
	var Test string = "Hello World"

	fmt.Println(a)
	fmt.Println(Test)
	
	// Short Declaration
	count := 5
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	// Primitve data
	// String
	var stringVar string = "Hello World"  // perlu di note bahwa string disini tidak bisa seperti PHP maupun JS dimana string disini harus kutip dua " "
	fmt.Println(stringVar)

	// Boolean
	var booleanVar bool = true // atau false
	fmt.Println(booleanVar)

	// Integer terdapat 3 tipe data yaitu itn, int32, int64 begantung pada arsitektur mesinnya 
	intVar1 := int(2)
	intVar2 := int32(2)
	intVar3 := int64(2)
	fmt.Printf("Type: %T, Value: %v\n", intVar1, intVar1)
	fmt.Printf("Type: %T, Value: %v\n", intVar2, intVar2)
	fmt.Printf("Type: %T, Value: %v\n", intVar3, intVar3)
	// %T = Mengetahui type nya, %v = Mengetahui value nya
	fmt.Printf("Type: %T, Value: %v\n", a, a)

	// Float terdapat 2 tipe data yaitu float32 dan float64
	// floatVar := float(2.2), itu tidak bisa pada GO
	floatVar1 := float32(2.2)
	floatVar2 := float64(2.2)
	fmt.Printf("Type: %T, Value: %v\n", floatVar1, floatVar1)
	fmt.Printf("Type: %T, Value: %v\n", floatVar2, floatVar2)

	// Uniknya disini Go menyediakan Unint yang merupakan tipe data yang hanya bisa diisi oleh bilangan bulat positif
	var unsignedVar uint = 2
	fmt.Printf("Type: %T, Value: %v\n", unsignedVar, unsignedVar)

	// dan ada juga bytes yang merupakan tipe data yang hanya bisa diisi oleh bilangan bulat positif yang lebih kecil
	// 
	byteVar := byte(255)
	fmt.Printf("Type: %T, Value: %v\n", byteVar, byteVar)

	byteVar2 := []byte("Hello World") // ini akan menghasilkan byte dari string tersebut, jadi ini akan menghasilkan byte dari setiap karakter yang ada di string tersebut. Dia mengubah string menjad ASCII
	fmt.Printf("Type: %T, Value: %v\n", byteVar2, byteVar2)

	// dan ada juga rune yang merupakan tipe data yang hanya bisa diisi oleh bilangan bulat positif yang lebih besar
	varRune := rune('a')
	fmt.Printf("Type: %T, Value: %v\n", varRune, varRune)

	// Complex, yakni tipe data yang digunakan untuk merepresentasikan bilangan kompleks. Bilangan kompleks terdiri dari dua bagian, yaitu bagian real dan imajiner.
	varComplex := complex(2, 3)
	fmt.Printf("Type: %T, Value: %v\n", varComplex, varComplex)

	// ada juga namanya valuable error, dimana 
	varError1 := errors.New("Error 1")
	fmt.Printf("Type: %T, Value: %v\n", varError1, varError1)

	// Interface dimana interface ini sebagai penentu metode dan penyimpan data kosong
	var myInterFaceVar interface{} // ini berupa data kosong, dia akan berupa any type nya 
	fmt.Printf("Type: %T, Value: %v\n", myInterFaceVar, myInterFaceVar)

	myInterFaceVar = 2 // dia akan berupa int
	fmt.Printf("Type: %T, Value: %v\n", myInterFaceVar, myInterFaceVar)
	myInterFaceVar = "Hello World" // dia akan berupa string
	fmt.Printf("Type: %T, Value: %v\n", myInterFaceVar, myInterFaceVar)


}
