package main

import (
	"fmt"
)

type Celcius struct {
	suhu float64
}
type Fahrenheit struct {
	suhu float64
}
type Kelvin struct {
	suhu float64
}

func (c Celcius) toCelcius() float64 {
	return c.suhu
}

func (c Celcius) toFahrenheit() float64 {
	return ((9.0 / 5.0) * c.suhu) + 32
}

func (c Celcius) toKelvin() float64 {
	return c.suhu + 273.15
}

func (f Fahrenheit) toFahrenheit() float64 {
	return f.suhu
}
func (f Fahrenheit) toCelcius() float64 {
	return (f.suhu - 32) * (5.0 / 9.0)
}

func (f Fahrenheit) toKelvin() float64 {
	return (f.suhu + 459.67) * (5.0 / 9.0)
}

func (k Kelvin) toKelvin() float64 {
	return k.suhu
}
func (k Kelvin) toCelcius() float64 {
	return k.suhu - 273.15
}

func (k Kelvin) toFahrenheit() float64 {
	return (k.suhu*(9.0/5.0) - 459.67)
}

type hitungSuhu interface {
	toCelcius() float64
	toFahrenheit() float64
	toKelvin() float64
}

func main() {
	fmt.Println("Masukkan suhu awal")
	fmt.Println("1.Celcius")
	fmt.Println("2.Fahrenheit")
	fmt.Println("3.Kelvin")
	fmt.Println("Masukkan suhu awal yang diinginkan: ")

	var suhuAwal int
	fmt.Scanf("%d\n", suhuAwal)
	for suhuAwal <1 || suhuAwal >3{
		fmt.Println(suhuAwal)
		fmt.Println("suhu awal tidak valid, masukkan suhu awal yang valid: ")
		fmt.Scanf("%d\n",&suhuAwal)
	}

	fmt.Println("Masukkan suhu akhir")
	fmt.Println("1.Celcius")
	fmt.Println("2.Fahrenheit")
	fmt.Println("3.Kelvin")
	fmt.Println("Masukkan suhu akhir yang diinginkan: ")

	var suhuAkhir int
	fmt.Scanf("%d\n", suhuAkhir)
	for suhuAkhir <1 || suhuAkhir >3{
		fmt.Println(suhuAkhir)
		fmt.Println("Suhu akhir tidak valid, masukkan suhu akhir yang diinginkan: ")
		fmt.Scanf("%d\n",&suhuAkhir)
	}
	var suhu float64
	fmt.Println("Masukkan suhu:")
	fmt.Scanf("%f", &suhu)

	var interfaceSuhu hitungSuhu
	switch suhuAwal{
	case 1:
		interfaceSuhu = Celcius{suhu}
	case 2:
		interfaceSuhu = Fahrenheit{suhu}
	case 3:
		interfaceSuhu =Kelvin{suhu}
	}

	var suhuFinal float64
	switch suhuAkhir{
	case 1:
		suhuFinal = interfaceSuhu.toCelcius()
	case 2:
		suhuFinal = interfaceSuhu.toFahrenheit()
	case 3: 
	suhuFinal = interfaceSuhu.toKelvin()
	}

	fmt.Printf("Suhu Akhir yang di dapat adalah: %.2f",suhuFinal )

}
