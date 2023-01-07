package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// HTTP represents HTTP transport.
type HTTP struct {
	client  *http.Client
	baseURL string
}

// NewHTTP creates new HTTP transport.
func NewHTTP(baseURL string) *HTTP {
	return &HTTP{
		client:  &http.Client{},
		baseURL: baseURL,
	}
}

// Do makes actual HTTP request.
func (t HTTP) Do(method, endpoint string, data interface{}) ([]byte, error) {
	var serializedData []byte
	if data != nil {
		var err error
		serializedData, err = json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "error marshaling data object")
		}
	}

	request, err := http.NewRequest(
		method, fmt.Sprintf("%s%s", t.baseURL, endpoint), bytes.NewBuffer(serializedData),
	)
	if err != nil {
		return nil, errors.Wrap(err, "error creating http request")
	}
	response, err := t.client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "error calling remote API")
	}
	if response.Body == nil {
		return nil, errors.New("response body is null")
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading response body")
	}
	if statusCode := response.StatusCode; statusCode >= http.StatusBadRequest {
		switch statusCode {
		case http.StatusBadRequest:
			err := ErrorResponse{Code: response.StatusCode}
			if err := json.Unmarshal(body, &err); err != nil {
				return nil, errors.Wrap(err, "error unmarshaling error body")
			}
			return nil, err
		case http.StatusNotFound:
			return nil, ErrorResponse{Code: http.StatusNotFound, Message: "entity not found"}
		default:
			return nil, errors.Errorf("server returned unexpected error with status code: %d", statusCode)
		}
	}

	return body, nil
}
