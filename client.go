package apaleo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-apaleo/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"

	requestTimestampsLimit         = 20
	requestTimestampsLimitDuration = time.Second
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "api.apaleo.com",
		Path:   "",
	}
)

// NewClient returns a new InvoiceXpress Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		http:              httpClient,
		requestTimestamps: []time.Time{},
	}

	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)
	client.SetRateLimit(ClientRateLimit{
		Limit:     "1m",
		Remaining: 100,
		Reset:     time.Now(),
	})

	return client
}

// Client manages communication with InvoiceXpress Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Rate limiting
	rateLimit         ClientRateLimit
	requestTimestamps []time.Time
	requestTimestampsMu sync.RWMutex

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c *Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c *Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetRateLimit(rateLimit ClientRateLimit) {
	c.rateLimit = rateLimit
}

func (c *Client) RateLimit() *ClientRateLimit {
	return &c.rateLimit
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) GetEndpointURL(relative string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()
	relativeURL, err := url.Parse(relative)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = path.Join(clientURL.Path, relativeURL.Path)

	query := url.Values{}
	for k, v := range clientURL.Query() {
		query[k] = append(query[k], v...)
	}
	for k, v := range relativeURL.Query() {
		query[k] = append(query[k], v...)
	}
	clientURL.RawQuery = query.Encode()

	tmpl, err := template.New("endpoint_url").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}
	clientURL.Path = buf.String()

	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if req.RequestBodyInterface() != nil {
		err := json.NewEncoder(buf).Encode(req.RequestBodyInterface())
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	r, err := http.NewRequest(req.Method(), req.URL().String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())

	return r, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	// Sleep for rate limits
	c.sleepUntilRequestLimit()

	// Register request timestamp
	c.registerRequestTimestamp(time.Now())

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	// Register request limit
	c.registerRequestLimit(httpResp)

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// Handle '429 - Too many requests' response
	if httpResp.StatusCode == 429 {
		c.sleepUntilRetryAfter(httpResp)
		return c.Do(req, responseBody)
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	if responseBody == nil {
		return httpResp, nil
	}

	errorResponse := &ErrorResponse{Response: httpResp}
	statusErrResponse := &StatusErrorResponse{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, []any{responseBody}, []any{errorResponse, statusErrResponse})
	if err != nil {
		return httpResp, err
	}

	if errorResponse.Error() != "" {
		return httpResp, errorResponse
	}

	if statusErrResponse.Error() != "" {
		return httpResp, statusErrResponse
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv []interface{}, optionalVv []interface{}) error {
	if len(vv) == 0 && len(optionalVv) == 0 {
		return nil
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			return err
		}
	}

	for _, v := range optionalVv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)

		_ = dec.Decode(v)
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; (c >= 200 && c <= 299) || c == 400 {
		return nil
	}

	err := checkContentType(r)
	if err != nil {
		return errors.New(r.Status)
	}

	// read data and copy it back
	data, err := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return nil
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

type StatusErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response
}

func (r *StatusErrorResponse) Error() string {
	if r.Response.StatusCode != 0 && (r.Response.StatusCode < 200 || r.Response.StatusCode > 299) {
		return fmt.Sprintf("%s", r.Response.Status)
	}

	return ""
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Messages []string `json:"messages"`
}

func (r *ErrorResponse) Error() string {
	if len(r.Messages) > 0 {
		return strings.Join(r.Messages, ", ")
	}

	return ""
}

type Message struct {
	MessageCode string `json:"message_code"`
	MessageType string `json:"message_type"`
	Message     string `json:"message"`
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

func (c *Client) sleepUntilRequestLimit() {
	// When remaining requests is more than 1, continue
	if c.RateLimit().Remaining > 1 {
		return
	}

	// Sleep for time until it reaches reset
	time.Sleep(time.Until(c.RateLimit().Reset))
}

func (c *Client) registerRequestLimit(req *http.Response) error {
	// Create variables
	var err error
	rateLimit := ClientRateLimit{}

	// Set limit
	rateLimit.Limit = req.Header.Get("X-Rate-Limit-Limit")

	// Set remaining
	rateLimit.Remaining, err = strconv.Atoi(req.Header.Get("X-Rate-Limit-Remaining"))
	if err != nil {
		return err
	}

	// Set reset timestamp
	rateLimit.Reset, err = time.Parse(time.RFC3339Nano, req.Header.Get("X-Rate-Limit-Reset"))
	if err != nil {
		return err
	}

	// Set rate limit
	c.SetRateLimit(rateLimit)

	return nil
}

func (c *Client) registerRequestTimestamp(t time.Time) {
	c.requestTimestampsMu.Lock()
	defer c.requestTimestampsMu.Unlock()

	if len(c.requestTimestamps) >= requestTimestampsLimit {
		c.requestTimestamps = c.requestTimestamps[1:requestTimestampsLimit]
	}
	c.requestTimestamps = append(c.requestTimestamps, t)
}

func (c *Client) sleepUntilRequestRate() {
	c.requestTimestampsMu.RLock()
	defer c.requestTimestampsMu.RUnlock()

	// Requestrate is <requestTimestampsLimit> requests per <requestTimestampsLimitDuration>

	// if there are less then <requestTimestampsLimit> registered requests: execute the request
	// immediately
	if len(c.requestTimestamps) < requestTimestampsLimit {
		return
	}

	// is the first item within <requestTimestampsLimitDuration>? If it's more than <requestTimestampsLimitDuration>: the request can be
	// executed immediately
	diff := time.Since(c.requestTimestamps[0])
	if diff >= requestTimestampsLimitDuration {
		return
	}

	// Sleep for the time it takes for the first item to be > <requestTimestampsLimitDuration> old
	time.Sleep(requestTimestampsLimitDuration - diff)
}

func (c *Client) sleepUntilRetryAfter(req *http.Response) error {
	// Get the "Retry-After" header
	retryAfter := req.Header.Get("Retry-After")

	// When the "Retry-After" header is not set, continue
	if retryAfter == "" {
		return nil
	}

	// Parse the "Retry-After" header to a duration in seconds
	diff, err := time.ParseDuration(fmt.Sprintf("%ss", retryAfter))
	if err != nil {
		return err
	}

	// Sleep for the duration of "Retry-After"
	time.Sleep(diff)

	return nil
}
