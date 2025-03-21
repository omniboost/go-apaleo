package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostBookingNSFWRequest() PostBookingNSFWRequest {
	return PostBookingNSFWRequest{
		client:      c,
		queryParams: c.NewPostBookingNSFWQueryParams(),
		pathParams:  c.NewPostBookingNSFWPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostBookingNSFWRequestBody(),
		force:       false,
	}
}

func (c *Client) NewPostBookingNSFWForceRequest() PostBookingNSFWRequest {
	return PostBookingNSFWRequest{
		client:      c,
		queryParams: c.NewPostBookingNSFWQueryParams(),
		pathParams:  c.NewPostBookingNSFWPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostBookingNSFWRequestBody(),
		force:       true,
	}
}

type PostBookingNSFWRequest struct {
	client      *Client
	queryParams *PostBookingNSFWQueryParams
	pathParams  *PostBookingNSFWPathParams
	method      string
	headers     http.Header
	requestBody PostBookingNSFWRequestBody
	force       bool
}

func (c *Client) NewPostBookingNSFWQueryParams() *PostBookingNSFWQueryParams {
	return &PostBookingNSFWQueryParams{}
}

type PostBookingNSFWQueryParams struct{}

func (p PostBookingNSFWQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostBookingNSFWRequest) QueryParams() *PostBookingNSFWQueryParams {
	return r.queryParams
}

func (c *Client) NewPostBookingNSFWPathParams() *PostBookingNSFWPathParams {
	return &PostBookingNSFWPathParams{}
}

type PostBookingNSFWPathParams struct {
}

func (p *PostBookingNSFWPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostBookingNSFWRequest) PathParams() *PostBookingNSFWPathParams {
	return r.pathParams
}

func (r *PostBookingNSFWRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostBookingNSFWRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostBookingNSFWRequest) Method() string {
	return r.method
}

func (s *Client) NewPostBookingNSFWRequestBody() PostBookingNSFWRequestBody {
	return PostBookingNSFWRequestBody{}
}

type PostBookingNSFWRequestBody struct {
	PaymentAccount       CreatePaymentAccountModel    `json:"paymentAccount,omitempty"`
	Booker               BookerModel                  `json:"booker"`
	Comment              string                       `json:"comment,omitempty"`
	BookerComment        string                       `json:"bookerComment,omitempty"`
	Reservations         []CreateReservationNSFWModel `json:"reservations,omitempty"`
	TransactionReference string                       `json:"transactionReference,omitempty"`
}

func (r PostBookingNSFWRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PostBookingNSFWRequest) RequestBody() *PostBookingNSFWRequestBody {
	return &r.requestBody
}

func (r *PostBookingNSFWRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostBookingNSFWRequest) SetRequestBody(body PostBookingNSFWRequestBody) {
	r.requestBody = body
}

func (r *PostBookingNSFWRequest) NewResponseBody() *PostBookingNSFWResponseBody {
	return &PostBookingNSFWResponseBody{}
}

type PostBookingNSFWResponseBody struct {
	ID             string `json:"id"`
	ReservationIDs []struct {
		ID string `json:"id"`
	} `json:"reservationIds"`
}

func (r *PostBookingNSFWRequest) URL() *url.URL {
	path := "booking/v0-nsfw/bookings"
	if r.force {
		path = path + "/$force"
	}
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PostBookingNSFWRequest) Do() (PostBookingNSFWResponseBody, error) {
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
