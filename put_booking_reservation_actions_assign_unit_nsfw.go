// Package for auto-assigning units to a reservation in NSFW API
package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPutReservationAssignUnitRequest() PutReservationAssignUnitRequest {
	return PutReservationAssignUnitRequest{
		client:      c,
		queryParams: c.NewPutReservationAssignUnitQueryParams(),
		pathParams:  c.NewPutReservationAssignUnitPathParams(),
		method:      http.MethodPut,
		headers:     http.Header{},
	}
}

type PutReservationAssignUnitRequest struct {
	client      *Client
	queryParams *PutReservationAssignUnitQueryParams
	pathParams  *PutReservationAssignUnitPathParams
	method      string
	headers     http.Header
}

func (c *Client) NewPutReservationAssignUnitQueryParams() *PutReservationAssignUnitQueryParams {
	return &PutReservationAssignUnitQueryParams{}
}

type PutReservationAssignUnitQueryParams struct {
	UnitConditions []string `schema:"unitConditions,omitempty"`
}

func (p PutReservationAssignUnitQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutReservationAssignUnitRequest) QueryParams() *PutReservationAssignUnitQueryParams {
	return r.queryParams
}

func (c *Client) NewPutReservationAssignUnitPathParams() *PutReservationAssignUnitPathParams {
	return &PutReservationAssignUnitPathParams{}
}

type PutReservationAssignUnitPathParams struct {
	ReservationID string `schema:"id"`
}

func (p *PutReservationAssignUnitPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ReservationID,
	}
}

func (r *PutReservationAssignUnitRequest) PathParams() *PutReservationAssignUnitPathParams {
	return r.pathParams
}

func (r *PutReservationAssignUnitRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PutReservationAssignUnitRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutReservationAssignUnitRequest) Method() string {
	return r.method
}

// Added RequestBody and RequestBodyInterface methods to implement Request interface
func (r *PutReservationAssignUnitRequest) RequestBody() *PutReservationAssignUnitRequestBody {
	return nil
}

func (r *PutReservationAssignUnitRequest) RequestBodyInterface() interface{} {
	return nil
}

// Defining empty request body struct
type PutReservationAssignUnitRequestBody struct{}

func (r *PutReservationAssignUnitRequest) NewResponseBody() *PutReservationAssignUnitResponseBody {
	return &PutReservationAssignUnitResponseBody{}
}

// PutReservationAssignUnitResponseBody matches AutoAssignedUnitListModel from the API
type PutReservationAssignUnitResponseBody struct {
	TimeSlices []AutoAssignedUnitItemModel `json:"timeSlices"`
}

// AutoAssignedUnitItemModel represents an assigned unit with time slice
type AutoAssignedUnitItemModel struct {
	Unit EmbeddedUnitModel `json:"unit"`
	From DateTime          `json:"from"`
	To   DateTime          `json:"to"`
}

func (r *PutReservationAssignUnitRequest) URL() *url.URL {
	path := "booking/v0-nsfw/reservation-actions/{{.id}}/assign-unit"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PutReservationAssignUnitRequest) Do() (PutReservationAssignUnitResponseBody, error) {
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
