package cbintxapi

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const Endpoint = "https://api.international.coinbase.com"

type Client struct {
	key, secret, passphase string
	subaccount             string
	client                 *http.Client
}

func NewClient(key, secret, passphase string) *Client {
	hc := &http.Client{
		Timeout: 10 * time.Second,
	}
	return &Client{
		key:       key,
		secret:    secret,
		passphase: passphase,
		client:    hc,
	}
}

func (c *Client) newRequest(method, path string, body []byte, sign bool) (*http.Request, error) {

	u, _ := url.ParseRequestURI(Endpoint)
	u.Path = u.Path + path
	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	if sign {
		//timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		//req.Header.Add("CB-ACCESS-KEY", c.key)
		//req.Header.Add("CB-ACCESS-PASSPHRASE", c.passphase)
		//req.Header.Add("CB-ACCESS-TIMESTAMP", timestamp)
		//req.Header.Add("CB-ACCESS-SIGN", c.sign(timestamp, method, path, body))
		//req.Header.Add("Content-Type", "application/json")
	}
	return req, nil
}

func (c *Client) sendRequest(method, path string, body []byte, sign bool) (response []byte, err error) {
	req, err := c.newRequest(method, path, body, sign)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}
	response, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return response, nil
}
