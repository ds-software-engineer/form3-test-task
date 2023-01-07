package accounts

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/form3-test-task/client"
	"github.com/form3-test-task/client/entities/accounts/objects"
	"github.com/form3-test-task/client/entities/transport"
	"github.com/form3-test-task/tests/data"
	"github.com/form3-test-task/tests/helpers"
)

// TestAccountCreate_Ok tests Create action happy path case.
func TestAccountCreate_Ok(t *testing.T) {
	accountID := uuid.New().String()

	fakeAPIClient := client.NewFakeAPIClient(os.Getenv("BASE_FAKE_API_URL"))
	actualAccount, err := fakeAPIClient.Accounts().Create(&objects.Account{
		ID:   accountID,
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
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	})
	assert.Nil(t, err)

	expectedAccount := data.ExpectedAccountObjectAfterCreation

	assert.Equal(t, accountID, actualAccount.ID)
	assert.Equal(t, expectedAccount.Type, actualAccount.Type)
	assert.Equal(t, expectedAccount.Attributes.Bic, actualAccount.Attributes.Bic)
	assert.Equal(t, expectedAccount.Attributes.Name, actualAccount.Attributes.Name)
	assert.Equal(t, expectedAccount.Attributes.BankID, actualAccount.Attributes.BankID)
	assert.Equal(t, *expectedAccount.Attributes.Country, *actualAccount.Attributes.Country)
	assert.Equal(t, expectedAccount.Attributes.BankIDCode, actualAccount.Attributes.BankIDCode)
	assert.Equal(t, expectedAccount.Attributes.BaseCurrency, actualAccount.Attributes.BaseCurrency)
	assert.Equal(t, expectedAccount.Attributes.AccountNumber, actualAccount.Attributes.AccountNumber)
	assert.Equal(t, *expectedAccount.Version, *actualAccount.Version)
}

// TestAccountCreate_Error tests all Create action error cases.
func TestAccountCreate_Error(t *testing.T) {
	tests := []struct {
		name    string
		account *objects.Account
		error   *transport.ErrorResponse
	}{
		{
			name: "CreateAccountWithEmptyID",
			error: &transport.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "validation failure list:\nvalidation failure list:\nattributes in body is required\nid in body is required\norganisation_id in body is required\ntype in body is required",
			},
			account: &objects.Account{
				ID: "",
			},
		},
		{
			name: "CreateAccountWithEmptyOrganizationID",
			error: &transport.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "validation failure list:\nvalidation failure list:\nattributes in body is required\norganisation_id in body is required\ntype in body is required",
			},
			account: &objects.Account{
				ID:             uuid.NewString(),
				OrganisationID: "",
			},
		},
		{
			name: "CreateAccountWithEmptyType",
			error: &transport.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "validation failure list:\nvalidation failure list:\nattributes in body is required\ntype in body is required",
			},
			account: &objects.Account{
				ID:             uuid.NewString(),
				OrganisationID: uuid.NewString(),
				Type:           "",
			},
		},
		{
			name: "CreateAccountWithNotSpecifiedAttributes",
			error: &transport.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "validation failure list:\nvalidation failure list:\nattributes in body is required",
			},
			account: &objects.Account{
				ID:             uuid.NewString(),
				OrganisationID: uuid.NewString(),
				Type:           "accounts",
				Attributes:     nil,
			},
		},
		{
			name: "CreateAccountWithEmptyCountryInAttributes",
			error: &transport.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "validation failure list:\nvalidation failure list:\nvalidation failure list:\ncountry in body is required\nname in body is required",
			},
			account: &objects.Account{
				ID:             uuid.NewString(),
				OrganisationID: uuid.NewString(),
				Type:           "accounts",
				Attributes: &objects.AccountAttributes{
					Country: nil,
				},
			},
		},
		{
			name: "CreateAccountWithIncorrectCountryInAttributes",
			error: &transport.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "validation failure list:\nvalidation failure list:\nvalidation failure list:\ncountry in body should match '^[A-Z]{2}$'\nname in body is required",
			},
			account: &objects.Account{
				ID:             uuid.NewString(),
				OrganisationID: uuid.NewString(),
				Type:           "accounts",
				Attributes: &objects.AccountAttributes{
					Country: helpers.GetStringPointer(""),
				},
			},
		},
		{
			name: "CreateAccountWithIncorrectCountryInAttributes",
			error: &transport.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "validation failure list:\nvalidation failure list:\nvalidation failure list:\nname in body is required",
			},
			account: &objects.Account{
				ID:             uuid.NewString(),
				OrganisationID: uuid.NewString(),
				Type:           "accounts",
				Attributes: &objects.AccountAttributes{
					Country: helpers.GetStringPointer("FR"),
					Name:    []string{},
				},
			},
		},
	}

	fakeAPIClient := client.NewFakeAPIClient(os.Getenv("BASE_FAKE_API_URL"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fakeAPIClient.Accounts().Create(tt.account)
			originalErr, ok := errors.Cause(err).(transport.ErrorResponse)
			if ok {
				assert.Equal(t, tt.error.GetCode(), originalErr.GetCode())
				assert.Equal(t, tt.error.GetMessage(), originalErr.GetMessage())
			}
		})
	}
}
