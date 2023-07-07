package Coinbase_RESTapi

import (
	"fmt"
	"net/http"
)

type InstrumentResponse struct {
	Instrumentid       string  `json:"instrument_id"`
	Instrumentuuid     string  `json:"instrument_uuid"`
	Symbol             string  `json:"symbol"`
	Type               string  `json:"type"`
	Baseassetid        string  `json:"base_asset_id"`
	Baseassetuuid      string  `json:"base_asset_uuid"`
	Baseassetname      string  `json:"base_asset_name"`
	Quoteassetid       string  `json:"quote_asset_id"`
	Quoteassetuuid     string  `json:"quote_asset_uuid"`
	Quoteassetname     string  `json:"quote_asset_name"`
	Baseincrement      string  `json:"base_increment"`
	Quoteincrement     string  `json:"quote_increment"`
	Marketorderpercent float32 `json:"market_order_percent"`
	Pricebandpercent   float32 `json:"price_band_percent"`
	Qty24hr            string  `json:"qty_24hr"`
	Notional24hr       string  `json:"notional_24hr"`
	Avgdailyqty        string  `json:"avg_daily_qty"`
	Avgdailynotional   string  `json:"avg_daily_notional"`
	Previousdayqty     string  `json:"previous_day_qty"`
	Positionlimitqty   string  `json:"position_limit_qty"`
	Positionlimitadv   float32 `json:"position_limit_adv"`
	Initialmarginadv   string  `json:"initial_margin_adv"`
	Replacementcost    string  `json:"replacement_cost"`
	Baseimf            float32 `json:"base_imf"`
	Minnotionalvalue   string  `json:"min_notional_value"`
	Fundinginterval    string  `json:"funding_interval"`
	Tradingstate       string  `json:"trading_state"`
}

type QuteperInstrumentResponse struct {
	Bestbidprice     string `json:"best_bid_price"`
	Bestbidsize      string `json:"best_bid_size"`
	Bestaskprice     string `json:"best_ask_price"`
	Bestasksize      string `json:"best_ask_size"`
	Tradeprice       string `json:"trade_price"`
	Tradeqty         string `json:"trade_qty"`
	Indexprice       string `json:"index_price"`
	Markprice        string `json:"mark_price"`
	Settlementprice  string `json:"settlement_price"`
	Limitup          string `json:"limit_up"`
	Limitdown        string `json:"limit_down"`
	Predictedfunding string `json:"predicted_funding"`
	Timestamp        string `json:"timestamp"`
}

func (c *Client) Instruments() (instruments []*InstrumentResponse, err error) {
	path := "/api/v1/instruments"
	resp, err := c.sendRequest(http.MethodGet, path, nil, false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &instruments)
	if err != nil {
		return nil, err
	}
	return instruments, nil
}

// Instrument symbol accepts symbol, instrument_id, instrument_uuid
func (c *Client) Instrument(symbol string) (instrument *InstrumentResponse, err error) {
	path := fmt.Sprintf("/api/v1/instruments/%s", symbol)
	resp, err := c.sendRequest(http.MethodGet, path, nil, false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &instrument)
	if err != nil {
		return nil, err
	}
	return instrument, nil
}

// QuteperInstrument symbol accepts symbol, instrument_id, instrument_uuid
func (c *Client) QuteperInstrument(symbol string) (quteperinstrument *QuteperInstrumentResponse, err error) {
	path := fmt.Sprintf("/api/v1/instruments/%s/quote", symbol)
	resp, err := c.sendRequest(http.MethodGet, path, nil, false)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &quteperinstrument)
	if err != nil {
		return nil, err
	}
	return quteperinstrument, nil
}
