package main

import "fmt"

func main(){
	var suhu float64
	fmt.Println("masukkan suhu:")
	fmt.Scanf("%f", &suhu)
	fmt.Printf("suhu anda adalah: %f \n ", suhu)

	var uang uint64
	fmt.Println("masukkan uang:")
	fmt.Scanf("%d",&uang)
	fmt.Printf("Uang yang diinput adalah %d \n",uang)

	var nama string
	fmt.Println("masukkan nama anda")
	fmt.Scanf("%s", &nama)
	fmt.Printf("Nama yang di input adalah : %s", nama)
}