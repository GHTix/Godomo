package bridge

import (
	"encoding/json"
)

type Device struct {
	Disabled  bool `json:"disabled"`
	Endpoints struct {
		Num1 struct {
			Bindings []interface{} `json:"bindings"`
			Clusters struct {
				Input  []string `json:"input"`
				Output []string `json:"output"`
			} `json:"clusters"`
			ConfiguredReportings []interface{} `json:"configured_reportings"`
			Scenes               []interface{} `json:"scenes"`
		} `json:"1"`
		Num242 struct {
			Bindings []interface{} `json:"bindings"`
			Clusters struct {
				Input  []interface{} `json:"input"`
				Output []string      `json:"output"`
			} `json:"clusters"`
			ConfiguredReportings []interface{} `json:"configured_reportings"`
			Scenes               []interface{} `json:"scenes"`
		} `json:"242"`
	} `json:"endpoints"`
	FriendlyName       string `json:"friendly_name"`
	IeeeAddress        string `json:"ieee_address"`
	InterviewCompleted bool   `json:"interview_completed"`
	Interviewing       bool   `json:"interviewing"`
	NetworkAddress     int    `json:"network_address"`
	Supported          bool   `json:"supported"`
	Type               string `json:"type"`
	DateCode           string `json:"date_code,omitempty"`
	Definition         struct {
		Description string `json:"description"`
		Exposes     []struct {
			Access      int      `json:"access"`
			Category    string   `json:"category,omitempty"`
			Description string   `json:"description"`
			Label       string   `json:"label"`
			Name        string   `json:"name"`
			Property    string   `json:"property"`
			Type        string   `json:"type"`
			Unit        string   `json:"unit,omitempty"`
			ValueMax    int      `json:"value_max,omitempty"`
			ValueMin    int      `json:"value_min,omitempty"`
			Values      []string `json:"values,omitempty"`
		} `json:"exposes"`
		Model   string `json:"model"`
		Options []struct {
			Access      int    `json:"access"`
			Description string `json:"description"`
			Features    []struct {
				Access      int    `json:"access"`
				Description string `json:"description"`
				Label       string `json:"label"`
				Name        string `json:"name"`
				Property    string `json:"property"`
				Type        string `json:"type"`
				ValueMin    int    `json:"value_min"`
				Unit        string `json:"unit,omitempty"`
			} `json:"features"`
			Label    string `json:"label"`
			Name     string `json:"name"`
			Property string `json:"property"`
			Type     string `json:"type"`
		} `json:"options"`
		SupportsOta bool   `json:"supports_ota"`
		Vendor      string `json:"vendor"`
	} `json:"definition,omitempty"`
	Manufacturer    string `json:"manufacturer,omitempty"`
	ModelID         string `json:"model_id,omitempty"`
	PowerSource     string `json:"power_source,omitempty"`
	SoftwareBuildID string `json:"software_build_id,omitempty"`
}

func NewDevices(data []byte) ([]Device, error) {
	var newValue []Device
	err := json.Unmarshal(data, &newValue)
	if err != nil {
		return nil, err
	}
	return newValue, nil
}
