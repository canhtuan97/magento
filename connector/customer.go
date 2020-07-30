package connector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Customer struct {
	Id                     int         `json:"id"`
	GroupId                int         `json:"group_id"`
	DefaultBilling         string      `json:"default_billing"`
	DefaultShipping        string      `json:"default_shipping"`
	CreatedAt              string      `json:"created_at"`
	UpdatedAt              string      `json:"updated_at"`
	CreatedIn              string      `json:"created_in"`
	Email                  string      `json:"email"`
	FirstName              string      `json:"firstname"`
	LastName               string      `json:"lastname"`
	StoreId                int         `json:"store_id"`
	WebsiteId              int         `json:"website_id"`
	Addresses              []Addresses `json:"addresses"`
	DisableAutoGroupChange int         `json:"disable_auto_group_change"`
}

type Addresses struct {
	Id         int      `json:"id"`
	CustomerId int      `json:"customer_id"`
	Region     Region   `json:"region"`
	RegionId   int      `json:"region_id"`
	CountryId  string   `json:"countryId"`
	Street     []string `json:"street"`
	Telephone  string   `json:"telephone"`
	Postcode   string   `json:"postcode"`
	City       string   `json:"city"`
	FirstName  string   `json:"firstname"`
	LastName   string   `json:"lastname"`

	DefaultShipping bool `json:"defaultShipping"`
	DefaultBilling  bool `json:"defaultBilling"`
}

type Region struct {
	RegionCode string `json:"regionCode"`
	Region     string `json:"region"`
	RegionId   int    `json:"regionId"`
}

type CreateCustomerRequest struct {
	Customer Customer `json:"customer"`
	Password string   `json:"password"`
}

type CustomersService interface {
	CreateCustomer(createCustomerRequest CreateCustomerRequest) (*Customer, error)
}

type Customers struct {
	client *Client
}

func (c Customers) CreateCustomer(createCustomerRequest CreateCustomerRequest) (*Customer, error) {

	url := c.client.UrlMagento + urlCreateCustomer
	dataConvert, err := json.Marshal(createCustomerRequest)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.CreateRequestPost(url, dataConvert)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	customer := Customer{}
	//-------------Đoạn này đang viết check lỗi
	//data, err := CheckResponse(resp)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//
	//if err == nil {
	//	log.Fatal(err)
	//}
	//--------------------------------------
	json.Unmarshal(data, &customer)
	fmt.Println(string(data))
	return &customer, nil

}
