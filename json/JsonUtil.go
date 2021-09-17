package _json

import "encoding/json"

func GetJsonString(obj interface{}) (jsonStr string) {
	bytes, err := json.Marshal(obj)
	if err == nil {
		jsonStr = string(bytes)
	}
	return jsonStr
}
