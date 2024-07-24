package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ValidateAccountResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Data    struct {
		Name     string `json:"name"`
		Accounts []struct {
			AccountID string `json:"account_id"`
			Currency  string `json:"currency"`
		} `json:"accounts"`
	} `json:"data"`
}

func validateMomoAccount(countryCode string, accountNo string, bankCode string) {
	url := fmt.Sprintf("%s/external/validate-account/%s/%s/%s", baseURL, countryCode, accountNo, bankCode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

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

	// Print HTTP status code
	fmt.Printf("HTTP Status Code: %d\n", resp.StatusCode)

	// Print raw response body
	fmt.Printf("Raw Response Body: %s\n", string(body))

	// Parse response JSON
	var response ValidateAccountResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error parsing response JSON: %v", err)
	}

	// Check for API request success
	if !response.Success {
		log.Fatalf("API request failed: %v\nError: %v", response.Message, response.Error)
	}

	// Print details if validation successful
	fmt.Printf("Validation successful for account: %s\n", accountNo)
	fmt.Printf("Account Name: %s\n", response.Data.Name)
	for _, acc := range response.Data.Accounts {
		fmt.Printf("Account ID: %s, Currency: %s\n", acc.AccountID, acc.Currency)
	}
}
