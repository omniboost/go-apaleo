package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetBookingByID(t *testing.T) {
	client := client()
	req := client.NewGetBookingByIDRequest()
	req.PathParams().ID = "CXFWEVVO"
	req.QueryParams().Expand = []string{"reservations"}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
