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

func TestPutReservationActionsBookService(t *testing.T) {
	client := client()
	req := client.NewPutReservationActionsBookServiceForceRequest()

	// Set reservation ID - replace with actual reservation ID for testing
	req.PathParams().ReservationID = "ZIAANEMS-1"

	// Optionally set city tax
	body := client.NewPutReservationActionsBookServiceRequestBody()
	body.ServiceID = "MID2-FB_FNB_500"
	body.Count = 1
	body.Amount = apaleo.MonetaryValueModel{
		Amount:   100,
		Currency: "EUR",
	}

	body.Dates = append(body.Dates, struct {
		ServiceDate apaleo.Date               `json:"serviceDate"`
		Count       int32                     `json:"count,omitempty"`
		Amount      apaleo.MonetaryValueModel `json:"amount,omitempty"`
	}{
		ServiceDate: apaleo.Date{Time: time.Date(2025, 4, 29, 0, 0, 0, 0, time.UTC)},
		Count:       1,
		Amount: apaleo.MonetaryValueModel{
			Amount:   100,
			Currency: "EUR",
		},
	})

	req.SetRequestBody(body)

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
