package handlers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/resend/resend-go/v2"
)

func CreateOTPHandler(resendClient *resend.Client) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){

	var otp string;
	for i := 0; i < 8; i++ {
		a, _ := rand.Int(rand.Reader, big.NewInt(10))
		str := strconv.FormatInt(a.Int64(), 10)
		otp += str
	}
	fmt.Println(otp)

	}
}