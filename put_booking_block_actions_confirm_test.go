package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestPutBookingBlockActionsConfirm(t *testing.T) {
	client := client()
	req := client.NewPutBookingBlockActionsConfirmRequest()
	req.PathParams().ID = "UTR-USJTES"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
