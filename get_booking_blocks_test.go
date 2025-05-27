package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetBookingBlocks(t *testing.T) {
	client := client()
	req := client.NewGetBookingBlocksRequest()

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetBookingBlocksAll(t *testing.T) {
	client := client()
	req := client.NewGetBookingBlocksRequest()

	resp, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
