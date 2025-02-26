package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostBookingReservationsNSFWRequest() PostBookingReservationsNSFWRequest {
	return PostBookingReservationsNSFWRequest{
		client:      c,
		queryParams: c.NewPostBookingReservationsNSFWQueryParams(),
		pathParams:  c.NewPostBookingReservationsNSFWPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostBookingReservationsNSFWRequestBody(),
		force:       false,
	}
}

func (c *Client) NewPostBookingReservationsNSFWForceRequest() PostBookingReservationsNSFWRequest {
	return PostBookingReservationsNSFWRequest{
		client:      c,
		queryParams: c.NewPostBookingReservationsNSFWQueryParams(),
		pathParams:  c.NewPostBookingReservationsNSFWPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostBookingReservationsNSFWRequestBody(),
		force:       true,
	}
}

type PostBookingReservationsNSFWRequest struct {
	client      *Client
	queryParams *PostBookingReservationsNSFWQueryParams
	pathParams  *PostBookingReservationsNSFWPathParams
	method      string
	headers     http.Header
	requestBody PostBookingReservationsNSFWRequestBody
	force       bool
}

func (c *Client) NewPostBookingReservationsNSFWQueryParams() *PostBookingReservationsNSFWQueryParams {
	return &PostBookingReservationsNSFWQueryParams{}
}

type PostBookingReservationsNSFWQueryParams struct{}

func (p PostBookingReservationsNSFWQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostBookingReservationsNSFWRequest) QueryParams() *PostBookingReservationsNSFWQueryParams {
	return r.queryParams
}

func (c *Client) NewPostBookingReservationsNSFWPathParams() *PostBookingReservationsNSFWPathParams {
	return &PostBookingReservationsNSFWPathParams{}
}

type PostBookingReservationsNSFWPathParams struct {
	BookingID string `schema:"id"`
}

func (p *PostBookingReservationsNSFWPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.BookingID,
	}
}

func (r *PostBookingReservationsNSFWRequest) PathParams() *PostBookingReservationsNSFWPathParams {
	return r.pathParams
}

func (r *PostBookingReservationsNSFWRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostBookingReservationsNSFWRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostBookingReservationsNSFWRequest) Method() string {
	return r.method
}

func (s *Client) NewPostBookingReservationsNSFWRequestBody() PostBookingReservationsNSFWRequestBody {
	return PostBookingReservationsNSFWRequestBody{}
}

type PostBookingReservationsNSFWRequestBody struct {
	Reservations []CreateReservationNSFWModel `json:"reservations,omitempty"`
}

func (r PostBookingReservationsNSFWRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PostBookingReservationsNSFWRequest) RequestBody() *PostBookingReservationsNSFWRequestBody {
	return &r.requestBody
}

func (r *PostBookingReservationsNSFWRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostBookingReservationsNSFWRequest) SetRequestBody(body PostBookingReservationsNSFWRequestBody) {
	r.requestBody = body
}

func (r *PostBookingReservationsNSFWRequest) NewResponseBody() *PostBookingReservationsNSFWResponseBody {
	return &PostBookingReservationsNSFWResponseBody{}
}

type PostBookingReservationsNSFWResponseBody struct {
	ReservationIDs []struct {
		ID string `json:"id"`
	} `json:"reservationIds"`
}

func (r *PostBookingReservationsNSFWRequest) URL() *url.URL {
	path := "booking/v0-nsfw/bookings/{{.id}}/reservations"
	if r.force {
		path = path + "/$force"
	}
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PostBookingReservationsNSFWRequest) Do() (PostBookingReservationsNSFWResponseBody, error) {
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
