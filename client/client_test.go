//go:build unit

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/form3-test-task/client/entities/accounts/objects"
	"github.com/form3-test-task/client/transport"
	"github.com/form3-test-task/tests/helpers"
)

type FakeClientTestSuite struct {
	suite.Suite
	fakeTransport *transport.MockHTTPProvider
	fakeAPIClient *FakeAPIClient
}

func TestFakeClientTestSuite(t *testing.T) {
	suite.Run(t, new(FakeClientTestSuite))
}

// Inject fresh mocks before each test.
func (s *FakeClientTestSuite) SetupTest() {
	s.fakeTransport = &transport.MockHTTPProvider{}
	s.fakeAPIClient = NewFakeAPIClient(
		WithTransport(s.fakeTransport),
	)
}

func (s *FakeClientTestSuite) TestCreateAccount_Ok() {
	// 1. configure fake transport.
	accountID := uuid.New().String()
	account := &objects.Account{
		ID:   accountID,
		Type: "accounts",
		Attributes: &objects.AccountAttributes{
			Bic:           "NWBKFR42",
			Iban:          "FR1420041010050500013M02606",
			Name:          []string{"Name1", "Name2"},
			BankID:        "20041",
			Country:       helpers.GetVariablePointer("FR"),
			BankIDCode:    "FR",
			BaseCurrency:  "EUR",
			AccountNumber: "0500013M026",
		},
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	}
	responseObject := objects.CreateAccountResponse{Data: *account}
	serializedResponseObject, err := json.Marshal(responseObject)
	assert.Nil(s.T(), err)
	s.fakeTransport.On(
		"Do", http.MethodPost, "/organisation/accounts", objects.CreateAccountRequest{Data: *account},
	).Return(serializedResponseObject, nil)

	// 2. call client.
	account, err = s.fakeAPIClient.Accounts().Create(account)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), accountID, account.ID)
	assert.Equal(s.T(), "accounts", account.Type)
	assert.Equal(s.T(), "NWBKFR42", account.Attributes.Bic)
	assert.Equal(s.T(), "FR1420041010050500013M02606", account.Attributes.Iban)
	assert.Equal(s.T(), []string{"Name1", "Name2"}, account.Attributes.Name)
	assert.Equal(s.T(), "20041", account.Attributes.BankID)
	assert.Equal(s.T(), "FR", *account.Attributes.Country)
	assert.Equal(s.T(), "FR", account.Attributes.BankIDCode)
	assert.Equal(s.T(), "EUR", account.Attributes.BaseCurrency)
	assert.Equal(s.T(), "0500013M026", account.Attributes.AccountNumber)
	assert.Equal(s.T(), "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c", account.OrganisationID)
}
func (s *FakeClientTestSuite) TestFetchAccount_Ok() {
	// 1. configure fake transport.
	accountID := uuid.New().String()
	account := &objects.Account{
		ID:   accountID,
		Type: "accounts",
		Attributes: &objects.AccountAttributes{
			Bic:           "NWBKFR42",
			Iban:          "FR1420041010050500013M02606",
			Name:          []string{"Name1", "Name2"},
			BankID:        "20041",
			Country:       helpers.GetVariablePointer("FR"),
			BankIDCode:    "FR",
			BaseCurrency:  "EUR",
			AccountNumber: "0500013M026",
		},
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	}
	responseObject := objects.CreateAccountResponse{Data: *account}
	serializedResponseObject, err := json.Marshal(responseObject)
	assert.Nil(s.T(), err)
	s.fakeTransport.On(
		"Do", http.MethodGet, fmt.Sprintf("/organisation/accounts/%s", accountID), nil,
	).Return(serializedResponseObject, nil)

	// 2. call client.
	account, err = s.fakeAPIClient.Accounts().Fetch(accountID)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), accountID, account.ID)
	assert.Equal(s.T(), "accounts", account.Type)
	assert.Equal(s.T(), "NWBKFR42", account.Attributes.Bic)
	assert.Equal(s.T(), "FR1420041010050500013M02606", account.Attributes.Iban)
	assert.Equal(s.T(), []string{"Name1", "Name2"}, account.Attributes.Name)
	assert.Equal(s.T(), "20041", account.Attributes.BankID)
	assert.Equal(s.T(), "FR", *account.Attributes.Country)
	assert.Equal(s.T(), "FR", account.Attributes.BankIDCode)
	assert.Equal(s.T(), "EUR", account.Attributes.BaseCurrency)
	assert.Equal(s.T(), "0500013M026", account.Attributes.AccountNumber)
	assert.Equal(s.T(), "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c", account.OrganisationID)
}
func (s *FakeClientTestSuite) TestDeleteAccount_Ok() {
	// 1. configure fake transport.
	version := int64(0)
	accountID := uuid.New().String()
	s.fakeTransport.On(
		"Do", http.MethodDelete, fmt.Sprintf("/organisation/accounts/%s?version=%d", accountID, version), nil,
	).Return([]byte{}, nil)

	// 2. call client.
	err := s.fakeAPIClient.Accounts().Delete(accountID, version)
	assert.Nil(s.T(), err)
}
