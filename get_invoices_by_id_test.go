package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetInvoicesByID(t *testing.T) {
	client := client()
	req := client.NewGetInvoicesByIDRequest()
	req.PathParams().ID = "TEST-20250200000001"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
