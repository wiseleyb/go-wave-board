package main

import (
	"math/rand"
	"time"
)

func newRnd(n int) int {
	rndSeed := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(rndSeed)
	return rnd.Intn(n)
}
