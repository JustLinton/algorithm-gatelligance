package utils

import (
	"crypto/sha256"
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
)

func GenerateUUID() string {
	var err error
	userUid := uuid.Must(uuid.NewV4(), err)
	strUsrUid := userUid.String()
	if err != nil {
		return "nil"
	}
	return strUsrUid
}

func GenerateSHA256(message []byte) string {
	hash := sha256.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}
