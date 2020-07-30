package customer

import (
	"encoding/json"
	"fmt"
	"github.com/canhtuan97/magentos/connector"
	"github.com/canhtuan97/magentos/proto/customer"
	"log"
	"net/http"
)

func CreateCustomer(request *customerPb.CreateCustomerRequest) (*connector.Customer, error) {
	var pbAddresses []connector.Addresses
	for _, requestedAddress := range request.Customer.Addresses {
		pbAddresses = append(pbAddresses, connector.Addresses{
			DefaultShipping: requestedAddress.DefaultShipping,
			DefaultBilling:  requestedAddress.DefaultBilling,
			FirstName:       requestedAddress.FirstName,
			LastName:        requestedAddress.LastName,
			Region: connector.Region{
				RegionCode: requestedAddress.Region.RegionCode,
				Region:     requestedAddress.Region.Region,
				RegionId:   int(requestedAddress.Region.RegionId),
			},
			Postcode:  requestedAddress.Postcode,
			Street:    requestedAddress.Street,
			City:      requestedAddress.City,
			Telephone: requestedAddress.Telephone,
			CountryId: requestedAddress.CountryId,
		})
	}

	var customer connector.Customer
	customer.Email = request.Customer.Email
	customer.FirstName = request.Customer.FirstName
	customer.LastName = request.Customer.LastName
	customer.Addresses = pbAddresses

	createCustomerRequest := connector.CreateCustomerRequest{
		Customer: customer,
		Password: request.Password,
	}

	client := connector.NewClient()
	data, err := client.Customers.CreateCustomer(createCustomerRequest)
	if err != nil {
		log.Fatalf(" loi cua minh%v", err)
	}
	return data, nil
}

func GetAccessTokenCustomer(request *customerPb.GetAccessTokenCustomerRequest) (*http.Response , error) {
	client := connector.NewClient()
	url := client.UrlMagento + connector.UrlGetAccessToken
	fmt.Println(url)
	login := Login{
		username: request.Username,
		password:  request.Password,
	}
	//login.username = request.UserName
	//login.username = request.Password


	dataConvert, err1 := json.Marshal(login)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(string(login.username))
	fmt.Println(string(login.password))
	fmt.Println(string(request.Password))
	fmt.Println(dataConvert)
	fmt.Println(string(dataConvert))
	//--------- đang k hiểu sao cái dataConvert null


	resp ,err := client.CreateRequestPost(url,dataConvert)
	if err != nil {
		log.Fatal(err)
	}
	return resp, nil
}

func GetQuoteIdCustomer(request *customerPb.GetQuoteIdCustomerRequest) (*customerPb.GetQuoteIdCustomerResponse,error)  {
	panic("implement me")
}

