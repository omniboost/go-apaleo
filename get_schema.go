package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetSchemaRequest() GetSchemaRequest {
	return GetSchemaRequest{
		client:      c,
		queryParams: c.NewGetSchemaQueryParams(),
		pathParams:  c.NewGetSchemaPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetSchemaRequestBody(),
	}
}

type GetSchemaRequest struct {
	client      *Client
	queryParams *GetSchemaQueryParams
	pathParams  *GetSchemaPathParams
	method      string
	headers     http.Header
	requestBody GetSchemaRequestBody
}

func (c *Client) NewGetSchemaQueryParams() *GetSchemaQueryParams {
	return &GetSchemaQueryParams{}
}

type GetSchemaQueryParams struct {
	PropertyID       string           `schema:"propertyId,omitempty"`
	Depth            int32            `schema:"depth,omitempty"`
	IncludeArchived  bool             `schema:"includeArchived,omitempty"`
	AccountingSchema AccountingSchema `schema:"accountingSchema,omitempty"`
	LanguageCode     string           `schema:"languageCode,omitempty"`
}

func (p GetSchemaQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetSchemaRequest) QueryParams() *GetSchemaQueryParams {
	return r.queryParams
}

func (c *Client) NewGetSchemaPathParams() *GetSchemaPathParams {
	return &GetSchemaPathParams{}
}

type GetSchemaPathParams struct {
}

func (p *GetSchemaPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetSchemaRequest) PathParams() *GetSchemaPathParams {
	return r.pathParams
}

func (r *GetSchemaRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetSchemaRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetSchemaRequest) Method() string {
	return r.method
}

func (s *Client) NewGetSchemaRequestBody() GetSchemaRequestBody {
	return GetSchemaRequestBody{}
}

type GetSchemaRequestBody struct {
}

func (r *GetSchemaRequest) RequestBody() *GetSchemaRequestBody {
	return nil
}

func (r *GetSchemaRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetSchemaRequest) SetRequestBody(body GetSchemaRequestBody) {
	r.requestBody = body
}

func (r *GetSchemaRequest) NewResponseBody() *GetSchemaResponseBody {
	return &GetSchemaResponseBody{}
}

type GetSchemaResponseBody GetSchemaModel

func (r *GetSchemaRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/schema", r.PathParams())
	return &u
}

func (r *GetSchemaRequest) Do() (GetSchemaResponseBody, error) {
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
