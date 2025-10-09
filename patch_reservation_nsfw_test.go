package apaleo_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/omniboost/go-apaleo"
)

func TestPatchReservationNSFW(t *testing.T) {
	client := client()
	req := client.NewPatchReservationNSFWRequest()

	// Set reservation ID - replace with actual reservation ID for testing
	req.PathParams().ReservationID = "VYYBJNWZ-1"

	// Optionally set city tax
	body := client.NewPatchReservationNSFWRequestBody()
	body = append(body, apaleo.Operation{
		Operation: "replace",
		Path:      "/isOpenForCharges",
		Value:     false,
	})

	body = append(body, apaleo.Operation{
		Operation: "add",
		Path:      "/marketSegment",
		Value: struct {
			ID string `json:"id"`
		}{ID: "BAR"},
	})

	req.SetRequestBody(body)

	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}

func TestPatchReservationPaymentAccountsNSFW(t *testing.T) {
	// Open the file
	jsonFile, err := os.Open("patch_payment_accounts.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened raw.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []struct {
		PropertyCode   string `json:"PropertyCode"`
		IDPMSFolio     string `json:"IDPMSFolio"`
		PaymentAccount struct {
			AccountNumber  string `json:"AccountNumber"`
			AccountHolder  string `json:"AccountHolder"`
			ExpiryMonth    int    `json:"ExpiryMonth"`
			ExpiryYear     int    `json:"ExpiryYear"`
			PayerReference string `json:"PayerReference"`
			IsVirtual      bool   `json:"IsVirtual"`
		} `json:"PaymentAccount"`
	}
	json.Unmarshal([]byte(byteValue), &result)

	// Create client
	client := client()

	for idx := range result {
		item := result[idx]

		// Fetch reservation details by ExternalReferences.GlobalDistributionSystemID
		req := client.NewGetReservationsNSFWRequest()
		req.QueryParams().ExternalReferences = apaleo.CommaSeparatedQueryParam([]string{item.IDPMSFolio})
		req.QueryParams().PropertyIDs = []string{item.PropertyCode}

		resp, err := req.All(context.Background())
		if err != nil {
			t.Error(err)
		}

		if len(resp) == 0 {
			log.Printf("No reservation found for PMSFolio '%s' and PropertyCode '%s'", item.IDPMSFolio, item.PropertyCode)
			continue
		}

		// Set reservation
		reservation := resp[0]

		// Create patch request
		patchReq := client.NewPatchReservationNSFWRequest()

		// Set reservation ID - replace with actual reservation ID for testing
		patchReq.PathParams().ReservationID = reservation.ID

		// Set payment account
		body := client.NewPatchReservationNSFWRequestBody()
		body = append(body, apaleo.Operation{
			Operation: "replace",
			Path:      "/paymentAccount",
			Value: map[string]interface{}{
				"accountNumber":  item.PaymentAccount.AccountNumber,
				"accountHolder":  item.PaymentAccount.AccountHolder,
				"expiryMonth":    item.PaymentAccount.ExpiryMonth,
				"expiryYear":     item.PaymentAccount.ExpiryYear,
				"payerReference": item.PaymentAccount.PayerReference,
				"isVirtual":      item.PaymentAccount.IsVirtual,
			},
		})

		// Set body
		patchReq.SetRequestBody(body)

		// Do request
		patchResp, err := patchReq.Do(context.Background())
		if err != nil {
			t.Error(err)
		}

		b, _ := json.MarshalIndent(patchResp, "", "  ")
		log.Println(string(b))
	}
}
