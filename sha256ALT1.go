package main

import (
	"fmt"
	"strconv"
)

//difference from normal sha-256:
/*
- constants have all same value
- initial hash value has all the same value

*/

func main() {
	// Prompt the user to enter the binary input
	fmt.Println("Please enter binary input:")

	var inputArg string
	fmt.Scanln(&inputArg)

	// Convert the binary string input into a byte slice
	var inputData []byte
	for i := 0; i < len(inputArg); i += 8 {
		if i+8 > len(inputArg) {
			fmt.Println("Error: Input binary string must be a multiple of 8 bits.")
			return
		}
		byteVal, err := strconv.ParseUint(inputArg[i:i+8], 2, 8)
		if err != nil {
			fmt.Println("Error: Invalid binary input.")
			return
		}
		inputData = append(inputData, byte(byteVal))
	}

	// Step 1: Initialize constants
	constantK := []uint32{
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
		0x00000000, 0x00000000, 0x00000000, 0x00000000,
	}

	// Initialize hash values
	var hash [8]uint32
	hash[0] = 0x00000000
	hash[1] = 0x00000000
	hash[2] = 0x00000000
	hash[3] = 0x00000000
	hash[4] = 0x00000000
	hash[5] = 0x00000000
	hash[6] = 0x00000000
	hash[7] = 0x00000000

	// Step 3: Pre-processing (Padding)
	inputLength := uint64(len(inputData)) * 8
	padding := []byte{0x80}
	var block []byte
	block = append(block, inputData...)
	block = append(block, padding...)
	for len(block)%64 != 56 {
		block = append(block, 0x00)
	}
	block = append(block,
		byte(inputLength>>56), byte(inputLength>>48), byte(inputLength>>40), byte(inputLength>>32),
		byte(inputLength>>24), byte(inputLength>>16), byte(inputLength>>8), byte(inputLength),
	)

	// Process the message in 512-bit blocks
	blockSize := 64
	numBlocks := len(block) / blockSize

	for i := 0; i < numBlocks; i++ {
		// Step 4a: Message schedule (W) computation
		W := make([]uint32, 64)
		for t := 0; t < 16; t++ {
			W[t] = uint32(block[i*blockSize+t*4])<<24 | uint32(block[i*blockSize+t*4+1])<<16 | uint32(block[i*blockSize+t*4+2])<<8 | uint32(block[i*blockSize+t*4+3])
		}
		for t := 16; t < 64; t++ {
			s0 := rightRotate(W[t-15], 7) ^ rightRotate(W[t-15], 18) ^ (W[t-15] >> 3)
			s1 := rightRotate(W[t-2], 17) ^ rightRotate(W[t-2], 19) ^ (W[t-2] >> 10)
			W[t] = W[t-16] + s0 + W[t-7] + s1
		}

		// Step 4b: Compression function
		a, b, c, d, e, f, g, h := hash[0], hash[1], hash[2], hash[3], hash[4], hash[5], hash[6], hash[7]

		for t := 0; t < 64; t++ {
			S1 := rightRotate(e, 6) ^ rightRotate(e, 11) ^ rightRotate(e, 25)
			ch := (e & f) ^ (^e & g)
			temp1 := h + S1 + ch + constantK[t] + W[t]
			S0 := rightRotate(a, 2) ^ rightRotate(a, 13) ^ rightRotate(a, 22)
			maj := (a & b) ^ (a & c) ^ (b & c)
			temp2 := S0 + maj

			h = g
			g = f
			f = e
			e = d + temp1
			d = c
			c = b
			b = a
			a = temp1 + temp2
		}

		// Step 4c: Update hash values (a-h) after each block
		hash[0] += a
		hash[1] += b
		hash[2] += c
		hash[3] += d
		hash[4] += e
		hash[5] += f
		hash[6] += g
		hash[7] += h
	}

	// Step 5: Output the final hash value
	fmt.Printf("SHA-256 Hash: %08x %08x %08x %08x %08x %08x %08x %08x\n", hash[0], hash[1], hash[2], hash[3], hash[4], hash[5], hash[6], hash[7])
}

func rightRotate(x uint32, n uint32) uint32 {
	return (x >> n) | (x << (32 - n))
}
