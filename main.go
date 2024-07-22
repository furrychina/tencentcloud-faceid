package faceid

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	IdCard = "123456789012345678"
	Name   = "张三"
	userId = "123"
)

var appId = os.Getenv("FACEID_APPID")
var appSecret = os.Getenv("FACEID_APP_SECRET")
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomNumber(length int) string {
	var result string
	for i := 0; i < length; i++ {
		result += strconv.Itoa(seededRand.Intn(10)) // 生成0到9之间的随机数，并转换为字符串
	}
	return result
}

func main() {
	ramNonce := generateRandomNumber(32)
	accessToken, err := GetAccessToken(appId, appSecret)
	if err != nil {
		panic(err)
	}

	signTicket, err := GetSignTicket(accessToken.AccessToken, appId, "SIGN")
	if err != nil {
		panic(err)
	}

	sign := Sign(appId, userId, signTicket.Tickets[0].Value, ramNonce)

	getFaceIdRequest := GetFaceIdRequest{
		AppId:  appId,
		Name:   Name,
		IdCard: IdCard,
		UserId: userId,
		Sign:   sign,
		Nonce:  ramNonce,
	}
	getFaceIdResponse, err := GetFaceId(&getFaceIdRequest)
	if err != nil {
		panic(err)
		return
	}
	if getFaceIdResponse.Code != "0" {
		slog.Error("实名认证不通过", getFaceIdResponse.Msg)
		return
	}

	fmt.Println(map[string]interface{}{
		"faceId":          getFaceIdResponse.Result.FaceId,
		"bizSeqNo":        getFaceIdResponse.Result.BizSeqNo,
		"transactionTime": getFaceIdResponse.Result.TransactionTime,
		"orderNo":         getFaceIdResponse.Result.OrderNo,
		"success":         getFaceIdResponse.Result.Success,
		"nonce":           ramNonce,
		"sign":            sign,
	})
}
