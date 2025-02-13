package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetBookingBlockByIDRequest() GetBookingBlockByIDRequest {
	return GetBookingBlockByIDRequest{
		client:      c,
		queryParams: c.NewGetBookingBlockByIDQueryParams(),
		pathParams:  c.NewGetBookingBlockByIDPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetBookingBlockByIDRequestBody(),
	}
}

type GetBookingBlockByIDRequest struct {
	client      *Client
	queryParams *GetBookingBlockByIDQueryParams
	pathParams  *GetBookingBlockByIDPathParams
	method      string
	headers     http.Header
	requestBody GetBookingBlockByIDRequestBody
}

func (c *Client) NewGetBookingBlockByIDQueryParams() *GetBookingBlockByIDQueryParams {
	return &GetBookingBlockByIDQueryParams{}
}

type GetBookingBlockByIDQueryParams struct{}

func (p GetBookingBlockByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetBookingBlockByIDRequest) QueryParams() *GetBookingBlockByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetBookingBlockByIDPathParams() *GetBookingBlockByIDPathParams {
	return &GetBookingBlockByIDPathParams{}
}

type GetBookingBlockByIDPathParams struct {
	BlockID string `schema:"id"`
}

func (p *GetBookingBlockByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.BlockID,
	}
}

func (r *GetBookingBlockByIDRequest) PathParams() *GetBookingBlockByIDPathParams {
	return r.pathParams
}

func (r *GetBookingBlockByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetBookingBlockByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetBookingBlockByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetBookingBlockByIDRequestBody() GetBookingBlockByIDRequestBody {
	return GetBookingBlockByIDRequestBody{}
}

type GetBookingBlockByIDRequestBody struct {
}

func (r *GetBookingBlockByIDRequest) RequestBody() *GetBookingBlockByIDRequestBody {
	return nil
}

func (r *GetBookingBlockByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetBookingBlockByIDRequest) SetRequestBody(body GetBookingBlockByIDRequestBody) {
	r.requestBody = body
}

func (r *GetBookingBlockByIDRequest) NewResponseBody() *GetBookingBlockByIDResponseBody {
	return &GetBookingBlockByIDResponseBody{}
}

type GetBookingBlockByIDResponseBody BlockItemModel

func (r *GetBookingBlockByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("booking/v1/blocks/{{.id}}", r.PathParams())
	return &u
}

func (r *GetBookingBlockByIDRequest) Do() (GetBookingBlockByIDResponseBody, error) {
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
