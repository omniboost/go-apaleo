package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetBookingsNSFWRequest() GetBookingsNSFWRequest {
	return GetBookingsNSFWRequest{
		client:      c,
		queryParams: c.NewGetBookingsNSFWQueryParams(),
		pathParams:  c.NewGetBookingsNSFWPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetBookingsNSFWRequestBody(),
	}
}

type GetBookingsNSFWRequest struct {
	client      *Client
	queryParams *GetBookingsNSFWQueryParams
	pathParams  *GetBookingsNSFWPathParams
	method      string
	headers     http.Header
	requestBody GetBookingsNSFWRequestBody
}

func (c *Client) NewGetBookingsNSFWQueryParams() *GetBookingsNSFWQueryParams {
	return &GetBookingsNSFWQueryParams{}
}

type GetBookingsNSFWQueryParams struct {
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

func (p GetBookingsNSFWQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetBookingsNSFWRequest) QueryParams() *GetBookingsNSFWQueryParams {
	return r.queryParams
}

func (c *Client) NewGetBookingsNSFWPathParams() *GetBookingsNSFWPathParams {
	return &GetBookingsNSFWPathParams{}
}

type GetBookingsNSFWPathParams struct {
}

func (p *GetBookingsNSFWPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetBookingsNSFWRequest) PathParams() *GetBookingsNSFWPathParams {
	return r.pathParams
}

func (r *GetBookingsNSFWRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetBookingsNSFWRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetBookingsNSFWRequest) Method() string {
	return r.method
}

func (s *Client) NewGetBookingsNSFWRequestBody() GetBookingsNSFWRequestBody {
	return GetBookingsNSFWRequestBody{}
}

type GetBookingsNSFWRequestBody struct {
}

func (r *GetBookingsNSFWRequest) RequestBody() *GetBookingsNSFWRequestBody {
	return nil
}

func (r *GetBookingsNSFWRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetBookingsNSFWRequest) SetRequestBody(body GetBookingsNSFWRequestBody) {
	r.requestBody = body
}

func (r *GetBookingsNSFWRequest) NewResponseBody() *GetBookingsNSFWResponseBody {
	return &GetBookingsNSFWResponseBody{}
}

type GetBookingsNSFWResponseBody struct {
	Count    int                `json:"count"`
	Bookings []BookingItemModel `json:"bookings"`
}

func (r *GetBookingsNSFWRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("booking/v0-nsfw/bookings", r.PathParams())
	return &u
}

func (r *GetBookingsNSFWRequest) Do() (GetBookingsNSFWResponseBody, error) {
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

func (r *GetBookingsNSFWRequest) All() ([]BookingItemModel, error) {
	bookings := []BookingItemModel{}
	for {
		resp, err := r.Do()
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
