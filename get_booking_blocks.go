package apaleo

import (
	"context"
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
	GroupID                string   `schema:"groupId,omitempty"`
	PropertyIDs            string   `schema:"propertyIds,omitempty"`
	Status                 string   `schema:"status,omitempty"`
	UnitGroupIDs           []string `schema:"unitGroupIds,omitempty"`
	RatePlanIDs            []string `schema:"ratePlanIds,omitempty"`
	TimeSliceDefinitionIDs []string `schema:"timeSliceDefinitionIds,omitempty"`
	UnitGroupTypes         []string `schema:"unitGroupTypes,omitempty"`
	TimeSliceTemplate      string   `schema:"timeSliceTemplate,omitempty"`
	From                   Date     `schema:"from,omitempty"`
	To                     Date     `schema:"to,omitempty"`
	PageNumber             int      `schema:"pageNumber,omitempty"`
	PageSize               int      `schema:"pageSize,omitempty"`
	Expand                 []string `schema:"expand,omitempty"`
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

func (r *GetBookingBlocksRequest) Do(ctx context.Context) (GetBookingBlocksResponseBody, error) {
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

func (r *GetBookingBlocksRequest) All(ctx context.Context) ([]BlockItemModel, error) {
	blocks := []BlockItemModel{}
	for {
		resp, err := r.Do(ctx)
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
