package blackbox

type UpgradeStatus struct {
	Upgradeable bool   `json:"updateable"`
	Installed   string `json:"installed"`
	Candidate   string `json:"candidate"`
}

func (c *Client) UpgradeStatus() (*UpgradeStatus, error) {
	result := new(UpgradeStatus)
	response, err := c.adminClient.R().SetResult(result).Get("info")
	if err := checkResponse(response, err); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Upgrade() (interface{}, error) {
	response, err := c.adminClient.R().Post("upgrade")

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}
	return response, nil
}
