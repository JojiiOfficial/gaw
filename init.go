package gaw

import (
	mrand "math/rand"
	"time"
)

// Init inits the gaw functions (generate random seeds, etc...)
func Init() {
	mrand.Seed(time.Now().UnixNano())
}
