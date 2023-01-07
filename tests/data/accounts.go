package data

import (
	"github.com/form3-test-task/client/entities/accounts/objects"
	"github.com/form3-test-task/tests/helpers"
)

var (
	ExpectedAccountObjectAfterCreation = &objects.Account{
		Type: "accounts",
		Attributes: &objects.AccountAttributes{
			Bic:           "NWBKFR42",
			Iban:          "FR1420041010050500013M02606",
			Name:          []string{"Name1", "Name2"},
			BankID:        "20041",
			Country:       helpers.GetStringPointer("FR"),
			BankIDCode:    "FR",
			BaseCurrency:  "EUR",
			AccountNumber: "0500013M026",
		},
		Version:        helpers.GetInt64Pointer(0),
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	}
)
