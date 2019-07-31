package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/ttacon/libphonenumber"
	"regexp"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

type Equatable interface {
	isEqual(another interface{}) bool
}

type Hashable interface {
	GetHash() string
}

func IsValidPhone(value string) bool {
	_, err := libphonenumber.Parse(value, "")
	return err == nil
}

func IsValidEmail(value string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(value)
}

func IsEqualArray(a, b []Equatable) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.isEqual(b[i]) {
			return false
		}
	}
	return true
}
