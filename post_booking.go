package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostBookingRequest() PostBookingRequest {
	return PostBookingRequest{
		client:      c,
		queryParams: c.NewPostBookingQueryParams(),
		pathParams:  c.NewPostBookingPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostBookingRequestBody(),
		force:       false,
	}
}

func (c *Client) NewPostBookingForceRequest() PostBookingRequest {
	return PostBookingRequest{
		client:      c,
		queryParams: c.NewPostBookingQueryParams(),
		pathParams:  c.NewPostBookingPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostBookingRequestBody(),
		force:       true,
	}
}

type PostBookingRequest struct {
	client      *Client
	queryParams *PostBookingQueryParams
	pathParams  *PostBookingPathParams
	method      string
	headers     http.Header
	requestBody PostBookingRequestBody
	force       bool
}

func (c *Client) NewPostBookingQueryParams() *PostBookingQueryParams {
	return &PostBookingQueryParams{}
}

type PostBookingQueryParams struct{}

func (p PostBookingQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostBookingRequest) QueryParams() *PostBookingQueryParams {
	return r.queryParams
}

func (c *Client) NewPostBookingPathParams() *PostBookingPathParams {
	return &PostBookingPathParams{}
}

type PostBookingPathParams struct {
}

func (p *PostBookingPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostBookingRequest) PathParams() *PostBookingPathParams {
	return r.pathParams
}

func (r *PostBookingRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostBookingRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostBookingRequest) Method() string {
	return r.method
}

func (s *Client) NewPostBookingRequestBody() PostBookingRequestBody {
	return PostBookingRequestBody{}
}

type PostBookingRequestBody struct {
	PaymentAccount       CreatePaymentAccountModel `json:"paymentAccount,omitempty"`
	Booker               BookerModel               `json:"booker"`
	Comment              string                    `json:"comment,omitempty"`
	BookerComment        string                    `json:"bookerComment,omitempty"`
	Reservations         []CreateReservationModel  `json:"reservations,omitempty"`
	TransactionReference string                    `json:"transactionReference,omitempty"`
}

func (r PostBookingRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PostBookingRequest) RequestBody() *PostBookingRequestBody {
	return &r.requestBody
}

func (r *PostBookingRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostBookingRequest) SetRequestBody(body PostBookingRequestBody) {
	r.requestBody = body
}

func (r *PostBookingRequest) NewResponseBody() *PostBookingResponseBody {
	return &PostBookingResponseBody{}
}

type PostBookingResponseBody struct {
	ID             string `json:"id"`
	ReservationIDs []struct {
		ID string `json:"id"`
	} `json:"reservationIds"`
}

func (r *PostBookingRequest) URL() *url.URL {
	path := "booking/v1/bookings"
	if r.force {
		path = path + "/$force"
	}
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PostBookingRequest) Do(ctx context.Context) (PostBookingResponseBody, error) {
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
