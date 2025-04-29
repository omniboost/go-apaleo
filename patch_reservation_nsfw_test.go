package apaleo_test

import (
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
	req.PathParams().ReservationID = "ZPVKPNGN-2"

	// Optionally set city tax
	body := client.NewPatchReservationNSFWRequestBody()
	body = append(body, apaleo.Operation{
		Operation: "replace",
		Path:      "/isOpenForCharges",
		Value:     true,
	})

	req.SetRequestBody(body)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
