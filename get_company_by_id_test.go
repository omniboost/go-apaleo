package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetCompanyByID(t *testing.T) {
	client := client()
	req := client.NewGetCompanyByIDRequest()
	req.PathParams().ID = "AMZ2-AMZ_C1041"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
