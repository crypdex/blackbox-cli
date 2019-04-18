package blackbox

import (
	"fmt"

	"github.com/logrusorgru/aurora"
	resty "gopkg.in/resty.v1"
)

type Client struct {
	client *resty.Client
	debug  bool
}

// NewClient ...
func NewClient(url string, token string, debug bool) (*Client, error) {

	if url == "" {
		fmt.Println("Defaulting to crypdex.local")
		url = "http://crypdex.local"
	}

	errorResponse := new(ErrorResponse)
	client := resty.
		SetDebug(debug).
		SetHostURL(url).
		SetHeader("Accept", "application/json").
		SetAuthToken(token).
		SetHeader("User-Agent", "blackbox-cli").
		SetError(errorResponse).
		OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
			// Now you have access to Client and current Request object
			// manipulate it as per your need
			if debug {
				fmt.Println(aurora.Cyan(fmt.Sprintf("%s %s%s\n", req.Method, c.HostURL, req.URL)))
			}
			return nil // if its success otherwise return error
		})

	return &Client{client: client, debug: debug}, nil
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

type LoginResponse struct {
	JWT string `json:"jwt"`
}

func (c *Client) Login(password string, save bool) (*LoginResponse, error) {
	res := new(LoginResponse)

	response, err := c.client.R().
		SetBody(map[string]interface{}{
			"password": password,
			"save":     save,
		}).
		SetResult(res).
		Post("/login")

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) Logout() error {

	response, err := c.client.R().Post("/logout")

	return checkResponse(response, err)
}

func (c *Client) SystemStatus() (interface{}, error) {
	response, err := c.client.R().Get("/system/status")

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SystemUpdate() (interface{}, error) {
	response, err := c.client.R().Post("/system/update")

	if err := checkResponse(response, err); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) MasternodeList(chain string) (interface{}, error) {
	response, err := c.client.R().Get(fmt.Sprintf("/v1/%s/masternodes", chain))

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
