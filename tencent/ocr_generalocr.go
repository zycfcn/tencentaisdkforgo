package tencent

import "bytes"


type GeneralOCRResponse struct {
	Code int 			`json:"ret"`
	Msg  string			`json:"msg"`
	Data map[string][]map[string]interface{} 	`json:"data"`
}

type GeneralOCRRequest struct {
	Image  []byte `json:"image"`
}

func (client *client) GeneralOCR(req *GeneralOCRRequest) (res *GeneralOCRResponse, err error) {
	params := client.genParams(req)
	res = &GeneralOCRResponse{}
	err = client.sendRequest("fcgi-bin/ocr/ocr_generalocr", bytes.NewReader(params), res)
	return
}
