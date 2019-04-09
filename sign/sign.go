package sign

import (
	"crypto/md5"
	"errors"
	"fmt"
	"sort"
	"strings"
	"strconv"
)

type kvPair struct {
	k, v string
}

type kvPairs []kvPair

func (t kvPairs) Sort() {
	sort.SliceStable(t, func(i, j int) bool {
		return t[i].k < t[j].k
	})
}

func (t kvPairs) Join() string {
	var args []string
	for _, kv := range t {
		args = append(args, kv.k+"="+kv.v)
	}
	return strings.Join(args, "&")
}

// 生成sign
func MakeSign(data map[string]interface{}, secretKey string) {
	p := kvPairs{}
	// 剔除空值 和 sign
	for k, v := range data {
		v := typeSwitcher(v)
		if !(v == "" || k == "sign") {
			p = append(p, kvPair{k, v})
		}
	}
	p.Sort()
	fmt.Println(p.Join())
	data["sign"] = md5Sign(p.Join(), secretKey)

}
func typeSwitcher(t interface{}) string {
	switch v := t.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	case int64:
		return strconv.Itoa(int(v))
	default:
		return ""
	}
}

// 验证sign
func VerifySign(data map[string]interface{}, secretKey string) error {
	p := kvPairs{}
	sign, ok := data["sign"]
	if !ok {
		return errors.New("sign not exist")
	}
	for k, v := range data {
		v := typeSwitcher(v)
		if !(v == "" || k == "sign") {
			p = append(p, kvPair{k, v})
		}
	}
	p.Sort()
	if sign == md5Sign(p.Join(), secretKey) {
		return nil
	} else {
		return errors.New("sign not same")
	}
}

func md5Sign(str, key string) string {
	h := md5.New()
	h.Write([]byte(str))
	h.Write([]byte("&" + key))
	return fmt.Sprintf("%x", h.Sum(nil))
}
