package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetBookingsRequest() GetBookingsRequest {
	return GetBookingsRequest{
		client:      c,
		queryParams: c.NewGetBookingsQueryParams(),
		pathParams:  c.NewGetBookingsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetBookingsRequestBody(),
	}
}

type GetBookingsRequest struct {
	client      *Client
	queryParams *GetBookingsQueryParams
	pathParams  *GetBookingsPathParams
	method      string
	headers     http.Header
	requestBody GetBookingsRequestBody
}

func (c *Client) NewGetBookingsQueryParams() *GetBookingsQueryParams {
	return &GetBookingsQueryParams{}
}

type GetBookingsQueryParams struct {
	ReservationID string   `schema:"reservationId,omitempty"`
	GroupID       string   `schema:"groupId,omitempty"`
	ChannelCode   []string `schema:"channelCode,omitempty"`
	ExternalCode  string   `schema:"externalCode,omitempty"`
	TextSearch    string   `schema:"textSearch,omitempty"`
	PageNumber    int      `schema:"pageNumber,omitempty"`
	PageSize      int      `schema:"pageSize,omitempty"`
	Expand        []string `schema:"expand,omitempty"`
	PropertyID    string   `schema:"propertyId,omitempty"`
	From          DateTime `schema:"from,omitempty"`
	To            DateTime `schema:"to,omitempty"`
}

func (p GetBookingsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetBookingsRequest) QueryParams() *GetBookingsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetBookingsPathParams() *GetBookingsPathParams {
	return &GetBookingsPathParams{}
}

type GetBookingsPathParams struct {
}

func (p *GetBookingsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetBookingsRequest) PathParams() *GetBookingsPathParams {
	return r.pathParams
}

func (r *GetBookingsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetBookingsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetBookingsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetBookingsRequestBody() GetBookingsRequestBody {
	return GetBookingsRequestBody{}
}

type GetBookingsRequestBody struct {
}

func (r *GetBookingsRequest) RequestBody() *GetBookingsRequestBody {
	return nil
}

func (r *GetBookingsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetBookingsRequest) SetRequestBody(body GetBookingsRequestBody) {
	r.requestBody = body
}

func (r *GetBookingsRequest) NewResponseBody() *GetBookingsResponseBody {
	return &GetBookingsResponseBody{}
}

type GetBookingsResponseBody struct {
	Count    int                `json:"count"`
	Bookings []BookingItemModel `json:"bookings"`
}

func (r *GetBookingsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("booking/v1/bookings", r.PathParams())
	return &u
}

func (r *GetBookingsRequest) Do(ctx context.Context) (GetBookingsResponseBody, error) {
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

func (r *GetBookingsRequest) All(ctx context.Context) ([]BookingItemModel, error) {
	bookings := []BookingItemModel{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return bookings, err
		}

		// Break out of loop when no bookings are found
		if len(resp.Bookings) == 0 {
			break
		}

		// Add bookings to list
		bookings = append(bookings, resp.Bookings...)

		if len(bookings) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return bookings, nil
}
