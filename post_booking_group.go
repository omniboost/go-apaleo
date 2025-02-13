package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostBookingGroupRequest() PostBookingGroupRequest {
	return PostBookingGroupRequest{
		client:      c,
		queryParams: c.NewPostBookingGroupQueryParams(),
		pathParams:  c.NewPostBookingGroupPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostBookingGroupRequestBody(),
	}
}

type PostBookingGroupRequest struct {
	client      *Client
	queryParams *PostBookingGroupQueryParams
	pathParams  *PostBookingGroupPathParams
	method      string
	headers     http.Header
	requestBody PostBookingGroupRequestBody
}

func (c *Client) NewPostBookingGroupQueryParams() *PostBookingGroupQueryParams {
	return &PostBookingGroupQueryParams{}
}

type PostBookingGroupQueryParams struct{}

func (p PostBookingGroupQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostBookingGroupRequest) QueryParams() *PostBookingGroupQueryParams {
	return r.queryParams
}

func (c *Client) NewPostBookingGroupPathParams() *PostBookingGroupPathParams {
	return &PostBookingGroupPathParams{}
}

type PostBookingGroupPathParams struct {
}

func (p *PostBookingGroupPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostBookingGroupRequest) PathParams() *PostBookingGroupPathParams {
	return r.pathParams
}

func (r *PostBookingGroupRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostBookingGroupRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostBookingGroupRequest) Method() string {
	return r.method
}

func (s *Client) NewPostBookingGroupRequestBody() PostBookingGroupRequestBody {
	return PostBookingGroupRequestBody{}
}

type PostBookingGroupRequestBody struct {
	Name           string                    `json:"name"`
	Booker         BookerModel               `json:"booker"`
	Comment        string                    `json:"comment,omitempty"`
	BookerComment  string                    `json:"bookerComment,omitempty"`
	PaymentAccount CreatePaymentAccountModel `json:"paymentAccount,omitempty"`
	PropertyIDs    []string                  `json:"propertyIds"`
}

func (r PostBookingGroupRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PostBookingGroupRequest) RequestBody() *PostBookingGroupRequestBody {
	return &r.requestBody
}

func (r *PostBookingGroupRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostBookingGroupRequest) SetRequestBody(body PostBookingGroupRequestBody) {
	r.requestBody = body
}

func (r *PostBookingGroupRequest) NewResponseBody() *PostBookingGroupResponseBody {
	return &PostBookingGroupResponseBody{}
}

type PostBookingGroupResponseBody struct {
	ID string `json:"id"`
}

func (r *PostBookingGroupRequest) URL() *url.URL {
	path := "booking/v1/groups"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PostBookingGroupRequest) Do() (PostBookingGroupResponseBody, error) {
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
