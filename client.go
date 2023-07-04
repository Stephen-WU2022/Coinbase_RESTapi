package Coinbase_RESTapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func (c *Client) newRequest(method, path string, body []byte) (*http.Request, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	u, _ := url.ParseRequestURI(Endpoint)
	u.Path = u.Path + path
	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("CB-ACCESS-KEY", c.key)
	req.Header.Add("CB-ACCESS-PASSPHRASE", c.passphase)
	req.Header.Add("CB-ACCESS-TIMESTAMP", timestamp)
	req.Header.Add("CB-ACCESS-SIGN", c.sign(timestamp, method, path, body))
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}
