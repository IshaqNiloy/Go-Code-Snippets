package unit

import (
	"context"
	"flag"
	"fmt"
	"gitlab.upay.dev/golang/bill_paymnet_lib/http_request"
	"testing"
)

type Payload struct {
	PolicyNumber string  `json:"policy_number"`
	PolicyType   int     `json:"policy_type"`
	PaymentFor   string  `json:"payment_for"`
	Amount       float64 `json:"amount"`
}

//The following lines of code are used to take inputs from the user as arguments with the command line "go test -v -args -policyNumber=1 -policyType=2 -premiumAmount=0.00
"
var policyNumber string
var policyType int
var premiumAmount float64

func init() {
	flag.StringVar(&policyNumber, "policyNumber", "", "help message for arg1")
	flag.IntVar(&policyType, "policyType", 1, "help message for arg2")
	flag.Float64Var(&premiumAmount, "premiumAmount", 1.0, "help message for arg3")
}

func TestMain(m *testing.M) {
	flag.Parse()
	m.Run()
}
//End
func TestInfoAPI(t *testing.T) {
	url := "http://localhost:8080/insurance/service/v1/policy-info"
	expectedJson := `{"code":"FIIO_BIS_200","message":"Bill info fetch success","lang":"en",
					"data":{"amount_info":[{"key":"amount","title":"Amount","value":"45000.00"},
					{"key":"service_charge","title":"Charge","value":"0.00"},{"key":"total_amount",
					"title":"Total Payment","value":"45000.00"}],"basic_info":[{"key":"policy_number",
					"title":"Policy number","value":"1"},{"key":"policy_holder_name","title":"Name",
					"value":"MST. SHATI KHATUN"}]}}`

	//Constructing Payload
	payload := Payload{PolicyNumber: policyNumber, PolicyType: policyType, PaymentFor: "insurance_akij_takaful_life", Amount: premiumAmount}
	requestPayload, err := json.Marshal(payload)

	//Making request to the API
	var bearer = "JWT " + uatLogIn()
	time.Sleep(1 * time.Second)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestPayload))
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	
	if err != nil || res.StatusCode != http.StatusOK {
		fmt.Println("Error occurred while getting response!")
	} else {
		bodyInBytes, err := ioutil.ReadAll(res.Body)
		// If the response body can not be read
		if err != nil {
			fmt.Println("Error occurred while converting the body in bytes!")
		} else {
			//Comparing two JSON objects
			var obj1, obj2 map[string]interface{}
			json.Unmarshal([]byte(expectedJson), &obj1)
			json.Unmarshal(bodyInBytes, &obj2)
	
			if !reflect.DeepEqual(obj1, obj2) {
				t.Errorf("Expected JSON is not same as actual JSON")
			}
		}
	}
}
