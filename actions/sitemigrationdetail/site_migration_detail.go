package sitemigrationdetail

import (
	"fmt"
	"github.com/pyntsociety/chargebee-go"
	"github.com/pyntsociety/chargebee-go/models/sitemigrationdetail"
)

func List(params *sitemigrationdetail.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/site_migration_details"), params)
}
