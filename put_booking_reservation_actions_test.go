package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestPutBookingReservationActions(t *testing.T) {
	client := client()
	req := client.NewPutBookingReservationActionsRequest()
	req.PathParams().ReservationID = "UTR-USJTES"
	req.PathParams().Action = "cancel"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
