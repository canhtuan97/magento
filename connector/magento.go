package connector

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const (
	apiKey              = "71qhxcusonwayh5jeynhupjci7v2ecob"
	urlMagento          = "https://magento23demo.connectpos.com"
	urlCreateCustomer   = "/rest/V1/customers"
	UrlGetAccessToken   = "/rest/V1/integration/customer/token"
	urlAddItemToCart    = "/rest/default/V1/carts/mine/items"
	urlEstimateShipping = "/rest/V1/carts/19333/estimate-shipping-methods"
	quoteId             = 19223
	dasd                = 19332
)

type Client struct {
	Client     *http.Client
	ApiKey     string
	UrlMagento string
	Customers  CustomersService
	Carts      CartServices
}

func NewClient() *Client {
	httpClient := http.DefaultClient
	c := &Client{Client: httpClient}
	c.ApiKey = apiKey
	c.ApiKey = apiKey
	c.UrlMagento = urlMagento
	c.Customers = &Customers{client: c}
	c.Carts = &Cart{client: c}

	return c
}
func (c Client) CreateRequestPost(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	bearer := "bearer " + apiKey
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil

}

func CheckResponse(res *http.Response) ([]byte, error) {
	if res.Status == "400" {
		return nil, nil
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
