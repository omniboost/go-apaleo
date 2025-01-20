package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostCompanyRequest() PostCompanyRequest {
	return PostCompanyRequest{
		client:      c,
		queryParams: c.NewPostCompanyQueryParams(),
		pathParams:  c.NewPostCompanyPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostCompanyRequestBody(),
	}
}

type PostCompanyRequest struct {
	client      *Client
	queryParams *PostCompanyQueryParams
	pathParams  *PostCompanyPathParams
	method      string
	headers     http.Header
	requestBody PostCompanyRequestBody
}

func (c *Client) NewPostCompanyQueryParams() *PostCompanyQueryParams {
	return &PostCompanyQueryParams{}
}

type PostCompanyQueryParams struct{}

func (p PostCompanyQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostCompanyRequest) QueryParams() *PostCompanyQueryParams {
	return r.queryParams
}

func (c *Client) NewPostCompanyPathParams() *PostCompanyPathParams {
	return &PostCompanyPathParams{}
}

type PostCompanyPathParams struct {
}

func (p *PostCompanyPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostCompanyRequest) PathParams() *PostCompanyPathParams {
	return r.pathParams
}

func (r *PostCompanyRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostCompanyRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostCompanyRequest) Method() string {
	return r.method
}

func (s *Client) NewPostCompanyRequestBody() PostCompanyRequestBody {
	return PostCompanyRequestBody{}
}

type PostCompanyRequestBody struct {
	CreateCompanyModel
}

func (r *PostCompanyRequest) RequestBody() *PostCompanyRequestBody {
	return &r.requestBody
}

func (r *PostCompanyRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostCompanyRequest) SetRequestBody(body PostCompanyRequestBody) {
	r.requestBody = body
}

func (r *PostCompanyRequest) NewResponseBody() *PostCompanyResponseBody {
	return &PostCompanyResponseBody{}
}

type PostCompanyResponseBody struct {
	ID string `json:"id"`
}

func (r *PostCompanyRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rateplan/v1/companies", r.PathParams())
	return &u
}

func (r *PostCompanyRequest) Do() (PostCompanyResponseBody, error) {
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
