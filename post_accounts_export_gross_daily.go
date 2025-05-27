package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostAccountsExportGrossDailyRequest() PostAccountsExportGrossDailyRequest {
	return PostAccountsExportGrossDailyRequest{
		client:      c,
		queryParams: c.NewPostAccountsExportGrossDailyQueryParams(),
		pathParams:  c.NewPostAccountsExportGrossDailyPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostAccountsExportGrossDailyRequestBody(),
	}
}

type PostAccountsExportGrossDailyRequest struct {
	client      *Client
	queryParams *PostAccountsExportGrossDailyQueryParams
	pathParams  *PostAccountsExportGrossDailyPathParams
	method      string
	headers     http.Header
	requestBody PostAccountsExportGrossDailyRequestBody
}

func (c *Client) NewPostAccountsExportGrossDailyQueryParams() *PostAccountsExportGrossDailyQueryParams {
	return &PostAccountsExportGrossDailyQueryParams{}
}

type PostAccountsExportGrossDailyQueryParams struct {
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
	// Allows to override the default accounting schema. Only specify this, when
	// you know what you are doing.
	AccountingSchema AccountingSchema `schema:"accountingSchema,omitempty"`
	// Unique key for safely retrying requests without accidentally performing
	// the same operation twice. We'll always send back the same response for
	// requests made with the same key, and keys can't be reused with different
	// request parameters. Keys expire after 24 hours.
	IdempotencyKey string `schema:"Idempotency-Key,omitempty"`
}

func (p PostAccountsExportGrossDailyQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostAccountsExportGrossDailyRequest) QueryParams() *PostAccountsExportGrossDailyQueryParams {
	return r.queryParams
}

func (c *Client) NewPostAccountsExportGrossDailyPathParams() *PostAccountsExportGrossDailyPathParams {
	return &PostAccountsExportGrossDailyPathParams{}
}

type PostAccountsExportGrossDailyPathParams struct {
}

func (p *PostAccountsExportGrossDailyPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostAccountsExportGrossDailyRequest) PathParams() *PostAccountsExportGrossDailyPathParams {
	return r.pathParams
}

func (r *PostAccountsExportGrossDailyRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostAccountsExportGrossDailyRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostAccountsExportGrossDailyRequest) Method() string {
	return r.method
}

func (s *Client) NewPostAccountsExportGrossDailyRequestBody() PostAccountsExportGrossDailyRequestBody {
	return PostAccountsExportGrossDailyRequestBody{}
}

type PostAccountsExportGrossDailyRequestBody struct {
	CreateCompanyModel
}

func (r *PostAccountsExportGrossDailyRequest) RequestBody() *PostAccountsExportGrossDailyRequestBody {
	return &r.requestBody
}

func (r *PostAccountsExportGrossDailyRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostAccountsExportGrossDailyRequest) SetRequestBody(body PostAccountsExportGrossDailyRequestBody) {
	r.requestBody = body
}

func (r *PostAccountsExportGrossDailyRequest) NewResponseBody() *PostAccountsExportGrossDailyResponseBody {
	return &PostAccountsExportGrossDailyResponseBody{}
}

type PostAccountsExportGrossDailyResponseBody PostAccountsExportGrossDailyModel

func (r *PostAccountsExportGrossDailyRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/export-gross-daily", r.PathParams())
	return &u
}

func (r *PostAccountsExportGrossDailyRequest) Do(ctx context.Context) (PostAccountsExportGrossDailyResponseBody, error) {
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
