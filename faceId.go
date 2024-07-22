package faceid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetFaceIdRequest struct {
	AppId   string `json:"appId"`
	Name    string `json:"name"`
	IdCard  string `json:"idNo"`
	UserId  string `json:"userId"`
	Sign    string `json:"sign"`
	Nonce   string `json:"nonce"`
	OrderNo string `json:"orderNo"`
	Version string `json:"version"`
}
type CommonResponse struct {
	Code            string `json:"code"`            // 状态码
	Msg             string `json:"msg"`             // 结果描述
	BizSeqNo        string `json:"bizSeqNo"`        // 业务流水号
	TransactionTime string `json:"transactionTime"` // 接口调取时间
}

type GetFaceIdResponse struct {
	CommonResponse
	Result faceIdResult `json:"result"`
}

type RecordResponse struct {
	CommonResponse
	Result PostVerifyResult `json:"result"`
}

type faceIdResult struct {
	BizSeqNo        string `json:"bizSeqNo"`
	TransactionTime string `json:"transactionTime"`
	OrderNo         string `json:"orderNo"`
	FaceId          string `json:"faceId"`
	Success         bool   `json:"success"`
}

type RecordRequest struct {
	AppId        string `json:"appId"`        // 应用ID
	Version      string `json:"version"`      // 版本号
	Nonce        string `json:"nonce"`        // 随机数
	OrderNo      string `json:"orderNo"`      // 订单号
	Sign         string `json:"sign"`         // 签名
	GetFile      string `json:"getFile"`      //是否需要获取人脸识别的视频和文件，值为1则返回视频和照片、值为2则返回照片、值为3则返回视频；其他则不返回
	QueryVersion string `json:"queryVersion"` // 查询接口版本号(传1.0则返回 sdk 版本号和 trtc 标识)
}

type PostVerifyResult struct {
	OrderNo      string `json:"orderNo"`      // 订单号
	LiveRate     string `json:"liveRate"`     // 活体检测通过率
	Similarity   string `json:"similarity"`   // 相似度
	OccurredTime string `json:"occurredTime"` //	刷脸时间
	Photo        string `json:"photo"`        //	刷脸照片
	Video        string `json:"video"`        // 刷脸视频
	SdkVersion   string `json:"sdkVersion"`   // SDK版本
	TrtcFlag     string `json:"trtcFlag"`     // Trtc 渠道刷脸则标识"Y"
	AppId        string `json:"appId"`        //	应用ID
}

func GetFaceId(data *GetFaceIdRequest) (GetFaceIdResponse, error) {
	var result GetFaceIdResponse
	var orderNo string = createOrderID()
	baseUrl := "https://miniprogram-kyc.tencentcloudapi.com/api/server/getfaceid"
	url := fmt.Sprintf("%s?orderNo=%s", baseUrl, orderNo)

	data.OrderNo = orderNo
	data.Version = "1.0.0"
	request, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(request))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	if result.Code != "0" {
		return result, nil
	}

	return result, nil
}

func GetResult(req *RecordRequest) (RecordResponse, error) {
	var result RecordResponse

	url := fmt.Sprintf("https://kyc1.qcloud.com/api/v2/base/queryfacerecord?orderNo=%s", req.OrderNo)
	request, err := json.Marshal(req)
	if err != nil {
		return RecordResponse{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(request))
	if err != nil {
		return RecordResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RecordResponse{}, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return RecordResponse{}, err
	}
	return result, nil
}
