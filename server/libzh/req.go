package libzh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ZenHubAPIRequest holds ZenHub API request details
type ZenHubAPIRequest struct {
	// url is the endpoint path
	url string

	// authToken is a ZenHub API auth token
	authToken string
}

// NewZenHubAPIRequest creates a new ZenHubAPIRequest
func NewZenHubAPIRequest(url, authToken string) ZenHubAPIRequest {
	return ZenHubAPIRequest{
		url:       url,
		authToken: authToken,
	}
}

// Do makes a ZenHub API request
func (r ZenHubAPIRequest) Do(respVar interface{}) error {
	// Setup request
	reqURL, err := url.Parse(r.url)

	if err != nil {
		return fmt.Errorf("error parsing request URL: %s", err.Error())
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    reqURL,
		Header: map[string][]string{
			"X-Authentication-Token": []string{r.authToken},
		},
	}

	// Make request
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Errorf("error making API request: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		// If response error, read body to print
		respBody, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return fmt.Errorf("response status not OK, and an error occurred reading the body: %s", err.Error())
		}

		return fmt.Errorf("response status not OK, status: %s, body: %s", resp.Status, string(respBody))
	}

	// Decode API response
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&respVar)
	if err != nil {
		return fmt.Errorf("error decoding response body: %s", err.Error())
	}

	return nil
}
