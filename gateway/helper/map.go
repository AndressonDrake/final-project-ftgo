package helper

import "encoding/json"

func MapToStruct(data map[string]interface{}, out interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, out)
}