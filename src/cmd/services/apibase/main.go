package main

import (
	"errors"
	"log/slog"
	"net/url"

	"github.com/spf13/cobra"

	"github.com/ghtix/gomodo/pkg/mosquitto"
)

func main() {
	slog.Info("ApiBase begin")

	var rootCommand = &cobra.Command{}
	var mosquittoUrlString, mosquittoUser, mosquittoPwd, mosquittoClientId, mosquittoTopic string

	var cmd = &cobra.Command{
		Use:   "start",
		Short: "start",
		Run: func(cmd *cobra.Command, args []string) {
			if len(mosquittoPwd) < 1 {
				panic(errors.New("mosquitto-password must have 1 character minimum"))
			}

			mosquittoUrl, err := url.Parse(mosquittoUrlString)
			if err != nil {
				panic(err)
			}

			mosquittoConfig := mosquitto.NewConfig(*mosquittoUrl, mosquittoUser, []byte(mosquittoPwd), mosquittoTopic, mosquittoClientId)

			StartService(mosquittoConfig)
		},
	}

	cmd.Flags().StringVarP(&mosquittoUrlString, "mosquitto-url", "", "mqtt://192.168.1.21:1883", "MQTT Server url")
	cmd.Flags().StringVarP(&mosquittoUser, "mosquitto-user", "", "homeassistant", "MQTT Server user")
	cmd.Flags().StringVarP(&mosquittoPwd, "mosquitto-password", "", "", "MQTT Server password")
	cmd.Flags().StringVarP(&mosquittoClientId, "mosquitto-client-id", "", "Godomo.Lixee", "MQTT Server client id")
	cmd.Flags().StringVarP(&mosquittoTopic, "mosquitto-topic", "", "zigbee2mqtt/LinkyLixee", "MQTT Server url")

	rootCommand.AddCommand(cmd)
	rootCommand.Execute()

}
