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

// TestAccountDelete_Ok tests Delete action happy path case.
func TestAccountDelete_Ok(t *testing.T) {
	fakeAPIClient := client.NewFakeAPIClient(os.Getenv("BASE_FAKE_API_URL"))
	// 1. create a new Account entity.
	account, err := fakeAPIClient.Accounts().Create(&objects.Account{
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

	// 2. delete newly created Account entity.
	err = fakeAPIClient.Accounts().Delete(account.ID, *account.Version)
	assert.Nil(t, err)
}

// TestAccountDelete_Error tests all Delete action error cases.
func TestAccountDelete_Error(t *testing.T) {
	tests := []struct {
		name      string
		accountID string
		version   int64
		error     *transport.ErrorResponse
	}{
		{
			name:      "DeleteAccountWithNotExistingID",
			accountID: uuid.NewString(),
			version:   0,
			error: &transport.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "entity not found",
			},
		},
		{
			name:      "DeleteAccountWithIncorrectIDFormat",
			accountID: "not_existing_account_id",
			version:   0,
			error: &transport.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "id is not a valid uuid",
			},
		},
	}

	fakeAPIClient := client.NewFakeAPIClient(os.Getenv("BASE_FAKE_API_URL"))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := fakeAPIClient.Accounts().Delete(tt.accountID, tt.version)
			originalErr, ok := errors.Cause(err).(transport.ErrorResponse)
			if ok {
				assert.Equal(t, tt.error.GetCode(), originalErr.GetCode())
				assert.Equal(t, tt.error.GetMessage(), originalErr.GetMessage())
			}
		})
	}
}
