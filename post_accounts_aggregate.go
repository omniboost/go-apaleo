package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostAccountsAggregateRequest() PostAccountsAggregateRequest {
	return PostAccountsAggregateRequest{
		client:      c,
		queryParams: c.NewPostAccountsAggregateQueryParams(),
		pathParams:  c.NewPostAccountsAggregatePathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostAccountsAggregateRequestBody(),
	}
}

type PostAccountsAggregateRequest struct {
	client      *Client
	queryParams *PostAccountsAggregateQueryParams
	pathParams  *PostAccountsAggregatePathParams
	method      string
	headers     http.Header
	requestBody PostAccountsAggregateRequestBody
}

func (c *Client) NewPostAccountsAggregateQueryParams() *PostAccountsAggregateQueryParams {
	return &PostAccountsAggregateQueryParams{}
}

type PostAccountsAggregateQueryParams struct {
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
	// Filter transactions by accounts number
	AccountsNumber string `schema:"accountsNumber,omitempty"`
	// Filter transactions by type
	AccountsType AccountType `schema:"accountsType,omitempty"`
	// Allows to override the default accountsing schema. Only specify this, when
	// you know what you are doing.
	AccountsingSchema AccountingSchema `schema:"accountsingSchema,omitempty"`
	//The language for the the report (2-letter ISO code)
	LanguageCode string `schema:"languageCode,omitempty"`
	// Unique key for safely retrying requests without accidentally performing
	// the same operation twice. We'll always send back the same response for
	// requests made with the same key, and keys can't be reused with different
	// request parameters. Keys expire after 24 hours.
	IdempotencyKey string `schema:"Idempotency-Key,omitempty"`
}

func (p PostAccountsAggregateQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostAccountsAggregateRequest) QueryParams() *PostAccountsAggregateQueryParams {
	return r.queryParams
}

func (c *Client) NewPostAccountsAggregatePathParams() *PostAccountsAggregatePathParams {
	return &PostAccountsAggregatePathParams{}
}

type PostAccountsAggregatePathParams struct {
}

func (p *PostAccountsAggregatePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostAccountsAggregateRequest) PathParams() *PostAccountsAggregatePathParams {
	return r.pathParams
}

func (r *PostAccountsAggregateRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostAccountsAggregateRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostAccountsAggregateRequest) Method() string {
	return r.method
}

func (s *Client) NewPostAccountsAggregateRequestBody() PostAccountsAggregateRequestBody {
	return PostAccountsAggregateRequestBody{}
}

type PostAccountsAggregateRequestBody struct {
	CreateCompanyModel
}

func (r *PostAccountsAggregateRequest) RequestBody() *PostAccountsAggregateRequestBody {
	return &r.requestBody
}

func (r *PostAccountsAggregateRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostAccountsAggregateRequest) SetRequestBody(body PostAccountsAggregateRequestBody) {
	r.requestBody = body
}

func (r *PostAccountsAggregateRequest) NewResponseBody() *PostAccountsAggregateResponseBody {
	return &PostAccountsAggregateResponseBody{}
}

type PostAccountsAggregateResponseBody PostAccountsAggregateModel

func (r *PostAccountsAggregateRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/aggregate", r.PathParams())
	return &u
}

func (r *PostAccountsAggregateRequest) Do(ctx context.Context) (PostAccountsAggregateResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
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
