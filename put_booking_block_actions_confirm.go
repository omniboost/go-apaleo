package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPutBookingBlockActionsConfirmRequest() PutBookingBlockActionsConfirmRequest {
	return PutBookingBlockActionsConfirmRequest{
		client:      c,
		queryParams: c.NewPutBookingBlockActionsConfirmQueryParams(),
		pathParams:  c.NewPutBookingBlockActionsConfirmPathParams(),
		method:      http.MethodPut,
		headers:     http.Header{},
		requestBody: c.NewPutBookingBlockActionsConfirmRequestBody(),
	}
}

type PutBookingBlockActionsConfirmRequest struct {
	client      *Client
	queryParams *PutBookingBlockActionsConfirmQueryParams
	pathParams  *PutBookingBlockActionsConfirmPathParams
	method      string
	headers     http.Header
	requestBody PutBookingBlockActionsConfirmRequestBody
}

func (c *Client) NewPutBookingBlockActionsConfirmQueryParams() *PutBookingBlockActionsConfirmQueryParams {
	return &PutBookingBlockActionsConfirmQueryParams{}
}

type PutBookingBlockActionsConfirmQueryParams struct{}

func (p PutBookingBlockActionsConfirmQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PutBookingBlockActionsConfirmRequest) QueryParams() *PutBookingBlockActionsConfirmQueryParams {
	return r.queryParams
}

func (c *Client) NewPutBookingBlockActionsConfirmPathParams() *PutBookingBlockActionsConfirmPathParams {
	return &PutBookingBlockActionsConfirmPathParams{}
}

type PutBookingBlockActionsConfirmPathParams struct {
	ID string `schema:"id"`
}

func (p *PutBookingBlockActionsConfirmPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *PutBookingBlockActionsConfirmRequest) PathParams() *PutBookingBlockActionsConfirmPathParams {
	return r.pathParams
}

func (r *PutBookingBlockActionsConfirmRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PutBookingBlockActionsConfirmRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutBookingBlockActionsConfirmRequest) Method() string {
	return r.method
}

func (s *Client) NewPutBookingBlockActionsConfirmRequestBody() PutBookingBlockActionsConfirmRequestBody {
	return PutBookingBlockActionsConfirmRequestBody{}
}

type PutBookingBlockActionsConfirmRequestBody struct{}

func (r PutBookingBlockActionsConfirmRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PutBookingBlockActionsConfirmRequest) RequestBody() *PutBookingBlockActionsConfirmRequestBody {
	return nil
}

func (r *PutBookingBlockActionsConfirmRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *PutBookingBlockActionsConfirmRequest) SetRequestBody(body PutBookingBlockActionsConfirmRequestBody) {
	r.requestBody = body
}

func (r *PutBookingBlockActionsConfirmRequest) NewResponseBody() *PutBookingBlockActionsConfirmResponseBody {
	return &PutBookingBlockActionsConfirmResponseBody{}
}

type PutBookingBlockActionsConfirmResponseBody struct{}

func (r *PutBookingBlockActionsConfirmRequest) URL() *url.URL {
	path := "booking/v1/block-actions/{{.id}}/confirm"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PutBookingBlockActionsConfirmRequest) Do() (PutBookingBlockActionsConfirmResponseBody, error) {
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
