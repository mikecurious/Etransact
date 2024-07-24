package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type InitiateMoMoPaymentRequest struct {
	Country     string `json:"country"`
	Currency    string `json:"currency"`
	Amount      string `json:"amount"`
	Reference   string `json:"reference"`
	DLCode      string `json:"dl_code"`
	BankCode    string `json:"bank_code"`
	AccountNum  string `json:"account_num"`
	AccountName string `json:"account_name"`
	Description string `json:"description"`
	WebhookURL  string `json:"webhook_url"`
}

type InitiateMoMoPaymentResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Reference string `json:"reference"`
		Status    string `json:"status"`
	} `json:"data"`
}

func initiateMoMoPayment() {
	url := fmt.Sprintf("%s/external/payment/momo", baseURL)

	requestBody := InitiateMoMoPaymentRequest{
		Country:     "UG",
		Currency:    "TVD",
		Amount:      "5.00",
		Reference:   "739903900657", //<32
		DLCode:      "e3e64659-4564-47ca-9654-f55fa5f77700",
		BankCode:    "AIRTEL",
		AccountNum:  "706218827",
		AccountName: "Account Name",
		Description: "Transaction description",
		WebhookURL:  webhookURL,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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

	var response InitiateMoMoPaymentResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error parsing response JSON: %v", err)
	}

	if response.Success {
		fmt.Printf("MoMo payment initiated successfully.\nReference: %s\nStatus: %s\n", response.Data.Reference, response.Data.Status)
	} else {
		log.Fatalf("MoMo payment initiation failed: %s\n", response.Message)
	}
}
