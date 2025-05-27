package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetAccountsSchemaRequest() GetAccountsSchemaRequest {
	return GetAccountsSchemaRequest{
		client:      c,
		queryParams: c.NewGetAccountsSchemaQueryParams(),
		pathParams:  c.NewGetAccountsSchemaPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetAccountsSchemaRequestBody(),
	}
}

type GetAccountsSchemaRequest struct {
	client      *Client
	queryParams *GetAccountsSchemaQueryParams
	pathParams  *GetAccountsSchemaPathParams
	method      string
	headers     http.Header
	requestBody GetAccountsSchemaRequestBody
}

func (c *Client) NewGetAccountsSchemaQueryParams() *GetAccountsSchemaQueryParams {
	return &GetAccountsSchemaQueryParams{}
}

type GetAccountsSchemaQueryParams struct {
	PropertyID       string           `schema:"propertyId,omitempty"`
	Depth            int32            `schema:"depth,omitempty"`
	IncludeArchived  bool             `schema:"includeArchived,omitempty"`
	AccountingSchema AccountingSchema `schema:"accountingSchema,omitempty"`
	LanguageCode     string           `schema:"languageCode,omitempty"`
}

func (p GetAccountsSchemaQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetAccountsSchemaRequest) QueryParams() *GetAccountsSchemaQueryParams {
	return r.queryParams
}

func (c *Client) NewGetAccountsSchemaPathParams() *GetAccountsSchemaPathParams {
	return &GetAccountsSchemaPathParams{}
}

type GetAccountsSchemaPathParams struct {
}

func (p *GetAccountsSchemaPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAccountsSchemaRequest) PathParams() *GetAccountsSchemaPathParams {
	return r.pathParams
}

func (r *GetAccountsSchemaRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetAccountsSchemaRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAccountsSchemaRequest) Method() string {
	return r.method
}

func (s *Client) NewGetAccountsSchemaRequestBody() GetAccountsSchemaRequestBody {
	return GetAccountsSchemaRequestBody{}
}

type GetAccountsSchemaRequestBody struct {
}

func (r *GetAccountsSchemaRequest) RequestBody() *GetAccountsSchemaRequestBody {
	return nil
}

func (r *GetAccountsSchemaRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetAccountsSchemaRequest) SetRequestBody(body GetAccountsSchemaRequestBody) {
	r.requestBody = body
}

func (r *GetAccountsSchemaRequest) NewResponseBody() *GetAccountsSchemaResponseBody {
	return &GetAccountsSchemaResponseBody{}
}

type GetAccountsSchemaResponseBody GetAccountsSchemaModel

func (r *GetAccountsSchemaRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/accounts/schema", r.PathParams())
	return &u
}

func (r *GetAccountsSchemaRequest) Do(ctx context.Context) (GetAccountsSchemaResponseBody, error) {
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
