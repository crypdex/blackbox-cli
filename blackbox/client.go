package blackbox

import (
	"fmt"
	"net"
	"net/url"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
)

type Client struct {
	client *resty.Client
}

// NewClient ...
func NewClient(config *viper.Viper) (*Client, error) {
	hostURL := config.GetString("url")

	if hostURL == "" {
		return nil, errors.New("missing address")
	}

	host, port, _ := net.SplitHostPort(hostURL)

	u := url.URL{Scheme: "http", Host: fmt.Sprintf("%s:%s", host, port)}

	errorResponse := new(ErrorResponse)
	client := resty.
		SetHostURL(u.String()).
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", "blackbox-cli").
		SetError(errorResponse)

	return &Client{client: client}, nil
}

func (c *Client) Init(request InitRequest) (*InitResponse, error) {
	result := new(InitResponse)

	response, err := c.client.R().
		SetBody(request).
		SetResult(result).
		Post("/initialize")

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) Login(password string) (interface{}, error) {
	response, err := c.client.R().
		SetBody(map[string]string{
			"password": password,
		}).
		Post("/login")

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) Logout() error {

	response, err := c.client.R().Post("/logout")

	return checkResponse(response, err)
}

func (c *Client) Status() (interface{}, error) {
	response, err := c.client.R().Get("/status")

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddressList(chain string) (interface{}, error) {
	response, err := c.client.R().Get(fmt.Sprintf("/v1/%s/addresses", chain))

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddressCreate(chain string, request CreateAddressRequest) (interface{}, error) {
	response, err := c.client.R().
		SetBody(request).
		Post(fmt.Sprintf("/v1/%s/addresses", chain))

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) AddressRecreate(chain string, count int) (interface{}, error) {
	response, err := c.client.R().
		SetBody(map[string]int{"count": count}).
		Put(fmt.Sprintf("/v1/%s/addresses", chain))

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return response, nil
}

func checkResponse(response *resty.Response, err error) error {
	if err != nil {
		return err
	}

	if response.Error() != nil {
		return response.Error().(*ErrorResponse)
	}
	return nil
}
