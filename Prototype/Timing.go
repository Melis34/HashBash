package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func measureExecutionTime(f func()) time.Duration {
	startTime := time.Now()

	// Call the provided function
	f()

	// Calculate the elapsed time
	elapsedTime := time.Since(startTime)
	return elapsedTime
}

func main() {
	n := 1 //hoeveel bytes worden gecheckt
	fmt.Println("amount of bytes checked:", n)
	for i := 0; i < 1000; i++ {
		startinput := hash(generateRandomBytes(10)) //Start
		durationtrad := measureExecutionTime(func() { traditional(startinput, n) })
		durationown := measureExecutionTime(func() { ownmethod(startinput, n) })
		fmt.Println("Tradtitional time", durationtrad)
		fmt.Println("Own time", durationown)
		// runTimesOwn = append(runTimesOwn, duration)
		// runTimesTrad = append(runTimesTrad, duration)
	}
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

func traditional(input []byte, n int) { //takes hash selects random inputs and compares them
	output := hash(input)
	collision := false
	for collision == false {
		check := generateRandomBytes(32) // 2 keer zo lang dan gebruik van md5
		checkoutput := hash(check)
		if areFirstNBytesEqual(checkoutput, output, n) {
			fmt.Println("traditional input")
			fmt.Println("input ", input)       //original input
			fmt.Println("first", output)       // original hash
			fmt.Println("check", check)        //input that leads to collision
			fmt.Println("found ", checkoutput) //output of said input
			collision = true
		}
	}
	return
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

func ownmethod(input []byte, n int) {
	output := hash(input)
	prevhash := output
	collision := false
	nexthash := hash(output)
	for collision == false {
		if areFirstNBytesEqual(nexthash, output, n) {
			fmt.Println("own method")
			fmt.Println("input ", input)   //original input
			fmt.Println("first", output)   // original hash
			fmt.Println("check", prevhash) //input that leads to collision
			fmt.Println("found ", output)  //output of said input
			collision = true
		} else {
			prevhash = nexthash
			nexthash = hash(nexthash)
		}
	}
	return
}

func hash(data []byte) []byte {
	hash := md5.Sum(data)
	return hash[:]
}
