package client

import (
	"github.com/form3-test-task/client/entities/accounts"
	"github.com/form3-test-task/client/entities/transport"
)

// FakeAPIClient represents client to work with fake API.
type FakeAPIClient struct {
	accounts *accounts.Accounts
}

// NewFakeAPIClient creates new instance of fake API client.
func NewFakeAPIClient(baseURL string) *FakeAPIClient {
	return &FakeAPIClient{
		accounts: accounts.NewAccounts(
			transport.NewHTTP(baseURL),
		),
	}
}

// Accounts returns object to work with Account entity.
func (c FakeAPIClient) Accounts() *accounts.Accounts {
	return c.accounts
}
