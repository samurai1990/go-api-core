package api_test

import (
	"encoding/json"
)

func SwapResponseToMapStruct(data []byte) map[string]interface{} {
	var res map[string]interface{}
	json.Unmarshal(data, &res)
	return res
}

func CheckBodyOkErrors(body map[string]interface{}) bool {
	if body["error"].(bool) == false && body["error_code"].(float64) == 0 {
		return true
	}
	return false
}
