package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostAggregateDailyRequest() PostAggregateDailyRequest {
	return PostAggregateDailyRequest{
		client:      c,
		queryParams: c.NewPostAggregateDailyQueryParams(),
		pathParams:  c.NewPostAggregateDailyPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostAggregateDailyRequestBody(),
	}
}

type PostAggregateDailyRequest struct {
	client      *Client
	queryParams *PostAggregateDailyQueryParams
	pathParams  *PostAggregateDailyPathParams
	method      string
	headers     http.Header
	requestBody PostAggregateDailyRequestBody
}

func (c *Client) NewPostAggregateDailyQueryParams() *PostAggregateDailyQueryParams {
	return &PostAggregateDailyQueryParams{}
}

type PostAggregateDailyQueryParams struct {
	// Specifies the property for which transactions will be exported
	PropertyID string `schema:"propertyId"`
	// The inclusive start time of the posting date. Either posting date or
	// business date interval should be specified.
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	From string `schema:"from"`
	// The exclusive end time of the posting date. Either posting date or
	// business date interval should be specified.
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	To string `schema:"to"`
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

func (p PostAggregateDailyQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostAggregateDailyRequest) QueryParams() *PostAggregateDailyQueryParams {
	return r.queryParams
}

func (c *Client) NewPostAggregateDailyPathParams() *PostAggregateDailyPathParams {
	return &PostAggregateDailyPathParams{}
}

type PostAggregateDailyPathParams struct {
}

func (p *PostAggregateDailyPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostAggregateDailyRequest) PathParams() *PostAggregateDailyPathParams {
	return r.pathParams
}

func (r *PostAggregateDailyRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostAggregateDailyRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostAggregateDailyRequest) Method() string {
	return r.method
}

func (s *Client) NewPostAggregateDailyRequestBody() PostAggregateDailyRequestBody {
	return PostAggregateDailyRequestBody{}
}

type PostAggregateDailyRequestBody struct {
	CreateCompanyModel
}

func (r *PostAggregateDailyRequest) RequestBody() *PostAggregateDailyRequestBody {
	return &r.requestBody
}

func (r *PostAggregateDailyRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostAggregateDailyRequest) SetRequestBody(body PostAggregateDailyRequestBody) {
	r.requestBody = body
}

func (r *PostAggregateDailyRequest) NewResponseBody() *PostAggregateDailyResponseBody {
	return &PostAggregateDailyResponseBody{}
}

type PostAggregateDailyResponseBody PostAggregateDailyModel

func (r *PostAggregateDailyRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/aggregate-daily", r.PathParams())
	return &u
}

func (r *PostAggregateDailyRequest) Do() (PostAggregateDailyResponseBody, error) {
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
