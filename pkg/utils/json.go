package utils

import (
	"bytes"
	"encoding/json"
)

// ConvertJsonToStr 将 JSON 数据转换为字符串
func ConvertJsonToStr(data interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)

	// 编码 JSON 数据到 buffer 中
	if err := encoder.Encode(data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func ConvertStrToList(data string, model []interface{}) error {
	dataBytes := []byte(data)
	return json.Unmarshal(dataBytes, &model)
}
