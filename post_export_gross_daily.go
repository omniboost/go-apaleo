package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostExportGrossDailyRequest() PostExportGrossDailyRequest {
	return PostExportGrossDailyRequest{
		client:      c,
		queryParams: c.NewPostExportGrossDailyQueryParams(),
		pathParams:  c.NewPostExportGrossDailyPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostExportGrossDailyRequestBody(),
	}
}

type PostExportGrossDailyRequest struct {
	client      *Client
	queryParams *PostExportGrossDailyQueryParams
	pathParams  *PostExportGrossDailyPathParams
	method      string
	headers     http.Header
	requestBody PostExportGrossDailyRequestBody
}

func (c *Client) NewPostExportGrossDailyQueryParams() *PostExportGrossDailyQueryParams {
	return &PostExportGrossDailyQueryParams{}
}

type PostExportGrossDailyQueryParams struct {
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
	// Allows to override the default accounting schema. Only specify this, when
	// you know what you are doing.
	AccountingSchema AccountingSchema `schema:"accountingSchema,omitempty"`
	// Unique key for safely retrying requests without accidentally performing
	// the same operation twice. We'll always send back the same response for
	// requests made with the same key, and keys can't be reused with different
	// request parameters. Keys expire after 24 hours.
	IdempotencyKey string `schema:"Idempotency-Key,omitempty"`
}

func (p PostExportGrossDailyQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostExportGrossDailyRequest) QueryParams() *PostExportGrossDailyQueryParams {
	return r.queryParams
}

func (c *Client) NewPostExportGrossDailyPathParams() *PostExportGrossDailyPathParams {
	return &PostExportGrossDailyPathParams{}
}

type PostExportGrossDailyPathParams struct {
}

func (p *PostExportGrossDailyPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostExportGrossDailyRequest) PathParams() *PostExportGrossDailyPathParams {
	return r.pathParams
}

func (r *PostExportGrossDailyRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostExportGrossDailyRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostExportGrossDailyRequest) Method() string {
	return r.method
}

func (s *Client) NewPostExportGrossDailyRequestBody() PostExportGrossDailyRequestBody {
	return PostExportGrossDailyRequestBody{}
}

type PostExportGrossDailyRequestBody struct {
	CreateCompanyModel
}

func (r *PostExportGrossDailyRequest) RequestBody() *PostExportGrossDailyRequestBody {
	return &r.requestBody
}

func (r *PostExportGrossDailyRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostExportGrossDailyRequest) SetRequestBody(body PostExportGrossDailyRequestBody) {
	r.requestBody = body
}

func (r *PostExportGrossDailyRequest) NewResponseBody() *PostExportGrossDailyResponseBody {
	return &PostExportGrossDailyResponseBody{}
}

type PostExportGrossDailyResponseBody PostExportGrossDailyModel

func (r *PostExportGrossDailyRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/export-gross-daily", r.PathParams())
	return &u
}

func (r *PostExportGrossDailyRequest) Do() (PostExportGrossDailyResponseBody, error) {
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
