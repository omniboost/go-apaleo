package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestDeleteCompanyByID(t *testing.T) {
	client := client()
	req := client.NewDeleteCompanyByIDRequest()
	req.PathParams().ID = "AMZ2-TEST"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestDeleteAllCompanies(t *testing.T) {
	client := client()
	req := client.NewGetCompaniesRequest()
	req.QueryParams().PropertyID = "MID2"
	companies, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	for _, company := range companies {
		req := client.NewDeleteCompanyByIDRequest()
		req.PathParams().ID = company.ID
		_, err := req.Do(context.Background())
		if err != nil {
			t.Error(err)
		}
	}

	b, _ := json.MarshalIndent(companies, "", "  ")
	log.Println(string(b))
}
