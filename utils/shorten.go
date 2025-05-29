package utils

import (
	"fmt"
	"time"
	"encoding/base64"
)

func GetShortCode() string {
	fmt.Println("Shortening URL")
	timeStamp := time.Now().UnixNano()
	fmt.Println("Timestamp: ", timeStamp)

	timeStamp_Bytes := []byte(fmt.Sprintf("%d", timeStamp))
	key := base64.StdEncoding.EncodeToString(timeStamp_Bytes)
	fmt.Println("Key: ", key)

	key = key[:len(key)-2]

	return key[16:]
}
