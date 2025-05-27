package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetReservations(t *testing.T) {
	client := client()
	req := client.NewGetReservationsRequest()
	req.QueryParams().BookingID = "LRNQSMTM"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetReservationsAll(t *testing.T) {
	client := client()
	req := client.NewGetReservationsRequest()
	req.QueryParams().PageSize = 1

	resp, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
