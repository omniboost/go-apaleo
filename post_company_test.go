package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TestPostCompany(t *testing.T) {
	client := client()
	req := client.NewPostCompanyRequest()
	req.RequestBody().Code = "OMNIBOOST"
	req.RequestBody().Name = "Omniboost"
	req.RequestBody().PropertyID = "OMNI"
	req.RequestBody().TaxID = "123456789"
	req.RequestBody().AdditionalTaxID = "987654321"
	req.RequestBody().CanCheckOutOnAr = true
	req.RequestBody().Address = apaleo.CompanyAddressModel{
		AddressLine1: "123 Main St",
		AddressLine2: "Apt 1",
		PostalCode:   "12345",
		City:         "New York",
		CountryCode:  "US",
	}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
