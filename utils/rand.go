package utils

import (
	"math/rand"
	"time"
)

const uidLen = 10
const accessTokenLen = 20
const uidPrefix = "bu_"
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charsetNumber = "0123456789"

func genRandomString(prefix string, strLen int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, strLen)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return prefix + string(b)
}

func GenerateAccessToken() string {
	return genRandomString("", accessTokenLen, charset)
}
func GenerateUID() string {
	return genRandomString(uidPrefix, uidLen, charset)
}
func GenerateVerifyCode() string {
	return genRandomString("", 6, charsetNumber)
}
func GenerateAppID() string {
	return genRandomString("app_", 10, charset)
}

func GenerateAuditID() string {
	return genRandomString("audit_", 10, charset)
}

func GenerateValuateID() string {
	return genRandomString("val_", 16, charset)
}

func GenerateGroupID() string {
	return genRandomString("grp_", 16, charset)
}

func GenStringWithPrefix(prefix string, strLen int) string {
	return genRandomString(prefix, strLen, charset)
}
