package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"gopkg.in/yaml.v3"
)

type RestClient struct {
	client    http.Client
	baseUrl   string
	timeout   time.Duration
	oAuthData *OAuthData
}

func New(baseUrl string, options ...func(*RestClient)) *RestClient {
	client := &RestClient{
		baseUrl: baseUrl,
		client:  *http.DefaultClient,
		timeout: 10 * time.Second,
	}
	for _, o := range options {
		o(client)
	}
	return client
}

func WithTimeout(timeout time.Duration) func(*RestClient) {
	return func(c *RestClient) {
		c.timeout = timeout
	}
}

func WithOAuth(oAuthData OAuthData) func(*RestClient) {
	oAuthData.RefreshBearer()
	return func(c *RestClient) {
		c.oAuthData = &oAuthData
	}
}

func (c *RestClient) Get(ctx context.Context, endpoint string) (*http.Response, error) {
	fullEndpoint, err := url.JoinPath(c.baseUrl, endpoint)
	if err != nil {
		return nil, fmt.Errorf("error building endpoint %q %q: %w", c.baseUrl, endpoint, err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fullEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error building http request %q: %w", fullEndpoint, err)
	}

	if c.oAuthData != nil {
		err = c.oAuthData.RefreshBearer()
		if err != nil {
			return nil, fmt.Errorf("error refreshing bearer %q: %w", c.baseUrl, err)
		}
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.oAuthData.State.AccessToken))
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error doing http request %q: %w", fullEndpoint, err)
	}

	if response.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if response.StatusCode > http.StatusBadRequest {
		return nil, fmt.Errorf("bad status %q", response.Status)
	}

	return response, nil
}

func (c *RestClient) PostJson(ctx context.Context, endpoint string, requestData interface{}) (*http.Response, error) {
	fullEndpoint, err := url.JoinPath(c.baseUrl, endpoint)
	if err != nil {
		return nil, fmt.Errorf("error building endpoint %q %q: %w", c.baseUrl, endpoint, err)
	}

	data, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling json %q %q: %w", c.baseUrl, endpoint, err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, fullEndpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error building http request %q: %w", fullEndpoint, err)
	}

	if c.oAuthData != nil {
		err = c.oAuthData.RefreshBearer()
		if err != nil {
			return nil, fmt.Errorf("error refreshing bearer %q: %w", c.baseUrl, err)
		}
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.oAuthData.State.AccessToken))
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error doing http request %q: %w", fullEndpoint, err)
	}

	if response.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if response.StatusCode > http.StatusBadRequest {
		return nil, fmt.Errorf("bad status %d %q", response.StatusCode, response.Status)
	}

	return response, nil
}

func JsonResponse[T any](resp *http.Response) (*T, error) {
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %w", err)
	}
	var jsonValue T
	err = json.Unmarshal(respData, &jsonValue)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling JSON error %s: %w", string(respData), err)
	}
	return &jsonValue, nil
}

func YamlResponse[T any](resp *http.Response) (*T, error) {
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %w", err)
	}
	var jsonValue T
	err = yaml.Unmarshal(respData, &jsonValue)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling YAML error %s: %w", string(respData), err)
	}
	return &jsonValue, nil
}
