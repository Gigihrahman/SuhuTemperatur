package main

import (
	"fmt"
	"time"
)
type Person struct {
	Name string
	Age  int
}

func SetAge(p *Person, start,end int, done chan bool)  {
	for i := 0; i< end; i++{
		p.Age= i
	}
	done <- true
}

func main() {
	start := time.Now()
	person := Person{Name: "Alice", Age: 20}
	// Deklarasi variabel dengan tipe pointer ke Person
	
	numWorkers := 4
	step := 100000000 / numWorkers
	done := make(chan bool)

	for i := 0; i < numWorkers; i++ {
		go SetAge(&person, i*step, (i+1)*step, done)
	}
	for i := 0; i < numWorkers; i++ {
		<-done
	}
	fmt.Println(person.Age) // Output: 30
	fmt.Println(person)

	duration := time.Since(start)
	fmt.Println("Execution speed:", duration)}