package main

import "fmt"

func main() {
	fmt.Println("masukkan kalkulator suhu yang ingin dipakai")
	fmt.Println("1.Celcius Ke Fahrenheit")
	fmt.Println("2.Celcius Ke Kelvin")
	fmt.Println("3.Fahrenheit Ke Celcius")
	fmt.Println("4.Fahrenheit Ke Kelvin")
	fmt.Println("5.Kelvin Ke Celcius")
	fmt.Println("6.Kelvin Ke Fahrenheit")
	fmt.Println("Masukkan pilihan anda: ")
	var pilihan int
	fmt.Scanf("%d",&pilihan)
	for pilihan <1 || pilihan>6{
		fmt.Println("kalkulator tidak tersedia, masukkan kalkulator pilihan anda: ")
		fmt.Scanf("%d",&pilihan)
	}

	var suhu float64
	fmt.Println("masukkan suhu awal : ")
	fmt.Scanf("%f\n", &suhu)

	var suhuAkhir float64
	if pilihan == 1{
		suhuAkhir = CelciusToFarenheit(suhu)
	} else if pilihan ==2 {
		suhuAkhir = CelciusToKelvin(suhu)
	} else if pilihan == 3 {
		suhuAkhir = FahrenheitToCelcius(suhu)
	} else if pilihan == 4 {
		suhuAkhir = FahrenheitToKelvin(suhu)
	} else if pilihan == 5{
		suhuAkhir =KelvinToCelcius(suhu)
	} else {
		suhuAkhir = KelvinToFahrenheit(suhu)
	}
	fmt.Printf("Suhu akhir hasil konversi adalah : %.2f", suhuAkhir)

}

func CelciusToFarenheit(suhuCelcius float64) float64 {
	suhuFahrenheit := ((9.0 / 5.0) * suhuCelcius) + 32
	return suhuFahrenheit
}

func CelciusToKelvin(suhuCelcius float64) float64 {
	suhuKelvin := suhuCelcius + 273.15
	return suhuKelvin
}
func FahrenheitToCelcius(suhuFahrenheit float64) float64 {
	suhuCelcius := (suhuFahrenheit - 32) * (5.0 / 9.0)
	return suhuCelcius
}

func FahrenheitToKelvin(suhuFahrenheit float64) float64 {
	suhuKelvin := (suhuFahrenheit + 459.67) * (5.0 / 9.0)
	return suhuKelvin

}

func KelvinToCelcius(suhuKelvin float64) float64 {
	suhuCelcius := suhuKelvin - 273.15
	return suhuCelcius

}

func KelvinToFahrenheit(suhuKelvin float64) float64 {
	suhuFahrenheit := (suhuKelvin * (9.0 / 5.0)) - 459.67
	return suhuFahrenheit
}