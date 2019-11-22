package estimate

import (
	"github.com/pyntsociety/chargebee-go/enum"
	"github.com/pyntsociety/chargebee-go/models/creditnoteestimate"
	"github.com/pyntsociety/chargebee-go/models/invoiceestimate"
	"github.com/pyntsociety/chargebee-go/models/subscriptionestimate"
	"github.com/pyntsociety/chargebee-go/models/unbilledcharge"
)

type Estimate struct {
	CreatedAt               int64                                      `json:"created_at"`
	SubscriptionEstimate    *subscriptionestimate.SubscriptionEstimate `json:"subscription_estimate"`
	InvoiceEstimate         *invoiceestimate.InvoiceEstimate           `json:"invoice_estimate"`
	InvoiceEstimates        []*invoiceestimate.InvoiceEstimate         `json:"invoice_estimates"`
	NextInvoiceEstimate     *invoiceestimate.InvoiceEstimate           `json:"next_invoice_estimate"`
	CreditNoteEstimates     []*creditnoteestimate.CreditNoteEstimate   `json:"credit_note_estimates"`
	UnbilledChargeEstimates []*unbilledcharge.UnbilledCharge           `json:"unbilled_charge_estimates"`
	Object                  string                                     `json:"object"`
}
type CreateSubscriptionRequestParams struct {
	Subscription            *CreateSubscriptionSubscriptionParams      `json:"subscription,omitempty"`
	BillingCycles           *int32                                     `json:"billing_cycles,omitempty"`
	Addons                  []*CreateSubscriptionAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*CreateSubscriptionEventBasedAddonParams `json:"event_based_addons,omitempty"`
	TermsToCharge           *int32                                     `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                  `json:"billing_alignment_mode,omitempty"`
	MandatoryAddonsToRemove []string                                   `json:"mandatory_addons_to_remove,omitempty"`
	CouponIds               []string                                   `json:"coupon_ids,omitempty"`
	BillingAddress          *CreateSubscriptionBillingAddressParams    `json:"billing_address,omitempty"`
	ShippingAddress         *CreateSubscriptionShippingAddressParams   `json:"shipping_address,omitempty"`
	Customer                *CreateSubscriptionCustomerParams          `json:"customer,omitempty"`
	InvoiceImmediately      *bool                                      `json:"invoice_immediately,omitempty"`
	ClientProfileId         string                                     `json:"client_profile_id,omitempty"`
}
type CreateSubscriptionSubscriptionParams struct {
	Id            string `json:"id,omitempty"`
	PlanId        string `json:"plan_id"`
	PlanQuantity  *int32 `json:"plan_quantity,omitempty"`
	PlanUnitPrice *int32 `json:"plan_unit_price,omitempty"`
	SetupFee      *int32 `json:"setup_fee,omitempty"`
	StartDate     *int64 `json:"start_date,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
	Coupon        string `json:"coupon,omitempty"`
}
type CreateSubscriptionAddonParams struct {
	Id            string `json:"id,omitempty"`
	Quantity      *int32 `json:"quantity,omitempty"`
	UnitPrice     *int32 `json:"unit_price,omitempty"`
	BillingCycles *int32 `json:"billing_cycles,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
}
type CreateSubscriptionEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type CreateSubscriptionBillingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubscriptionShippingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateSubscriptionCustomerParams struct {
	VatNumber        string                   `json:"vat_number,omitempty"`
	RegisteredForGst *bool                    `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability          `json:"taxability,omitempty"`
	EntityCode       enum.EntityCode          `json:"entity_code,omitempty"`
	ExemptNumber     string                   `json:"exempt_number,omitempty"`
	ExemptionDetails []map[string]interface{} `json:"exemption_details,omitempty"`
	CustomerType     enum.CustomerType        `json:"customer_type,omitempty"`
}
type CreateSubForCustomerEstimateRequestParams struct {
	UseExistingBalances     *bool                                                `json:"use_existing_balances,omitempty"`
	Subscription            *CreateSubForCustomerEstimateSubscriptionParams      `json:"subscription,omitempty"`
	InvoiceImmediately      *bool                                                `json:"invoice_immediately,omitempty"`
	BillingCycles           *int32                                               `json:"billing_cycles,omitempty"`
	Addons                  []*CreateSubForCustomerEstimateAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*CreateSubForCustomerEstimateEventBasedAddonParams `json:"event_based_addons,omitempty"`
	TermsToCharge           *int32                                               `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                            `json:"billing_alignment_mode,omitempty"`
	MandatoryAddonsToRemove []string                                             `json:"mandatory_addons_to_remove,omitempty"`
	ShippingAddress         *CreateSubForCustomerEstimateShippingAddressParams   `json:"shipping_address,omitempty"`
	CouponIds               []string                                             `json:"coupon_ids,omitempty"`
}
type CreateSubForCustomerEstimateSubscriptionParams struct {
	Id            string `json:"id,omitempty"`
	PlanId        string `json:"plan_id"`
	PlanQuantity  *int32 `json:"plan_quantity,omitempty"`
	PlanUnitPrice *int32 `json:"plan_unit_price,omitempty"`
	SetupFee      *int32 `json:"setup_fee,omitempty"`
	StartDate     *int64 `json:"start_date,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
}
type CreateSubForCustomerEstimateAddonParams struct {
	Id            string `json:"id,omitempty"`
	Quantity      *int32 `json:"quantity,omitempty"`
	UnitPrice     *int32 `json:"unit_price,omitempty"`
	BillingCycles *int32 `json:"billing_cycles,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
}
type CreateSubForCustomerEstimateEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
}
type CreateSubForCustomerEstimateShippingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionRequestParams struct {
	Subscription            *UpdateSubscriptionSubscriptionParams      `json:"subscription,omitempty"`
	BillingCycles           *int32                                     `json:"billing_cycles,omitempty"`
	Addons                  []*UpdateSubscriptionAddonParams           `json:"addons,omitempty"`
	EventBasedAddons        []*UpdateSubscriptionEventBasedAddonParams `json:"event_based_addons,omitempty"`
	ReplaceAddonList        *bool                                      `json:"replace_addon_list,omitempty"`
	MandatoryAddonsToRemove []string                                   `json:"mandatory_addons_to_remove,omitempty"`
	TermsToCharge           *int32                                     `json:"terms_to_charge,omitempty"`
	ReactivateFrom          *int64                                     `json:"reactivate_from,omitempty"`
	BillingAlignmentMode    enum.BillingAlignmentMode                  `json:"billing_alignment_mode,omitempty"`
	CouponIds               []string                                   `json:"coupon_ids,omitempty"`
	ReplaceCouponList       *bool                                      `json:"replace_coupon_list,omitempty"`
	Prorate                 *bool                                      `json:"prorate,omitempty"`
	EndOfTerm               *bool                                      `json:"end_of_term,omitempty"`
	ForceTermReset          *bool                                      `json:"force_term_reset,omitempty"`
	Reactivate              *bool                                      `json:"reactivate,omitempty"`
	IncludeDelayedCharges   *bool                                      `json:"include_delayed_charges,omitempty"`
	UseExistingBalances     *bool                                      `json:"use_existing_balances,omitempty"`
	BillingAddress          *UpdateSubscriptionBillingAddressParams    `json:"billing_address,omitempty"`
	ShippingAddress         *UpdateSubscriptionShippingAddressParams   `json:"shipping_address,omitempty"`
	Customer                *UpdateSubscriptionCustomerParams          `json:"customer,omitempty"`
	InvoiceImmediately      *bool                                      `json:"invoice_immediately,omitempty"`
}
type UpdateSubscriptionSubscriptionParams struct {
	Id            string `json:"id"`
	PlanId        string `json:"plan_id,omitempty"`
	PlanQuantity  *int32 `json:"plan_quantity,omitempty"`
	PlanUnitPrice *int32 `json:"plan_unit_price,omitempty"`
	SetupFee      *int32 `json:"setup_fee,omitempty"`
	StartDate     *int64 `json:"start_date,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
	Coupon        string `json:"coupon,omitempty"`
}
type UpdateSubscriptionAddonParams struct {
	Id            string `json:"id,omitempty"`
	Quantity      *int32 `json:"quantity,omitempty"`
	UnitPrice     *int32 `json:"unit_price,omitempty"`
	BillingCycles *int32 `json:"billing_cycles,omitempty"`
	TrialEnd      *int64 `json:"trial_end,omitempty"`
}
type UpdateSubscriptionEventBasedAddonParams struct {
	Id                  string        `json:"id,omitempty"`
	Quantity            *int32        `json:"quantity,omitempty"`
	UnitPrice           *int32        `json:"unit_price,omitempty"`
	ServicePeriodInDays *int32        `json:"service_period_in_days,omitempty"`
	ChargeOn            enum.ChargeOn `json:"charge_on,omitempty"`
	OnEvent             enum.OnEvent  `json:"on_event,omitempty"`
	ChargeOnce          *bool         `json:"charge_once,omitempty"`
}
type UpdateSubscriptionBillingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionShippingAddressParams struct {
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateSubscriptionCustomerParams struct {
	VatNumber        string          `json:"vat_number,omitempty"`
	RegisteredForGst *bool           `json:"registered_for_gst,omitempty"`
	Taxability       enum.Taxability `json:"taxability,omitempty"`
}
type RenewalEstimateRequestParams struct {
	IncludeDelayedCharges       *bool `json:"include_delayed_charges,omitempty"`
	UseExistingBalances         *bool `json:"use_existing_balances,omitempty"`
	IgnoreScheduledCancellation *bool `json:"ignore_scheduled_cancellation,omitempty"`
	IgnoreScheduledChanges      *bool `json:"ignore_scheduled_changes,omitempty"`
}
type ChangeTermEndRequestParams struct {
	TermEndsAt         *int64 `json:"term_ends_at,omitempty"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
}
type CancelSubscriptionRequestParams struct {
	EndOfTerm                         *bool                                  `json:"end_of_term,omitempty"`
	CreditOptionForCurrentTermCharges enum.CreditOptionForCurrentTermCharges `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             enum.UnbilledChargesOption             `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        enum.AccountReceivablesHandling        `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         enum.RefundableCreditsHandling         `json:"refundable_credits_handling,omitempty"`
}
type PauseSubscriptionRequestParams struct {
	PauseOption             enum.PauseOption                     `json:"pause_option,omitempty"`
	Subscription            *PauseSubscriptionSubscriptionParams `json:"subscription,omitempty"`
	UnbilledChargesHandling enum.UnbilledChargesHandling         `json:"unbilled_charges_handling,omitempty"`
}
type PauseSubscriptionSubscriptionParams struct {
	PauseDate  *int64 `json:"pause_date,omitempty"`
	ResumeDate *int64 `json:"resume_date,omitempty"`
}
type ResumeSubscriptionRequestParams struct {
	ResumeOption    enum.ResumeOption                     `json:"resume_option,omitempty"`
	Subscription    *ResumeSubscriptionSubscriptionParams `json:"subscription,omitempty"`
	ChargesHandling enum.ChargesHandling                  `json:"charges_handling,omitempty"`
}
type ResumeSubscriptionSubscriptionParams struct {
	ResumeDate *int64 `json:"resume_date,omitempty"`
}
type GiftSubscriptionRequestParams struct {
	Gift            *GiftSubscriptionGiftParams            `json:"gift,omitempty"`
	Gifter          *GiftSubscriptionGifterParams          `json:"gifter,omitempty"`
	GiftReceiver    *GiftSubscriptionGiftReceiverParams    `json:"gift_receiver,omitempty"`
	Subscription    *GiftSubscriptionSubscriptionParams    `json:"subscription,omitempty"`
	Addons          []*GiftSubscriptionAddonParams         `json:"addons,omitempty"`
	ShippingAddress *GiftSubscriptionShippingAddressParams `json:"shipping_address,omitempty"`
	CouponIds       []string                               `json:"coupon_ids,omitempty"`
}
type GiftSubscriptionGiftParams struct {
	ScheduledAt     *int64 `json:"scheduled_at,omitempty"`
	AutoClaim       *bool  `json:"auto_claim,omitempty"`
	ClaimExpiryDate *int64 `json:"claim_expiry_date,omitempty"`
}
type GiftSubscriptionGifterParams struct {
	CustomerId   string `json:"customer_id"`
	Signature    string `json:"signature"`
	Note         string `json:"note,omitempty"`
	PaymentSrcId string `json:"payment_src_id,omitempty"`
}
type GiftSubscriptionGiftReceiverParams struct {
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
}
type GiftSubscriptionSubscriptionParams struct {
	PlanId       string `json:"plan_id"`
	PlanQuantity *int32 `json:"plan_quantity,omitempty"`
}
type GiftSubscriptionAddonParams struct {
	Id       string `json:"id,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
}
type GiftSubscriptionShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateInvoiceRequestParams struct {
	Invoice                    *CreateInvoiceInvoiceParams         `json:"invoice,omitempty"`
	CurrencyCode               string                              `json:"currency_code,omitempty"`
	Addons                     []*CreateInvoiceAddonParams         `json:"addons,omitempty"`
	Charges                    []*CreateInvoiceChargeParams        `json:"charges,omitempty"`
	Coupon                     string                              `json:"coupon,omitempty"`
	AuthorizationTransactionId string                              `json:"authorization_transaction_id,omitempty"`
	PaymentSourceId            string                              `json:"payment_source_id,omitempty"`
	AutoCollection             enum.AutoCollection                 `json:"auto_collection,omitempty"`
	ShippingAddress            *CreateInvoiceShippingAddressParams `json:"shipping_address,omitempty"`
}
type CreateInvoiceInvoiceParams struct {
	CustomerId string `json:"customer_id"`
	PoNumber   string `json:"po_number,omitempty"`
}
type CreateInvoiceAddonParams struct {
	Id        string `json:"id,omitempty"`
	Quantity  *int32 `json:"quantity,omitempty"`
	UnitPrice *int32 `json:"unit_price,omitempty"`
	DateFrom  *int64 `json:"date_from,omitempty"`
	DateTo    *int64 `json:"date_to,omitempty"`
}
type CreateInvoiceChargeParams struct {
	Amount                 *int32               `json:"amount,omitempty"`
	Description            string               `json:"description,omitempty"`
	AvalaraSaleType        enum.AvalaraSaleType `json:"avalara_sale_type,omitempty"`
	AvalaraTransactionType *int32               `json:"avalara_transaction_type,omitempty"`
	AvalaraServiceType     *int32               `json:"avalara_service_type,omitempty"`
	DateFrom               *int64               `json:"date_from,omitempty"`
	DateTo                 *int64               `json:"date_to,omitempty"`
}
type CreateInvoiceShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
