package main

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Select a hash function:")
    fmt.Println("1. SHA-256")
    fmt.Println("2. MD5")
    fmt.Println("3. SHA-1")

    var choice int
    fmt.Print("Enter your choice: ")
    fmt.Scan(&choice)

    var hashFunc func(data []byte) []byte
    var hashName string

    switch choice {
    case 1:
        hashFunc = func(data []byte) []byte {
            hash := sha256.Sum256(data)
            return hash[:]
        }
        hashName = "SHA-256"
    case 2:
        hashFunc = func(data []byte) []byte {
            hash := md5.Sum(data)
            return hash[:]
        }
        hashName = "MD5"
    case 3:
        hashFunc = func(data []byte) []byte {
            hash := sha1.Sum(data)
            return hash[:]
        }
        hashName = "SHA-1"
    default:
        fmt.Println("Invalid choice. Exiting.")
        return
    }

    hashMap := make(map[string]int)
    foundDuplicate := false

    for i := 0; i <= 10000; i += 8 {
        zeros := strings.Repeat("0", i)
        hash := hashFunc([]byte(zeros))
        hashHex := hex.EncodeToString(hash)

        if count, exists := hashMap[hashHex]; exists {
            fmt.Printf("Hash Function: %s\nNumber of zeros: %d\nHash: %s (Duplicate of %d zeros)\n\n", hashName, i, hashHex, count)
            foundDuplicate = true
            break
        } else {
            hashMap[hashHex] = i
            fmt.Printf("Hash Function: %s\nNumber of zeros: %d\nHash: %s\n\n", hashName, i, hashHex)
        }
    }

    if foundDuplicate {
        fmt.Println("Duplicate found. Exiting.")
    } else {
        fmt.Println("No duplicates found.")
    }
}
