package overkiz

type CommandActionCommand struct {
	Type                       int      `json:"type,omitempty"`
	Name                       string   `json:"name"`
	Parameters                 []string `json:"parameters"`
	SensitiveParametersIndexes []int    `json:"sensitiveParametersIndexes,omitempty"`
	Authentication             string   `json:"authentication,omitempty"`
	Delay                      int      `json:"delay,omitempty"`
}

type CommandAction struct {
	DeviceURL string                 `json:"deviceURL"`
	Commands  []CommandActionCommand `json:"commands"`
}

type Command struct {
	Label                   string          `json:"label"`
	Metadata                string          `json:"metadata,omitempty"`
	Shortcut                bool            `json:"shortcut,omitempty"`
	NotificationTypeMask    int             `json:"notificationTypeMask,omitempty"`
	NotificationCondition   string          `json:"notificationCondition,omitempty"`
	NotificationText        string          `json:"notificationText,omitempty"`
	NotificationTitle       string          `json:"notificationTitle,omitempty"`
	TargetEmailAddresses    []string        `json:"targetEmailAddresses,omitempty"`
	TargetPhoneNumbers      []string        `json:"targetPhoneNumbers,omitempty"`
	TargetPushSubscriptions []string        `json:"targetPushSubscriptions,omitempty"`
	Actions                 []CommandAction `json:"actions"`
}

type CommandResponse struct {
	ExecId string `json:"execId"`
}

func NewSingleCommand(label string, deviceUrl string, name string, parameters []string) *Command {
	return &Command{
		Label: label,
		Actions: []CommandAction{
			{
				DeviceURL: deviceUrl,
				Commands: []CommandActionCommand{
					{
						Name:       name,
						Parameters: parameters,
					},
				},
			},
		},
	}
}
