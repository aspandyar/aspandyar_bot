package util

import (
	"math/rand"
	"time"
)

func RandomNumber1To5() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(5) + 1
}
