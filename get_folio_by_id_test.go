package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestGetFolioByID(t *testing.T) {
	client := client()
	req := client.NewGetFolioByIDRequest()
	req.PathParams().ID = "AAJPJLNI-1-1"

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
