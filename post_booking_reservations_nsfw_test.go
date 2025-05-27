package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TestPostBookingReservationsNSFW(t *testing.T) {
	client := client()
	req := client.NewPostBookingReservationsNSFWRequest()

	reservations := []apaleo.CreateReservationNSFWModel{}
	reservations = append(reservations, apaleo.CreateReservationNSFWModel{
		Arrival:      apaleo.Date{time.Now()},
		Departure:    apaleo.Date{time.Now().Add(time.Hour * 24)},
		Adults:       1,
		ChildrenAges: []int32{5},
		Comment:      "Don't provide help",
		GuestComment: "Please provide help",
		ExternalCode: "",
		ChannelCode:  "Direct",
		PrimaryGuest: apaleo.GuestModel{
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
		},
		GuaranteeType: "CreditCard",
		TravelPurpose: "Business",
		TimeSlices: []apaleo.CreateReservationTimeSliceModel{
			{
				RatePlanID: "OMNI-RP_OVER",
				TotalAmount: apaleo.MonetaryValueModel{
					Currency: "USD",
					Amount:   1000,
				},
			},
		},
	})

	req.RequestBody().Reservations = reservations

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
