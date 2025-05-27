package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TestPatchReservationNSFW(t *testing.T) {
	client := client()
	req := client.NewPatchReservationNSFWRequest()

	// Set reservation ID - replace with actual reservation ID for testing
	req.PathParams().ReservationID = "VYYBJNWZ-1"

	// Optionally set city tax
	body := client.NewPatchReservationNSFWRequestBody()
	body = append(body, apaleo.Operation{
		Operation: "replace",
		Path:      "/isOpenForCharges",
		Value:     false,
	})

	body = append(body, apaleo.Operation{
		Operation: "add",
		Path:      "/marketSegment",
		Value: struct {
			ID string `json:"id"`
		}{ID: "BAR"},
	})

	req.SetRequestBody(body)

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
