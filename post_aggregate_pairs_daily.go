package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostAggregatePairsDailyRequest() PostAggregatePairsDailyRequest {
	return PostAggregatePairsDailyRequest{
		client:      c,
		queryParams: c.NewPostAggregatePairsDailyQueryParams(),
		pathParams:  c.NewPostAggregatePairsDailyPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostAggregatePairsDailyRequestBody(),
	}
}

type PostAggregatePairsDailyRequest struct {
	client      *Client
	queryParams *PostAggregatePairsDailyQueryParams
	pathParams  *PostAggregatePairsDailyPathParams
	method      string
	headers     http.Header
	requestBody PostAggregatePairsDailyRequestBody
}

func (c *Client) NewPostAggregatePairsDailyQueryParams() *PostAggregatePairsDailyQueryParams {
	return &PostAggregatePairsDailyQueryParams{}
}

type PostAggregatePairsDailyQueryParams struct {
	// Specifies the property for which transactions will be exported
	PropertyID string `schema:"propertyId"`
	// The inclusive start time of the posting date. Either posting date or
	// business date interval should be specified.
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	From Date `schema:"from"`
	// The exclusive end time of the posting date. Either posting date or
	// business date interval should be specified.
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	To Date `schema:"to"`
	// Filter transactions by reference (reservation id/external folio id/property id for house folio).
	Reference string `schema:"reference,omitempty"`
	// Filter transactions by account number
	AccountNumber string `schema:"accountNumber,omitempty"`
	// Filter transactions by type
	AccountType AccountType `schema:"accountType,omitempty"`
	// Allows to override the default accounting schema. Only specify this, when
	// you know what you are doing.
	AccountingSchema AccountingSchema `schema:"accountingSchema,omitempty"`
	//The language for the the report (2-letter ISO code)
	LanguageCode string `schema:"languageCode,omitempty"`
	// Unique key for safely retrying requests without accidentally performing
	// the same operation twice. We'll always send back the same response for
	// requests made with the same key, and keys can't be reused with different
	// request parameters. Keys expire after 24 hours.
	IdempotencyKey string `schema:"Idempotency-Key,omitempty"`
}

func (p PostAggregatePairsDailyQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostAggregatePairsDailyRequest) QueryParams() *PostAggregatePairsDailyQueryParams {
	return r.queryParams
}

func (c *Client) NewPostAggregatePairsDailyPathParams() *PostAggregatePairsDailyPathParams {
	return &PostAggregatePairsDailyPathParams{}
}

type PostAggregatePairsDailyPathParams struct {
}

func (p *PostAggregatePairsDailyPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostAggregatePairsDailyRequest) PathParams() *PostAggregatePairsDailyPathParams {
	return r.pathParams
}

func (r *PostAggregatePairsDailyRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostAggregatePairsDailyRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostAggregatePairsDailyRequest) Method() string {
	return r.method
}

func (s *Client) NewPostAggregatePairsDailyRequestBody() PostAggregatePairsDailyRequestBody {
	return PostAggregatePairsDailyRequestBody{}
}

type PostAggregatePairsDailyRequestBody struct {
	CreateCompanyModel
}

func (r *PostAggregatePairsDailyRequest) RequestBody() *PostAggregatePairsDailyRequestBody {
	return &r.requestBody
}

func (r *PostAggregatePairsDailyRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostAggregatePairsDailyRequest) SetRequestBody(body PostAggregatePairsDailyRequestBody) {
	r.requestBody = body
}

func (r *PostAggregatePairsDailyRequest) NewResponseBody() *PostAggregatePairsDailyResponseBody {
	return &PostAggregatePairsDailyResponseBody{}
}

type PostAggregatePairsDailyResponseBody PostAggregatePairsDailyModel

func (r *PostAggregatePairsDailyRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/aggregate-pairs-daily", r.PathParams())
	return &u
}

func (r *PostAggregatePairsDailyRequest) Do() (PostAggregatePairsDailyResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, true)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
