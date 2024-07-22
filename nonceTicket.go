package faceid

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type NonceTicket struct {
	Code            string   `json:"code"`
	Msg             string   `json:"msg"`
	TransactionTime string   `json:"transaction_time"`
	Tickets         []ticket `json:"tickets"`
}

type ticket struct {
	Value      string `json:"value"`
	ExpireTime string `json:"expire_time"`
	ExpireIn   int    `json:"expire_in"`
}

func GetNonceTicket(appId, accessToken, userId string) (NonceTicket, error) {
	baseUrl := "https://kyc.qcloud.com/api/oauth2/api_ticket"
	url := fmt.Sprintf("%s?app_id=%s&access_token=%s&type=NONCE&version=1.0.0&user_id=%s", baseUrl, appId, accessToken, userId)
	resp, err := http.Get(url)
	if err != nil {
		return NonceTicket{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NonceTicket{}, err
	}

	// 将响应体转换为结构体
	var result NonceTicket
	if err := json.Unmarshal(body, &result); err != nil {
		return NonceTicket{}, err
	}

	if result.Code != "0" {
		return NonceTicket{}, errors.New(result.Msg)
	}
	return result, nil
}
