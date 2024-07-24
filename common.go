package main

import (
	"io/ioutil"
	"net/http"
)

const (
	baseURL      = "https://sandbox-api.xcelapp.com/xas/v1"
	merchantID   = "8e9745acfe"
	apiKey       = "XCL_DEV_83ap3OVtTEpbhOT56h22PZU4sAkPbVSCnNajl5CSxo1CYraqygZGSbofGxo2fsJs"
	configPreset = "SgKOr2mJp"
	ESAURL       = "https://sandbox-api.xcelapp.com/esa-api/api"
	webhookURL   = "https://webhook.site/8bd6abfd-6452-42fd-a8d2-58089977b171"
)

func sendRequest(req *http.Request) ([]byte, int, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}
