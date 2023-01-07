package client

import (
	"github.com/form3-test-task/client/entities/accounts"
	"github.com/form3-test-task/client/transport"
)

// Let's point by default client to fake api.
const (
	baseAPIHost = "http://account-api:8080/v1"
)

// FakeAPIClient represents client to work with fake API.
type FakeAPIClient struct {
	accounts  *accounts.Accounts
	transport transport.HTTPProvider
}

// NewFakeAPIClient creates new instance of fake API client.
func NewFakeAPIClient(options ...func(*FakeAPIClient)) *FakeAPIClient {
	fakeAPIClient := FakeAPIClient{
		transport: transport.NewHTTP(baseAPIHost),
	}
	for _, option := range options {
		option(&fakeAPIClient)
	}
	return &fakeAPIClient
}

// Accounts returns object to work with Account entity.
func (c FakeAPIClient) Accounts() *accounts.Accounts {
	if c.accounts == nil {
		c.accounts = accounts.NewAccounts(c.transport)
		return c.accounts
	}
	return c.accounts
}

// WithTransport makes it possible to configure Fake API client with custom client.
// In current case it only needs for unit tests, when we should use fake transport instead of real one.
func WithTransport(transport transport.HTTPProvider) func(*FakeAPIClient) {
	return func(s *FakeAPIClient) {
		s.transport = transport
	}
}
