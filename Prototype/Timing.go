package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("test")
	n := 2 //hoeveel bytes worden gecheckt
	fmt.Println("amount of bytes checked:", n)
	totalownmethod := 0
	totaltraditional := 0
	for i := 0; i < 1000; i++ {
		startinput := hash(generateRandomBytes(10)) //Start

		traditional := traditional(startinput, n)
		own := ownmethod(startinput, n)
		totalownmethod = totalownmethod + own
		totaltraditional = totaltraditional + traditional
	}
	fmt.Println("Results for this test are as follows:")
	fmt.Println("Total of traditional: ", totaltraditional)
	fmt.Println("Total of ownmethod: ", totalownmethod)
	fmt.Println("This gives an avarage of:")
	fmt.Println("Traditional: ", totaltraditional/1000)
	fmt.Println("Ownmethod: ", totalownmethod/1000)
}

/* func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(randomString)
} */

func generateRandomBytes(length int) []byte {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// Handle the error
	}
	return randomBytes
}

func traditional(input []byte, n int) int { //takes hash selects random inputs and compares them
	result := 1
	output := hash(input)
	collision := false
	for collision == false {
		result++
		check := generateRandomBytes(32) // 2 keer zo lang dan gebruik van md5
		result++
		checkoutput := hash(check)
		result++
		if areFirstNBytesEqual(checkoutput, output, n) {
			collision = true
		}
	}
	return result
}

func areByteArraysEqual(arr1, arr2 []byte) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func areFirstNBytesEqual(arr1, arr2 []byte, n int) bool {
	if len(arr1) < n || len(arr2) < n {
		return false
	}

	for i := 0; i < n; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func ownmethod(input []byte, n int) int {
	result := 1
	output := hash(input)
	collision := false
	for collision == false {
		result++
		nexthash := hash(output)
		result++
		if areFirstNBytesEqual(nexthash, output, n) {
			collision = true
		} else {
			output = nexthash
		}
	}
	return result
}

func hash(data []byte) []byte {
	hash := md5.Sum(data)
	return hash[:]
}
