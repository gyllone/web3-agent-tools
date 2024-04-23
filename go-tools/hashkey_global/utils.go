package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	query "github.com/google/go-querystring/query"
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

func requestWithSignature(reqObj interface{}, deFunc deserializeFunc, apiEndPoint string, method string, auth *HashKeyApiAuth) error {
	return generalRequest(reqObj, deFunc, apiEndPoint, method, auth)
}

func requestWithoutSignature(reqObj interface{}, deFunc deserializeFunc, apiEndPoint string, method string) error {
	return generalRequest(reqObj, deFunc, apiEndPoint, method, nil)
}

func generalRequest(reqObj interface{}, deFunc deserializeFunc, apiEndPoint string, method string, auth *HashKeyApiAuth) error {
	urlValues, err := query.Values(reqObj)
	if err != nil {
		return err
	}

	urlString := urlValues.Encode()

	url := ""
	if auth != nil {
		url = fmt.Sprintf("%s?%s&signature=%s", apiEndPoint, urlString, create_signature(urlString, auth.Secret))
	} else {
		url = fmt.Sprintf("%s?%s", apiEndPoint, urlString)
	}
	fmt.Printf("req: url: %s\n", url)

	req, _ := http.NewRequest(method, url, nil)

	req.Header.Add("accept", "application/json")
	if auth != nil {
		req.Header.Add("X-HK-APIKEY", auth.ApiKey)
	}
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Printf("get body:%s\n", body)

	if res.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}

	return deFunc(body)
}

type deserializeFunc func([]byte) error

func getDeserializeJsonFunc(ret interface{}) deserializeFunc {
	deserializeJson := func(content []byte) error {
		return json.Unmarshal(content, ret)
	}

	return deserializeJson
}

func getPtr[T any](value T) *T {
	return &value
}
