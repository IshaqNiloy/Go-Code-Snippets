package unit

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	PolicyNumber string  `json:"policy_number"`
	PolicyType   int     `json:"policy_type"`
	PaymentFor   string  `json:"payment_for"`
	amount       float64 `json:"amount"`
}

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Lang    string      `json:"lang"`
	Data    *InsideData `json:"data"`
}

type InsideData struct {
	AmountInfo []ResponseInfo `json:"amount_info"`
	BasicInfo  []ResponseInfo `json:"basic_info"`
}

type ResponseInfo struct {
	Key   string `json:"key"`
	Title string `json:"title"`
	Value string `json:"value"`
}

var response Response

func getCode() string {
	return response.Code
}

func parser(policyNumber string, policyType int, premiumAmount float64) []byte {
	//API url
	url := "http://localhost:8080/insurance/service/v1/policy-info"

	//Constructing Payload
	payload := Payload{PolicyNumber: policyNumber, PolicyType: policyType, PaymentFor: "insurance_bengal_islami_life", amount: premiumAmount}
	requestPayload, err := json.Marshal(payload)

	if err != nil {
		log.Error().Err(err).Msg("error occurred when payload was unmarshalled.")
		return nil
	}
	//Making request to the API
	var bearer = "JWT " + uatLogIn()
	//time.Sleep(1 * time.Second)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestPayload))
	if err != nil {
		log.Error().Err(err).Msg("error occurred while constructing a new request object.")
		return nil
	}
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("error occurred while getting a response.")
		return nil
	}
	defer res.Body.Close()

	bodyInBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Error().Err(err).Msg("error occurred when response was marshalled.")
		return nil
	}

	err = json.Unmarshal(bodyInBytes, &response)

	if err != nil {
		log.Error().Err(err).Msg("error occurred when bodyInBytes was unmarshalled.")
		return nil
	}

	return bodyInBytes
}
