package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetBookingBlocksRequest() GetBookingBlocksRequest {
	return GetBookingBlocksRequest{
		client:      c,
		queryParams: c.NewGetBookingBlocksQueryParams(),
		pathParams:  c.NewGetBookingBlocksPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetBookingBlocksRequestBody(),
	}
}

type GetBookingBlocksRequest struct {
	client      *Client
	queryParams *GetBookingBlocksQueryParams
	pathParams  *GetBookingBlocksPathParams
	method      string
	headers     http.Header
	requestBody GetBookingBlocksRequestBody
}

func (c *Client) NewGetBookingBlocksQueryParams() *GetBookingBlocksQueryParams {
	return &GetBookingBlocksQueryParams{}
}

type GetBookingBlocksQueryParams struct {
	GroupID                string   `json:"groupId,omitempty"`
	PropertyIDs            string   `json:"propertyIds,omitempty"`
	Status                 string   `json:"status,omitempty"`
	UnitGroupIDs           []string `json:"unitGroupIds,omitempty"`
	RatePlanIDs            []string `json:"ratePlanIds,omitempty"`
	TimeSliceDefinitionIDs []string `json:"timeSliceDefinitionIds,omitempty"`
	UnitGroupTypes         []string `json:"unitGroupTypes,omitempty"`
	TimeSliceTemplate      string   `json:"timeSliceTemplate,omitempty"`
	From                   Date     `json:"from,omitempty"`
	To                     Date     `json:"to,omitempty"`
	PageNumber             int      `json:"pageNumber,omitempty"`
	PageSize               int      `json:"pageSize,omitempty"`
	Expand                 []string `json:"expand,omitempty"`
}

func (p GetBookingBlocksQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetBookingBlocksRequest) QueryParams() *GetBookingBlocksQueryParams {
	return r.queryParams
}

func (c *Client) NewGetBookingBlocksPathParams() *GetBookingBlocksPathParams {
	return &GetBookingBlocksPathParams{}
}

type GetBookingBlocksPathParams struct {
}

func (p *GetBookingBlocksPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetBookingBlocksRequest) PathParams() *GetBookingBlocksPathParams {
	return r.pathParams
}

func (r *GetBookingBlocksRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetBookingBlocksRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetBookingBlocksRequest) Method() string {
	return r.method
}

func (s *Client) NewGetBookingBlocksRequestBody() GetBookingBlocksRequestBody {
	return GetBookingBlocksRequestBody{}
}

type GetBookingBlocksRequestBody struct {
}

func (r *GetBookingBlocksRequest) RequestBody() *GetBookingBlocksRequestBody {
	return nil
}

func (r *GetBookingBlocksRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetBookingBlocksRequest) SetRequestBody(body GetBookingBlocksRequestBody) {
	r.requestBody = body
}

func (r *GetBookingBlocksRequest) NewResponseBody() *GetBookingBlocksResponseBody {
	return &GetBookingBlocksResponseBody{}
}

type GetBookingBlocksResponseBody struct {
	Count  int              `json:"count"`
	Blocks []BlockItemModel `json:"groups"`
}

func (r *GetBookingBlocksRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("booking/v1/blocks", r.PathParams())
	return &u
}

func (r *GetBookingBlocksRequest) Do() (GetBookingBlocksResponseBody, error) {
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

func (r *GetBookingBlocksRequest) All() ([]BlockItemModel, error) {
	blocks := []BlockItemModel{}
	for {
		resp, err := r.Do()
		if err != nil {
			return blocks, err
		}

		// Break out of loop when no blocks are found
		if len(resp.Blocks) == 0 {
			break
		}

		// Add blocks to list
		blocks = append(blocks, resp.Blocks...)

		if len(blocks) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return blocks, nil
}
