package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TesPostBookingReservations(t *testing.T) {
	client := client()
	req := client.NewPostBookingReservationsRequest()

	reservations := []apaleo.CreateReservationModel{}
	reservations = append(reservations, apaleo.CreateReservationModel{
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

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
