package objects

// CreateAccountRequest represents request object for `POST /v1/organisation/accounts` endpoint.
type CreateAccountRequest struct {
	Data Account `json:"data"`
}
