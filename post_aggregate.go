package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostAggregateRequest() PostAggregateRequest {
	return PostAggregateRequest{
		client:      c,
		queryParams: c.NewPostAggregateQueryParams(),
		pathParams:  c.NewPostAggregatePathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostAggregateRequestBody(),
	}
}

type PostAggregateRequest struct {
	client      *Client
	queryParams *PostAggregateQueryParams
	pathParams  *PostAggregatePathParams
	method      string
	headers     http.Header
	requestBody PostAggregateRequestBody
}

func (c *Client) NewPostAggregateQueryParams() *PostAggregateQueryParams {
	return &PostAggregateQueryParams{}
}

type PostAggregateQueryParams struct {
	// Specifies the property for which transactions will be exported
	PropertyID string `schema:"propertyId"`
	// The inclusive start time of the posting date. Either posting date or
	// business date interval should be specified.
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	From DateTime `schema:"from"`
	// The exclusive end time of the posting date. Either posting date or
	// business date interval should be specified.
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	To DateTime `schema:"to"`
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

func (p PostAggregateQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostAggregateRequest) QueryParams() *PostAggregateQueryParams {
	return r.queryParams
}

func (c *Client) NewPostAggregatePathParams() *PostAggregatePathParams {
	return &PostAggregatePathParams{}
}

type PostAggregatePathParams struct {
}

func (p *PostAggregatePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostAggregateRequest) PathParams() *PostAggregatePathParams {
	return r.pathParams
}

func (r *PostAggregateRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostAggregateRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostAggregateRequest) Method() string {
	return r.method
}

func (s *Client) NewPostAggregateRequestBody() PostAggregateRequestBody {
	return PostAggregateRequestBody{}
}

type PostAggregateRequestBody struct {
	CreateCompanyModel
}

func (r *PostAggregateRequest) RequestBody() *PostAggregateRequestBody {
	return &r.requestBody
}

func (r *PostAggregateRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostAggregateRequest) SetRequestBody(body PostAggregateRequestBody) {
	r.requestBody = body
}

func (r *PostAggregateRequest) NewResponseBody() *PostAggregateResponseBody {
	return &PostAggregateResponseBody{}
}

type PostAggregateResponseBody PostAggregateModel

func (r *PostAggregateRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/aggregate", r.PathParams())
	return &u
}

func (r *PostAggregateRequest) Do() (PostAggregateResponseBody, error) {
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
