package common

import (
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// GetRandomNum 获取随机数  纯数字
func GetRandomNum(n int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// Sha1En 加密
func Sha1En(data string) string {
	t := sha1.New() ///产生一个散列值得方式
	_, _ = io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// InProtectedRoutes 查找某值是否在数组中
func InProtectedRoutes(v string, m *[]string) bool {
	for _, value := range *m {
		if value == v {
			return true
		}
	}
	return false
}

// HashPassword Combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password
// as a base64 encoded string
func HashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

// DoPasswordsMatch Check if two passwords match
func DoPasswordsMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = HashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}
