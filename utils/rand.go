package utils

import (
	"math/rand"
	"sync"
	"time"
)

var gMut sync.Mutex

const randKey = "QWERTYUIOPASDFGHJKLZXCVBNM1234567890QWERTYUIOPASDFGHJKLZXCVBNM1234567890QWERTYUIOPASDFGHJKLZXCVBNM1234567890"

func RandKey(ln int) string {
	gMut.Lock()
	defer gMut.Unlock()

	if ln == 0 {
		ln = 9
	}
	var result string
	pln := len(randKey)
	for i := 0; i < ln; i++ {
		time.Sleep(time.Nanosecond * 250)
		rand.Seed(time.Now().UnixNano())
		result += string(randKey[rand.Intn(pln)])
	}

	return result
}
