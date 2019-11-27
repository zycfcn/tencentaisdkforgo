package tencent

import "bytes"


type IdCardOCRResponse struct {
	Code int 			`json:"ret"`
	Msg  string			`json:"msg"`
	Data map[string]interface{} 	`json:"data"`
}

type IdCardOCRRequest struct {
	Image  		[]byte `json:"image"`
	CardType 	int `json:"card_type"`		// 身份证图片类型，0-正面，1-反面
}

func (client *client) IdCardOCR(req *IdCardOCRRequest) (res *IdCardOCRResponse, err error) {
	params := client.genParams(req)
	res = &IdCardOCRResponse{}
	err = client.sendRequest("fcgi-bin/ocr/ocr_idcardocr", bytes.NewReader(params), res)
	return
}
