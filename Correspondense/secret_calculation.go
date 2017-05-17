package Correspondense

import (
	"crypto/rand"
	"math/big"
	"math"
	"fmt"
)

type PrimeStruct struct {
	Prime1 *big.Int
	Prime2 *big.Int
	Prime3 *big.Int

	LocalSecret float64
	ForeignSecret *big.Int
}
func MakePrime() *big.Int {

	test,_ := rand.Prime(rand.Reader, 22)


	// Note that err == nil only if we read len(b) bytes.


	return test
}

func GetPrime(prime *big.Int) *big.Int {

	primes := PrimeStruct{}
	prime = primes.Prime1
	primes.Prime2 = MakePrime()
	return primes.Prime2
}

func VerifyPrime() {

}

func GenerateSecret(g float64, p float64, a float64 ) float64{

	//primes := PrimeStruct{}

	/*
	primes.Prime3 = MakePrime()
	g := primes.Prime1
	p := primes.Prime2
	a := primes.Prime3
	*/

	/*
	g := new(big.Float).SetInt(primes.Prime1)
	fmt.Print(g)
	p := new(big.Float).SetInt(primes.Prime2)
	fmt.Print(p)
	a := new(big.Float).SetInt(primes.Prime3)
	fmt.Print(a)
	*/


	fin := math.Pow(g, a)
	fmt.Println(fin)

	secret := math.Mod(fin, p)
	fmt.Println(secret)

	return secret
}

func CombineSecret(ForeignSecret float64, localRand float64, p float64) float64{
	fin := math.Pow(ForeignSecret, localRand)
	fmt.Println(fin)

	secret := math.Mod(fin, p)
	fmt.Println(secret)
	return secret

}
