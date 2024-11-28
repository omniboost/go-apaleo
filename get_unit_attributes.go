package apaleo

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetUnitAttributesRequest() GetUnitAttributesRequest {
	return GetUnitAttributesRequest{
		client:      c,
		queryParams: c.NewGetUnitAttributesQueryParams(),
		pathParams:  c.NewGetUnitAttributesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetUnitAttributesRequestBody(),
	}
}

type GetUnitAttributesRequest struct {
	client      *Client
	queryParams *GetUnitAttributesQueryParams
	pathParams  *GetUnitAttributesPathParams
	method      string
	headers     http.Header
	requestBody GetUnitAttributesRequestBody
}

func (c *Client) NewGetUnitAttributesQueryParams() *GetUnitAttributesQueryParams {
	return &GetUnitAttributesQueryParams{}
}

type GetUnitAttributesQueryParams struct {
	PageNumber int `schema:"pageNumber,omitempty"`
	PageSize   int `schema:"pageSize,omitempty"`
}

func (p GetUnitAttributesQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetUnitAttributesRequest) QueryParams() *GetUnitAttributesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetUnitAttributesPathParams() *GetUnitAttributesPathParams {
	return &GetUnitAttributesPathParams{}
}

type GetUnitAttributesPathParams struct {
}

func (p *GetUnitAttributesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetUnitAttributesRequest) PathParams() *GetUnitAttributesPathParams {
	return r.pathParams
}

func (r *GetUnitAttributesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetUnitAttributesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetUnitAttributesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetUnitAttributesRequestBody() GetUnitAttributesRequestBody {
	return GetUnitAttributesRequestBody{}
}

type GetUnitAttributesRequestBody struct {
}

func (r *GetUnitAttributesRequest) RequestBody() *GetUnitAttributesRequestBody {
	return nil
}

func (r *GetUnitAttributesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetUnitAttributesRequest) SetRequestBody(body GetUnitAttributesRequestBody) {
	r.requestBody = body
}

func (r *GetUnitAttributesRequest) NewResponseBody() *GetUnitAttributesResponseBody {
	return &GetUnitAttributesResponseBody{}
}

type GetUnitAttributesResponseBody struct {
	Count          int            `json:"count"`
	UnitAttributes UnitAttributes `json:"unitAttributes"`
}

func (r *GetUnitAttributesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("inventory/v1/unit-attributes", r.PathParams())
	return &u
}

func (r *GetUnitAttributesRequest) Do() (GetUnitAttributesResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *GetUnitAttributesRequest) All() (UnitAttributes, error) {
	unitAttributes := UnitAttributes{}
	for {
		resp, err := r.Do()
		if err != nil {
			return unitAttributes, err
		}

		// Break out of loop when no unitAttributes are found
		if len(resp.UnitAttributes) == 0 {
			break
		}

		// Add unitAttributes to list
		unitAttributes = append(unitAttributes, resp.UnitAttributes...)

		if len(unitAttributes) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return unitAttributes, nil
}
