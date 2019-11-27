package tencent

import (
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func TestClient_GeneralOCR(t *testing.T) {
	client := NewClient(appId, key, 3, 9, 5, 1)
	data, err := ioutil.ReadFile("../test_materials/imgtotxt1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(data))
	req := &GeneralOCRRequest{[]byte(base64.StdEncoding.EncodeToString(data))}
	res, err := client.GeneralOCR(req)
	if err != nil {
		t.Fatal(err)
	}

	if res.Code != 0 {
		t.Fatalf("code: %d, msg: %s", res.Code, res.Msg)
	}

	for _, itemInfo := range res.Data["item_list"] {
		for _, item := range itemInfo {
			t.Log(item)
		}
	}
	//t.Log(res)
}
