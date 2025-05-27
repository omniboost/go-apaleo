// Package for auto-assigning units to a reservation in NSFW API
package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPutReservationAssignUnitByIDRequest() PutReservationAssignUnitByIDRequest {
	return PutReservationAssignUnitByIDRequest{
		client:      c,
		queryParams: c.NewPutReservationAssignUnitByIDQueryParams(),
		pathParams:  c.NewPutReservationAssignUnitByIDPathParams(),
		method:      http.MethodPut,
		headers:     http.Header{},
	}
}

type PutReservationAssignUnitByIDRequest struct {
	client      *Client
	queryParams *PutReservationAssignUnitByIDQueryParams
	pathParams  *PutReservationAssignUnitByIDPathParams
	method      string
	headers     http.Header
}

func (c *Client) NewPutReservationAssignUnitByIDQueryParams() *PutReservationAssignUnitByIDQueryParams {
	return &PutReservationAssignUnitByIDQueryParams{}
}

type PutReservationAssignUnitByIDQueryParams struct {
	UnitConditions []string `schema:"unitConditions,omitempty"`
}

func (p PutReservationAssignUnitByIDQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PutReservationAssignUnitByIDRequest) QueryParams() *PutReservationAssignUnitByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewPutReservationAssignUnitByIDPathParams() *PutReservationAssignUnitByIDPathParams {
	return &PutReservationAssignUnitByIDPathParams{}
}

type PutReservationAssignUnitByIDPathParams struct {
	ReservationID string `schema:"id"`
	UnitID        string `schema:"unitId"`
}

func (p *PutReservationAssignUnitByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id":     p.ReservationID,
		"unitId": p.UnitID,
	}
}

func (r *PutReservationAssignUnitByIDRequest) PathParams() *PutReservationAssignUnitByIDPathParams {
	return r.pathParams
}

func (r *PutReservationAssignUnitByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PutReservationAssignUnitByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *PutReservationAssignUnitByIDRequest) Method() string {
	return r.method
}

// Added RequestBody and RequestBodyInterface methods to implement Request interface
func (r *PutReservationAssignUnitByIDRequest) RequestBody() *PutReservationAssignUnitByIDRequestBody {
	return nil
}

func (r *PutReservationAssignUnitByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

// Defining empty request body struct
type PutReservationAssignUnitByIDRequestBody struct{}

func (r *PutReservationAssignUnitByIDRequest) NewResponseBody() *PutReservationAssignUnitByIDResponseBody {
	return &PutReservationAssignUnitByIDResponseBody{}
}

// PutReservationAssignUnitByIDResponseBody matches AutoAssignedUnitListModel from the API
type PutReservationAssignUnitByIDResponseBody struct {
	Unit EmbeddedUnitModel `json:"unit"`
}

func (r *PutReservationAssignUnitByIDRequest) URL() *url.URL {
	path := "booking/v0-nsfw/reservation-actions/{{.id}}/assign-unit/{{.unitId}}"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PutReservationAssignUnitByIDRequest) Do(ctx context.Context) (PutReservationAssignUnitByIDResponseBody, error) {
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
