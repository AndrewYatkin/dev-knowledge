package commonTesting

import (
	"crypto/rand"
	"dev-knowledge/infrastructure/errors"
	"fmt"
	"github.com/google/uuid"
	"math/big"
	"strconv"
)

const (
	defaultMinNumber = int64(1000000)
	defaultMaxNumber = int64(9999999)

	floatNumberFrom0To1Base = 1000000
)

var (
	lettersForRandom = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandomEmail() string {
	login := RandomDefaultStr()
	mail := RandomStr(7)
	return fmt.Sprintf("%s@%s.test", login, mail)
}

func RandomUUID() string {
	return uuid.NewString()
}

func RandomURL() string {
	return "https://host-stub.ru/random-url/" + RandomDefaultStr()
}

func RandomBytesWithSize(bytesCount int64) []byte {
	bytes := make([]byte, bytesCount)
	for i := int64(0); i < bytesCount; i++ {
		randomByte := byte(RandomNumber(1, 100) % 2)
		bytes[i] = randomByte
	}

	return bytes
}

func RandomBytes() []byte {
	bytesCount := RandomNumber(5, 25)

	bytes := make([]byte, bytesCount)
	for i := 0; i < bytesCount; i++ {
		randomByte := byte(RandomNumber(1, 100) % 2)
		bytes[i] = randomByte
	}

	return bytes
}

func RandomDefaultStr() string {
	return RandomStr(15)
}

func RandomStr(length int64) string {
	randomLetters := make([]rune, length)
	for i := int64(0); i < length; i++ {
		randomLetterIndex := RandomNumber(0, int64(len(lettersForRandom)-1))
		randomLetters[i] = lettersForRandom[randomLetterIndex]
	}

	return string(randomLetters)
}

func RandomDefaultNumberStr() string {
	randNumber := RandomDefaultNumber()
	return strconv.Itoa(randNumber)
}

func RandomDefaultNumber() int {
	return RandomNumber(defaultMinNumber, defaultMaxNumber)
}

func RandomNumber(minNumber int64, maxNumber int64) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(maxNumber-minNumber))
	if err != nil {
		panic(err)
	}

	return int(nBig.Int64() + minNumber)
}

func RandomFloatNumberFrom0To1() float64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(floatNumberFrom0To1Base))
	if err != nil {
		panic(err)
	}

	return float64(nBig.Int64()) / floatNumberFrom0To1Base
}

func RandomFloatNumber(minFloatNumber float64, maxFloatNumber float64) float64 {
	return RandomFloatNumberFrom0To1()*(maxFloatNumber-minFloatNumber) + minFloatNumber
}

func RandomDefaultFloatNumber() float64 {
	return RandomFloatNumber(float64(defaultMinNumber), float64(defaultMaxNumber))
}

func RandomError() *errors.Error {
	randomErrorCode := errors.ErrorCode(RandomDefaultNumberStr())
	randomErrorText := RandomDefaultStr()
	return errors.NewError(randomErrorCode, randomErrorText)
}
