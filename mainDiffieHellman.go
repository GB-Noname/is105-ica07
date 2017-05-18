package main

import (
	"./Correspondense"
	"fmt"
)

func main() {

	/*
	Due to problems with assignment this is an ilustration of
	how we would implement diffie hellman key exhcange
	Fixed numbers: 10, 541
	Person 1 secret number: 5
	Person 2 secret number: 7
	 */
	var g float64 = 10
	var p float64 = 541
	var sn1 float64 = 5
	var sn2 float64 = 7

	fmt.Println("Fixed numbers: ", g, " and ", p)
	fmt.Println("Person 1 random number: ", sn1)
	fmt.Println("Person 2 random number: ", sn2)

	//Generating secret for both correspondents based on diffie-hellman formula
	secret1 := Correspondense.GenerateSecret(g,p,sn1)
	secret2 := Correspondense.GenerateSecret(g,p,sn2)
	fmt.Println("Person 1 has the secret: ", secret1)
	fmt.Println("Person 2 has the secret: ", secret2)

	//Calculating the secret based on know fixed number, other persons secret and own rand number
	secret2Calc := Correspondense.CombineSecret(secret1,7, 541)
	secret1Calc := Correspondense.CombineSecret(secret2,5, 541)
	fmt.Println("Person 1 has calculated the secret from person 2 to: ", secret2Calc)
	fmt.Println("Person 2 has calculated the secret from person 1 to: ", secret1Calc)

}
