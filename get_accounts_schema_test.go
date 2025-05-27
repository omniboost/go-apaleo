package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetSchema(t *testing.T) {
	client := client()
	req := client.NewGetAccountsSchemaRequest()
	req.QueryParams().PropertyID = "AMZ2"
	req.QueryParams().Depth = 4
	req.QueryParams().AccountingSchema = "Extended"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
