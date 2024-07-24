package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MomoMNO struct {
	BankName              string `json:"bank_name"`
	BankCode              string `json:"bank_code"`
	CountryCode           string `json:"country_code"`
	CountryCodeText       string `json:"country_code_text"`
	CountryCurrencySymbol string `json:"country_currency_symbol"`
	Type                  string `json:"type"`
}

type GetMomoMNOsResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []MomoMNO `json:"data"`
}

func getMomoMNOs(countryCode string) {
	url := fmt.Sprintf("%s/external/momo-mnos/%s", baseURL, countryCode)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("X-MERCHANT-ID", merchantID)
	req.Header.Set("X-API-KEY", apiKey)

	body, statusCode, err := sendRequest(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	fmt.Printf("HTTP Status Code: %d\n", statusCode)
	fmt.Printf("Raw Response Body: %s\n", string(body))

	var response GetMomoMNOsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Error parsing response JSON: %v", err)
	}

	if !response.Success {
		log.Fatalf("API request failed: %v", response.Message)
	}

	if len(response.Data) == 0 {
		fmt.Println("No Momo MNOs found for the specified country code.")
		return
	}

	fmt.Println("Momo MNO List:")
	for _, mno := range response.Data {
		fmt.Printf("Bank Name: %s, Bank Code: %s, Country Code: %s, Country Code Text: %s, Currency Symbol: %s, Type: %s\n",
			mno.BankName, mno.BankCode, mno.CountryCode, mno.CountryCodeText, mno.CountryCurrencySymbol, mno.Type)
	}
}
