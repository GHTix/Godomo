package rest

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"
)

type OAuthState struct {
	AccessToken  string
	RefreshToken string
	ValidityTime time.Time
}

type OAuthConfig struct {
	LoginEndPoint string
	ClientId      string
	ClientSecret  string
	UserName      string
	Password      string
}

type OAuthData struct {
	Config OAuthConfig
	State  OAuthState
}

func NewOAuthData(oAuthConfig OAuthConfig) OAuthData {
	return OAuthData{
		Config: OAuthConfig{
			LoginEndPoint: oAuthConfig.LoginEndPoint,
			ClientId:      oAuthConfig.ClientId,
			ClientSecret:  oAuthConfig.ClientSecret,
			UserName:      oAuthConfig.UserName,
			Password:      oAuthConfig.Password,
		},
		State: newOAuthStateReset(),
	}
}

type OAuthResponseToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

func newOAuthStateReset() OAuthState {
	return OAuthState{
		AccessToken:  "",
		RefreshToken: "",
		ValidityTime: time.Now().Add(-24 * time.Hour),
	}
}

func newOAuthState(oAuthResponseToken OAuthResponseToken) OAuthState {
	validityDuration := time.Duration(int64(float64(oAuthResponseToken.ExpiresIn)*0.05)) * time.Second
	return OAuthState{
		AccessToken:  oAuthResponseToken.AccessToken,
		RefreshToken: oAuthResponseToken.RefreshToken,
		ValidityTime: time.Now().Add(validityDuration),
	}
}

func (o *OAuthData) RefreshBearer() error {
	slog.Info("GetAuthorizationBearer Begin", "Validity", o.State.ValidityTime, "AccessToken", o.State.AccessToken, "RefreshToken", o.State.RefreshToken)
	if time.Now().After(o.State.ValidityTime) {
		data := url.Values{}
		data.Set("client_id", o.Config.ClientId)
		data.Set("client_secret", o.Config.ClientSecret)

		if o.State.RefreshToken != "" {
			slog.Info("GetAuthorizationBearer", "Use Refresh Token", o.State.ValidityTime)
			data.Set("grant_type", "refresh_token")
			data.Set("refresh_token", o.State.RefreshToken)
		} else {
			slog.Info("GetAuthorizationBearer", "Get Access Token", o.State.ValidityTime)
			data.Set("grant_type", "password")
			data.Set("username", o.Config.UserName)
			data.Set("password", o.Config.Password)
		}
		o.State = newOAuthStateReset()

		resp, err := http.DefaultClient.PostForm(o.Config.LoginEndPoint, data)
		if err != nil {
			return fmt.Errorf("error retrieving access_token: %w", err)
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("bad status %q", resp.Status)
		}
		res, err := JsonResponse[OAuthResponseToken](resp)
		if err != nil {
			return fmt.Errorf("error decoding access_token: %w", err)
		}
		o.State = newOAuthState(*res)
	}
	slog.Info("GetAuthorizationBearer End", "Validity", o.State.ValidityTime, "AccessToken", o.State.AccessToken, "RefreshToken", o.State.RefreshToken)
	return nil
}
