package unit

import (
	"encoding/json"
	"flag"
	"reflect"
	"testing"
)

var policyNumber string
var policyType int
var premiumAmount float64

func init() {
	flag.StringVar(&policyNumber, "policyNumber", "", "help message for arg1")
	flag.IntVar(&policyType, "policyType", 1, "help message for arg2")
	flag.Float64Var(&premiumAmount, "premiumAmount", 1, "help message for arg3")
}

func TestMain(m *testing.M) {
	flag.Parse()
	m.Run()
}

func TestInfoAPIDynamic(t *testing.T) {
	//Inputs from user
	bodyInBytes := parser(policyNumber, policyType, premiumAmount)
	expectedJson := `{"code":"` + response.Code + `","message":"` + response.Message +
		`","lang":"` + response.Lang + `","data":{"amount_info":[{"key":"` +
		response.Data.AmountInfo[0].Key + `","title":"` +
		response.Data.AmountInfo[0].Title + `","value":"` +
		response.Data.AmountInfo[0].Value + `"},{"key":"` +
		response.Data.AmountInfo[1].Key + `","title":"` +
		response.Data.AmountInfo[1].Title + `","value":"` +
		response.Data.AmountInfo[1].Value + `"},{"key":"` +
		response.Data.AmountInfo[2].Key + `","title":"` +
		response.Data.AmountInfo[2].Title + `","value":"` +
		response.Data.AmountInfo[2].Value + `"}],"basic_info":[{"key":"` +
		response.Data.BasicInfo[0].Key + `","title":"` +
		response.Data.BasicInfo[0].Title + `","value":"` +
		response.Data.BasicInfo[0].Value + `"},{"key":"` +
		response.Data.BasicInfo[1].Key + `","title":"` +
		response.Data.BasicInfo[1].Title + `","value":"` +
		response.Data.BasicInfo[1].Value + `"},{"key":"` +
		response.Data.BasicInfo[2].Key + `","title":"` +
		response.Data.BasicInfo[2].Title + `","value":"` +
		response.Data.BasicInfo[2].Value + `"},{"key":"` +
		response.Data.BasicInfo[3].Key + `","title":"` +
		response.Data.BasicInfo[3].Title + `","value":"` +
		response.Data.BasicInfo[3].Value + `"}]}}`

	//Comparing two JSON objects
	var obj1, obj2 map[string]interface{}
	json.Unmarshal([]byte(expectedJson), &obj1)
	json.Unmarshal(bodyInBytes, &obj2)

	if !reflect.DeepEqual(obj1, obj2) {
		t.Errorf("Expected JSON is not same as actual JSON")
	}
}
