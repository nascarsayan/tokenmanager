package server

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

// computeArgMinHash finds the value in the range [start, end), for which Hash is minimum.
func computeArgMinHash(name string, start uint64, end uint64) uint64 {
	partial := start
	minVal := Hash(name, start)
	for i := start + 1; i < end; i++ {
		currVal := Hash(name, i)
		if currVal < minVal {
			minVal = currVal
			partial = i
		}
	}
	return partial
}
