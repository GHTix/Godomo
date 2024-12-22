package overkiz

import "fmt"

type Setup struct {
	CreationTime   int64  `json:"creationTime"`
	LastUpdateTime int64  `json:"lastUpdateTime"`
	ID             string `json:"id"`
	Location       struct {
		CreationTime              int64   `json:"creationTime"`
		LastUpdateTime            int64   `json:"lastUpdateTime"`
		City                      string  `json:"city"`
		Country                   string  `json:"country"`
		PostalCode                string  `json:"postalCode"`
		AddressLine1              string  `json:"addressLine1"`
		AddressLine2              string  `json:"addressLine2"`
		Timezone                  string  `json:"timezone"`
		Longitude                 float64 `json:"longitude"`
		Latitude                  float64 `json:"latitude"`
		TwilightMode              int     `json:"twilightMode"`
		TwilightAngle             string  `json:"twilightAngle"`
		TwilightCity              string  `json:"twilightCity"`
		SummerSolsticeDuskMinutes int     `json:"summerSolsticeDuskMinutes"`
		WinterSolsticeDuskMinutes int     `json:"winterSolsticeDuskMinutes"`
		TwilightOffsetEnabled     bool    `json:"twilightOffsetEnabled"`
		DawnOffset                int     `json:"dawnOffset"`
		DuskOffset                int     `json:"duskOffset"`
		CountryCode               string  `json:"countryCode"`
	} `json:"location"`
	Gateways []struct {
		GatewayID         string `json:"gatewayId"`
		Type              int    `json:"type"`
		SubType           int    `json:"subType"`
		PlaceOID          string `json:"placeOID"`
		AutoUpdateEnabled bool   `json:"autoUpdateEnabled"`
		Alive             bool   `json:"alive"`
		TimeReliable      bool   `json:"timeReliable"`
		Connectivity      struct {
			Status          string `json:"status"`
			ProtocolVersion string `json:"protocolVersion"`
		} `json:"connectivity"`
		UpToDate       bool   `json:"upToDate"`
		UpdateStatus   string `json:"updateStatus"`
		SyncInProgress bool   `json:"syncInProgress"`
		Mode           string `json:"mode"`
		Functions      string `json:"functions"`
	} `json:"gateways"`
	Devices                    []Device      `json:"devices"`
	Zones                      []interface{} `json:"zones"`
	ResellerDelegationType     string        `json:"resellerDelegationType"`
	Metadata                   string        `json:"metadata"`
	DisconnectionConfiguration struct {
		NotificationTitle       string   `json:"notificationTitle"`
		NotificationText        string   `json:"notificationText"`
		TargetPushSubscriptions []string `json:"targetPushSubscriptions"`
		NotificationType        string   `json:"notificationType"`
	} `json:"disconnectionConfiguration"`
	Oid       string `json:"oid"`
	RootPlace struct {
		CreationTime   int64         `json:"creationTime"`
		LastUpdateTime int64         `json:"lastUpdateTime"`
		Label          string        `json:"label"`
		Type           int           `json:"type"`
		Oid            string        `json:"oid"`
		SubPlaces      []interface{} `json:"subPlaces"`
	} `json:"rootPlace"`
	Features []struct {
		Name   string `json:"name"`
		Source string `json:"source"`
	} `json:"features"`
}

func (s *Setup) DevicesFilterByClass(class string) (*[]Device, error) {
	var devices []Device

	for _, device := range s.Devices {
		if device.UIClass == class {
			deviceCopy, err := device.Clone()
			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}
			devices = append(devices, *deviceCopy)
		}
	}
	return &devices, nil
}

func (s *Setup) DevicesFilterByClassAndPlace(class string, placeOid string) (*[]Device, error) {
	var devices []Device

	for _, device := range s.Devices {
		if device.UIClass == class && device.PlaceOID == placeOid {
			deviceCopy, err := device.Clone()
			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}
			devices = append(devices, *deviceCopy)
		}
	}
	return &devices, nil
}

func (s *Setup) DeviceFilterByOid(oid string) (*Device, error) {
	for _, device := range s.Devices {
		if device.Oid == oid {
			deviceCopy, err := device.Clone()
			if err != nil {
				return nil, fmt.Errorf("%w", err)
			}
			return deviceCopy, nil
		}
	}
	return nil, nil
}
