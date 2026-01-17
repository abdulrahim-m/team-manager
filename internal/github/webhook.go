package github

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/abdulrahim-m/team-manager/internal/config"
)

func VerifySignature(body []byte, signature string) bool {
	const prefix = "sha256="
	if len(signature) < len(prefix) || signature[:len(prefix)] != prefix {
		return false
	}

	mac := hmac.New(sha256.New, []byte(config.GetGithubSecret()))
	mac.Write(body)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(signature[len(prefix):]), []byte(expectedMAC))
}
