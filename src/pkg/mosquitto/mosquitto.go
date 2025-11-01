package mosquitto

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
)

type MosquittoClient struct {
	config            *Config
	Router            *paho.StandardRouter
	ConnectionManager *autopaho.ConnectionManager
}

func New(mosquittoConfig *Config) *MosquittoClient {
	return &MosquittoClient{
		config: mosquittoConfig,
	}
}

func (mosquittoClient *MosquittoClient) Start(context context.Context, handler paho.MessageHandler) error {
	mosquittoClient.Router = paho.NewStandardRouter()

	cliCfg := autopaho.ClientConfig{
		ConnectUsername: mosquittoClient.config.User,
		ConnectPassword: mosquittoClient.config.Pwd,
		ServerUrls:      []*url.URL{&mosquittoClient.config.Url},
		KeepAlive:       20, // Keepalive message should be sent every 20 seconds
		// CleanStartOnInitialConnection defaults to false. Setting this to true will clear the session on the first connection.
		CleanStartOnInitialConnection: false,
		// SessionExpiryInterval - Seconds that a session will survive after disconnection.
		// It is important to set this because otherwise, any queued messages will be lost if the connection drops and
		// the server will not queue messages while it is down. The specific setting will depend upon your needs
		// (60 = 1 minute, 3600 = 1 hour, 86400 = one day, 0xFFFFFFFE = 136 years, 0xFFFFFFFF = don't expire)
		SessionExpiryInterval: 60,
		OnConnectionUp: func(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
			slog.Info("Mqtt connection up", slog.String("hostname", mosquittoClient.config.Url.Hostname()))
			// Subscribing in the OnConnectionUp callback is recommended (ensures the subscription is reestablished if
			// the connection drops)
			if _, err := cm.Subscribe(context, &paho.Subscribe{
				Subscriptions: []paho.SubscribeOptions{
					{Topic: mosquittoClient.config.Topic, QoS: 1},
				},
			}); err != nil {
				fmt.Printf("failed to subscribe (%s). This is likely to mean no messages will be received.", err)
			}
			slog.Info("Mqtt subscription made", slog.String("topic", mosquittoClient.config.Topic))
		},
		OnConnectError: func(err error) { fmt.Printf("error whilst attempting connection: %s\n", err) },
		// eclipse/paho.golang/paho provides base mqtt functionality, the below config will be passed in for each connection
		ClientConfig: paho.ClientConfig{
			// If you are using QOS 1/2, then it's important to specify a client id (which must be unique)
			ClientID: mosquittoClient.config.ClientId,
			// OnPublishReceived is a slice of functions that will be called when a message is received.
			// You can write the function(s) yourself or use the supplied Router
			OnPublishReceived: []func(paho.PublishReceived) (bool, error){
				func(pr paho.PublishReceived) (bool, error) {
					mosquittoClient.Router.Route(pr.Packet.Packet())
					return true, nil // we assume that the router handles all messages (todo: amend router API)
				}},
			OnClientError: func(err error) { fmt.Printf("client error: %s\n", err) },
			OnServerDisconnect: func(d *paho.Disconnect) {
				slog.Info("c.Disconnecting")
				if d.Properties != nil {
					fmt.Printf("server requested disconnect: %s\n", d.Properties.ReasonString)
				} else {
					fmt.Printf("server requested disconnect; reason code: %d\n", d.ReasonCode)
				}
			},
		},
	}

	mosquittoClient.Router.DefaultHandler(func(p *paho.Publish) { fmt.Printf("defaulthandler received message with topic: %s\n", p.Topic) })
	mosquittoClient.Router.RegisterHandler(mosquittoClient.config.Topic, handler)

	slog.Info("autopaho.NewConnection")
	c, err := autopaho.NewConnection(context, cliCfg)
	if err != nil {
		return err
	}
	mosquittoClient.ConnectionManager = c

	if err = mosquittoClient.ConnectionManager.AwaitConnection(context); err != nil {
		return err
	}

	return nil
}

func (mosquittoClient *MosquittoClient) Stop(context context.Context) error {
	slog.Info("mqttConnection.UnregisterHandler")
	mosquittoClient.Router.UnregisterHandler(mosquittoClient.config.Topic)
	slog.Info("mqttConnection.Disconnect")
	err := mosquittoClient.ConnectionManager.Disconnect(context)
	if err != nil {
		slog.Error("mqttConnection.Disconnect", slog.String("err", err.Error()))
	}
	return nil
}
