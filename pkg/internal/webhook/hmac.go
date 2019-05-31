package webhook

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func VerifySha1(payload []byte, hmacSecret []byte, signature []byte) bool {
	mac := hmac.New(sha1.New, hmacSecret)
	mac.Write(payload)
	expectedMAC := mac.Sum(nil)
	fullComputedHash := fmt.Sprintf("sha1=%s", hex.EncodeToString(expectedMAC))
	return hmac.Equal(signature, []byte(fullComputedHash))
}
