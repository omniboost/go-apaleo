package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostExportRequest() PostExportRequest {
	return PostExportRequest{
		client:      c,
		queryParams: c.NewPostExportQueryParams(),
		pathParams:  c.NewPostExportPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostExportRequestBody(),
	}
}

type PostExportRequest struct {
	client      *Client
	queryParams *PostExportQueryParams
	pathParams  *PostExportPathParams
	method      string
	headers     http.Header
	requestBody PostExportRequestBody
}

func (c *Client) NewPostExportQueryParams() *PostExportQueryParams {
	return &PostExportQueryParams{}
}

type PostExportQueryParams struct {
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

func (p PostExportQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostExportRequest) QueryParams() *PostExportQueryParams {
	return r.queryParams
}

func (c *Client) NewPostExportPathParams() *PostExportPathParams {
	return &PostExportPathParams{}
}

type PostExportPathParams struct {
}

func (p *PostExportPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostExportRequest) PathParams() *PostExportPathParams {
	return r.pathParams
}

func (r *PostExportRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostExportRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostExportRequest) Method() string {
	return r.method
}

func (s *Client) NewPostExportRequestBody() PostExportRequestBody {
	return PostExportRequestBody{}
}

type PostExportRequestBody struct {
	CreateCompanyModel
}

func (r *PostExportRequest) RequestBody() *PostExportRequestBody {
	return &r.requestBody
}

func (r *PostExportRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostExportRequest) SetRequestBody(body PostExportRequestBody) {
	r.requestBody = body
}

func (r *PostExportRequest) NewResponseBody() *PostExportResponseBody {
	return &PostExportResponseBody{}
}

type PostExportResponseBody PostExportModel

func (r *PostExportRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/export", r.PathParams())
	return &u
}

func (r *PostExportRequest) Do() (PostExportResponseBody, error) {
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
