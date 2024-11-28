package apaleo

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetReservationsRequest() GetReservationsRequest {
	return GetReservationsRequest{
		client:      c,
		queryParams: c.NewGetReservationsQueryParams(),
		pathParams:  c.NewGetReservationsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetReservationsRequestBody(),
	}
}

type GetReservationsRequest struct {
	client      *Client
	queryParams *GetReservationsQueryParams
	pathParams  *GetReservationsPathParams
	method      string
	headers     http.Header
	requestBody GetReservationsRequestBody
}

func (c *Client) NewGetReservationsQueryParams() *GetReservationsQueryParams {
	return &GetReservationsQueryParams{}
}

type GetReservationsQueryParams struct {
}

func (p GetReservationsQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetReservationsRequest) QueryParams() *GetReservationsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetReservationsPathParams() *GetReservationsPathParams {
	return &GetReservationsPathParams{}
}

type GetReservationsPathParams struct {
}

func (p *GetReservationsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetReservationsRequest) PathParams() *GetReservationsPathParams {
	return r.pathParams
}

func (r *GetReservationsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetReservationsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetReservationsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetReservationsRequestBody() GetReservationsRequestBody {
	return GetReservationsRequestBody{}
}

type GetReservationsRequestBody struct {
}

func (r *GetReservationsRequest) RequestBody() *GetReservationsRequestBody {
	return nil
}

func (r *GetReservationsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetReservationsRequest) SetRequestBody(body GetReservationsRequestBody) {
	r.requestBody = body
}

func (r *GetReservationsRequest) NewResponseBody() *GetReservationsResponseBody {
	return &GetReservationsResponseBody{}
}

type GetReservationsResponseBody struct {
	Count        int   `json:"count"`
	Reservations []any `json:"reservations"`
}

func (r *GetReservationsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("booking/v1/reservations", r.PathParams())
	return &u
}

func (r *GetReservationsRequest) Do() (GetReservationsResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
