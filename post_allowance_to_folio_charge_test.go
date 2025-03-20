package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TestPostAllowanceToFolioCharge(t *testing.T) {
	client := client()
	req := client.NewPostAllowanceToFolioChargeRequest()
	req.PathParams().FolioID = "IMJWEMLG-AMZ2-1"
	req.PathParams().ChargeID = "IMJWEMLG-AMZ2-1-D-1"

	body := req.NewRequestBody()
	body.Reason = "Refund: AMZ-F65001-2333866"
	body.Amount = apaleo.MonetaryValueModel{
		Currency: "EUR",
		Amount:   3.5,
	}

	req.SetRequestBody(body)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
