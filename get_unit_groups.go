package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetUnitGroupsRequest() GetUnitGroupsRequest {
	return GetUnitGroupsRequest{
		client:      c,
		queryParams: c.NewGetUnitGroupsQueryParams(),
		pathParams:  c.NewGetUnitGroupsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetUnitGroupsRequestBody(),
	}
}

type GetUnitGroupsRequest struct {
	client      *Client
	queryParams *GetUnitGroupsQueryParams
	pathParams  *GetUnitGroupsPathParams
	method      string
	headers     http.Header
	requestBody GetUnitGroupsRequestBody
}

func (c *Client) NewGetUnitGroupsQueryParams() *GetUnitGroupsQueryParams {
	return &GetUnitGroupsQueryParams{}
}

type GetUnitGroupsQueryParams struct {
	PropertyID     string   `schema:"propertyId,omitempty"`
	UnitGroupTypes []string `schema:"unitGroupTypes,omitempty"`
	PageNumber     int      `schema:"pageNumber,omitempty"`
	PageSize       int      `schema:"pageSize,omitempty"`
	Expand         []string `schema:"expand,omitempty"`
}

func (p GetUnitGroupsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetUnitGroupsRequest) QueryParams() *GetUnitGroupsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetUnitGroupsPathParams() *GetUnitGroupsPathParams {
	return &GetUnitGroupsPathParams{}
}

type GetUnitGroupsPathParams struct {
}

func (p *GetUnitGroupsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetUnitGroupsRequest) PathParams() *GetUnitGroupsPathParams {
	return r.pathParams
}

func (r *GetUnitGroupsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetUnitGroupsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetUnitGroupsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetUnitGroupsRequestBody() GetUnitGroupsRequestBody {
	return GetUnitGroupsRequestBody{}
}

type GetUnitGroupsRequestBody struct {
}

func (r *GetUnitGroupsRequest) RequestBody() *GetUnitGroupsRequestBody {
	return nil
}

func (r *GetUnitGroupsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetUnitGroupsRequest) SetRequestBody(body GetUnitGroupsRequestBody) {
	r.requestBody = body
}

func (r *GetUnitGroupsRequest) NewResponseBody() *GetUnitGroupsResponseBody {
	return &GetUnitGroupsResponseBody{}
}

type GetUnitGroupsResponseBody struct {
	Count      int        `json:"count"`
	UnitGroups UnitGroups `json:"unitGroups"`
}

func (r *GetUnitGroupsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("inventory/v1/unit-groups", r.PathParams())
	return &u
}

func (r *GetUnitGroupsRequest) Do() (GetUnitGroupsResponseBody, error) {
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

func (r *GetUnitGroupsRequest) All() (UnitGroups, error) {
	unitGroups := UnitGroups{}
	for {
		resp, err := r.Do()
		if err != nil {
			return unitGroups, err
		}

		// Break out of loop when no unitGroups are found
		if len(resp.UnitGroups) == 0 {
			break
		}

		// Add unitGroups to list
		unitGroups = append(unitGroups, resp.UnitGroups...)

		if len(unitGroups) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return unitGroups, nil
}
