package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TestPostGroupBookingNSFW(t *testing.T) {
	client := client()
	req := client.NewPostBookingGroupNSFWRequest()
	req.RequestBody().Name = "API test"
	req.RequestBody().PaymentAccount = apaleo.CreatePaymentAccountModel{
		AccountNumber:  "123456789",
		AccountHolder:  "John Doe",
		ExpiryMonth:    "1",
		ExpiryYear:     "2025",
		PaymentMethod:  "CreditCard",
		PayerEmail:     "a@b.c",
		PayerReference: "12345678910",
		IsVirtual:      false,
	}

	req.RequestBody().Booker = apaleo.BookerModel{
		Title:     "Mr",
		Gender:    "Male",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "a@b.c",
		Phone:     "123456789",
		Address: apaleo.PersonAddressModel{
			AddressLine1: "123 Main St",
			AddressLine2: "Apt 1",
			PostalCode:   "12345",
			City:         "New York",
			CountryCode:  "US",
		},
	}

	req.RequestBody().PropertyIDs = []string{"UTR"}
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
