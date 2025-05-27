package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostAccountsExportRequest() PostAccountsExportRequest {
	return PostAccountsExportRequest{
		client:      c,
		queryParams: c.NewPostAccountsExportQueryParams(),
		pathParams:  c.NewPostAccountsExportPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostAccountsExportRequestBody(),
	}
}

type PostAccountsExportRequest struct {
	client      *Client
	queryParams *PostAccountsExportQueryParams
	pathParams  *PostAccountsExportPathParams
	method      string
	headers     http.Header
	requestBody PostAccountsExportRequestBody
}

func (c *Client) NewPostAccountsExportQueryParams() *PostAccountsExportQueryParams {
	return &PostAccountsExportQueryParams{}
}

type PostAccountsExportQueryParams struct {
	// Specifies the property for which transactions will be Accountsexported
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

func (p PostAccountsExportQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostAccountsExportRequest) QueryParams() *PostAccountsExportQueryParams {
	return r.queryParams
}

func (c *Client) NewPostAccountsExportPathParams() *PostAccountsExportPathParams {
	return &PostAccountsExportPathParams{}
}

type PostAccountsExportPathParams struct {
}

func (p *PostAccountsExportPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostAccountsExportRequest) PathParams() *PostAccountsExportPathParams {
	return r.pathParams
}

func (r *PostAccountsExportRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostAccountsExportRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostAccountsExportRequest) Method() string {
	return r.method
}

func (s *Client) NewPostAccountsExportRequestBody() PostAccountsExportRequestBody {
	return PostAccountsExportRequestBody{}
}

type PostAccountsExportRequestBody struct {
	CreateCompanyModel
}

func (r *PostAccountsExportRequest) RequestBody() *PostAccountsExportRequestBody {
	return &r.requestBody
}

func (r *PostAccountsExportRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostAccountsExportRequest) SetRequestBody(body PostAccountsExportRequestBody) {
	r.requestBody = body
}

func (r *PostAccountsExportRequest) NewResponseBody() *PostAccountsExportResponseBody {
	return &PostAccountsExportResponseBody{}
}

type PostAccountsExportResponseBody PostAccountsExportModel

func (r *PostAccountsExportRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/export", r.PathParams())
	return &u
}

func (r *PostAccountsExportRequest) Do(ctx context.Context) (PostAccountsExportResponseBody, error) {
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
