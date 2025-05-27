package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostAccountsAggregateDailyRequest() PostAccountsAggregateDailyRequest {
	return PostAccountsAggregateDailyRequest{
		client:      c,
		queryParams: c.NewPostAccountsAggregateDailyQueryParams(),
		pathParams:  c.NewPostAccountsAggregateDailyPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostAccountsAggregateDailyRequestBody(),
	}
}

type PostAccountsAggregateDailyRequest struct {
	client      *Client
	queryParams *PostAccountsAggregateDailyQueryParams
	pathParams  *PostAccountsAggregateDailyPathParams
	method      string
	headers     http.Header
	requestBody PostAccountsAggregateDailyRequestBody
}

func (c *Client) NewPostAccountsAggregateDailyQueryParams() *PostAccountsAggregateDailyQueryParams {
	return &PostAccountsAggregateDailyQueryParams{}
}

type PostAccountsAggregateDailyQueryParams struct {
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

func (p PostAccountsAggregateDailyQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostAccountsAggregateDailyRequest) QueryParams() *PostAccountsAggregateDailyQueryParams {
	return r.queryParams
}

func (c *Client) NewPostAccountsAggregateDailyPathParams() *PostAccountsAggregateDailyPathParams {
	return &PostAccountsAggregateDailyPathParams{}
}

type PostAccountsAggregateDailyPathParams struct {
}

func (p *PostAccountsAggregateDailyPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostAccountsAggregateDailyRequest) PathParams() *PostAccountsAggregateDailyPathParams {
	return r.pathParams
}

func (r *PostAccountsAggregateDailyRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostAccountsAggregateDailyRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostAccountsAggregateDailyRequest) Method() string {
	return r.method
}

func (s *Client) NewPostAccountsAggregateDailyRequestBody() PostAccountsAggregateDailyRequestBody {
	return PostAccountsAggregateDailyRequestBody{}
}

type PostAccountsAggregateDailyRequestBody struct {
	CreateCompanyModel
}

func (r *PostAccountsAggregateDailyRequest) RequestBody() *PostAccountsAggregateDailyRequestBody {
	return &r.requestBody
}

func (r *PostAccountsAggregateDailyRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostAccountsAggregateDailyRequest) SetRequestBody(body PostAccountsAggregateDailyRequestBody) {
	r.requestBody = body
}

func (r *PostAccountsAggregateDailyRequest) NewResponseBody() *PostAccountsAggregateDailyResponseBody {
	return &PostAccountsAggregateDailyResponseBody{}
}

type PostAccountsAggregateDailyResponseBody PostAccountsAggregateDailyModel

func (r *PostAccountsAggregateDailyRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/aggregate-daily", r.PathParams())
	return &u
}

func (r *PostAccountsAggregateDailyRequest) Do(ctx context.Context) (PostAccountsAggregateDailyResponseBody, error) {
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
