package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func create_signature(content string, secret string) string {
	// 待签名的消息
	// HMAC使用的密钥

	// 创建一个新的HMAC SHA256哈希使用指定的密钥
	h := hmac.New(sha256.New, []byte(secret))
	// 写入消息以计算其哈希值
	h.Write([]byte(content))
	// 计算最终的HMAC值
	signature := h.Sum(nil)
	return hex.EncodeToString(signature)
}
