package pkg

import (
	"crypto/rand"
	"math/big"
)

type CryptStruct struct{}

func (c CryptStruct) GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)

	var availRetries = 3
	var i = 0
	for i < n {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err == nil {
			ret = append(ret, letters[num.Int64()])
			i++
			continue
		}

		if availRetries != 0 {
			availRetries--
			continue
		}

		return "", err

	}

	return string(ret), nil
}

var Crypt = CryptStruct{}
