// Package for auto-assigning units to a reservation in NSFW API
package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPutReservationCheckInRequest() PutReservationCheckInRequest {
	return PutReservationCheckInRequest{
		client:      c,
		queryParams: c.NewPutReservationCheckInQueryParams(),
		pathParams:  c.NewPutReservationCheckInPathParams(),
		method:      http.MethodPut,
		headers:     http.Header{},
	}
}

type PutReservationCheckInRequest struct {
	client      *Client
	queryParams *PutReservationCheckInQueryParams
	pathParams  *PutReservationCheckInPathParams
	method      string
	headers     http.Header
}

func (c *Client) NewPutReservationCheckInQueryParams() *PutReservationCheckInQueryParams {
	return &PutReservationCheckInQueryParams{}
}

type PutReservationCheckInQueryParams struct {
	WithCityTax bool `schema:"withCityTax,omitempty"`
}

func (p PutReservationCheckInQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutReservationCheckInRequest) QueryParams() *PutReservationCheckInQueryParams {
	return r.queryParams
}

func (c *Client) NewPutReservationCheckInPathParams() *PutReservationCheckInPathParams {
	return &PutReservationCheckInPathParams{}
}

type PutReservationCheckInPathParams struct {
	ReservationID string `schema:"id"`
}

func (p *PutReservationCheckInPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ReservationID,
	}
}

func (r *PutReservationCheckInRequest) PathParams() *PutReservationCheckInPathParams {
	return r.pathParams
}

func (r *PutReservationCheckInRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PutReservationCheckInRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutReservationCheckInRequest) Method() string {
	return r.method
}

// Added RequestBody and RequestBodyInterface methods to implement Request interface
func (r *PutReservationCheckInRequest) RequestBody() *PutReservationCheckInRequestBody {
	return nil
}

func (r *PutReservationCheckInRequest) RequestBodyInterface() interface{} {
	return nil
}

// Defining empty request body struct
type PutReservationCheckInRequestBody struct{}

func (r *PutReservationCheckInRequest) NewResponseBody() *PutReservationCheckInResponseBody {
	return &PutReservationCheckInResponseBody{}
}

// PutReservationCheckInResponseBody matches AutoAssignedUnitListModel from the API
type PutReservationCheckInResponseBody struct{}

func (r *PutReservationCheckInRequest) URL() *url.URL {
	path := "booking/v0-nsfw/reservation-actions/{{.id}}/checkin"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PutReservationCheckInRequest) Do() (PutReservationCheckInResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(context.TODO(), r)
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
