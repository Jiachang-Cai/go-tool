package sign

import (
	"testing"
	"fmt"
)

// 私钥
const secretKey = "0gGAfdaxAgQoYJPG"

// 生成签名
func TestMakeSign(t *testing.T) {
	data := map[string]interface{}{
		"name":  "zhang",
		"color": "red",
	}
	MakeSign(data, secretKey)
	fmt.Println(data["sign"])
}

// 验证签名
func TestVerifySign(t *testing.T) {
	datas := []map[string]interface{}{
		{"name": "zhang",
			"color": "red",
			"sign": "4ea8e4c4f6d518b0b03174771f63a84e"},
		{"name": "zhang",
			"color": "red",
		},
		{"name": "zhang",
			"color_change": "red",
			"sign": "4ea8e4c4f6d518b0b03174771f63a84e"},
	}
	for _, data := range datas {
		if err := VerifySign(data, secretKey); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("verify success")
	}
}
