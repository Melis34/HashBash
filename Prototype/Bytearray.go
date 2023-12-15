package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Read input from the terminal
	fmt.Print("Enter the array of bytes (space-separated): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Convert the input string to a byte array
	byteArray, err := parseInput(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return
	}

	// Hash the byte array using MD5
	hasher := sha256.New()
	hasher.Write(byteArray)
	hash := hasher.Sum(nil)

	// Display the hash as a space-separated list of decimal values
	for i, b := range hash {
		fmt.Print(b)
		if i < len(hash)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func parseInput(input string) ([]byte, error) {
	// Split the input string into space-separated values
	values := strings.Fields(input)

	// Convert the string values to integers
	var byteArray []byte
	for _, value := range values {
		b, err := parseByte(value)
		if err != nil {
			return nil, err
		}
		byteArray = append(byteArray, b)
	}

	return byteArray, nil
}

func parseByte(value string) (byte, error) {
	// Convert the string to an integer
	var num byte
	_, err := fmt.Sscanf(value, "%d", &num)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %s", value)
	}

	return num, nil
}
