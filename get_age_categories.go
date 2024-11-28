package apaleo

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetAgeCategoriesRequest() GetAgeCategoriesRequest {
	return GetAgeCategoriesRequest{
		client:      c,
		queryParams: c.NewGetAgeCategoriesQueryParams(),
		pathParams:  c.NewGetAgeCategoriesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetAgeCategoriesRequestBody(),
	}
}

type GetAgeCategoriesRequest struct {
	client      *Client
	queryParams *GetAgeCategoriesQueryParams
	pathParams  *GetAgeCategoriesPathParams
	method      string
	headers     http.Header
	requestBody GetAgeCategoriesRequestBody
}

func (c *Client) NewGetAgeCategoriesQueryParams() *GetAgeCategoriesQueryParams {
	return &GetAgeCategoriesQueryParams{}
}

type GetAgeCategoriesQueryParams struct {
	PropertyID string `schema:"propertyId,omitempty"`
	PageNumber int    `schema:"pageNumber,omitempty"`
	PageSize   int    `schema:"pageSize,omitempty"`
}

func (p GetAgeCategoriesQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetAgeCategoriesRequest) QueryParams() *GetAgeCategoriesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetAgeCategoriesPathParams() *GetAgeCategoriesPathParams {
	return &GetAgeCategoriesPathParams{}
}

type GetAgeCategoriesPathParams struct {
}

func (p *GetAgeCategoriesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetAgeCategoriesRequest) PathParams() *GetAgeCategoriesPathParams {
	return r.pathParams
}

func (r *GetAgeCategoriesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetAgeCategoriesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetAgeCategoriesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetAgeCategoriesRequestBody() GetAgeCategoriesRequestBody {
	return GetAgeCategoriesRequestBody{}
}

type GetAgeCategoriesRequestBody struct {
}

func (r *GetAgeCategoriesRequest) RequestBody() *GetAgeCategoriesRequestBody {
	return nil
}

func (r *GetAgeCategoriesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetAgeCategoriesRequest) SetRequestBody(body GetAgeCategoriesRequestBody) {
	r.requestBody = body
}

func (r *GetAgeCategoriesRequest) NewResponseBody() *GetAgeCategoriesResponseBody {
	return &GetAgeCategoriesResponseBody{}
}

type GetAgeCategoriesResponseBody struct {
	Count         int           `json:"count"`
	AgeCategories AgeCategories `json:"ageCategories"`
}

func (r *GetAgeCategoriesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("settings/v1/age-categories", r.PathParams())
	return &u
}

func (r *GetAgeCategoriesRequest) Do() (GetAgeCategoriesResponseBody, error) {
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

func (r *GetAgeCategoriesRequest) All() (AgeCategories, error) {
	ageCategories := AgeCategories{}
	for {
		resp, err := r.Do()
		if err != nil {
			return ageCategories, err
		}

		// Break out of loop when no ageCategories are found
		if len(resp.AgeCategories) == 0 {
			break
		}

		// Add ageCategories to list
		ageCategories = append(ageCategories, resp.AgeCategories...)

		if len(ageCategories) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return ageCategories, nil
}
