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
	"github.com/form3-test-task/tests/helpers"
)

// TestAccountFetch_Ok tests Fetch action happy path case.
func TestAccountFetch_Ok(t *testing.T) {
	fakeAPIClient := client.NewFakeAPIClient(os.Getenv("BASE_FAKE_API_URL"))
	// 1. create a new Account entity.
	expectedAccount, err := fakeAPIClient.Accounts().Create(&objects.Account{
		ID:   uuid.New().String(),
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

	// 2. fetch existing Account entity by its id.
	actualAccount, err := fakeAPIClient.Accounts().Fetch(expectedAccount.ID)
	assert.Nil(t, err)
	assert.Equal(t, expectedAccount.ID, actualAccount.ID)
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

// TestAccountFetch_Error tests all Fetch action error cases.
func TestAccountFetch_Error(t *testing.T) {
	tests := []struct {
		name      string
		accountID string
		error     *transport.ErrorResponse
	}{
		{
			name:      "FetchAccountWithNotExistingID",
			accountID: uuid.NewString(),
			error: &transport.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "entity not found",
			},
		},
		{
			name:      "FetchAccountWithIncorrectIDFormat",
			accountID: "not_existing_account_id",
			error: &transport.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "id is not a valid uuid",
			},
		},
	}

	fakeAPIClient := client.NewFakeAPIClient(os.Getenv("BASE_FAKE_API_URL"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := fakeAPIClient.Accounts().Fetch(tt.accountID)
			originalErr, ok := errors.Cause(err).(transport.ErrorResponse)
			if ok {
				assert.Equal(t, tt.error.GetCode(), originalErr.GetCode())
				assert.Equal(t, tt.error.GetMessage(), originalErr.GetMessage())
			}
		})
	}
}
