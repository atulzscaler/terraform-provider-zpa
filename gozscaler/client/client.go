package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/SecurityGeekIO/terraform-provider-zpa/gozscaler"
	"github.com/google/go-querystring/query"
)

type AuthToken struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	//ExpiresIn   string `json:"expires_in"`
}

type Client struct {
	Config    *gozscaler.Config
	AuthToken *AuthToken
	Context   context.Context
}

// NewClient returns a new client for the specified apiKey.
func NewClient(config *gozscaler.Config) (c *Client) {
	if config == nil {
		config, _ = gozscaler.NewConfig("", "", "", "")
	}
	c = &Client{Config: config}
	return
}

func (client *Client) NewRequestDo(method, url string, options, body, v interface{}) (*http.Response, error) {
	if client.AuthToken == nil {
		formData := []byte(fmt.Sprintf("client_id=%s&client_secret=%s",
			client.Config.ClientID,
			client.Config.ClientSecret,
		))

		req, err := http.NewRequest("POST", client.Config.BaseURL.String()+"/signin", bytes.NewBuffer(formData))
		if err != nil {
			return nil, err
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Config.HTTPClient.Do(req)

		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var a AuthToken
		err = json.Unmarshal(respBody, &a)
		if err != nil {
			return nil, err
		}

		// we need keep auth token for future http request
		client.AuthToken = &a
	}

	req, err := client.newRequest(method, url, options, body)
	if err != nil {
		return nil, err
	}
	client.logRequest(req)
	return client.do(req, v)
}

// Generating the Http request
func (client *Client) newRequest(method, urlPath string, options, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// Join the path to the base-url
	u := *client.Config.BaseURL
	unescaped, err := url.PathUnescape(urlPath)
	if err != nil {
		return nil, err
	}

	// Set the encoded path data
	u.RawPath = client.Config.BaseURL.Path + urlPath
	u.Path = client.Config.BaseURL.Path + unescaped

	// Set the query parameters
	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// We are using JWT not Basic auth
	// req.SetBasicAuth(client.Config.ClientID, client.Config.ClientSecret)
	// req.Header.Add("Accept", "application/x-www-form-urlencoded")
	// if body != nil {
	// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// }

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.AuthToken.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Accept", "application/json")
	return req, nil
}

func (client *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := client.Config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// response is close before parsing response body due to  this function defer the body close
	// defer func() {
	// 	if rerr := resp.Body.Close(); err == nil {
	// 		err = rerr
	// 	}
	// }()

	if err := checkErrorInResponse(resp); err != nil {
		return resp, err
	}

	if v != nil {
		if err := decodeJSON(resp, v); err != nil {
			return resp, err
		}
	}
	client.logResponse(resp)

	return resp, nil
}

func decodeJSON(res *http.Response, v interface{}) error {
	return json.NewDecoder(res.Body).Decode(&v)
}
