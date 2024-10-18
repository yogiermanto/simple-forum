package refresh_token

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateRefreshToken() string {
	randByte := make([]byte, 18)
	_, err := rand.Read(randByte)
	if err != nil {
		return ""
	}

	nowByte := []byte(fmt.Sprintf("%d", time.Now().Unix()))
	return hex.EncodeToString(randByte) + "." + base64.StdEncoding.EncodeToString(nowByte)
}
