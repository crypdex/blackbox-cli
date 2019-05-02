package blackbox

func (c *Client) Status() (interface{}, error) {

	response, err := c.client.R().Get("/status")

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return string(response.Body()), nil
}

type PivxStatus struct {
	Blockchain struct {
		Balance       float64 `json:"balance"`
		StakingStatus string  `json:"staking status" mapstructure:"staking status"`
	} `json:"blockchain"`
	SyncProgress string `json:"sync_progress" mapstructure:"sync_progress"`
}

// "blockchain": {
// "version": 3020000,
// "protocolverion": 0,
// "walletversion": 61000,
// "balance": 13346.75256383,
// "zerocoinbalance": 0,
// "blocks": 1753817,
// "timeoffset": 0,
// "connections": 10,
// "proxy": "",
// "difficulty": 165142.8767528207,
// "testnet": false,
// "moneysupply": 59432665.5323948,
// "keypoololdest": 1550026205,
// "keypoolsize": 1001,
// "paytxfee": 0,
// "relayfee": 0.0001,
// "staking status": "Staking Active",
// "errors": ""
// },
