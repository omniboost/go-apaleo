package apaleo

import (
	"net/http"
	"net/url"
)

func (c *Client) NewGetCompaniesRequest() GetCompaniesRequest {
	return GetCompaniesRequest{
		client:      c,
		queryParams: c.NewGetCompaniesQueryParams(),
		pathParams:  c.NewGetCompaniesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetCompaniesRequestBody(),
	}
}

type GetCompaniesRequest struct {
	client      *Client
	queryParams *GetCompaniesQueryParams
	pathParams  *GetCompaniesPathParams
	method      string
	headers     http.Header
	requestBody GetCompaniesRequestBody
}

func (c *Client) NewGetCompaniesQueryParams() *GetCompaniesQueryParams {
	return &GetCompaniesQueryParams{}
}

type GetCompaniesQueryParams struct {
	PropertyID       string   `schema:"propertyId,omitempty"`
	TextSearch       string   `schema:"textSearch,omitempty"`
	OnlySoldAsExtras bool     `schema:"onlySoldAsExtras,omitempty"`
	ChannelCodes     []string `schema:"channelCodes,omitempty"`
	ServiceTypes     []string `schema:"serviceTypes,omitempty"`
	PageNumber       int      `schema:"pageNumber,omitempty"`
	PageSize         int      `schema:"pageSize,omitempty"`
	Expand           []string `schema:"expand,omitempty"`
}

func (p GetCompaniesQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetCompaniesRequest) QueryParams() *GetCompaniesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCompaniesPathParams() *GetCompaniesPathParams {
	return &GetCompaniesPathParams{}
}

type GetCompaniesPathParams struct {
}

func (p *GetCompaniesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCompaniesRequest) PathParams() *GetCompaniesPathParams {
	return r.pathParams
}

func (r *GetCompaniesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCompaniesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCompaniesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCompaniesRequestBody() GetCompaniesRequestBody {
	return GetCompaniesRequestBody{}
}

type GetCompaniesRequestBody struct {
}

func (r *GetCompaniesRequest) RequestBody() *GetCompaniesRequestBody {
	return nil
}

func (r *GetCompaniesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetCompaniesRequest) SetRequestBody(body GetCompaniesRequestBody) {
	r.requestBody = body
}

func (r *GetCompaniesRequest) NewResponseBody() *GetCompaniesResponseBody {
	return &GetCompaniesResponseBody{}
}

type GetCompaniesResponseBody struct {
	Count     int       `json:"count"`
	Companies Companies `json:"companies"`
}

func (r *GetCompaniesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rateplan/v1/companies", r.PathParams())
	return &u
}

func (r *GetCompaniesRequest) Do() (GetCompaniesResponseBody, error) {
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

func (r *GetCompaniesRequest) All() (Companies, error) {
	companies := Companies{}
	for {
		resp, err := r.Do()
		if err != nil {
			return companies, err
		}

		// Break out of loop when no companies are found
		if len(resp.Companies) == 0 {
			break
		}

		// Add companies to list
		companies = append(companies, resp.Companies...)

		if len(companies) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return companies, nil
}
