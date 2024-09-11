package utils

import (
	"encoding/json"
	"fmt"
)

func ValueJsonString(value any) string {
	resultJSON, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(resultJSON)
}
