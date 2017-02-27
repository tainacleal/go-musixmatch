package musixmatch

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	//BaseURL is the API URL
	BaseURL            = "http://api.musixmatch.com/ws/1.1"
	defaultHTTPTimeout = 30 * time.Second
)

// Key is the MusixMatch API key
var Key string
var httpClient = &http.Client{Timeout: defaultHTTPTimeout}

// BackendService is an interface to be implemented by MusixMatch clients
type BackendService interface {
	Call(method string, path string, params string, v *Return) error
}

// Backend holds the user API key and the HTTP client
type Backend struct {
	Key        string
	HTTPClient *http.Client
}

// GetBackend returns a Backend with the default HTTPClient
func GetBackend() Backend {
	return Backend{
		Key:        Key,
		HTTPClient: httpClient,
	}
}

// SetBackend sets a Backend type with a custom HTTPClient
func SetBackend(client *http.Client) Backend {
	return Backend{
		Key:        Key,
		HTTPClient: client,
	}
}

// Call is a helper function to execute an API request
func (b Backend) Call(method string, path string, params string, v *Return) error {
	if params == "" {
		return errors.New("params are required")
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s/%s?%s&apikey=%s", BaseURL, path, params, b.Key),
		nil,
	)
	if err != nil {
		return err
	}

	resp, err := b.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}

	return CheckStatusCode(v.Header.StatusCode)
}

// Return is the type that will hold the response from the API request
type Return struct {
	Message `json:"message"`
}

// Message represents the JSON object returned from the API request
type Message struct {
	Header Header           `json:"header"`
	Body   *json.RawMessage `json:"body"`
}

// Header contains the Status Code returned from the API request
type Header struct {
	StatusCode int `json:"status_code"`
}

// CheckStatusCode checks for response errors
func CheckStatusCode(code int) error {
	switch code {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return errors.New("400: bad request")
	case http.StatusUnauthorized:
		return errors.New("401: authetication failed")
	case http.StatusPaymentRequired:
		return errors.New("402: usage limit has been reached")
	case http.StatusForbidden:
		return errors.New("403: not authorized to perform this operation")
	case http.StatusNotFound:
		return errors.New("404: resource not found")
	case http.StatusMethodNotAllowed:
		return errors.New("405: requested method not found")
	case http.StatusServiceUnavailable:
		return errors.New("503: service unavailable")
	default:
		return errors.New("500: internal server error")
	}
}
