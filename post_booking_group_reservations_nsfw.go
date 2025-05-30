package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostBookingGroupReservationsNSFWRequest() PostBookingGroupReservationsNSFWRequest {
	return PostBookingGroupReservationsNSFWRequest{
		client:      c,
		queryParams: c.NewPostBookingGroupReservationsNSFWQueryParams(),
		pathParams:  c.NewPostBookingGroupReservationsNSFWPathParams(),
		method:      http.MethodPost,
		headers:     c.NewPostBookingGroupReservationsNSFWHeaders(),
		requestBody: c.NewPostBookingGroupReservationsNSFWRequestBody(),
	}
}

type PostBookingGroupReservationsNSFWRequest struct {
	client      *Client
	queryParams *PostBookingGroupReservationsNSFWQueryParams
	pathParams  *PostBookingGroupReservationsNSFWPathParams
	method      string
	headers     *PostBookingGroupReservationsNSFWHeaders
	requestBody PostBookingGroupReservationsNSFWRequestBody
}

func (c *Client) NewPostBookingGroupReservationsNSFWQueryParams() *PostBookingGroupReservationsNSFWQueryParams {
	return &PostBookingGroupReservationsNSFWQueryParams{}
}

type PostBookingGroupReservationsNSFWQueryParams struct{}

func (p PostBookingGroupReservationsNSFWQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostBookingGroupReservationsNSFWRequest) QueryParams() *PostBookingGroupReservationsNSFWQueryParams {
	return r.queryParams
}

func (c *Client) NewPostBookingGroupReservationsNSFWHeaders() *PostBookingGroupReservationsNSFWHeaders {
	return &PostBookingGroupReservationsNSFWHeaders{}
}

type PostBookingGroupReservationsNSFWHeaders struct {
	IdempotencyKey string `schema:"Idempotency-Key,omitempty"`
}

func (r *PostBookingGroupReservationsNSFWRequest) Headers() *PostBookingGroupReservationsNSFWHeaders {
	return r.headers
}

func (c *Client) NewPostBookingGroupReservationsNSFWPathParams() *PostBookingGroupReservationsNSFWPathParams {
	return &PostBookingGroupReservationsNSFWPathParams{}
}

type PostBookingGroupReservationsNSFWPathParams struct {
	GroupID string `schema:"id"`
}

func (p *PostBookingGroupReservationsNSFWPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.GroupID,
	}
}

func (r *PostBookingGroupReservationsNSFWRequest) PathParams() *PostBookingGroupReservationsNSFWPathParams {
	return r.pathParams
}

func (r *PostBookingGroupReservationsNSFWRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostBookingGroupReservationsNSFWRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostBookingGroupReservationsNSFWRequest) Method() string {
	return r.method
}

func (s *Client) NewPostBookingGroupReservationsNSFWRequestBody() PostBookingGroupReservationsNSFWRequestBody {
	return PostBookingGroupReservationsNSFWRequestBody{}
}

type PostBookingGroupReservationsNSFWRequestBody struct {
	Reservations []PickUpReservationModel `json:"reservations,omitempty"`
}

func (r PostBookingGroupReservationsNSFWRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PostBookingGroupReservationsNSFWRequest) RequestBody() *PostBookingGroupReservationsNSFWRequestBody {
	return &r.requestBody
}

func (r *PostBookingGroupReservationsNSFWRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostBookingGroupReservationsNSFWRequest) SetRequestBody(body PostBookingGroupReservationsNSFWRequestBody) {
	r.requestBody = body
}

func (r *PostBookingGroupReservationsNSFWRequest) NewResponseBody() *PostBookingGroupReservationsNSFWResponseBody {
	return &PostBookingGroupReservationsNSFWResponseBody{}
}

type PostBookingGroupReservationsNSFWResponseBody struct {
	ReservationIDs []struct {
		ID string `json:"id"`
	} `json:"reservationIds"`
}

func (r *PostBookingGroupReservationsNSFWRequest) URL() *url.URL {
	path := "booking/v0-nsfw/groups/{{.id}}/reservations"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PostBookingGroupReservationsNSFWRequest) Do(ctx context.Context) (PostBookingGroupReservationsNSFWResponseBody, error) {
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
