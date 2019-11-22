package paymentintent

import (
	"fmt"
	"github.com/pyntsociety/chargebee-go"
	"github.com/pyntsociety/chargebee-go/models/paymentintent"
)

func Create(params *paymentintent.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_intents"), params)
}
func Update(id string, params *paymentintent.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_intents/%v", id), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/payment_intents/%v", id), nil)
}
