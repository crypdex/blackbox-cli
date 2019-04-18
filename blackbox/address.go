package blackbox

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Balance struct {
	Available decimal.Decimal `json:"available"`
	Pending   decimal.Decimal `json:"pending"`
	Locked    decimal.Decimal `json:"locked"`
}

type Address struct {
	PublicKey string  `json:"public_key"`
	Balance   Balance `json:"balance"`
}

func (c *Client) AddressList(chain string) ([]Address, error) {
	result := new([]Address)
	response, err := c.client.R().SetResult(result).Get(fmt.Sprintf("/v1/%s/addresses", chain))

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}
	return *result, nil
}

func (c *Client) AddressCreate(chain string, request CreateAddressRequest) (*Address, error) {
	result := new(Address)
	response, err := c.client.R().
		SetResult(result).
		SetBody(request).
		Post(fmt.Sprintf("/v1/%s/addresses", chain))

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) AddressRecreate(chain string, count int) (map[string]string, error) {
	result := new(map[string]string)
	response, err := c.client.R().
		SetResult(result).
		SetBody(map[string]int{"count": count}).
		Put(fmt.Sprintf("/v1/%s/addresses", chain))

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return *result, nil
}
