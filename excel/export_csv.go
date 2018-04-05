package excel

import (
	"bytes"
	"encoding/csv"
	"golang.org/x/text/transform"
	"strings"
	"golang.org/x/text/encoding/simplifiedchinese"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"io/ioutil"
)

// 导出CSV
func ExportCsv(c *gin.Context) {
	fileName := "test.csv"
	record := []string{"张三", "李四", "王五"}
	b := &bytes.Buffer{}
	wr := csv.NewWriter(b)
	for i := 0; i < 100; i++ {
		wr.Write(record)
	}
	wr.Flush()
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	csv, _ := UTF82GBK(b.Bytes())
	c.String(http.StatusOK, csv)
}

// 中文转码
func UTF82GBK(src []byte) (string, error) {
	reader := transform.NewReader(strings.NewReader(string(src)), simplifiedchinese.GBK.NewEncoder())
	if buf, err := ioutil.ReadAll(reader); err != nil {
		return "", err
	} else {
		return string(buf), nil
	}
}
