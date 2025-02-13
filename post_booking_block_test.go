package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TestPostBookingBlock(t *testing.T) {
	client := client()
	req := client.NewPostBookingBlockRequest()
	req.RequestBody().GroupID = "TEST"
	req.RequestBody().RatePlanID = "TEST"
	req.RequestBody().From = apaleo.Date{time.Now()}
	req.RequestBody().To = apaleo.Date{time.Now().Add(time.Hour * 24)}
	req.RequestBody().GrossDailyRate = apaleo.MonetaryValueModel{
		Amount:   1000,
		Currency: "EUR",
	}

	req.RequestBody().TimeSlices = []apaleo.CreateBlockTimeSliceModel{
		apaleo.CreateBlockTimeSliceModel{
			BlockedUnits: 5,
		},
	}

	req.RequestBody().BlockedUnits = 5
	req.RequestBody().PromoCode = ""
	req.RequestBody().CorporateCode = ""

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
