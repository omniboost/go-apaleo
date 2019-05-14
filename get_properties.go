package apaleo

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetPropertiesRequest() GetPropertiesRequest {
	return GetPropertiesRequest{
		client:      c,
		queryParams: c.NewGetPropertiesQueryParams(),
		pathParams:  c.NewGetPropertiesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetPropertiesRequestBody(),
	}
}

type GetPropertiesRequest struct {
	client      *Client
	queryParams *GetPropertiesQueryParams
	pathParams  *GetPropertiesPathParams
	method      string
	headers     http.Header
	requestBody GetPropertiesRequestBody
}

func (c *Client) NewGetPropertiesQueryParams() *GetPropertiesQueryParams {
	return &GetPropertiesQueryParams{}
}

type GetPropertiesQueryParams struct {
}

func (p GetPropertiesQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetPropertiesRequest) QueryParams() *GetPropertiesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetPropertiesPathParams() *GetPropertiesPathParams {
	return &GetPropertiesPathParams{}
}

type GetPropertiesPathParams struct {
}

func (p *GetPropertiesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetPropertiesRequest) PathParams() *GetPropertiesPathParams {
	return r.pathParams
}

func (r *GetPropertiesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPropertiesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetPropertiesRequestBody() GetPropertiesRequestBody {
	return GetPropertiesRequestBody{}
}

type GetPropertiesRequestBody struct {
}

func (r *GetPropertiesRequest) RequestBody() *GetPropertiesRequestBody {
	return &r.requestBody
}

func (r *GetPropertiesRequest) SetRequestBody(body GetPropertiesRequestBody) {
	r.requestBody = body
}

func (r *GetPropertiesRequest) NewResponseBody() *GetPropertiesResponseBody {
	return &GetPropertiesResponseBody{}
}

type GetPropertiesResponseBody struct {
	Count      int                 `json:"count"`
	Properties []PropertyItemModel `json:"properties"`
}

func (r *GetPropertiesRequest) URL() url.URL {
	return r.client.GetEndpointURL("inventory/v1/properties", r.PathParams())
}

func (r *GetPropertiesRequest) Do() (GetPropertiesResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
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
