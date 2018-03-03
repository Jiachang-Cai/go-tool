package sign

import (
	"testing"
	"fmt"
)

// 私钥
const secretKey = "8b489108c2aacce6db991e4cf625b9c7"

// 生成签名
func TestMakeSign(t *testing.T) {
	data := map[string]string{
		"name":  "zhang",
		"color": "red",
	}
	MakeSign(data, secretKey)
	fmt.Println(data["sign"])
}

// 验证签名
func TestVerifySign(t *testing.T) {
	datas := []map[string]string{
		{"name": "zhang",
			"color": "red",
			"sign": "09326c3a2d8c447f3f1d6bca04c10ce1"},
		{"name": "zhang",
			"color": "red",
		},
		{"name": "zhang",
			"color_change": "red",
			"sign": "09326c3a2d8c447f3f1d6bca04c10ce1"},
	}
	for _, data := range datas {
		if err := VerifySign(data, secretKey); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("verify success")
	}
}
