package deltaexchapigo

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	_ "fmt"
	"net/http"
	"strconv"
	"time"
)

var (
	requestTimeout time.Duration = 7000 * time.Millisecond
	baseURI        string        = "https://api.india.delta.exchange"
)

type Client struct {
	baseURL    string
	APIKey     string
	APISecret  string
	UserAgent  string
	httpClient HTTPClient
	debug      bool
}

// New initializes and returns a new Client instance
func New(baseURL, apiKey, apiSecret, userAgent string) *Client {

	client := &Client{
		baseURL:   baseURI,
		APIKey:    apiKey,
		APISecret: apiSecret,
		UserAgent: userAgent,
	}

	// Create a default http handler with default timeout.
	client.SetHTTPClient(&http.Client{
		Timeout:   requestTimeout,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	})

	return client
}

// SetHTTPClient overrides default http handler with a custom one.
// This can be used to set custom timeouts and transport.
func (c *Client) SetHTTPClient(h *http.Client) {
	c.httpClient = NewHTTPClient(h, nil, c.debug)
}

// SetDebug sets debug mode to enable HTTP logs.
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
	c.httpClient.GetClient().debug = debug
}

// SetBaseURI overrides the base norenAPI endpoint with custom url.
func (c *Client) SetBaseURI(baseURI string) {
	c.baseURL = baseURI
}

// GenerateSignature generates the HMAC SHA256 signature
func (c *Client) GenerateSignature(message string) string {
	h := hmac.New(sha256.New, []byte(c.APISecret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// GetTimestamp gets the current Unix timestamp as a string
func GetTimestamp() string {
	return strconv.FormatInt(time.Now().UTC().Unix(), 10)
}

func (c *Client) doEnvelope(method, uri string, params map[string]interface{}, headers http.Header, v interface{}, authorization ...bool) error {
	if params == nil {
		params = map[string]interface{}{}
	}

	// Send custom headers set
	if headers == nil {
		headers = map[string][]string{}
	}

	// Add Kite Connect version to header
	headers.Add("Content-Type", "application/json")

	// headers.Add("charset", "utf-8")

	fmt.Printf("\n--> Method : %s \nURL : %s \nParam : %v\n Header : %v\n, V: %v\n",
		method, c.baseURL+uri, params, headers, v)

	return c.httpClient.DoEnvelope(method, c.baseURL+uri, params, headers, v)
}
