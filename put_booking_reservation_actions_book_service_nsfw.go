// Package for auto-assigning units to a reservation in NSFW API
package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPutReservationActionsBookServiceRequest() PutReservationActionsBookServiceRequest {
	return PutReservationActionsBookServiceRequest{
		client:      c,
		queryParams: c.NewPutReservationActionsBookServiceQueryParams(),
		pathParams:  c.NewPutReservationActionsBookServicePathParams(),
		method:      http.MethodPut,
		headers:     http.Header{},
		force:       false,
	}
}

func (c *Client) NewPutReservationActionsBookServiceForceRequest() PutReservationActionsBookServiceRequest {
	return PutReservationActionsBookServiceRequest{
		client:      c,
		queryParams: c.NewPutReservationActionsBookServiceQueryParams(),
		pathParams:  c.NewPutReservationActionsBookServicePathParams(),
		method:      http.MethodPut,
		headers:     http.Header{},
		requestBody: c.NewPutReservationActionsBookServiceRequestBody(),
		force:       true,
	}
}

type PutReservationActionsBookServiceRequest struct {
	client      *Client
	queryParams *PutReservationActionsBookServiceQueryParams
	pathParams  *PutReservationActionsBookServicePathParams
	method      string
	headers     http.Header
	requestBody PutReservationActionsBookServiceRequestBody
	force       bool
}

func (c *Client) NewPutReservationActionsBookServiceQueryParams() *PutReservationActionsBookServiceQueryParams {
	return &PutReservationActionsBookServiceQueryParams{}
}

type PutReservationActionsBookServiceQueryParams struct{}

func (p PutReservationActionsBookServiceQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutReservationActionsBookServiceRequest) QueryParams() *PutReservationActionsBookServiceQueryParams {
	return r.queryParams
}

func (c *Client) NewPutReservationActionsBookServicePathParams() *PutReservationActionsBookServicePathParams {
	return &PutReservationActionsBookServicePathParams{}
}

type PutReservationActionsBookServicePathParams struct {
	ReservationID string `schema:"id"`
}

func (p *PutReservationActionsBookServicePathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ReservationID,
	}
}

func (r *PutReservationActionsBookServiceRequest) PathParams() *PutReservationActionsBookServicePathParams {
	return r.pathParams
}

func (r *PutReservationActionsBookServiceRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PutReservationActionsBookServiceRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutReservationActionsBookServiceRequest) Method() string {
	return r.method
}

func (s *Client) NewPutReservationActionsBookServiceRequestBody() PutReservationActionsBookServiceRequestBody {
	return PutReservationActionsBookServiceRequestBody{}
}

type PutReservationActionsBookServiceRequestBody struct {
	BookReservationServiceModel
}

func (r *PutReservationActionsBookServiceRequest) RequestBody() *PutReservationActionsBookServiceRequestBody {
	return &r.requestBody
}

func (r *PutReservationActionsBookServiceRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PutReservationActionsBookServiceRequest) SetRequestBody(body PutReservationActionsBookServiceRequestBody) {
	r.requestBody = body
}

func (r *PutReservationActionsBookServiceRequest) NewResponseBody() *PutReservationActionsBookServiceResponseBody {
	return &PutReservationActionsBookServiceResponseBody{}
}

// PutReservationActionsBookServiceResponseBody matches AutoAssignedUnitListModel from the API
type PutReservationActionsBookServiceResponseBody struct{}

func (r *PutReservationActionsBookServiceRequest) URL() *url.URL {
	path := "booking/v0-nsfw/reservation-actions/{{.id}}/book-service"
	if r.force {
		path = path + "/$force"
	}
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PutReservationActionsBookServiceRequest) Do() (PutReservationActionsBookServiceResponseBody, error) {
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
