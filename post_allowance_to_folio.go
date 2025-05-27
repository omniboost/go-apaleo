package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostAllowanceToFolioRequest() PostAllowanceToFolioRequest {
	return PostAllowanceToFolioRequest{
		client:      c,
		queryParams: c.NewPostAllowanceToFolioQueryParams(),
		pathParams:  c.NewPostAllowanceToFolioPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostAllowanceToFolioRequestBody(),
	}
}

type PostAllowanceToFolioRequest struct {
	client      *Client
	queryParams *PostAllowanceToFolioQueryParams
	pathParams  *PostAllowanceToFolioPathParams
	method      string
	headers     http.Header
	requestBody PostAllowanceToFolioRequestBody
}

func (c *Client) NewPostAllowanceToFolioQueryParams() *PostAllowanceToFolioQueryParams {
	return &PostAllowanceToFolioQueryParams{}
}

type PostAllowanceToFolioQueryParams struct{}

func (p PostAllowanceToFolioQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostAllowanceToFolioRequest) QueryParams() *PostAllowanceToFolioQueryParams {
	return r.queryParams
}

func (c *Client) NewPostAllowanceToFolioPathParams() *PostAllowanceToFolioPathParams {
	return &PostAllowanceToFolioPathParams{}
}

type PostAllowanceToFolioPathParams struct {
	ID string `schema:"id"`
}

func (p *PostAllowanceToFolioPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *PostAllowanceToFolioRequest) PathParams() *PostAllowanceToFolioPathParams {
	return r.pathParams
}

func (r *PostAllowanceToFolioRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostAllowanceToFolioRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostAllowanceToFolioRequest) Method() string {
	return r.method
}

func (s *Client) NewPostAllowanceToFolioRequestBody() PostAllowanceToFolioRequestBody {
	return PostAllowanceToFolioRequestBody{}
}

type PostAllowanceToFolioRequestBody struct {
	ServiceType  string             `json:"serviceType"`
	VatType      string             `json:"vatType"`
	SubAccountID string             `json:"subAccountId,omitempty"`
	Reason       string             `json:"reason,omitempty"`
	Amount       MonetaryValueModel `json:"amount"`
	BusinessDate *Date              `json:"businessDate,omitempty"`
}

func (r PostAllowanceToFolioRequest) NewRequestBody() PostAllowanceToFolioRequestBody {
	return PostAllowanceToFolioRequestBody{}
}

func (r *PostAllowanceToFolioRequest) RequestBody() *PostAllowanceToFolioRequestBody {
	return &r.requestBody
}

func (r *PostAllowanceToFolioRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *PostAllowanceToFolioRequest) SetRequestBody(body PostAllowanceToFolioRequestBody) {
	r.requestBody = body
}

func (r *PostAllowanceToFolioRequest) NewResponseBody() *PostAllowanceToFolioResponseBody {
	return &PostAllowanceToFolioResponseBody{}
}

type PostAllowanceToFolioResponseBody struct {
	ID string `json:"id"`
}

func (r *PostAllowanceToFolioRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/folio-actions/{{.id}}/allowances", r.PathParams())
	return &u
}

func (r *PostAllowanceToFolioRequest) Do(ctx context.Context) (PostAllowanceToFolioResponseBody, error) {
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
