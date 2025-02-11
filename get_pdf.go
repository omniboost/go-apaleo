package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetPdfRequest() GetPdfRequest {
	return GetPdfRequest{
		client:      c,
		queryParams: c.NewGetPdfQueryParams(),
		pathParams:  c.NewGetPdfPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetPdfRequestBody(),
	}
}

type GetPdfRequest struct {
	client      *Client
	queryParams *GetPdfQueryParams
	pathParams  *GetPdfPathParams
	method      string
	headers     http.Header
	requestBody GetPdfRequestBody
}

func (c *Client) NewGetPdfQueryParams() *GetPdfQueryParams {
	return &GetPdfQueryParams{}
}

type GetPdfQueryParams struct{}

func (p GetPdfQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetPdfRequest) QueryParams() *GetPdfQueryParams {
	return r.queryParams
}

func (c *Client) NewGetPdfPathParams() *GetPdfPathParams {
	return &GetPdfPathParams{}
}

type GetPdfPathParams struct {
	ID string
}

func (p *GetPdfPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetPdfRequest) PathParams() *GetPdfPathParams {
	return r.pathParams
}

func (r *GetPdfRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetPdfRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPdfRequest) Method() string {
	return r.method
}

func (s *Client) NewGetPdfRequestBody() GetPdfRequestBody {
	return GetPdfRequestBody{}
}

type GetPdfRequestBody struct {
}

func (r *GetPdfRequest) RequestBody() *GetPdfRequestBody {
	return nil
}

func (r *GetPdfRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetPdfRequest) SetRequestBody(body GetPdfRequestBody) {
	r.requestBody = body
}

func (r *GetPdfRequest) NewResponseBody() *GetPdfResponseBody {
	return &GetPdfResponseBody{}
}

type GetPdfResponseBody struct {
	Invoices []InvoiceItemModel `json:"invoices"`
}

func (r *GetPdfRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/invoices/{{.id}}/pdf", r.PathParams())
	return &u
}

func (r *GetPdfRequest) Do() (GetPdfResponseBody, error) {
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
