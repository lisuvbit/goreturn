package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "Man"
	fmt.Println("原字符串是：", str)
	enStr := Base64EncodeString(str)
	fmt.Println("编码后字符串是：", enStr)
	deStrBytes, deStr := Base64DecodeString(enStr)
	fmt.Printf("解码后字符串：字符串%s 字节%v\n", deStrBytes, deStr)

}

// Base64EncodeString 编码
func Base64EncodeString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64DecodeString 解码
func Base64DecodeString(str string) (string, []byte) {
	resBytes, _ := base64.StdEncoding.DecodeString(str)
	return string(resBytes), resBytes
}
