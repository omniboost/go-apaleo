package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewDeleteCompanyByIDRequest() DeleteCompanyByIDRequest {
	return DeleteCompanyByIDRequest{
		client:      c,
		queryParams: c.NewDeleteCompanyByIDQueryParams(),
		pathParams:  c.NewDeleteCompanyByIDPathParams(),
		method:      http.MethodDelete,
		headers:     http.Header{},
		requestBody: c.NewDeleteCompanyByIDRequestBody(),
	}
}

type DeleteCompanyByIDRequest struct {
	client      *Client
	queryParams *DeleteCompanyByIDQueryParams
	pathParams  *DeleteCompanyByIDPathParams
	method      string
	headers     http.Header
	requestBody DeleteCompanyByIDRequestBody
}

func (c *Client) NewDeleteCompanyByIDQueryParams() *DeleteCompanyByIDQueryParams {
	return &DeleteCompanyByIDQueryParams{}
}

type DeleteCompanyByIDQueryParams struct{}

func (p DeleteCompanyByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DeleteCompanyByIDRequest) QueryParams() *DeleteCompanyByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewDeleteCompanyByIDPathParams() *DeleteCompanyByIDPathParams {
	return &DeleteCompanyByIDPathParams{}
}

type DeleteCompanyByIDPathParams struct {
	ID string `schema:"id"`
}

func (p *DeleteCompanyByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *DeleteCompanyByIDRequest) PathParams() *DeleteCompanyByIDPathParams {
	return r.pathParams
}

func (r *DeleteCompanyByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DeleteCompanyByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *DeleteCompanyByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewDeleteCompanyByIDRequestBody() DeleteCompanyByIDRequestBody {
	return DeleteCompanyByIDRequestBody{}
}

type DeleteCompanyByIDRequestBody struct {
}

func (r *DeleteCompanyByIDRequest) RequestBody() *DeleteCompanyByIDRequestBody {
	return nil
}

func (r *DeleteCompanyByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *DeleteCompanyByIDRequest) SetRequestBody(body DeleteCompanyByIDRequestBody) {
	r.requestBody = body
}

func (r *DeleteCompanyByIDRequest) NewResponseBody() *DeleteCompanyByIDResponseBody {
	return &DeleteCompanyByIDResponseBody{}
}

type DeleteCompanyByIDResponseBody struct{}

func (r *DeleteCompanyByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rateplan/v1/companies/{{.id}}", r.PathParams())
	return &u
}

func (r *DeleteCompanyByIDRequest) Do(ctx context.Context) (DeleteCompanyByIDResponseBody, error) {
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
