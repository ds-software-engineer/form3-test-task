package objects

// CreateAccountResponse represents response object for `POST /v1/organisation/accounts` endpoint.
type CreateAccountResponse struct {
	Data Account `json:"data"`
}

// FetchAccountResponse represents response object for `GET /v1/organisation/accounts/{id}` endpoint.
type FetchAccountResponse struct {
	Data Account `json:"data"`
}
