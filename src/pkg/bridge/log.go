package bridge

import "encoding/json"

type Log struct {
	Message string `json:"message"`
	Meta    struct {
		FriendlyName string `json:"friendly_name"`
	} `json:"meta"`
	Type string `json:"type"`
}

func NewLog(data []byte) (*Log, error) {
	var newValue Log
	err := json.Unmarshal(data, &newValue)
	if err != nil {

		return nil, err
	}
	return &newValue, nil
}
