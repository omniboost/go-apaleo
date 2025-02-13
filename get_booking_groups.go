package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetBookingGroupsRequest() GetBookingGroupsRequest {
	return GetBookingGroupsRequest{
		client:      c,
		queryParams: c.NewGetBookingGroupsQueryParams(),
		pathParams:  c.NewGetBookingGroupsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetBookingGroupsRequestBody(),
	}
}

type GetBookingGroupsRequest struct {
	client      *Client
	queryParams *GetBookingGroupsQueryParams
	pathParams  *GetBookingGroupsPathParams
	method      string
	headers     http.Header
	requestBody GetBookingGroupsRequestBody
}

func (c *Client) NewGetBookingGroupsQueryParams() *GetBookingGroupsQueryParams {
	return &GetBookingGroupsQueryParams{}
}

type GetBookingGroupsQueryParams struct {
	TextSearch  string   `json:"textSearch,omitempty"`
	PropertyIDs string   `json:"propertyIds,omitempty"`
	From        DateTime `json:"from,omitempty"`
	To          DateTime `json:"to,omitempty"`
	PageNumber  int      `json:"pageNumber,omitempty"`
	PageSize    int      `json:"pageSize,omitempty"`
	Expand      []string `json:"expand,omitempty"`
}

func (p GetBookingGroupsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetBookingGroupsRequest) QueryParams() *GetBookingGroupsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetBookingGroupsPathParams() *GetBookingGroupsPathParams {
	return &GetBookingGroupsPathParams{}
}

type GetBookingGroupsPathParams struct {
}

func (p *GetBookingGroupsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetBookingGroupsRequest) PathParams() *GetBookingGroupsPathParams {
	return r.pathParams
}

func (r *GetBookingGroupsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetBookingGroupsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetBookingGroupsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetBookingGroupsRequestBody() GetBookingGroupsRequestBody {
	return GetBookingGroupsRequestBody{}
}

type GetBookingGroupsRequestBody struct {
}

func (r *GetBookingGroupsRequest) RequestBody() *GetBookingGroupsRequestBody {
	return nil
}

func (r *GetBookingGroupsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetBookingGroupsRequest) SetRequestBody(body GetBookingGroupsRequestBody) {
	r.requestBody = body
}

func (r *GetBookingGroupsRequest) NewResponseBody() *GetBookingGroupsResponseBody {
	return &GetBookingGroupsResponseBody{}
}

type GetBookingGroupsResponseBody struct {
	Count  int              `json:"count"`
	Groups []GroupItemModel `json:"groups"`
}

func (r *GetBookingGroupsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("booking/v1/groups", r.PathParams())
	return &u
}

func (r *GetBookingGroupsRequest) Do() (GetBookingGroupsResponseBody, error) {
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

func (r *GetBookingGroupsRequest) All() ([]GroupItemModel, error) {
	groups := []GroupItemModel{}
	for {
		resp, err := r.Do()
		if err != nil {
			return groups, err
		}

		// Break out of loop when no groups are found
		if len(resp.Groups) == 0 {
			break
		}

		// Add groups to list
		groups = append(groups, resp.Groups...)

		if len(groups) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return groups, nil
}
