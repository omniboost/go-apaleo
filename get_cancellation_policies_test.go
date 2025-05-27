package apaleo_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetCancellationPolicies(t *testing.T) {
	client := client()
	req := client.NewGetCancellationPoliciesRequest()
	req.QueryParams().PropertyID = "OMNI"

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestGetCancellationPoliciesAll(t *testing.T) {
	client := client()
	req := client.NewGetCancellationPoliciesRequest()
	req.QueryParams().PropertyID = "OMNI"

	resp, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
