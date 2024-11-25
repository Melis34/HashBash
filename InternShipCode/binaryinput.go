package main

import (
    "crypto/sha256"
    "fmt"
    "os"
)

func main() {
    // gets input, scans if there are 2 values supplied, one to call the command, the second is the binary data that will be hashed
    if len(os.Args) != 2 {
        fmt.Println("Usage: binary_hasher <binary_input>")
        os.Exit(1)
    }
    binaryInput := os.Args[1]

    // parses the text of 1 and 0 into the binary 1 and 0
    binaryBytes, err := parseBinaryString(binaryInput)
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(2)
    }

    // Hash the binary input
    hash := sha256.Sum256(binaryBytes)

    // Convert the hash to a hexadecimal string
    hexHash := fmt.Sprintf("%x", hash)

    // Print the hexadecimal hash
    fmt.Println("Hexadecimal Hash:", hexHash)
}

func parseBinaryString(s string) ([]byte, error) {
    // Ensures that the binary input exists out of full bytes
    if len(s)%8 != 0 {
        return nil, fmt.Errorf("Input length must be a multiple of 8")
    }

    // Initialize a []byte to store the binary data
    binaryData := make([]byte, len(s)/8)

    //takes the bit's provited and turns it into a array of bytes. 
    for i := 0; i < len(s)/8; i++ {
        byteStr := s[i*8 : (i+1)*8]
        var byteValue byte
        for j := 0; j < 8; j++ {
            if byteStr[j] == '1' { //if a 1 is found adds it into the array
                byteValue |= 1 << (7 - j) 
            } else if byteStr[j] != '0' { //if the value isn't either 1 or 0 gives an error
                return nil, fmt.Errorf("Invalid binary input: %s", s)
            }
        }
        binaryData[i] = byteValue
    }

    return binaryData, nil //returns the array of bytes, and no error since none was found
}
