package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostBookingGroupNSFWRequest() PostBookingGroupNSFWRequest {
	return PostBookingGroupNSFWRequest{
		client:      c,
		queryParams: c.NewPostBookingGroupNSFWQueryParams(),
		pathParams:  c.NewPostBookingGroupNSFWPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostBookingGroupNSFWRequestBody(),
	}
}

type PostBookingGroupNSFWRequest struct {
	client      *Client
	queryParams *PostBookingGroupNSFWQueryParams
	pathParams  *PostBookingGroupNSFWPathParams
	method      string
	headers     http.Header
	requestBody PostBookingGroupNSFWRequestBody
}

func (c *Client) NewPostBookingGroupNSFWQueryParams() *PostBookingGroupNSFWQueryParams {
	return &PostBookingGroupNSFWQueryParams{}
}

type PostBookingGroupNSFWQueryParams struct{}

func (p PostBookingGroupNSFWQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostBookingGroupNSFWRequest) QueryParams() *PostBookingGroupNSFWQueryParams {
	return r.queryParams
}

func (c *Client) NewPostBookingGroupNSFWPathParams() *PostBookingGroupNSFWPathParams {
	return &PostBookingGroupNSFWPathParams{}
}

type PostBookingGroupNSFWPathParams struct {
}

func (p *PostBookingGroupNSFWPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostBookingGroupNSFWRequest) PathParams() *PostBookingGroupNSFWPathParams {
	return r.pathParams
}

func (r *PostBookingGroupNSFWRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostBookingGroupNSFWRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostBookingGroupNSFWRequest) Method() string {
	return r.method
}

func (s *Client) NewPostBookingGroupNSFWRequestBody() PostBookingGroupNSFWRequestBody {
	return PostBookingGroupNSFWRequestBody{}
}

type PostBookingGroupNSFWRequestBody struct {
	Name           string                    `json:"name"`
	Booker         BookerModel               `json:"booker"`
	Comment        string                    `json:"comment,omitempty"`
	BookerComment  string                    `json:"bookerComment,omitempty"`
	PaymentAccount CreatePaymentAccountModel `json:"paymentAccount,omitempty"`
	PropertyIDs    []string                  `json:"propertyIds"`
}

func (r PostBookingGroupNSFWRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PostBookingGroupNSFWRequest) RequestBody() *PostBookingGroupNSFWRequestBody {
	return &r.requestBody
}

func (r *PostBookingGroupNSFWRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostBookingGroupNSFWRequest) SetRequestBody(body PostBookingGroupNSFWRequestBody) {
	r.requestBody = body
}

func (r *PostBookingGroupNSFWRequest) NewResponseBody() *PostBookingGroupNSFWResponseBody {
	return &PostBookingGroupNSFWResponseBody{}
}

type PostBookingGroupNSFWResponseBody struct {
	ID string `json:"id"`
}

func (r *PostBookingGroupNSFWRequest) URL() *url.URL {
	path := "booking/v0-nsfw/groups"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PostBookingGroupNSFWRequest) Do(ctx context.Context) (PostBookingGroupNSFWResponseBody, error) {
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
