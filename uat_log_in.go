package unit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Info struct {
	WalletNumber string   `json:"wallet_number"`
	DeviceUuid   string   `json:"device_uuid"`
	PinNumber    string   `json:"pin_number"`
	GeoLocation  *Latlong `json:"geo_location"`
}

type Latlong struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Authentication struct {
	Code    string `json:"code"`
	Lang    string `json:"lang"`
	Message string `json:"message"`
	Data    *Data  `json:"data"`
}

type Data struct {
	AccessToken              string `json:"access_token"`
	RefreshToken             string `json:"refresh_token"`
	TransactionToken         string `json:"transaction_token"`
	DeviceAuthorizationToken string `json:"device_authorization_token"`
	FirebaseToken            string `json:"firebase_token"`
	WalletStatus             string `json:"wallet_status"`
	WalletLockTimer          string `json:"wallet_lock_timer"`
	WalletLockCounter        int    `json:"wallet_lock_counter"`
}

func uatLogIn() (token string) {
	url := "https://uat-api.upay.systems/dfsc/oam/app/v1/login/"
	info := Info{WalletNumber: "01676766793", DeviceUuid: "c65m117a98bc0cea44e15137", PinNumber: "1122", GeoLocation: &Latlong{Lat: 190.00, Long: 35.00}}
	requestPayload, _ := json.Marshal(info)
	var authentication Authentication

	res, err := http.Post(url, "application/json", bytes.NewBuffer(requestPayload))

	if err != nil {
		fmt.Println("error occurred while getting the response!")
	} else {
		bodyInBytes, _ := ioutil.ReadAll(res.Body)
		err := json.Unmarshal(bodyInBytes, &authentication)
		if err != nil {
			return
		}
	}
	return authentication.Data.AccessToken
}
