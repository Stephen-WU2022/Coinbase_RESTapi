package cbintxapi

import (
	"fmt"
	"net/http"
)

type AssetResponse struct {
	Assetid          string  `json:"asset_id"`
	Assetuuid        string  `json:"asset_uuid"`
	Assetname        string  `json:"asset_name"`
	Status           string  `json:"status"`
	Collateralweight float32 `json:"collateral_weight"`
}

type NetworksperAssetResponse struct {
	Assetid             string `json:"asset_id"`
	Assetuuid           string `json:"asset_uuid"`
	Assetname           string `json:"asset_name"`
	Isdefault           bool   `json:"is_default"`
	Networkname         string `json:"network_name"`
	Displayname         string `json:"display_name"`
	Networkarnid        string `json:"network_arn_id"`
	Minwithdrawalamount string `json:"min_withdrawal_amt"`
	Maxwithdrawalamount string `json:"max_withdrawal_amt"`
	Networkconfirms     int    `json:"network_confirms"`
	Processingtime      int    `json:"processing_time"`
}

func (c *Client) Assets() (assets []*AssetResponse, err error) {
	path := "/api/v1/assets"
	resp, err := c.sendRequest(http.MethodGet, path, nil, false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &assets)
	if err != nil {
		return nil, err
	}
	return assets, nil
}

// Asset symbol accepts asset_name, asset_id, asset_uuid
func (c *Client) Asset(symbol string) (asset *AssetResponse, err error) {
	path := fmt.Sprintf("/api/v1/assets/%s", symbol)
	resp, err := c.sendRequest(http.MethodGet, path, nil, false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &asset)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

// NetworksperAsset symbol accepts asset_name, asset_id, asset_uuid
func (c *Client) NetworksperAsset(symbol string) (networks []*NetworksperAssetResponse, err error) {
	path := fmt.Sprintf("/api/v1/assets/%s/networks", symbol)
	resp, err := c.sendRequest(http.MethodGet, path, nil, false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &networks)
	if err != nil {
		return nil, err
	}
	return networks, nil
}
