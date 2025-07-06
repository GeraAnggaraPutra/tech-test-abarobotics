package logger

import (
	"encoding/json"
	"fmt"
)

// ParseJSON will transform struct data as json string.
func ParseJSON(data interface{}) string {
	JSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(JSON)
}

// ParsePrettyJSON will transform struct data as json indent string.
func ParsePrettyJSON(data interface{}) string {
	JSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(JSON)
}

// generateMessage will transform message string and fields struct data as json string.
func generateMessage(msg string, fields []interface{}) string {
	if len(fields) == 0 {
		return msg
	}

	var json string

	if len(fields) > 1 {
		for i := 1; i < len(fields); i++ {
			json = fmt.Sprintf("%s\n%s : %s", json, fields[i-1], ParseJSON(fields[i]))
			i++
		}
	}

	return fmt.Sprintf("%s%s", msg, json)
}
