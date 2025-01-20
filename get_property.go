package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetPropertyRequest() GetPropertyRequest {
	return GetPropertyRequest{
		client:      c,
		queryParams: c.NewGetPropertyQueryParams(),
		pathParams:  c.NewGetPropertyPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetPropertyRequestBody(),
	}
}

type GetPropertyRequest struct {
	client      *Client
	queryParams *GetPropertyQueryParams
	pathParams  *GetPropertyPathParams
	method      string
	headers     http.Header
	requestBody GetPropertyRequestBody
}

func (c *Client) NewGetPropertyQueryParams() *GetPropertyQueryParams {
	return &GetPropertyQueryParams{}
}

type GetPropertyQueryParams struct {
	// 'all' or comma separated list of two-letter language codes (ISO Alpha-2)
	Languages []string `json:"languages,omitempty"`
}

func (p GetPropertyQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetPropertyRequest) QueryParams() *GetPropertyQueryParams {
	return r.queryParams
}

func (c *Client) NewGetPropertyPathParams() *GetPropertyPathParams {
	return &GetPropertyPathParams{}
}

type GetPropertyPathParams struct {
	ID string
}

func (p *GetPropertyPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetPropertyRequest) PathParams() *GetPropertyPathParams {
	return r.pathParams
}

func (r *GetPropertyRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetPropertyRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPropertyRequest) Method() string {
	return r.method
}

func (s *Client) NewGetPropertyRequestBody() GetPropertyRequestBody {
	return GetPropertyRequestBody{}
}

type GetPropertyRequestBody struct{}

func (r *GetPropertyRequest) RequestBody() *GetPropertyRequestBody {
	return nil
}

func (r *GetPropertyRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetPropertyRequest) SetRequestBody(body GetPropertyRequestBody) {
	r.requestBody = body
}

func (r *GetPropertyRequest) NewResponseBody() *GetPropertyResponseBody {
	return &GetPropertyResponseBody{}
}

type GetPropertyResponseBody struct {
	PropertyModel
}

func (r *GetPropertyRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("inventory/v1/properties/{{.id}}", r.PathParams())
	return &u
}

func (r *GetPropertyRequest) Do() (GetPropertyResponseBody, error) {
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
