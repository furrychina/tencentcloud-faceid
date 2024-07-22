package faceid

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

func Sign(WBAppId, userId, ticket, nonce string) string {
	const version string = "1.0.0"
	//将 WBAppId、userId、version 连同 ticket、nonce 共五个参数的值进行字典序排序。
	arrayList := []string{version, WBAppId, ticket, nonce, userId}
	// 字典序排序
	sort.Strings(arrayList)
	// 将排序后的参数值拼接成一个字符串
	combinedRequestBody := strings.Join(arrayList, "")
	// 对拼接后的字符串进行 SHA1 计算
	h := sha1.New()
	h.Write([]byte(combinedRequestBody))
	sha1Hash := h.Sum(nil)
	// 将计算后的 SHA1 值转为十六进制字符串并返回
	return hex.EncodeToString(sha1Hash)
}

func SignRecord(WBAppId, orderNo, ticket, nonce string) string {
	const version string = "1.0.0"
	//将 WBAppId、userId、version 连同 ticket、nonce 共五个参数的值进行字典序排序。
	arrayList := []string{version, WBAppId, ticket, nonce, orderNo}
	// 字典序排序
	sort.Strings(arrayList)
	// 将排序后的参数值拼接成一个字符串
	combinedRequestBody := strings.Join(arrayList, "")
	// 对拼接后的字符串进行 SHA1 计算
	h := sha1.New()
	h.Write([]byte(combinedRequestBody))
	sha1Hash := h.Sum(nil)
	// 将计算后的 SHA1 值转为十六进制字符串并返回
	return hex.EncodeToString(sha1Hash)
}
