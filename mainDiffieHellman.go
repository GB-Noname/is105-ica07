package main

import (

	"./Correspondense"
	"fmt"
)

func main() {
	/*
	prime1 := Correspondense.MakePrime()
	fmt.Println(prime1)
	prime2 := Correspondense.GetPrime(prime1)
	fmt.Print(prime2)
	*/

	/*
	Due to problems with assignment this is an ilustration of
	how we would implement diffie hellman key exhcange
	Fixed numbers: 10, 541
	Person 1 secret number: 5
	Person 2 secret number: 7
	 */

	//Generating secret for both correspondents based on diffie-hellman formula
	secret1 := Correspondense.GenerateSecret(10,541,5)
	secret2 := Correspondense.GenerateSecret(10,541,7)
	fmt.Println("Person 1 has the secret: ", secret1)
	fmt.Println("Person 2 has the secret: ", secret2)

	//Calculating the secret based on know fixed number, other persons secret and own rand number
	secret2Calc := Correspondense.CombineSecret(secret1,7, 541)
	secret1Calc := Correspondense.CombineSecret(secret2,5, 541)
	fmt.Println("Person 1 has calculated the secret from person 2 to: ", secret2Calc)
	fmt.Println("Person 2 has calculated the secret from person 2 to: ", secret1Calc)


}
