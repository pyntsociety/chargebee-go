package subscription

import (
	"fmt"
	"github.com/pyntsociety/chargebee-go"
	"github.com/pyntsociety/chargebee-go/models/subscription"
)

func Create(params *subscription.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions"), params)
}
func CreateForCustomer(id string, params *subscription.CreateForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/subscriptions", id), params)
}
func List(params *subscription.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/subscriptions"), params)
}
func SubscriptionsForCustomer(id string, params *subscription.SubscriptionsForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/customers/%v/subscriptions", id), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v", id), nil)
}
func RetrieveWithScheduledChanges(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/subscriptions/%v/retrieve_with_scheduled_changes", id), nil)
}
func RemoveScheduledChanges(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_changes", id), nil)
}
func RemoveScheduledCancellation(id string, params *subscription.RemoveScheduledCancellationRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_cancellation", id), params)
}
func RemoveCoupons(id string, params *subscription.RemoveCouponsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_coupons", id), params)
}
func Update(id string, params *subscription.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v", id), params)
}
func ChangeTermEnd(id string, params *subscription.ChangeTermEndRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/change_term_end", id), params)
}
func Cancel(id string, params *subscription.CancelRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/cancel", id), params)
}
func Reactivate(id string, params *subscription.ReactivateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/reactivate", id), params)
}
func AddChargeAtTermEnd(id string, params *subscription.AddChargeAtTermEndRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/add_charge_at_term_end", id), params)
}
func ChargeAddonAtTermEnd(id string, params *subscription.ChargeAddonAtTermEndRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/charge_addon_at_term_end", id), params)
}
func ChargeFutureRenewals(id string, params *subscription.ChargeFutureRenewalsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/charge_future_renewals", id), params)
}
func ImportSubscription(params *subscription.ImportSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/import_subscription"), params)
}
func ImportForCustomer(id string, params *subscription.ImportForCustomerRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/customers/%v/import_subscription", id), params)
}
func OverrideBillingProfile(id string, params *subscription.OverrideBillingProfileRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/override_billing_profile", id), params)
}
func Delete(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/delete", id), nil)
}
func Pause(id string, params *subscription.PauseRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/pause", id), params)
}
func Resume(id string, params *subscription.ResumeRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/resume", id), params)
}
func RemoveScheduledPause(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_pause", id), nil)
}
func RemoveScheduledResumption(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/subscriptions/%v/remove_scheduled_resumption", id), nil)
}
