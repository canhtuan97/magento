package cart

import (
	"github.com/canhtuan97/magentos/connector"
	"github.com/canhtuan97/magentos/proto/cart"
	"log"
)

func AddItemProductSimple(request *cartPb.AddItemProductSimpleRequest) (*connector.CartItem, error) {
	addItemSimpleRequest := connector.AddItemSimpleRequest{
		CartItem: connector.CartRequest{
			Sku:     request.CartItem.Sku,
			Qty:     int(request.CartItem.Qty),
			QuoteId: request.CartItem.QuoteId,
		},
	}
	client := connector.NewClient()

	data, err := client.Carts.AddProductSimple(addItemSimpleRequest)
	if err != nil {
		log.Fatal(err)
	}
	return data, nil
}

func AddItemProductConfigurable(request *cartPb.AddItemProductConfigurableRequest) (*cartPb.AddItemProductConfigurableResponse, error) {
	var pbConfigurableItemOptions []connector.ConfigurableItemOptions
	for _, requestConfigurableItemOptions := range request.CartItem.ProductOption.ExtensionAttributes.ConfigurableItemOptions {
		pbConfigurableItemOptions = append(pbConfigurableItemOptions, connector.ConfigurableItemOptions{
			OptionId:    requestConfigurableItemOptions.OptionId,
			OptionValue: int(requestConfigurableItemOptions.OptionValue),
		})
	}

	addProductConfigurableRequest := connector.AddProductConfigurableRequest{
		CartItem: connector.CartAddProductConfigurableRequest{
			Sku:           request.CartItem.Sku,
			Qty:           int(request.CartItem.Qty),
			QuoteId:       request.CartItem.QuoteId,
			ProductOption: connector.ProductOption{
				ExtensionAttributes: connector.ExtensionAttributes{
					ConfigurableIteOptions: pbConfigurableItemOptions,
				},
			},
		},
	}

	client := connector.NewClient()
	data , err := client.Carts.AddProductConfigurable(addProductConfigurableRequest)
	if err != nil {
		log.Fatalf(" loi cua minh%v", err)
	}

	resp := &cartPb.AddItemProductConfigurableResponse{
		ItemId:        int32(data.ItemId),
		Sku:           data.Sku,
		Qty:           int32(data.Qty),
		Name:          data.Name,
		Price:         int32(data.Price),
		ProductType:   data.ProductType,
		QuoteId:       data.QuoteId,

	}
	return resp , nil
}


func EstimateShipping(request *cartPb.EstimateShippingRequest) (*cartPb.EstimateShippingResponse,error)  {
	estimateShippingRequest := connector.EstimateShippingRequest{
		Address: connector.Address{
			Region:        request.Address.Region,
			RegionId:      request.Address.RegionId,
			RegionCode:    request.Address.RegionCode,
			CountryId:     request.Address.CountryId,
			Street:        request.Address.Street,
			Postcode:      request.Address.Postcode,
			City:          request.Address.City,
			FirstName:     request.Address.FirstName,
			LastName:      request.Address.LastName,
			CustomerId:    int(request.Address.CustomerId),
			Email:         request.Address.Email,
			Telephone:     request.Address.Telephone,
			SameAsBilling: 0,
		},
	}

	
	client := connector.NewClient()
	
	
	data , err := client.Carts.EstimateShipping(estimateShippingRequest)
	if err != nil {
		log.Fatalf(" loi cua minh%v", err)
	}

	resp := prepareDataToResponse(data)
	

	return &cartPb.EstimateShippingResponse{
		Data: resp,
	}, err
}

func prepareDataToResponse(items []*connector.DataResponse) []*cartPb.Data {
	dataResponse := make([]*cartPb.Data, len(items))

	for idx, item := range items {
		data := &cartPb.Data{
			CarrierCode:  item.CarrierCode,
			MethodCode:   item.MethodCode,
			CarrierTitle: item.CarrierTitle,
			MethodTitle:  item.MethodTitle,
			Amount:       int32(item.Amount),
			BaseAmount:   int32(item.BaseAmount),
			Available:    item.Available,
			ErrorMessage: item.ErrorMessage,
			PriceExclTax: int32(item.PriceExclTax),
			PriceInclTax: int32(item.PriceInclTax),
		}

		dataResponse[idx] = data
 	}

 	return dataResponse
}