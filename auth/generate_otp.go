package auth

import (
	"crypto/rand"
	"math/big"
)

func GenerateOtp() (string, error) {
	//specify which characters to include in the one time password
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// set the length to 8 to indicate that it will be an 8 character byte
	b := make([]byte, 8)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		b[i] = chars[num.Int64()]
	}
	return string(b), nil
}
