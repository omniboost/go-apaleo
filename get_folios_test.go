package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TestGetFolios(t *testing.T) {
	client := client()
	req := client.NewGetFoliosRequest()

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetFoliosAll(t *testing.T) {
	client := client()
	req := client.NewGetFoliosRequest()
	req.QueryParams().ReservationIDs = apaleo.CommaSeparatedQueryParam{"AAJPJLNI-1"}

	resp, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
