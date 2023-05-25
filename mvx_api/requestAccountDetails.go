package mvx_api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func RequestAccountDetail(address string) (Accounts, error) {
	//Encode the data
	var (
		url = "https://" + API_URL + "/accounts/" + address
	)

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Err(err).Msg("Error when client create GET request to " + url)
		return Accounts{}, err
	}

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Err(err).Msg("Error when client do request to " + url)
		return Accounts{}, err
	}
	defer resp.Body.Close()

	// Check error code
	if resp.StatusCode != http.StatusOK {
		var errorMvx ErrorMvx
		body, _ := io.ReadAll(resp.Body)
		_ = json.Unmarshal(body, &errorMvx)
		return Accounts{}, errors.New(errorMvx.Message)
	}

	// Convert JSON into struct
	var accounts Accounts
	err = json.NewDecoder(resp.Body).Decode(&accounts)
	if err != nil {
		log.Err(err).Msg("Error when Decode JSON")
		return Accounts{}, err
	}

	return accounts, nil
}
