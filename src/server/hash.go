package main

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// Hash concatentates a message and a nonce and generates a hash value.
func Hash(name string, nonce uint64) uint64 {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s %d", name, nonce)))
	return binary.BigEndian.Uint64(hasher.Sum(nil))
}

// getArgMin finds the value in the range [start, end), for which the Hash is minimum.
func getArgMin(name string, start uint64, end uint64) uint64 {
	partial := start
	minV := Hash(name, start)
	for i := start + 1; i < end; i++ {
		currVal := Hash(name, i)
		if currVal < minV {
			minV = currVal
			partial = i
		}
	}
	return partial
}
