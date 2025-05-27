// Package for auto-assigning units to a reservation in NSFW API
package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPatchReservationNSFWRequest() PatchReservationNSFWRequest {
	return PatchReservationNSFWRequest{
		client:      c,
		queryParams: c.NewPatchReservationNSFWQueryParams(),
		pathParams:  c.NewPatchReservationNSFWPathParams(),
		method:      http.MethodPatch,
		headers:     http.Header{},
	}
}

type PatchReservationNSFWRequest struct {
	client      *Client
	queryParams *PatchReservationNSFWQueryParams
	pathParams  *PatchReservationNSFWPathParams
	method      string
	headers     http.Header
	requestBody PatchReservationNSFWRequestBody
}

func (c *Client) NewPatchReservationNSFWQueryParams() *PatchReservationNSFWQueryParams {
	return &PatchReservationNSFWQueryParams{}
}

type PatchReservationNSFWQueryParams struct{}

func (p PatchReservationNSFWQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PatchReservationNSFWRequest) QueryParams() *PatchReservationNSFWQueryParams {
	return r.queryParams
}

func (c *Client) NewPatchReservationNSFWPathParams() *PatchReservationNSFWPathParams {
	return &PatchReservationNSFWPathParams{}
}

type PatchReservationNSFWPathParams struct {
	ReservationID string `schema:"id"`
}

func (p *PatchReservationNSFWPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ReservationID,
	}
}

func (r *PatchReservationNSFWRequest) PathParams() *PatchReservationNSFWPathParams {
	return r.pathParams
}

func (r *PatchReservationNSFWRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PatchReservationNSFWRequest) SetMethod(method string) {
	r.method = method
}

func (r *PatchReservationNSFWRequest) Method() string {
	return r.method
}

func (s *Client) NewPatchReservationNSFWRequestBody() PatchReservationNSFWRequestBody {
	return PatchReservationNSFWRequestBody{}
}

type PatchReservationNSFWRequestBody []Operation

func (r *PatchReservationNSFWRequest) RequestBody() *PatchReservationNSFWRequestBody {
	return &r.requestBody
}

func (r *PatchReservationNSFWRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PatchReservationNSFWRequest) SetRequestBody(body PatchReservationNSFWRequestBody) {
	r.requestBody = body
}

func (r *PatchReservationNSFWRequest) NewResponseBody() *PatchReservationNSFWResponseBody {
	return &PatchReservationNSFWResponseBody{}
}

type PatchReservationNSFWResponseBody struct{}

func (r *PatchReservationNSFWRequest) URL() *url.URL {
	path := "booking/v0-nsfw/reservations/{{.id}}"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PatchReservationNSFWRequest) Do(ctx context.Context) (PatchReservationNSFWResponseBody, error) {
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
