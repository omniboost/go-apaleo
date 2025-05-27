package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetReservationsNSFW(t *testing.T) {
	client := client()
	req := client.NewGetReservationsNSFWRequest()
	req.QueryParams().ExternalReferences = []string{"AMZ-F65539", "AMZ-FX154905"}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetReservationsAllNSFW(t *testing.T) {
	client := client()
	req := client.NewGetReservationsNSFWRequest()
	req.QueryParams().PageSize = 1

	resp, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
