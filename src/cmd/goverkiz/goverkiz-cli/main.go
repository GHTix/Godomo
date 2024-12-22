package main

import (
	"context"
	"io"
	"log/slog"

	"github.com/ghtix/gomodo/internal/config"
	"github.com/ghtix/gomodo/internal/rest"
	overkiz "github.com/ghtix/gomodo/pkg/overkiz/model"

	"gopkg.in/yaml.v3"
)

func main() {
	config, err := config.New("../private/config.yml")
	if err != nil {
		slog.Error("main", "error loading config", err.Error())
	}

	c := rest.New(
		config.Overkiz.BaseUrl,
		rest.WithOAuth(rest.NewOAuthData(
			rest.OAuthConfig{
				LoginEndPoint: config.Overkiz.OAuthLoginEndpoint,
				ClientId:      config.Overkiz.OAuthClientId,
				ClientSecret:  config.Overkiz.OAuthClientSecret,
				UserName:      config.Overkiz.UserName,
				Password:      config.Overkiz.Password,
			}),
		),
	)

	// resp, err := c.Get(context.Background(), "setup")
	// if err != nil {
	// 	slog.Error("goverkiz-cli", "error getting devices", err.Error())
	// }

	// setup, err := rest.JsonResponse[overkiz.Setup](resp)
	// if err != nil {
	// 	slog.Error("goverkiz-cli", "error json unmarshaling setup", err.Error())
	// }

	// for _, device := range setup.Devices {
	// 	println(device.Label, " - ", device.UIClass)
	// }

	// resp, err = c.Get(context.Background(), "/actionGroups")
	// if err != nil {
	// 	slog.Error("goverkiz-cli", "error getting exec current", err.Error())
	// }

	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	slog.Error("goverkiz-cli", "error reading response exec current", err.Error())
	// }

	// println(string(data))

	command := overkiz.NewSingleCommand("chambre 1", "io://1003-5304-7279/15435125", "open", nil)

	resp, err := c.PostJson(context.Background(), "exec/apply", command)
	if err != nil {
		slog.Error("goverkiz-cli", "error ", err.Error())
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("goverkiz-cli", "error reading response exec current", err.Error())
	} else {
		println(string(data))
	}

	LogToYaml(command)
}

func LogToYaml(data any) {
	out, err := yaml.Marshal(data)
	if err != nil {
		slog.Error("goberkiz-cli", "error yaml marshaling setup", err.Error())
	}
	println(string(out))
}
