package verificationPrimitive

import (
	"crypto/rand"
	"math/big"
	"strconv"
)

const (
	minCode int64 = 100000
	maxCode int64 = 999999
)

func GenerateCode() (string, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(maxCode-minCode))
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(minCode+nBig.Int64(), 10), nil
}
