package promotionalcredit

import (
	"github.com/pyntsociety/chargebee-go/enum"
	"github.com/pyntsociety/chargebee-go/filter"
	promotionalCreditEnum "github.com/pyntsociety/chargebee-go/models/promotionalcredit/enum"
)

type PromotionalCredit struct {
	Id             string                     `json:"id"`
	CustomerId     string                     `json:"customer_id"`
	Type           promotionalCreditEnum.Type `json:"type"`
	Amount         int32                      `json:"amount"`
	CurrencyCode   string                     `json:"currency_code"`
	Description    string                     `json:"description"`
	CreditType     enum.CreditType            `json:"credit_type"`
	Reference      string                     `json:"reference"`
	ClosingBalance int32                      `json:"closing_balance"`
	DoneBy         string                     `json:"done_by"`
	CreatedAt      int64                      `json:"created_at"`
	Object         string                     `json:"object"`
}
type AddRequestParams struct {
	CustomerId   string          `json:"customer_id"`
	Amount       *int32          `json:"amount"`
	CurrencyCode string          `json:"currency_code,omitempty"`
	Description  string          `json:"description"`
	CreditType   enum.CreditType `json:"credit_type,omitempty"`
	Reference    string          `json:"reference,omitempty"`
}
type DeductRequestParams struct {
	CustomerId   string          `json:"customer_id"`
	Amount       *int32          `json:"amount"`
	CurrencyCode string          `json:"currency_code,omitempty"`
	Description  string          `json:"description"`
	CreditType   enum.CreditType `json:"credit_type,omitempty"`
	Reference    string          `json:"reference,omitempty"`
}
type SetRequestParams struct {
	CustomerId   string          `json:"customer_id"`
	Amount       *int32          `json:"amount"`
	CurrencyCode string          `json:"currency_code,omitempty"`
	Description  string          `json:"description"`
	CreditType   enum.CreditType `json:"credit_type,omitempty"`
	Reference    string          `json:"reference,omitempty"`
}
type ListRequestParams struct {
	Limit      *int32                  `json:"limit,omitempty"`
	Offset     string                  `json:"offset,omitempty"`
	Id         *filter.StringFilter    `json:"id,omitempty"`
	CreatedAt  *filter.TimestampFilter `json:"created_at,omitempty"`
	Type       *filter.EnumFilter      `json:"type,omitempty"`
	CustomerId *filter.StringFilter    `json:"customer_id,omitempty"`
}
