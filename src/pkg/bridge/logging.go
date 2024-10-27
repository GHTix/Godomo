package bridge

import "encoding/json"

type Logging struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func NewLogging(data []byte) (*Logging, error) {
	var newValue Logging
	err := json.Unmarshal(data, &newValue)
	if err != nil {

		return nil, err
	}
	return &newValue, nil
}
