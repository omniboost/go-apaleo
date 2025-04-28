package apaleo_test

import (
	"encoding/json"
	"log"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestPutReservationAssignUnitByID(t *testing.T) {
	client := client()
	req := client.NewPutReservationAssignUnitByIDRequest()

	// Set reservation ID - replace with actual reservation ID for testing
	req.PathParams().ReservationID = "reservation-id-here"

	// Optionally set unit conditions
	req.QueryParams().UnitConditions = []string{"Clean", "CleanToBeInspected"}

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
