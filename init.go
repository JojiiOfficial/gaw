package gaw

import (
	"math/rand"
	"time"
)

//Init inits the gaw functions (generate random seeds, etc...)
func Init() {
	rand.Seed(time.Now().UnixNano())
}
