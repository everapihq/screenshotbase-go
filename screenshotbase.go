package screenshotbase

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var apiKey string
var baseURL = "https://api.screenshotbase.com"

// Init sets the API key used for all requests.
func Init(key string) {
	apiKey = key
}

func SetBaseURL(u string) { // for testing
	baseURL = u
}

func client() *http.Client { return http.DefaultClient }

// Status calls GET /v1/status and returns the raw body bytes.
func Status() ([]byte, error) {
	if apiKey == "" {
		return nil, errors.New("API key not set, call screenshotbase.Init")
	}
	req, _ := http.NewRequest("GET", baseURL+"/v1/status", nil)
	req.Header.Set("apikey", apiKey)
	res, err := client().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("HTTP %d: %s", res.StatusCode, string(b))
	}
	return io.ReadAll(res.Body)
}

// Take calls GET /v1/take and returns image bytes.
// Params:
//   - url: required website URL
//   - format, quality, full_page, viewport_width, viewport_height: optional
func Take(params map[string]string) ([]byte, error) {
	if apiKey == "" {
		return nil, errors.New("API key not set, call screenshotbase.Init")
	}
	u, err := url.Parse(baseURL + "/v1/take")
	if err != nil {
		return nil, err
	}
	q := u.Query()
	if params == nil || params["url"] == "" {
		return nil, errors.New("missing required parameter 'url'")
	}
	for k, v := range params {
		if v != "" {
			q.Set(k, v)
		}
	}
	u.RawQuery = q.Encode()
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("apikey", apiKey)
	res, err := client().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("HTTP %d: %s", res.StatusCode, string(b))
	}
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, res.Body)
	return buf.Bytes(), err
}
