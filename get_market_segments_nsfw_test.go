package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetMarketSegmentsNSFW(t *testing.T) {
	client := client()
	req := client.NewGetMarketSegmentsNSFWRequest()

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetMarketSegmentsNSFWAll(t *testing.T) {
	client := client()
	req := client.NewGetMarketSegmentsNSFWRequest()

	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
