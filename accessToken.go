package faceid

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type AccessToken struct {
	Code            string `json:"code"`
	Msg             string `json:"msg"`
	TransactionTime string `json:"transaction_time"`
	AccessToken     string `json:"access_token"`
	ExpireTime      string `json:"expire_time"`
	ExpireIn        int    `json:"expire_in"`
}

func GetAccessToken(appId, secret string) (AccessToken, error) {
	url := fmt.Sprintf("https://kyc.qcloud.com/api/oauth2/access_token?app_id=%s&secret=%s&grant_type=client_credential&version=1.0.0", appId, secret)
	resp, err := http.Get(url)
	if err != nil {
		return AccessToken{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AccessToken{}, err
	}

	// 将响应体转换为结构体
	var result AccessToken
	if err := json.Unmarshal(body, &result); err != nil {
		return AccessToken{}, err
	}

	if result.Code != "0" {
		return AccessToken{}, errors.New(result.Msg)
	}

	return result, nil
}
