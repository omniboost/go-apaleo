package apaleo_test

import (
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetPdf(t *testing.T) {
	client := client()
	req := client.NewGetPdfRequest()

	req.PathParams().ID = "TEST-20190700000001"
	// req.QueryParams().Expand = []string{"invoices"}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	// b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(resp))
}
