package main

import (

	"./Correspondense"
	"fmt"
)

func main() {

	prime1 := Correspondense.MakePrime()
	fmt.Println(prime1)
	prime2 := Correspondense.GetPrime(prime1)
	fmt.Print(prime2)

	Correspondense.GenerateSecret(10,541,5)
}
