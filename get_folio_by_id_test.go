package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetFolioByID(t *testing.T) {
	client := client()
	req := client.NewGetFolioByIDRequest()
	req.PathParams().ID = "UZFKVQBI-1-1"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
