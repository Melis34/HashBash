package main

import (
    "fmt"
    "regexp"
    "strconv"
    "strings"
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

    // Remove leading zeros and convert to integer
    intValue, err := strconv.ParseInt(strings.TrimPrefix(lastHexValue, "0x"), 16, 64)
    if err != nil {
        fmt.Println("Error converting hexadecimal to integer:", err)
        return
    }

    // Simplify the last digit
    simplifiedHex := fmt.Sprintf("0x%X", intValue%16)

    fmt.Printf("Original Message: %s\n", message)
    fmt.Printf("Simplified Hex Value: %s\n", simplifiedHex)
}

func findHexValues(input string) []string {
    re := regexp.MustCompile(`0[xX][0-9A-Fa-f]+`)
    return re.FindAllString(input, -1)
}
