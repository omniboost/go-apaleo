package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TestPostAllowanceToFolio(t *testing.T) {
	client := client()
	req := client.NewPostAllowanceToFolioRequest()
	req.PathParams().ID = "IMJWEMLG-AMZ2-1"

	body := req.NewRequestBody()
	body.ServiceType = "Other"
	body.VatType = "Null"
	body.SubAccountID = "AMZ2-OTHERMIGRATION"
	body.Reason = "Refund: AMZ-F65001-2333866"
	body.Amount = apaleo.MonetaryValueModel{
		Currency: "EUR",
		Amount:   3.5,
	}

	req.SetRequestBody(body)

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
