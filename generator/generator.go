package generator

import (
	"errors"
	"math/rand"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" //All the possible characters that can be used in the ids
	letterIdxBits = 6                                                                // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1                                             // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits                                               // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

//GenerateID - Creates a random sequence of characters into a string and prepends with the prefix, thus creating an ID
//This is the most efficient way according to https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func GenerateID(n int, prefix string) (string, error) {
	if len(prefix) != 3 {
		return "", errors.New("Prefix not 3 characters")
	}

	if n < 1 || n > 30 {
		return "", errors.New("You probably dont want to generate a string that doesnt exist or is that long")
	}

	bytes := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache = src.Int63()
			remain = letterIdxMax
		}

		idx := int(cache & letterIdxMask)
		if idx < len(letterBytes) {
			bytes[i] = letterBytes[idx]
			i--
		}

		cache >>= letterIdxBits
		remain--
	}

	randomChars := string(bytes)
	charsWithPrefix := prefix + "_" + randomChars
	return charsWithPrefix, nil
}
