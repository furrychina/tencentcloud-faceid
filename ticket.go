package faceid

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type SignTicket struct {
	Code            string   `json:"code"`
	Msg             string   `json:"msg"`
	TransactionTime string   `json:"transaction_time"`
	Tickets         []Ticket `json:"tickets"`
}

type Ticket struct {
	Value      string `json:"value"`
	ExpireIn   int    `json:"expire_in"`
	ExpireTime string `json:"expire_time"`
}

func GetSignTicket(accessToken, AppId, ticketType string) (SignTicket, error) {
	url := fmt.Sprintf("https://kyc.qcloud.com/api/oauth2/api_ticket?access_token=%s&app_id=%s&&type=%s&version=1.0.0", accessToken, AppId, ticketType)
	resp, err := http.Get(url)
	if err != nil {
		return SignTicket{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SignTicket{}, err
	}

	// 将响应体转换为结构体
	var result SignTicket
	if err := json.Unmarshal(body, &result); err != nil {
		return SignTicket{}, err
	}

	if result.Code != "0" {
		return SignTicket{}, errors.New(result.Msg)
	}
	return result, nil
}
