package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
)

func RNGGen(max int) int {

	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure RNG")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	rng := math_rand.Intn(max)
	return rng
}
