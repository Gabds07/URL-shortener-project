package utils

import (
	"fmt"
	"time"
	"encoding/base64"
)

func GetShortCode() string {

	timeStamp := time.Now().UnixNano()
	fmt.Println("Timestamp: ", timeStamp)

	timeStamp_Bytes := []byte(fmt.Sprintf("%d", timeStamp))
	key := base64.StdEncoding.EncodeToString(timeStamp_Bytes)

	key = key[:len(key)-2]

	return key[16:]
}
