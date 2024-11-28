package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetCancellationPolicies(t *testing.T) {
	client := client()
	req := client.NewGetCancellationPoliciesRequest()
	req.QueryParams().PropertyID = "OMNI"

	resp, err := req.Do()
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

	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
