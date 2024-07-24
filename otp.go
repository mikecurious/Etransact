package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GenerateOTPRequest struct {
	ConfigPreset      string                       `json:"config_preset"`
	TransactionParams GenerateOTPTransactionParams `json:"transaction_params"`
}

type GenerateOTPTransactionParams struct {
	SrcAmount string `json:"src_amount"`
	DesAmount string `json:"des_amount"`
	PayeeID   string `json:"payeeId"`
	PayerID   string `json:"payerId"`
}

type GenerateOTPResponse struct {
	Status bool `json:"status"`
	Data   struct {
		Status  bool   `json:"status"`
		OTPSID  string `json:"otp_sid"`
		Message string `json:"message"`
	} `json:"data"`
}

func generateOTP() {
	url := fmt.Sprintf("%s/otp/generate/dynamic-link", ESAURL)

	requestBody := GenerateOTPRequest{
		ConfigPreset: configPreset,
		TransactionParams: GenerateOTPTransactionParams{
			SrcAmount: "5.00",
			DesAmount: "5.00",
			PayeeID:   "AIRTEL706218827",
			PayerID:   merchantID,
		},
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

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Parse response JSON
	var response GenerateOTPResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error parsing response JSON: %v", err)
	}

	// Check if OTP link generation was successful
	if response.Data.Status {
		fmt.Printf("OTP generated successfully.\nOTP SID: %s\nMessage: %s\n", response.Data.OTPSID, response.Data.Message)
	} else {
		log.Fatalf("OTP generation failed: %v\n", response.Data.Message)
	}
}
