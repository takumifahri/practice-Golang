package main

import "fmt"

func greet(nama string, NIM int) {
	fmt.Println("Hello", nama, "dengan NIM : ",NIM)
}

func main() {
	greet("Rafi", 18219001)
}