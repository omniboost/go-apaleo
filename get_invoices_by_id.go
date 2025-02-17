package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetInvoicesByIDRequest() GetInvoicesByIDRequest {
	return GetInvoicesByIDRequest{
		client:      c,
		queryParams: c.NewGetInvoicesByIDQueryParams(),
		pathParams:  c.NewGetInvoicesByIDPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetInvoicesByIDRequestBody(),
	}
}

type GetInvoicesByIDRequest struct {
	client      *Client
	queryParams *GetInvoicesByIDQueryParams
	pathParams  *GetInvoicesByIDPathParams
	method      string
	headers     http.Header
	requestBody GetInvoicesByIDRequestBody
}

func (c *Client) NewGetInvoicesByIDQueryParams() *GetInvoicesByIDQueryParams {
	return &GetInvoicesByIDQueryParams{}
}

type GetInvoicesByIDQueryParams struct {
	Expand []string `schema:"expand,omitempty"`
}

func (p GetInvoicesByIDQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInvoicesByIDRequest) QueryParams() *GetInvoicesByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetInvoicesByIDPathParams() *GetInvoicesByIDPathParams {
	return &GetInvoicesByIDPathParams{}
}

type GetInvoicesByIDPathParams struct {
	ID string `schema:"id"`
}

func (p *GetInvoicesByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetInvoicesByIDRequest) PathParams() *GetInvoicesByIDPathParams {
	return r.pathParams
}

func (r *GetInvoicesByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetInvoicesByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoicesByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetInvoicesByIDRequestBody() GetInvoicesByIDRequestBody {
	return GetInvoicesByIDRequestBody{}
}

type GetInvoicesByIDRequestBody struct {
}

func (r *GetInvoicesByIDRequest) RequestBody() *GetInvoicesByIDRequestBody {
	return nil
}

func (r *GetInvoicesByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetInvoicesByIDRequest) SetRequestBody(body GetInvoicesByIDRequestBody) {
	r.requestBody = body
}

func (r *GetInvoicesByIDRequest) NewResponseBody() *GetInvoicesByIDResponseBody {
	return &GetInvoicesByIDResponseBody{}
}

type GetInvoicesByIDResponseBody InvoicesByIDModel

func (r *GetInvoicesByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/invoices/{{.id}}", r.PathParams())
	return &u
}

func (r *GetInvoicesByIDRequest) Do() (GetInvoicesByIDResponseBody, error) {
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
