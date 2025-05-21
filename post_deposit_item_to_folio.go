package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostDepositItemToFolioRequest() PostDepositItemToFolioRequest {
	return PostDepositItemToFolioRequest{
		client:      c,
		queryParams: c.NewPostDepositItemToFolioQueryParams(),
		pathParams:  c.NewPostDepositItemToFolioPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostDepositItemToFolioRequestBody(),
	}
}

type PostDepositItemToFolioRequest struct {
	client      *Client
	queryParams *PostDepositItemToFolioQueryParams
	pathParams  *PostDepositItemToFolioPathParams
	method      string
	headers     http.Header
	requestBody PostDepositItemToFolioRequestBody
}

func (c *Client) NewPostDepositItemToFolioQueryParams() *PostDepositItemToFolioQueryParams {
	return &PostDepositItemToFolioQueryParams{}
}

type PostDepositItemToFolioQueryParams struct{}

func (p PostDepositItemToFolioQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostDepositItemToFolioRequest) QueryParams() *PostDepositItemToFolioQueryParams {
	return r.queryParams
}

func (c *Client) NewPostDepositItemToFolioPathParams() *PostDepositItemToFolioPathParams {
	return &PostDepositItemToFolioPathParams{}
}

type PostDepositItemToFolioPathParams struct {
	ID string `schema:"id"`
}

func (p *PostDepositItemToFolioPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *PostDepositItemToFolioRequest) PathParams() *PostDepositItemToFolioPathParams {
	return r.pathParams
}

func (r *PostDepositItemToFolioRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostDepositItemToFolioRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostDepositItemToFolioRequest) Method() string {
	return r.method
}

func (s *Client) NewPostDepositItemToFolioRequestBody() PostDepositItemToFolioRequestBody {
	return PostDepositItemToFolioRequestBody{}
}

type PostDepositItemToFolioRequestBody struct {
	Amount               MonetaryValueModel `json:"amount"`
	Name                 map[string]string  `json:"name"`
	VatType              string             `json:"vatType,omitempty"`
	Quantity             int32              `json:"quantity,omitempty"`
	ExpectedDeliveryDate *Date              `json:"expectedDeliveryDate,omitempty"`
	ServiceDate          *Date              `json:"serviceDate,omitempty"`
}

func (r PostDepositItemToFolioRequest) NewRequestBody() PostDepositItemToFolioRequestBody {
	return PostDepositItemToFolioRequestBody{}
}

func (r *PostDepositItemToFolioRequest) RequestBody() *PostDepositItemToFolioRequestBody {
	return &r.requestBody
}

func (r *PostDepositItemToFolioRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *PostDepositItemToFolioRequest) SetRequestBody(body PostDepositItemToFolioRequestBody) {
	r.requestBody = body
}

func (r *PostDepositItemToFolioRequest) NewResponseBody() *PostDepositItemToFolioResponseBody {
	return &PostDepositItemToFolioResponseBody{}
}

type PostDepositItemToFolioResponseBody struct {
	ID string `json:"id"`
}

func (r *PostDepositItemToFolioRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v0-nsfw/folio-actions/{{.id}}/deposit-items", r.PathParams())
	return &u
}

func (r *PostDepositItemToFolioRequest) Do() (PostDepositItemToFolioResponseBody, error) {
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
