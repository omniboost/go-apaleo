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

func TestPostAggregateDaily(t *testing.T) {
	client := client()

	today := time.Now()
	yesterday := today.AddDate(0, 0, -9)

	req := client.NewPostAccountsAggregateDailyRequest()
	req.QueryParams().PropertyID = "BER"
	req.QueryParams().From = apaleo.Date{yesterday}
	req.QueryParams().To = apaleo.Date{today}
	// req.QueryParams().Reference = "1"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
