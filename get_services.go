package apaleo

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetServicesRequest() GetServicesRequest {
	return GetServicesRequest{
		client:      c,
		queryParams: c.NewGetServicesQueryParams(),
		pathParams:  c.NewGetServicesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetServicesRequestBody(),
	}
}

type GetServicesRequest struct {
	client      *Client
	queryParams *GetServicesQueryParams
	pathParams  *GetServicesPathParams
	method      string
	headers     http.Header
	requestBody GetServicesRequestBody
}

func (c *Client) NewGetServicesQueryParams() *GetServicesQueryParams {
	return &GetServicesQueryParams{}
}

type GetServicesQueryParams struct {
	PropertyID       string   `schema:"propertyId,omitempty"`
	TextSearch       string   `schema:"textSearch,omitempty"`
	OnlySoldAsExtras bool     `schema:"onlySoldAsExtras,omitempty"`
	ChannelCodes     []string `schema:"channelCodes,omitempty"`
	ServiceTypes     []string `schema:"serviceTypes,omitempty"`
	PageNumber       int      `schema:"pageNumber,omitempty"`
	PageSize         int      `schema:"pageSize,omitempty"`
	Expand           []string `schema:"expand,omitempty"`
}

func (p GetServicesQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetServicesRequest) QueryParams() *GetServicesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetServicesPathParams() *GetServicesPathParams {
	return &GetServicesPathParams{}
}

type GetServicesPathParams struct {
}

func (p *GetServicesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetServicesRequest) PathParams() *GetServicesPathParams {
	return r.pathParams
}

func (r *GetServicesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetServicesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetServicesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetServicesRequestBody() GetServicesRequestBody {
	return GetServicesRequestBody{}
}

type GetServicesRequestBody struct {
}

func (r *GetServicesRequest) RequestBody() *GetServicesRequestBody {
	return nil
}

func (r *GetServicesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetServicesRequest) SetRequestBody(body GetServicesRequestBody) {
	r.requestBody = body
}

func (r *GetServicesRequest) NewResponseBody() *GetServicesResponseBody {
	return &GetServicesResponseBody{}
}

type GetServicesResponseBody struct {
	Count    int      `json:"count"`
	Services Services `json:"services"`
}

func (r *GetServicesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rateplan/v1/services", r.PathParams())
	return &u
}

func (r *GetServicesRequest) Do() (GetServicesResponseBody, error) {
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

func (r *GetServicesRequest) All() (Services, error) {
	services := Services{}
	for {
		resp, err := r.Do()
		if err != nil {
			return services, err
		}

		// Break out of loop when no services are found
		if len(resp.Services) == 0 {
			break
		}

		// Add services to list
		services = append(services, resp.Services...)

		if len(services) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return services, nil
}
