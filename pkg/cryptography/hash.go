package cryptography

import (
	"encoding/base64"
	"golang.org/x/crypto/argon2"
)

type CryptParams struct {
	Times     uint32
	Memory    uint32
	Threads   uint8
	KeyLength uint32
	Salt      string
}

func HashPassword(password string, cryptParams CryptParams) string {

	hash := argon2.IDKey(
		[]byte(password),
		[]byte(cryptParams.Salt),
		cryptParams.Times,
		cryptParams.Memory,
		cryptParams.Threads,
		cryptParams.KeyLength,
	)
	
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	return b64Hash
}
