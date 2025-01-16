package main

import (
	"fmt"
)

func Hello_Tes(){
	fmt.Println("Hello Tes")
}

func main(){
	fmt.Println("Conditional Syntax")

	// If, Else if, else

	/*
	var umur int
	fmt.Print("Masukkan umur anda: ")
	_, err := fmt.Scan(&umur)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else if umur > 17 {
		fmt.Printf("Umurmu %d, sudah cukup umur\n", umur)
	} else if umur < 17 {
		fmt.Printf("Umurmu %d, belum cukup umur\n", umur)
	} else {
		fmt.Printf("Umurmu %d, terlalu tua\n", umur)
	}
	*/


	// Switch Case
	// var Coffe string
	// fmt.Printf("Masukkan kopi yang ingin dipesan: ")
	// fmt.Scan(&Coffe)
	// switch Coffe {
	// 	case "Americano":
	// 		fmt.Println("Americano")
	// 	case "Latte":
	// 		fmt.Println("Latte")
	// 	case "Cappucino":
	// 		fmt.Println("Cappucino")
	// 	case "Mocha":
	// 		fmt.Printf("Tidak ada kopi")
	// 		fallthrough
	// 	default:
	// 		fmt.Println("Tidak ada kopi")
	// }

	// Looping
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Looping ke-",i)
	// }

	// While, bedanya di golang tidak ada penggunaan kata kunci while. kita tetap bisa menggunakan for
	// halo := -1
	// for halo < 5 {
	// 	fmt.Println(halo)
	// 	halo++
	// }

	// Looping dengan break dan continue
	Looping := 0 
	for {
		Looping += 1
		if Looping == 1 {
			continue
		}
		fmt.Printf("Looping ke-%d\n", Looping)
		if Looping == 5 {
			break
		}
	}

	// Defer
	// Defer adalah sebuah keyword yang membuat function yang kita panggil menggunakan defer akan dipanggil diakhir.
	// Biasanya digunakan untuk clean up data, close file, close database, dll.
	defer Hello_Tes()
}