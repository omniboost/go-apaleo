package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPutBookingReservationActionsRequest() PutBookingReservationActionsRequest {
	return PutBookingReservationActionsRequest{
		client:      c,
		queryParams: c.NewPutBookingReservationActionsQueryParams(),
		pathParams:  c.NewPutBookingReservationActionsPathParams(),
		method:      http.MethodPut,
		headers:     http.Header{},
		requestBody: c.NewPutBookingReservationActionsRequestBody(),
	}
}

type PutBookingReservationActionsRequest struct {
	client      *Client
	queryParams *PutBookingReservationActionsQueryParams
	pathParams  *PutBookingReservationActionsPathParams
	method      string
	headers     http.Header
	requestBody PutBookingReservationActionsRequestBody
}

func (c *Client) NewPutBookingReservationActionsQueryParams() *PutBookingReservationActionsQueryParams {
	return &PutBookingReservationActionsQueryParams{}
}

type PutBookingReservationActionsQueryParams struct{}

func (p PutBookingReservationActionsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PutBookingReservationActionsRequest) QueryParams() *PutBookingReservationActionsQueryParams {
	return r.queryParams
}

func (c *Client) NewPutBookingReservationActionsPathParams() *PutBookingReservationActionsPathParams {
	return &PutBookingReservationActionsPathParams{}
}

type PutBookingReservationActionsPathParams struct {
	ReservationID string `schema:"id"`
	Action        string `schema:"action"`
}

func (p *PutBookingReservationActionsPathParams) Params() map[string]string {
	return map[string]string{
		"id":     p.ReservationID,
		"action": p.Action,
	}
}

func (r *PutBookingReservationActionsRequest) PathParams() *PutBookingReservationActionsPathParams {
	return r.pathParams
}

func (r *PutBookingReservationActionsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PutBookingReservationActionsRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutBookingReservationActionsRequest) Method() string {
	return r.method
}

func (s *Client) NewPutBookingReservationActionsRequestBody() PutBookingReservationActionsRequestBody {
	return PutBookingReservationActionsRequestBody{}
}

type PutBookingReservationActionsRequestBody struct{}

func (r PutBookingReservationActionsRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PutBookingReservationActionsRequest) RequestBody() *PutBookingReservationActionsRequestBody {
	return nil
}

func (r *PutBookingReservationActionsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *PutBookingReservationActionsRequest) SetRequestBody(body PutBookingReservationActionsRequestBody) {
	r.requestBody = body
}

func (r *PutBookingReservationActionsRequest) NewResponseBody() *PutBookingReservationActionsResponseBody {
	return &PutBookingReservationActionsResponseBody{}
}

type PutBookingReservationActionsResponseBody struct{}

func (r *PutBookingReservationActionsRequest) URL() *url.URL {
	path := "booking/v1/reservation-actions/{{.id}}/{{.action}}"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PutBookingReservationActionsRequest) Do(ctx context.Context) (PutBookingReservationActionsResponseBody, error) {
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
