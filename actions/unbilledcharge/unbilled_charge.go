package unbilledcharge

import (
	"fmt"
	"github.com/pyntsociety/chargebee-go"
	"github.com/pyntsociety/chargebee-go/models/unbilledcharge"
)

func InvoiceUnbilledCharges(params *unbilledcharge.InvoiceUnbilledChargesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/invoice_unbilled_charges"), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/%v/delete", id), nil)
}
func List(params *unbilledcharge.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/unbilled_charges"), params)
}
func InvoiceNowEstimate(params *unbilledcharge.InvoiceNowEstimateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/unbilled_charges/invoice_now_estimate"), params)
}
