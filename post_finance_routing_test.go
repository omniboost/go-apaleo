package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestPostFinanceRouting(t *testing.T) {
	client := client()
	req := client.NewPostFinanceRoutingRequest()

	body := req.NewRequestBody()
	body.PropertyID = "MID2"
	body.BookingID = "CDDEFNPF"
	body.DestinationFolioID = "CDDEFNPF-MID2-1"
	body.Filter.ServiceTypes = []string{"Accommodation", "FoodAndBeverages", "NoShow", "CancellationFees", "Other"}

	req.SetRequestBody(body)

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
