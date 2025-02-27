package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostAccountsExportDailyRequest() PostAccountsExportDailyRequest {
	return PostAccountsExportDailyRequest{
		client:      c,
		queryParams: c.NewPostAccountsExportDailyQueryParams(),
		pathParams:  c.NewPostAccountsExportDailyPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostAccountsExportDailyRequestBody(),
	}
}

type PostAccountsExportDailyRequest struct {
	client      *Client
	queryParams *PostAccountsExportDailyQueryParams
	pathParams  *PostAccountsExportDailyPathParams
	method      string
	headers     http.Header
	requestBody PostAccountsExportDailyRequestBody
}

func (c *Client) NewPostAccountsExportDailyQueryParams() *PostAccountsExportDailyQueryParams {
	return &PostAccountsExportDailyQueryParams{}
}

type PostAccountsExportDailyQueryParams struct {
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

func (p PostAccountsExportDailyQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostAccountsExportDailyRequest) QueryParams() *PostAccountsExportDailyQueryParams {
	return r.queryParams
}

func (c *Client) NewPostAccountsExportDailyPathParams() *PostAccountsExportDailyPathParams {
	return &PostAccountsExportDailyPathParams{}
}

type PostAccountsExportDailyPathParams struct {
}

func (p *PostAccountsExportDailyPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostAccountsExportDailyRequest) PathParams() *PostAccountsExportDailyPathParams {
	return r.pathParams
}

func (r *PostAccountsExportDailyRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostAccountsExportDailyRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostAccountsExportDailyRequest) Method() string {
	return r.method
}

func (s *Client) NewPostAccountsExportDailyRequestBody() PostAccountsExportDailyRequestBody {
	return PostAccountsExportDailyRequestBody{}
}

type PostAccountsExportDailyRequestBody struct {
	CreateCompanyModel
}

func (r *PostAccountsExportDailyRequest) RequestBody() *PostAccountsExportDailyRequestBody {
	return &r.requestBody
}

func (r *PostAccountsExportDailyRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostAccountsExportDailyRequest) SetRequestBody(body PostAccountsExportDailyRequestBody) {
	r.requestBody = body
}

func (r *PostAccountsExportDailyRequest) NewResponseBody() *PostAccountsExportDailyResponseBody {
	return &PostAccountsExportDailyResponseBody{}
}

type PostAccountsExportDailyResponseBody PostAccountsExportDailyModel

func (r *PostAccountsExportDailyRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/export-daily", r.PathParams())
	return &u
}

func (r *PostAccountsExportDailyRequest) Do() (PostAccountsExportDailyResponseBody, error) {
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
