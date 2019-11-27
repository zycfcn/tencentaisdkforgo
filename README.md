

## 使用方法

使用该SDK做的图片识别在线工具：<https://tuchuang.richuyun.com/pictotxt.html>

##### 安装
`go get github.com/zycfcn/tencentaisdkforgo`

##### 使用
> 请事先申请腾讯AI开放平台的开发者，并准备好要识别的图片，并命名为"imtotxt.jpg 或自行修改为其它 


```golang
import (
	T "github.com/zycfcn/tencentaisdkforgo"
	"encoding/base64"
	"io/ioutil"
	"fmt"
)
func main() {
	// appId :type int
	client := T.NewClient(appId, key, 3, 9, 5, 1)
	data, err := ioutil.ReadFile("imgtotxt.jpg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	req := &T.GeneralOCRRequest{[]byte(base64.StdEncoding.EncodeToString(data))}
	res, err := client.GeneralOCR(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if res.Code != 0 {
		fmt.Println("code: %d, msg: %s", res.Code, res.Msg)
	}
	fmt.Println(res)
}
```

