package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostPaymentToFolioRequest() PostPaymentToFolioRequest {
	return PostPaymentToFolioRequest{
		client:      c,
		queryParams: c.NewPostPaymentToFolioQueryParams(),
		pathParams:  c.NewPostPaymentToFolioPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostPaymentToFolioRequestBody(),
	}
}

type PostPaymentToFolioRequest struct {
	client      *Client
	queryParams *PostPaymentToFolioQueryParams
	pathParams  *PostPaymentToFolioPathParams
	method      string
	headers     http.Header
	requestBody PostPaymentToFolioRequestBody
}

func (c *Client) NewPostPaymentToFolioQueryParams() *PostPaymentToFolioQueryParams {
	return &PostPaymentToFolioQueryParams{}
}

type PostPaymentToFolioQueryParams struct{}

func (p PostPaymentToFolioQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostPaymentToFolioRequest) QueryParams() *PostPaymentToFolioQueryParams {
	return r.queryParams
}

func (c *Client) NewPostPaymentToFolioPathParams() *PostPaymentToFolioPathParams {
	return &PostPaymentToFolioPathParams{}
}

type PostPaymentToFolioPathParams struct {
	ID string `schema:"id"`
}

func (p *PostPaymentToFolioPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *PostPaymentToFolioRequest) PathParams() *PostPaymentToFolioPathParams {
	return r.pathParams
}

func (r *PostPaymentToFolioRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostPaymentToFolioRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostPaymentToFolioRequest) Method() string {
	return r.method
}

func (s *Client) NewPostPaymentToFolioRequestBody() PostPaymentToFolioRequestBody {
	return PostPaymentToFolioRequestBody{}
}

type PostPaymentToFolioRequestBody struct {
	Method       string             `json:"method"`
	Receipt      string             `json:"receipt,omitempty"`
	BusinessDate *Date              `json:"businessDate,omitempty"`
	Amount       MonetaryValueModel `json:"amount"`
	PaidCharges  []struct {
		ChargeID string  `json:"chargeId"`
		Amount   float64 `json:"amount"`
	} `json:"paidCharges,omitempty"`
}

func (r PostPaymentToFolioRequest) NewRequestBody() PostPaymentToFolioRequestBody {
	return PostPaymentToFolioRequestBody{}
}

func (r *PostPaymentToFolioRequest) RequestBody() *PostPaymentToFolioRequestBody {
	return &r.requestBody
}

func (r *PostPaymentToFolioRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *PostPaymentToFolioRequest) SetRequestBody(body PostPaymentToFolioRequestBody) {
	r.requestBody = body
}

func (r *PostPaymentToFolioRequest) NewResponseBody() *PostPaymentToFolioResponseBody {
	return &PostPaymentToFolioResponseBody{}
}

type PostPaymentToFolioResponseBody struct {
	ID string `json:"id"`
}

func (r *PostPaymentToFolioRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/folios/{{.id}}/payments", r.PathParams())
	return &u
}

func (r *PostPaymentToFolioRequest) Do(ctx context.Context) (PostPaymentToFolioResponseBody, error) {
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
