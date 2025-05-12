package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetMarketSegmentsNSFWRequest() GetMarketSegmentsNSFWRequest {
	return GetMarketSegmentsNSFWRequest{
		client:      c,
		queryParams: c.NewGetMarketSegmentsNSFWQueryParams(),
		pathParams:  c.NewGetMarketSegmentsNSFWPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetMarketSegmentsNSFWRequestBody(),
	}
}

type GetMarketSegmentsNSFWRequest struct {
	client      *Client
	queryParams *GetMarketSegmentsNSFWQueryParams
	pathParams  *GetMarketSegmentsNSFWPathParams
	method      string
	headers     http.Header
	requestBody GetMarketSegmentsNSFWRequestBody
}

func (c *Client) NewGetMarketSegmentsNSFWQueryParams() *GetMarketSegmentsNSFWQueryParams {
	return &GetMarketSegmentsNSFWQueryParams{}
}

type GetMarketSegmentsNSFWQueryParams struct {
	// Return market segments with any of the specified property ids
	PropertyIDs CommaSeparatedQueryParam `schema:"propertyIds,omitempty"`

	// Page number, 1-based. Default value is 1 (if this is not set or not positive). Results in 204 if there are no items on that page.
	PageNumber int `schema:"pageNumber,omitempty"`

	// Page size. If this is not set or not positive, the pageNumber is ignored and all items are returned.
	PageSize int `schema:"pageSize,omitempty"`
}

func (p GetMarketSegmentsNSFWQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetMarketSegmentsNSFWRequest) QueryParams() *GetMarketSegmentsNSFWQueryParams {
	return r.queryParams
}

func (c *Client) NewGetMarketSegmentsNSFWPathParams() *GetMarketSegmentsNSFWPathParams {
	return &GetMarketSegmentsNSFWPathParams{}
}

type GetMarketSegmentsNSFWPathParams struct {
}

func (p *GetMarketSegmentsNSFWPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetMarketSegmentsNSFWRequest) PathParams() *GetMarketSegmentsNSFWPathParams {
	return r.pathParams
}

func (r *GetMarketSegmentsNSFWRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetMarketSegmentsNSFWRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetMarketSegmentsNSFWRequest) Method() string {
	return r.method
}

func (s *Client) NewGetMarketSegmentsNSFWRequestBody() GetMarketSegmentsNSFWRequestBody {
	return GetMarketSegmentsNSFWRequestBody{}
}

type GetMarketSegmentsNSFWRequestBody struct {
}

func (r *GetMarketSegmentsNSFWRequest) RequestBody() *GetMarketSegmentsNSFWRequestBody {
	return nil
}

func (r *GetMarketSegmentsNSFWRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetMarketSegmentsNSFWRequest) SetRequestBody(body GetMarketSegmentsNSFWRequestBody) {
	r.requestBody = body
}

func (r *GetMarketSegmentsNSFWRequest) NewResponseBody() *GetMarketSegmentsNSFWResponseBody {
	return &GetMarketSegmentsNSFWResponseBody{}
}

type GetMarketSegmentsNSFWResponseBody struct {
	Count          int            `json:"count"`
	MarketSegments MarketSegments `json:"marketSegments"`
}

func (r *GetMarketSegmentsNSFWRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("settings/v0-nsfw/market-segments", r.PathParams())
	return &u
}

func (r *GetMarketSegmentsNSFWRequest) Do() (GetMarketSegmentsNSFWResponseBody, error) {
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

func (r *GetMarketSegmentsNSFWRequest) All() (MarketSegments, error) {
	values := MarketSegments{}
	for {
		resp, err := r.Do()
		if err != nil {
			return values, err
		}

		// Break out of loop when no market segments are found
		if len(resp.MarketSegments) == 0 {
			break
		}

		// Add market segments to list
		values = append(values, resp.MarketSegments...)

		if len(values) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return values, nil
}
