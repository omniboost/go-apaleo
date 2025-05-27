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

func TestPostDepositItemToFolio(t *testing.T) {
	client := client()
	req := client.NewPostDepositItemToFolioRequest()
	req.PathParams().ID = "BOICRCXO-MID2-1"

	body := req.NewRequestBody()
	body.Name = map[string]string{
		"en": "Deposit",
	}
	body.ServiceDate = &apaleo.Date{
		Time: time.Now(),
	}
	body.Amount = apaleo.MonetaryValueModel{
		Currency: "EUR",
		Amount:   50.21,
	}

	req.SetRequestBody(body)

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
