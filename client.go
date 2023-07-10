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

// Client is a class for interacting with the Coinbase international exchange restAPI, subaccount for now empty
type Client struct {
	key, secret, passphase string
	subaccount             string
	client                 *http.Client //http client
}

// Create a client struct
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

// form a request, sign it if needed, sign for now empty
func (c *Client) newRequest(method, path string, body []byte, sign bool) (*http.Request, error) {

	u, _ := url.ParseRequestURI(Endpoint) // parse endpoint and form a url
	u.Path = u.Path + path                // add parm to url
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

// send request and get response
func (c *Client) sendRequest(method, path string, body []byte, sign bool) (response []byte, err error) {
	req, err := c.newRequest(method, path, body, sign)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req) // if response is error message and err is nil, may not trigger the defer func
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() // close need to be done before error(close the do func), otherwise it would be ignored
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}
	response, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return response, nil
}
