package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetAccountsExportRequest() GetAccountsExportRequest {
	return GetAccountsExportRequest{
		client:      c,
		queryParams: c.NewGetAccountsExportQueryParams(),
		pathParams:  c.NewGetAccountsExportPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewGetAccountsExportRequestBody(),
	}
}

type GetAccountsExportRequest struct {
	client      *Client
	queryParams *GetAccountsExportQueryParams
	pathParams  *GetAccountsExportPathParams
	method      string
	headers     http.Header
	requestBody GetAccountsExportRequestBody
}

func (c *Client) NewGetAccountsExportQueryParams() *GetAccountsExportQueryParams {
	return &GetAccountsExportQueryParams{}
}

type GetAccountsExportQueryParams struct {
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
	// Unique key for safely retrying requests without accidentally performing
	// the same operation twice. We'll always send back the same response for
	// requests made with the same key, and keys can't be reused with different
	// request parameters. Keys expire after 24 hours.
	IdempotencyKey string `schema:"Idempotency-Key,omitempty"`
}

func (p GetAccountsExportQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetAccountsExportRequest) QueryParams() *GetAccountsExportQueryParams {
	return r.queryParams
}

func (c *Client) NewGetAccountsExportPathParams() *GetAccountsExportPathParams {
	return &GetAccountsExportPathParams{}
}

type GetAccountsExportPathParams struct {
}

func (p *GetAccountsExportPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAccountsExportRequest) PathParams() *GetAccountsExportPathParams {
	return r.pathParams
}

func (r *GetAccountsExportRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetAccountsExportRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAccountsExportRequest) Method() string {
	return r.method
}

func (s *Client) NewGetAccountsExportRequestBody() GetAccountsExportRequestBody {
	return GetAccountsExportRequestBody{}
}

type GetAccountsExportRequestBody struct {
}

func (r *GetAccountsExportRequest) RequestBody() *GetAccountsExportRequestBody {
	return nil
}

func (r *GetAccountsExportRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetAccountsExportRequest) SetRequestBody(body GetAccountsExportRequestBody) {
	r.requestBody = body
}

func (r *GetAccountsExportRequest) NewResponseBody() *GetAccountsExportResponseBody {
	return &GetAccountsExportResponseBody{}
}

type GetAccountsExportResponseBody struct {
	Transactions []ExportTransactionItemModel `json:"transactions"`
}

func (r *GetAccountsExportRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/export", r.PathParams())
	return &u
}

func (r *GetAccountsExportRequest) Do() (GetAccountsExportResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
