package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CheckPaymentStatusResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Status          string `json:"status"`
		Description     string `json:"description"`
		Reference       string `json:"reference"`
		ClientReference string `json:"clientReference"`
		TransDate       string `json:"transDate"`
	} `json:"data"`
}

func checkPaymentStatus(reference string) {
	url := fmt.Sprintf("%s/external/payment/status/%s", baseURL, reference)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-MERCHANT-ID", merchantID)
	req.Header.Set("X-API-KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var response CheckPaymentStatusResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error parsing response JSON: %v", err)
	}

	if response.Success {
		fmt.Printf("Status enquiry successful.\nStatus: %s\nDescription: %s\nReference: %s\nClient Reference: %s\nTransaction Date: %s\n",
			response.Data.Status, response.Data.Description, response.Data.Reference, response.Data.ClientReference, response.Data.TransDate)
	} else {
		log.Fatalf("Status enquiry failed: %s\n", response.Message)
	}
}
