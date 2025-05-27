package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostBookingBlockRequest() PostBookingBlockRequest {
	return PostBookingBlockRequest{
		client:      c,
		queryParams: c.NewPostBookingBlockQueryParams(),
		pathParams:  c.NewPostBookingBlockPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostBookingBlockRequestBody(),
	}
}

func (c *Client) NewPostBookingBlockForceRequest() PostBookingBlockRequest {
	return PostBookingBlockRequest{
		client:      c,
		queryParams: c.NewPostBookingBlockQueryParams(),
		pathParams:  c.NewPostBookingBlockPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostBookingBlockRequestBody(),
	}
}

type PostBookingBlockRequest struct {
	client      *Client
	queryParams *PostBookingBlockQueryParams
	pathParams  *PostBookingBlockPathParams
	method      string
	headers     http.Header
	requestBody PostBookingBlockRequestBody
	force       bool
}

func (c *Client) NewPostBookingBlockQueryParams() *PostBookingBlockQueryParams {
	return &PostBookingBlockQueryParams{}
}

type PostBookingBlockQueryParams struct{}

func (p PostBookingBlockQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostBookingBlockRequest) QueryParams() *PostBookingBlockQueryParams {
	return r.queryParams
}

func (c *Client) NewPostBookingBlockPathParams() *PostBookingBlockPathParams {
	return &PostBookingBlockPathParams{}
}

type PostBookingBlockPathParams struct {
}

func (p *PostBookingBlockPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostBookingBlockRequest) PathParams() *PostBookingBlockPathParams {
	return r.pathParams
}

func (r *PostBookingBlockRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostBookingBlockRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostBookingBlockRequest) Method() string {
	return r.method
}

func (s *Client) NewPostBookingBlockRequestBody() PostBookingBlockRequestBody {
	return PostBookingBlockRequestBody{}
}

type PostBookingBlockRequestBody struct {
	GroupID        string                      `json:"groupId"`
	RatePlanID     string                      `json:"ratePlanId"`
	From           Date                        `json:"from"`
	To             Date                        `json:"to"`
	GrossDailyRate MonetaryValueModel          `json:"grossDailyRate"`
	TimeSlices     []CreateBlockTimeSliceModel `json:"timeSlices,omitempty"`
	BlockedUnits   int32                       `json:"blockedUnits,omitempty"`
	PromoCode      string                      `json:"promoCode,omitempty"`
	CorporateCode  string                      `json:"corporateCode,omitempty"`
}

func (r PostBookingBlockRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PostBookingBlockRequest) RequestBody() *PostBookingBlockRequestBody {
	return &r.requestBody
}

func (r *PostBookingBlockRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostBookingBlockRequest) SetRequestBody(body PostBookingBlockRequestBody) {
	r.requestBody = body
}

func (r *PostBookingBlockRequest) NewResponseBody() *PostBookingBlockResponseBody {
	return &PostBookingBlockResponseBody{}
}

type PostBookingBlockResponseBody struct {
	ID string `json:"id"`
}

func (r *PostBookingBlockRequest) URL() *url.URL {
	path := "booking/v1/blocks"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PostBookingBlockRequest) Do(ctx context.Context) (PostBookingBlockResponseBody, error) {
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
