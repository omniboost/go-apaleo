package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TestPostAggregatePairsDaily(t *testing.T) {
	client := client()

	today := time.Now()
	yesterday := today.AddDate(0, 0, -9)

	req := client.NewPostAccountsAggregatePairsDailyRequest()
	req.QueryParams().PropertyID = "BER"
	req.QueryParams().From = apaleo.Date{yesterday}
	req.QueryParams().To = apaleo.Date{today}
	// req.QueryParams().Reference = "1"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
