package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/omitempty"
	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostGroupBookingRequest() PostGroupBookingRequest {
	return PostGroupBookingRequest{
		client:      c,
		queryParams: c.NewPostGroupBookingQueryParams(),
		pathParams:  c.NewPostGroupBookingPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostGroupBookingRequestBody(),
	}
}

type PostGroupBookingRequest struct {
	client      *Client
	queryParams *PostGroupBookingQueryParams
	pathParams  *PostGroupBookingPathParams
	method      string
	headers     http.Header
	requestBody PostGroupBookingRequestBody
}

func (c *Client) NewPostGroupBookingQueryParams() *PostGroupBookingQueryParams {
	return &PostGroupBookingQueryParams{}
}

type PostGroupBookingQueryParams struct{}

func (p PostGroupBookingQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PostGroupBookingRequest) QueryParams() *PostGroupBookingQueryParams {
	return r.queryParams
}

func (c *Client) NewPostGroupBookingPathParams() *PostGroupBookingPathParams {
	return &PostGroupBookingPathParams{}
}

type PostGroupBookingPathParams struct {
}

func (p *PostGroupBookingPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostGroupBookingRequest) PathParams() *PostGroupBookingPathParams {
	return r.pathParams
}

func (r *PostGroupBookingRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostGroupBookingRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostGroupBookingRequest) Method() string {
	return r.method
}

func (s *Client) NewPostGroupBookingRequestBody() PostGroupBookingRequestBody {
	return PostGroupBookingRequestBody{}
}

type PostGroupBookingRequestBody struct {
	Name           string                    `json:"name"`
	Booker         BookerModel               `json:"booker"`
	Comment        string                    `json:"comment,omitempty"`
	BookerComment  string                    `json:"bookerComment,omitempty"`
	PaymentAccount CreatePaymentAccountModel `json:"paymentAccount,omitempty"`
	PropertyIDs    []string                  `json:"propertyIds"`
}

func (r PostGroupBookingRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *PostGroupBookingRequest) RequestBody() *PostGroupBookingRequestBody {
	return &r.requestBody
}

func (r *PostGroupBookingRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *PostGroupBookingRequest) SetRequestBody(body PostGroupBookingRequestBody) {
	r.requestBody = body
}

func (r *PostGroupBookingRequest) NewResponseBody() *PostGroupBookingResponseBody {
	return &PostGroupBookingResponseBody{}
}

type PostGroupBookingResponseBody struct {
	ID string `json:"id"`
}

func (r *PostGroupBookingRequest) URL() *url.URL {
	path := "booking/v1/groups"
	u := r.client.GetEndpointURL(path, r.PathParams())
	return &u
}

func (r *PostGroupBookingRequest) Do() (PostGroupBookingResponseBody, error) {
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
