package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostRefundToFolioRequest() PostRefundToFolioRequest {
	return PostRefundToFolioRequest{
		client:      c,
		queryParams: c.NewPostRefundToFolioQueryParams(),
		pathParams:  c.NewPostRefundToFolioPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostRefundToFolioRequestBody(),
	}
}

type PostRefundToFolioRequest struct {
	client      *Client
	queryParams *PostRefundToFolioQueryParams
	pathParams  *PostRefundToFolioPathParams
	method      string
	headers     http.Header
	requestBody PostRefundToFolioRequestBody
}

func (c *Client) NewPostRefundToFolioQueryParams() *PostRefundToFolioQueryParams {
	return &PostRefundToFolioQueryParams{}
}

type PostRefundToFolioQueryParams struct{}

func (p PostRefundToFolioQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostRefundToFolioRequest) QueryParams() *PostRefundToFolioQueryParams {
	return r.queryParams
}

func (c *Client) NewPostRefundToFolioPathParams() *PostRefundToFolioPathParams {
	return &PostRefundToFolioPathParams{}
}

type PostRefundToFolioPathParams struct {
	ID string `schema:"id"`
}

func (p *PostRefundToFolioPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *PostRefundToFolioRequest) PathParams() *PostRefundToFolioPathParams {
	return r.pathParams
}

func (r *PostRefundToFolioRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostRefundToFolioRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostRefundToFolioRequest) Method() string {
	return r.method
}

func (s *Client) NewPostRefundToFolioRequestBody() PostRefundToFolioRequestBody {
	return PostRefundToFolioRequestBody{}
}

type PostRefundToFolioRequestBody struct {
	Method       string             `json:"method"`
	Amount       MonetaryValueModel `json:"amount"`
	Receipt      string             `json:"receipt,omitempty"`
	BusinessDate *Date              `json:"businessDate,omitempty"`
	Reason       string             `json:"reason,omitempty"`
}

func (r PostRefundToFolioRequest) NewRequestBody() PostRefundToFolioRequestBody {
	return PostRefundToFolioRequestBody{}
}

func (r *PostRefundToFolioRequest) RequestBody() *PostRefundToFolioRequestBody {
	return &r.requestBody
}

func (r *PostRefundToFolioRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *PostRefundToFolioRequest) SetRequestBody(body PostRefundToFolioRequestBody) {
	r.requestBody = body
}

func (r *PostRefundToFolioRequest) NewResponseBody() *PostRefundToFolioResponseBody {
	return &PostRefundToFolioResponseBody{}
}

type PostRefundToFolioResponseBody struct {
	ID string `json:"id"`
}

func (r *PostRefundToFolioRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/folios/{{.id}}/refunds", r.PathParams())
	return &u
}

func (r *PostRefundToFolioRequest) Do() (PostRefundToFolioResponseBody, error) {
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
