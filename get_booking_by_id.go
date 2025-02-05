package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetBookingByIDRequest() GetBookingByIDRequest {
	return GetBookingByIDRequest{
		client:      c,
		queryParams: c.NewGetBookingByIDQueryParams(),
		pathParams:  c.NewGetBookingByIDPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetBookingByIDRequestBody(),
	}
}

type GetBookingByIDRequest struct {
	client      *Client
	queryParams *GetBookingByIDQueryParams
	pathParams  *GetBookingByIDPathParams
	method      string
	headers     http.Header
	requestBody GetBookingByIDRequestBody
}

func (c *Client) NewGetBookingByIDQueryParams() *GetBookingByIDQueryParams {
	return &GetBookingByIDQueryParams{}
}

type GetBookingByIDQueryParams struct {
	Expand []string `schema:"expand,omitempty"`
}

func (p GetBookingByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetBookingByIDRequest) QueryParams() *GetBookingByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetBookingByIDPathParams() *GetBookingByIDPathParams {
	return &GetBookingByIDPathParams{}
}

type GetBookingByIDPathParams struct {
	ID string `schema:"id"`
}

func (p *GetBookingByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetBookingByIDRequest) PathParams() *GetBookingByIDPathParams {
	return r.pathParams
}

func (r *GetBookingByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetBookingByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetBookingByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetBookingByIDRequestBody() GetBookingByIDRequestBody {
	return GetBookingByIDRequestBody{}
}

type GetBookingByIDRequestBody struct {
}

func (r *GetBookingByIDRequest) RequestBody() *GetBookingByIDRequestBody {
	return nil
}

func (r *GetBookingByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetBookingByIDRequest) SetRequestBody(body GetBookingByIDRequestBody) {
	r.requestBody = body
}

func (r *GetBookingByIDRequest) NewResponseBody() *GetBookingByIDResponseBody {
	return &GetBookingByIDResponseBody{}
}

type GetBookingByIDResponseBody BookingItemModel

func (r *GetBookingByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("booking/v1/bookings/{{.id}}", r.PathParams())
	return &u
}

func (r *GetBookingByIDRequest) Do() (GetBookingByIDResponseBody, error) {
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
