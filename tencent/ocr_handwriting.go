package tencent

import "bytes"


type HandWritingOCRResponse struct {
	Code int 			`json:"ret"`
	Msg  string			`json:"msg"`
	Data map[string][]map[string]interface{} 	`json:"data"`
}

type HandWritingOCRRequest struct {
	Image  []byte `json:"image"`
	ImageUrl string `json:"image_url"`
}

func (client *client) HandWritingOCR(req *HandWritingOCRRequest) (res *HandWritingOCRResponse, err error) {
	params := client.genParams(req)
	res = &HandWritingOCRResponse{}
	err = client.sendRequest("fcgi-bin/ocr/ocr_handwritingocr", bytes.NewReader(params), res)
	return
}
