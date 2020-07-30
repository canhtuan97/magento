package connector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type CartItem struct {
	ItemId        int           `json:"item_id"`
	Sku           string        `json:"sku"`
	Qty           int           `json:"qty"`
	Name          string        `json:"name"`
	Price         float32       `json:"price"`
	ProductType   string        `json:"product_type"`
	QuoteId       string        `json:"quote_id"`
	ProductOption ProductOption `json:"productOption"`
}

type CartRequest struct {
	Sku     string `json:"sku"`
	Qty     int    `json:"qty"`
	QuoteId string `json:"quote_id"`
}

type ConfigurableItemOptions struct {
	OptionId    string `json:"option_id"`
	OptionValue int    `json:"option_value"`
}

type Address struct {
	Region        string   `json:"region"`
	RegionId      string   `json:"region_id"`
	RegionCode    string   `json:"region_code"`
	CountryId     string   `json:"country_id"`
	Street        []string `json:"street"`
	Postcode      string   `json:"postcode"`
	City          string   `json:"city"`
	FirstName     string   `json:"first_name"`
	LastName      string   `json:"lastname"`
	CustomerId    int      `json:"customer_id"`
	Email         string   `json:"email"`
	Telephone     string   `json:"telephone"`
	SameAsBilling int      `json:"same_as_billing"`
}

type EstimateShippingRequest struct {
	Address Address `json:"address"`
}

type DataResponse struct {
	CarrierCode  string `json:"carrier_code"`
	MethodCode   string `json:"method_code"`
	CarrierTitle string `json:"carrier_title"`
	MethodTitle  string `json:"method_title"`
	Amount       int    `json:"amount"`
	BaseAmount   int    `json:"base_amount"`
	Available    bool   `json:"available"`
	ErrorMessage string `json:"error_message"`
	PriceExclTax int    `json:"price_excl_tax"`
	PriceInclTax int    `json:"price_incl_tax"`
}

type EstimateShippingResponse struct {
	DataResponse []DataResponse `json:"Data"`
}

type ExtensionAttributes struct {
	ConfigurableIteOptions []ConfigurableItemOptions `json:"configurableItemOptions"`
}
type ProductOption struct {
	ExtensionAttributes ExtensionAttributes `json:"extensionAttributes"`
}
type CartAddProductConfigurableRequest struct {
	Sku           string        `json:"sku"`
	Qty           int           `json:"qty"`
	QuoteId       string        `json:"quote_id"`
	ProductOption ProductOption `json:"product_option"`
}

type AddItemSimpleRequest struct {
	CartItem CartRequest `json:"cartItem"`
}
type AddProductConfigurableRequest struct {
	CartItem CartAddProductConfigurableRequest `json:"cartItem"`
}
type Cart struct {
	client *Client
}
type CartServices interface {
	AddProductSimple(addItemSimpleRequest AddItemSimpleRequest) (*CartItem, error)
	AddProductConfigurable(addProductConfigurableRequest AddProductConfigurableRequest) (*CartItem, error)
	EstimateShipping(estimateShippingRequest EstimateShippingRequest) ([]*DataResponse, error)
}

func (c Cart) AddProductSimple(addItemSimpleRequest AddItemSimpleRequest) (*CartItem, error) {
	url := c.client.UrlMagento + urlAddItemToCart
	dataConvert, err := json.Marshal(addItemSimpleRequest)
	fmt.Println(string(url))
	fmt.Println("day la test ", string(dataConvert))
	if err != nil {
		return nil, err
	}
	resp, err := c.client.CreateRequestPost(url, dataConvert)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	cartItem := CartItem{}
	json.Unmarshal(data, &cartItem)
	fmt.Println(string(data))
	return &cartItem, nil
}

func (c Cart) AddProductConfigurable(addProductConfigurableRequest AddProductConfigurableRequest) (*CartItem, error) {
	url := c.client.UrlMagento + urlAddItemToCart
	dataConvert, err := json.Marshal(addProductConfigurableRequest)
	fmt.Println(string(url))
	fmt.Println("day la test ", string(dataConvert))
	if err != nil {
		return nil, err
	}
	resp, err := c.client.CreateRequestPost(url, dataConvert)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	cartItem := CartItem{}
	json.Unmarshal(data, &cartItem)
	fmt.Println(string(data))
	return &cartItem, nil
}

func (c Cart) EstimateShipping(estimateShippingRequest EstimateShippingRequest) ([]*DataResponse, error) {
	url := c.client.UrlMagento + urlEstimateShipping
	dataConvert, err := json.Marshal(estimateShippingRequest)
	fmt.Println(string(url))
	fmt.Println("day la test ", string(dataConvert))

	resp, err := c.client.CreateRequestPost(url, dataConvert)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	var dataResponse []*DataResponse
	json.Unmarshal(data, &dataResponse)
	fmt.Println(string(data))
	return dataResponse, nil
}
