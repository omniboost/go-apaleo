package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestPostExportGrossDaily(t *testing.T) {
	client := client()

	// today := time.Now()
	// yesterday := today.AddDate(0, 0, -9)

	req := client.NewPostExportGrossDailyRequest()
	req.QueryParams().PropertyID = "BER"
	req.QueryParams().From = "2024-01-01"
	req.QueryParams().To = "2024-01-29"
	// req.QueryParams().Reference = "1"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
