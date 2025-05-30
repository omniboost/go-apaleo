package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostBookingReservationsRequest() PostBookingReservationsRequest {
	return PostBookingReservationsRequest{
		client:      c,
		queryParams: c.NewPostBookingReservationsQueryParams(),
		pathParams:  c.NewPostBookingReservationsPathParams(),
		method:      http.MethodPost,
		headers:     c.NewPostBookingReservationsHeaders(),
		requestBody: c.NewPostBookingReservationsRequestBody(),
		force:       false,
	}
}

func (c *Client) NewPostBookingReservationsForceRequest() PostBookingReservationsRequest {
	return PostBookingReservationsRequest{
		client:      c,
		queryParams: c.NewPostBookingReservationsQueryParams(),
		pathParams:  c.NewPostBookingReservationsPathParams(),
		method:      http.MethodPost,
		headers:     c.NewPostBookingReservationsHeaders(),
		requestBody: c.NewPostBookingReservationsRequestBody(),
		force:       true,
	}
}

type PostBookingReservationsRequest struct {
	client      *Client
	queryParams *PostBookingReservationsQueryParams
	pathParams  *PostBookingReservationsPathParams
	method      string
	headers     *PostBookingReservationsHeaders
	requestBody PostBookingReservationsRequestBody
	force       bool
}

func (c *Client) NewPostBookingReservationsQueryParams() *PostBookingReservationsQueryParams {
	return &PostBookingReservationsQueryParams{}
}

type PostBookingReservationsQueryParams struct{}

func (p PostBookingReservationsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostBookingReservationsRequest) QueryParams() *PostBookingReservationsQueryParams {
	return r.queryParams
}

func (c *Client) NewPostBookingReservationsHeaders() *PostBookingReservationsHeaders {
	return &PostBookingReservationsHeaders{}
}

type PostBookingReservationsHeaders struct {
	IdempotencyKey string `schema:"Idempotency-Key,omitempty"`
}

func (r *PostBookingReservationsRequest) Headers() *PostBookingReservationsHeaders {
	return r.headers
}

func (c *Client) NewPostBookingReservationsPathParams() *PostBookingReservationsPathParams {
	return &PostBookingReservationsPathParams{}
}

type PostBookingReservationsPathParams struct {
	BookingID string `schema:"id"`
}

func (p *PostBookingReservationsPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.BookingID,
	}
}

func (r *PostBookingReservationsRequest) PathParams() *PostBookingReservationsPathParams {
	return r.pathParams
}

func (r *PostBookingReservationsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostBookingReservationsRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostBookingReservationsRequest) Method() string {
	return r.method
}

func (s *Client) NewPostBookingReservationsRequestBody() PostBookingReservationsRequestBody {
	return PostBookingReservationsRequestBody{}
}

type PostBookingReservationsRequestBody struct {
	Reservations []CreateReservationModel `json:"reservations,omitempty"`
}

func (r PostBookingReservationsRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PostBookingReservationsRequest) RequestBody() *PostBookingReservationsRequestBody {
	return &r.requestBody
}

func (r *PostBookingReservationsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostBookingReservationsRequest) SetRequestBody(body PostBookingReservationsRequestBody) {
	r.requestBody = body
}

func (r *PostBookingReservationsRequest) NewResponseBody() *PostBookingReservationsResponseBody {
	return &PostBookingReservationsResponseBody{}
}

type PostBookingReservationsResponseBody struct {
	ReservationIDs []struct {
		ID string `json:"id"`
	} `json:"reservationIds"`
}

func (r *PostBookingReservationsRequest) URL() *url.URL {
	path := "booking/v1/bookings/{{.id}}/reservations"
	if r.force {
		path = path + "/$force"
	}
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PostBookingReservationsRequest) Do(ctx context.Context) (PostBookingReservationsResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Add headers
	if r.Headers().IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", r.Headers().IdempotencyKey)
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
