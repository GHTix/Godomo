package aqara

import "encoding/json"

type Aqara struct {
	Battery          int     `json:"battery"`
	Humidity         float64 `json:"humidity"`
	Linkquality      int     `json:"linkquality"`
	PowerOutageCount int     `json:"power_outage_count"`
	Pressure         int     `json:"pressure"`
	Temperature      float64 `json:"temperature"`
	Voltage          int     `json:"voltage"`
}

func New(data []byte) (*Aqara, error) {
	var newValue Aqara
	err := json.Unmarshal(data, &newValue)
	if err != nil {
		return nil, err
	}
	return &newValue, nil
}
