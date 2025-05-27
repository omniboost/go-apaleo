package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetInvoicesByID(t *testing.T) {
	client := client()
	req := client.NewGetInvoicesByIDRequest()
	req.PathParams().ID = "AMZ-202500000018"
	req.QueryParams().Expand = []string{"company"}

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
