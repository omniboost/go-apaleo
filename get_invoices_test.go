package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetInvoices(t *testing.T) {
	client := client()
	req := client.NewGetInvoicesRequest()

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetInvoicesAll(t *testing.T) {
	client := client()
	req := client.NewGetInvoicesRequest()
	req.QueryParams().CheckedOutOnAccountsReceivable = true
	req.QueryParams().DateFilter = []string{"gte_2025-03-24", "lte_2025-03-25"}

	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
