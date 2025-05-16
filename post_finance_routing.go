package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostFinanceRoutingRequest() PostFinanceRoutingRequest {
	return PostFinanceRoutingRequest{
		client:      c,
		queryParams: c.NewPostFinanceRoutingQueryParams(),
		pathParams:  c.NewPostFinanceRoutingPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostFinanceRoutingRequestBody(),
	}
}

type PostFinanceRoutingRequest struct {
	client      *Client
	queryParams *PostFinanceRoutingQueryParams
	pathParams  *PostFinanceRoutingPathParams
	method      string
	headers     http.Header
	requestBody PostFinanceRoutingRequestBody
}

func (c *Client) NewPostFinanceRoutingQueryParams() *PostFinanceRoutingQueryParams {
	return &PostFinanceRoutingQueryParams{}
}

type PostFinanceRoutingQueryParams struct{}

func (p PostFinanceRoutingQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostFinanceRoutingRequest) QueryParams() *PostFinanceRoutingQueryParams {
	return r.queryParams
}

func (c *Client) NewPostFinanceRoutingPathParams() *PostFinanceRoutingPathParams {
	return &PostFinanceRoutingPathParams{}
}

type PostFinanceRoutingPathParams struct {
}

func (p *PostFinanceRoutingPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PostFinanceRoutingRequest) PathParams() *PostFinanceRoutingPathParams {
	return r.pathParams
}

func (r *PostFinanceRoutingRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostFinanceRoutingRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostFinanceRoutingRequest) Method() string {
	return r.method
}

func (s *Client) NewPostFinanceRoutingRequestBody() PostFinanceRoutingRequestBody {
	return PostFinanceRoutingRequestBody{}
}

type PostFinanceRoutingRequestBody struct {
	BookingID  string `json:"bookingId"`
	PropertyID string `json:"propertyId"`
	Filter     struct {
		FolioIDs      []string `json:"folioIds,omitempty"`
		SubAccountIDs []string `json:"subAccountIds,omitempty"`
		ServiceTypes  []string `json:"serviceTypes,omitempty"`
		From          *Date    `json:"from,omitempty"`
		To            *Date    `json:"to,omitempty"`
	} `json:"filter,omitempty"`
	DestinationFolioID string `json:"destinationFolioId"`
}

func (r PostFinanceRoutingRequest) NewRequestBody() PostFinanceRoutingRequestBody {
	return PostFinanceRoutingRequestBody{}
}

func (r *PostFinanceRoutingRequest) RequestBody() *PostFinanceRoutingRequestBody {
	return &r.requestBody
}

func (r *PostFinanceRoutingRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *PostFinanceRoutingRequest) SetRequestBody(body PostFinanceRoutingRequestBody) {
	r.requestBody = body
}

func (r *PostFinanceRoutingRequest) NewResponseBody() *PostFinanceRoutingResponseBody {
	return &PostFinanceRoutingResponseBody{}
}

type PostFinanceRoutingResponseBody struct {
	ID string `json:"id"`
}

func (r *PostFinanceRoutingRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v0-nsfw/routings", r.PathParams())
	return &u
}

func (r *PostFinanceRoutingRequest) Do() (PostFinanceRoutingResponseBody, error) {
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
