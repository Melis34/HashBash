package main

import (
    "fmt"
    "regexp"
    "strconv"
	"strings"
	"math/big"
)

func main() {
    // Get user input for the message
    var message string
    fmt.Print("Enter a message: ")
    fmt.Scanln(&message)
    // Find all hexadecimal values in the input message
    hexValues := findHexValues(message)

    // Count the total length of all hexadecimal values
    totalHexLength := 0
    for _, hexValue := range hexValues {
        totalHexLength += len(hexValue)
    }

    if totalHexLength == 0 {
        fmt.Println("No hexadecimal values found in the input message.")
        return
    }

    // Extract the last hexadecimal value
    lastHexValue := hexValues[len(hexValues)-1]

    // Convert the last hexadecimal value to an integer
    lastHexInt, err := strconv.ParseInt(lastHexValue[2:], 16, 64)
    if err != nil {
        fmt.Println("Error parsing hexadecimal value:", err)
        return
    }

    // Calculate the number of hexadecimal digits
    numHexDigits := totalHexLength - len(hexValues) -1 // Subtract the '0x' prefix
    // Add the count of hexadecimal digits to the last hexadecimal value
	
	for i:= 1; i < numHexDigits; i++ {
		hexValue := strings.TrimPrefix(lastHexValue, "0x")
		intValue := new(big.Int)
		intValue.SetString(strings.TrimPrefix(hexValue, "0x"), 16)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		intValue.Add(intValue, big.NewInt(1))
		if intValue.Cmp(big.NewInt(16)) == 0 {
			intValue.SetInt64(0)
		}
		lastHexValue = fmt.Sprintf("0x%X", intValue)
	}
    lastHexInt += int64(numHexDigits)

    // Print the results
    fmt.Printf("Original Message: %s\n", message)
    fmt.Printf("output of hashfunction: %s\n", lastHexValue)
}

func findHexValues(input string) []string {
    re := regexp.MustCompile(`0[xX][0-9A-Fa-f]+`)
    return re.FindAllString(input, -1)
}
