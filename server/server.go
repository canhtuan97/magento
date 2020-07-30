package main

import (
	"context"
	"fmt"
	"github.com/canhtuan97/magento/application/cart"
	"github.com/canhtuan97/magento/application/customer"
	"github.com/canhtuan97/magento/proto/cart"
	"github.com/canhtuan97/magento/proto/customer"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
)

type server struct{}

func (s *server) EstimateShipping(ctx context.Context, request *cartPb.EstimateShippingRequest) (*cartPb.EstimateShippingResponse, error) {
	data , err := cart.EstimateShipping(request)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *server) AddItemProductConfigurable(ctx context.Context, request *cartPb.AddItemProductConfigurableRequest) (*cartPb.AddItemProductConfigurableResponse, error) {
	data ,err := cart.AddItemProductConfigurable(request)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *server) GetQuoteIdCustomer(ctx context.Context, request *customerPb.GetQuoteIdCustomerRequest) (*customerPb.GetQuoteIdCustomerResponse, error) {
	panic("implement me")
}

func (s *server) GetAccessTokenCustomer(ctx context.Context, request *customerPb.GetAccessTokenCustomerRequest) (*customerPb.GetAccessTokenCustomerResponse, error) {
	accessToken, err := customer.GetAccessTokenCustomer(request)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(accessToken.Body)
	if err != nil {
		return nil, err
	}
	resp := &customerPb.GetAccessTokenCustomerResponse{
		AccessToken: string(data),
	}
	return resp, nil
}

func (s *server) AddItemProductSimple(ctx context.Context, request *cartPb.AddItemProductSimpleRequest) (*cartPb.AddItemProductSimpleResponse, error) {
	data, err := cart.AddItemProductSimple(request)
	if err != nil {
		return nil, err
	}

	resp := &cartPb.AddItemProductSimpleResponse{
		ItemId:      int32(data.ItemId),
		Sku:         data.Sku,
		Qty:         int32(data.Qty),
		Name:        data.Name,
		Price:       data.Price,
		ProductType: data.ProductType,
		QuoteId:     data.QuoteId,
	}
	return resp, nil
}

func (s *server) CreateCustomer(ctx context.Context, req *customerPb.CreateCustomerRequest) (*customerPb.CreateCustomerResponse, error) {
	log.Println("Create customer running...")
	data, err := customer.CreateCustomer(req)
	if err != nil {
		return nil, err
	}

	resp := &customerPb.CreateCustomerResponse{
		Id:                     int32(data.Id),
		GroupId:                int32(data.GroupId),
		DefaultBilling:         data.DefaultBilling,
		DefaultShipping:        data.DefaultShipping,
		CreatedAt:              data.CreatedAt,
		UpdatedAt:              data.UpdatedAt,
		CreatedIn:              data.CreatedIn,
		Email:                  data.Email,
		FirstName:              data.FirstName,
		LastName:               data.LastName,
		StoreId:                int32(data.StoreId),
		WebsiteId:              int32(data.WebsiteId),
		DisableAutoGroupChange: int32(data.DisableAutoGroupChange),
	}
	return resp, nil

}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50069")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	customerPb.RegisterCustomerServer(s, &server{})
	cartPb.RegisterAddItemProductServer(s, &server{})
	fmt.Println("Server running ...")

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
