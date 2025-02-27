package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetCompanyByIDRequest() GetCompanyByIDRequest {
	return GetCompanyByIDRequest{
		client:      c,
		queryParams: c.NewGetCompanyByIDQueryParams(),
		pathParams:  c.NewGetCompanyByIDPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetCompanyByIDRequestBody(),
	}
}

type GetCompanyByIDRequest struct {
	client      *Client
	queryParams *GetCompanyByIDQueryParams
	pathParams  *GetCompanyByIDPathParams
	method      string
	headers     http.Header
	requestBody GetCompanyByIDRequestBody
}

func (c *Client) NewGetCompanyByIDQueryParams() *GetCompanyByIDQueryParams {
	return &GetCompanyByIDQueryParams{}
}

type GetCompanyByIDQueryParams struct{}

func (p GetCompanyByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetCompanyByIDRequest) QueryParams() *GetCompanyByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCompanyByIDPathParams() *GetCompanyByIDPathParams {
	return &GetCompanyByIDPathParams{}
}

type GetCompanyByIDPathParams struct {
	ID string `schema:"id"`
}

func (p *GetCompanyByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetCompanyByIDRequest) PathParams() *GetCompanyByIDPathParams {
	return r.pathParams
}

func (r *GetCompanyByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCompanyByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCompanyByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCompanyByIDRequestBody() GetCompanyByIDRequestBody {
	return GetCompanyByIDRequestBody{}
}

type GetCompanyByIDRequestBody struct {
}

func (r *GetCompanyByIDRequest) RequestBody() *GetCompanyByIDRequestBody {
	return nil
}

func (r *GetCompanyByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetCompanyByIDRequest) SetRequestBody(body GetCompanyByIDRequestBody) {
	r.requestBody = body
}

func (r *GetCompanyByIDRequest) NewResponseBody() *GetCompanyByIDResponseBody {
	return &GetCompanyByIDResponseBody{}
}

type GetCompanyByIDResponseBody CompanyModel

func (r *GetCompanyByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rateplan/v1/companies/{{.id}}", r.PathParams())
	return &u
}

func (r *GetCompanyByIDRequest) Do() (GetCompanyByIDResponseBody, error) {
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
