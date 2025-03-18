package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TestPostChargeToFolio(t *testing.T) {
	client := client()
	req := client.NewPostChargeToFolioRequest()
	req.PathParams().ID = "AAJPJLNI-1-1"

	body := req.NewRequestBody()
	body.ServiceType = "Other"
	body.VatType = "Null"
	body.SubAccountID = "AMZ2-OTHERMIGRATION"
	body.Name = "Other Migration"
	body.Amount = apaleo.MonetaryValueModel{
		Currency: "EUR",
		Amount:   50.21,
	}

	req.SetRequestBody(body)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
