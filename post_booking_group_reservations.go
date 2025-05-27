package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostBookingGroupReservationsRequest() PostBookingGroupReservationsRequest {
	return PostBookingGroupReservationsRequest{
		client:      c,
		queryParams: c.NewPostBookingGroupReservationsQueryParams(),
		pathParams:  c.NewPostBookingGroupReservationsPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostBookingGroupReservationsRequestBody(),
	}
}

type PostBookingGroupReservationsRequest struct {
	client      *Client
	queryParams *PostBookingGroupReservationsQueryParams
	pathParams  *PostBookingGroupReservationsPathParams
	method      string
	headers     http.Header
	requestBody PostBookingGroupReservationsRequestBody
}

func (c *Client) NewPostBookingGroupReservationsQueryParams() *PostBookingGroupReservationsQueryParams {
	return &PostBookingGroupReservationsQueryParams{}
}

type PostBookingGroupReservationsQueryParams struct{}

func (p PostBookingGroupReservationsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostBookingGroupReservationsRequest) QueryParams() *PostBookingGroupReservationsQueryParams {
	return r.queryParams
}

func (c *Client) NewPostBookingGroupReservationsPathParams() *PostBookingGroupReservationsPathParams {
	return &PostBookingGroupReservationsPathParams{}
}

type PostBookingGroupReservationsPathParams struct {
	GroupID string `schema:"id"`
}

func (p *PostBookingGroupReservationsPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.GroupID,
	}
}

func (r *PostBookingGroupReservationsRequest) PathParams() *PostBookingGroupReservationsPathParams {
	return r.pathParams
}

func (r *PostBookingGroupReservationsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostBookingGroupReservationsRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostBookingGroupReservationsRequest) Method() string {
	return r.method
}

func (s *Client) NewPostBookingGroupReservationsRequestBody() PostBookingGroupReservationsRequestBody {
	return PostBookingGroupReservationsRequestBody{}
}

type PostBookingGroupReservationsRequestBody struct {
	Reservations []PickUpReservationModel `json:"reservations,omitempty"`
}

func (r PostBookingGroupReservationsRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PostBookingGroupReservationsRequest) RequestBody() *PostBookingGroupReservationsRequestBody {
	return &r.requestBody
}

func (r *PostBookingGroupReservationsRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostBookingGroupReservationsRequest) SetRequestBody(body PostBookingGroupReservationsRequestBody) {
	r.requestBody = body
}

func (r *PostBookingGroupReservationsRequest) NewResponseBody() *PostBookingGroupReservationsResponseBody {
	return &PostBookingGroupReservationsResponseBody{}
}

type PostBookingGroupReservationsResponseBody struct {
	ReservationIDs []struct {
		ID string `json:"id"`
	} `json:"reservationIds"`
}

func (r *PostBookingGroupReservationsRequest) URL() *url.URL {
	path := "booking/v1/groups/{{.id}}/reservations"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PostBookingGroupReservationsRequest) Do(ctx context.Context) (PostBookingGroupReservationsResponseBody, error) {
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
