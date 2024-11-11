package overkiz

import (
	"encoding/json"
	"fmt"
)

// https://ha101-1.overkiz.com/enduser-mobile-web/enduserAPI/setup/devices
type Device struct {
	CreationTime     int64  `json:"creationTime"`
	LastUpdateTime   int64  `json:"lastUpdateTime"`
	Label            string `json:"label"`
	DeviceURL        string `json:"deviceURL"`
	Shortcut         bool   `json:"shortcut"`
	ControllableName string `json:"controllableName"`
	Definition       struct {
		Commands []struct {
			CommandName string `json:"commandName"`
			Nparams     int    `json:"nparams"`
		} `json:"commands"`
		States []struct {
			Type          string   `json:"type"`
			Values        []string `json:"values,omitempty"`
			QualifiedName string   `json:"qualifiedName"`
			EventBased    bool     `json:"eventBased,omitempty"`
		} `json:"states"`
		DataProperties []interface{} `json:"dataProperties"`
		WidgetName     string        `json:"widgetName"`
		UIProfiles     []string      `json:"uiProfiles"`
		UIClass        string        `json:"uiClass"`
		QualifiedName  string        `json:"qualifiedName"`
		Type           string        `json:"type"`
	} `json:"definition"`
	States []struct {
		Name  string      `json:"name"`
		Type  int         `json:"type"`
		Value interface{} `json:"value"`
	} `json:"states,omitempty"`
	Available  bool   `json:"available"`
	Enabled    bool   `json:"enabled"`
	PlaceOID   string `json:"placeOID"`
	Type       int    `json:"type"`
	Widget     string `json:"widget"`
	Oid        string `json:"oid"`
	UIClass    string `json:"uiClass"`
	Attributes []struct {
		Name  string      `json:"name"`
		Type  int         `json:"type"`
		Value interface{} `json:"value"`
	} `json:"attributes,omitempty"`
}

func (d Device) Clone() (*Device, error) {
	var deviceCopy Device
	data, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	err = json.Unmarshal(data, &deviceCopy)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &deviceCopy, nil
}
