package apaleo

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetPropertiesRequest() GetPropertiesRequest {
	return GetPropertiesRequest{
		client:      c,
		queryParams: c.NewGetPropertiesQueryParams(),
		pathParams:  c.NewGetPropertiesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetPropertiesRequestBody(),
	}
}

type GetPropertiesRequest struct {
	client      *Client
	queryParams *GetPropertiesQueryParams
	pathParams  *GetPropertiesPathParams
	method      string
	headers     http.Header
	requestBody GetPropertiesRequestBody
}

func (c *Client) NewGetPropertiesQueryParams() *GetPropertiesQueryParams {
	return &GetPropertiesQueryParams{}
}

type GetPropertiesQueryParams struct {
	// Filter result by property status
	Status []string `schema:"status,omitempty"`

	// Include archived properties in the result. If not set, or set to false, it only returns non-archived properties
	IncludeArchived bool `schema:"includeArchived,omitempty"`

	// Filter result by country code
	CountryCode []string `schema:"countryCode,omitempty"`

	// Page number, 1-based. Default value is 1 (if this is not set or not positive). Results in 204 if there are no items on that page.
	PageNumber int `schema:"pageNumber,omitempty"`

	// Page size. If this is not set or not positive, the pageNumber is ignored and all items are returned.
	PageSize int `schema:"pageSize,omitempty"`

	// List of all embedded resources that should be expanded in the response. Possible values are: actions. All other values will be silently ignored.
	Expand []string `schema:"expand,omitempty"`
}

func (p GetPropertiesQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetPropertiesRequest) QueryParams() *GetPropertiesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetPropertiesPathParams() *GetPropertiesPathParams {
	return &GetPropertiesPathParams{}
}

type GetPropertiesPathParams struct {
}

func (p *GetPropertiesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetPropertiesRequest) PathParams() *GetPropertiesPathParams {
	return r.pathParams
}

func (r *GetPropertiesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetPropertiesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetPropertiesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetPropertiesRequestBody() GetPropertiesRequestBody {
	return GetPropertiesRequestBody{}
}

type GetPropertiesRequestBody struct {
}

func (r *GetPropertiesRequest) RequestBody() *GetPropertiesRequestBody {
	return &r.requestBody
}

func (r *GetPropertiesRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *GetPropertiesRequest) SetRequestBody(body GetPropertiesRequestBody) {
	r.requestBody = body
}

func (r *GetPropertiesRequest) NewResponseBody() *GetPropertiesResponseBody {
	return &GetPropertiesResponseBody{}
}

type GetPropertiesResponseBody struct {
	Count      int                 `json:"count"`
	Properties []PropertyListModel `json:"properties"`
}

func (r *GetPropertiesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("inventory/v1/properties", r.PathParams())
	return &u
}

func (r *GetPropertiesRequest) Do() (GetPropertiesResponseBody, error) {
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

func (r *GetPropertiesRequest) All() (PropertyList, error) {
	properties := PropertyList{}
	for {
		resp, err := r.Do()
		if err != nil {
			return properties, err
		}

		// Break out of loop when no properties are found
		if len(resp.Properties) == 0 {
			break
		}

		// Add properties to list
		properties = append(properties, resp.Properties...)

		if len(properties) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return properties, nil
}
