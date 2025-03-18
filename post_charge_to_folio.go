package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostChargeToFolioRequest() PostChargeToFolioRequest {
	return PostChargeToFolioRequest{
		client:      c,
		queryParams: c.NewPostChargeToFolioQueryParams(),
		pathParams:  c.NewPostChargeToFolioPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostChargeToFolioRequestBody(),
	}
}

type PostChargeToFolioRequest struct {
	client      *Client
	queryParams *PostChargeToFolioQueryParams
	pathParams  *PostChargeToFolioPathParams
	method      string
	headers     http.Header
	requestBody PostChargeToFolioRequestBody
}

func (c *Client) NewPostChargeToFolioQueryParams() *PostChargeToFolioQueryParams {
	return &PostChargeToFolioQueryParams{}
}

type PostChargeToFolioQueryParams struct{}

func (p PostChargeToFolioQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostChargeToFolioRequest) QueryParams() *PostChargeToFolioQueryParams {
	return r.queryParams
}

func (c *Client) NewPostChargeToFolioPathParams() *PostChargeToFolioPathParams {
	return &PostChargeToFolioPathParams{}
}

type PostChargeToFolioPathParams struct {
	ID string `schema:"id"`
}

func (p *PostChargeToFolioPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *PostChargeToFolioRequest) PathParams() *PostChargeToFolioPathParams {
	return r.pathParams
}

func (r *PostChargeToFolioRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostChargeToFolioRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostChargeToFolioRequest) Method() string {
	return r.method
}

func (s *Client) NewPostChargeToFolioRequestBody() PostChargeToFolioRequestBody {
	return PostChargeToFolioRequestBody{}
}

type PostChargeToFolioRequestBody struct {
	ServiceType  string             `json:"serviceType"`
	VatType      string             `json:"vatType"`
	SubAccountID string             `json:"subAccountId,omitempty"`
	Name         string             `json:"name"`
	Amount       MonetaryValueModel `json:"amount"`
	Receipt      string             `json:"receipt,omitempty"`
	Quantity     int32              `json:"quantity,omitempty"`
}

func (r PostChargeToFolioRequest) NewRequestBody() PostChargeToFolioRequestBody {
	return PostChargeToFolioRequestBody{}
}

func (r *PostChargeToFolioRequest) RequestBody() *PostChargeToFolioRequestBody {
	return &r.requestBody
}

func (r *PostChargeToFolioRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *PostChargeToFolioRequest) SetRequestBody(body PostChargeToFolioRequestBody) {
	r.requestBody = body
}

func (r *PostChargeToFolioRequest) NewResponseBody() *PostChargeToFolioResponseBody {
	return &PostChargeToFolioResponseBody{}
}

type PostChargeToFolioResponseBody struct {
	ID string `json:"id"`
}

func (r *PostChargeToFolioRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/folios/{{.id}}/charges", r.PathParams())
	return &u
}

func (r *PostChargeToFolioRequest) Do() (PostChargeToFolioResponseBody, error) {
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
