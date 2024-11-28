package apaleo

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetInvoicesRequest() GetInvoicesRequest {
	return GetInvoicesRequest{
		client:      c,
		queryParams: c.NewGetInvoicesQueryParams(),
		pathParams:  c.NewGetInvoicesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetInvoicesRequestBody(),
	}
}

type GetInvoicesRequest struct {
	client      *Client
	queryParams *GetInvoicesQueryParams
	pathParams  *GetInvoicesPathParams
	method      string
	headers     http.Header
	requestBody GetInvoicesRequestBody
}

func (c *Client) NewGetInvoicesQueryParams() *GetInvoicesQueryParams {
	return &GetInvoicesQueryParams{}
}

type GetInvoicesQueryParams struct {
}

func (p GetInvoicesQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetInvoicesRequest) QueryParams() *GetInvoicesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetInvoicesPathParams() *GetInvoicesPathParams {
	return &GetInvoicesPathParams{}
}

type GetInvoicesPathParams struct {
}

func (p *GetInvoicesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetInvoicesRequest) PathParams() *GetInvoicesPathParams {
	return r.pathParams
}

func (r *GetInvoicesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetInvoicesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoicesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetInvoicesRequestBody() GetInvoicesRequestBody {
	return GetInvoicesRequestBody{}
}

type GetInvoicesRequestBody struct {
}

func (r *GetInvoicesRequest) RequestBody() *GetInvoicesRequestBody {
	return nil
}

func (r *GetInvoicesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetInvoicesRequest) SetRequestBody(body GetInvoicesRequestBody) {
	r.requestBody = body
}

func (r *GetInvoicesRequest) NewResponseBody() *GetInvoicesResponseBody {
	return &GetInvoicesResponseBody{}
}

type GetInvoicesResponseBody struct {
}

func (r *GetInvoicesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/invoices", r.PathParams())
	return &u
}

func (r *GetInvoicesRequest) Do() (GetInvoicesResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
