package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetNoShowPoliciesRequest() GetNoShowPoliciesRequest {
	return GetNoShowPoliciesRequest{
		client:      c,
		queryParams: c.NewGetNoShowPoliciesQueryParams(),
		pathParams:  c.NewGetNoShowPoliciesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetNoShowPoliciesRequestBody(),
	}
}

type GetNoShowPoliciesRequest struct {
	client      *Client
	queryParams *GetNoShowPoliciesQueryParams
	pathParams  *GetNoShowPoliciesPathParams
	method      string
	headers     http.Header
	requestBody GetNoShowPoliciesRequestBody
}

func (c *Client) NewGetNoShowPoliciesQueryParams() *GetNoShowPoliciesQueryParams {
	return &GetNoShowPoliciesQueryParams{}
}

type GetNoShowPoliciesQueryParams struct {
	PropertyID string `schema:"propertyId,omitempty"`
	PageNumber int    `schema:"pageNumber,omitempty"`
	PageSize   int    `schema:"pageSize,omitempty"`
}

func (p GetNoShowPoliciesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetNoShowPoliciesRequest) QueryParams() *GetNoShowPoliciesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetNoShowPoliciesPathParams() *GetNoShowPoliciesPathParams {
	return &GetNoShowPoliciesPathParams{}
}

type GetNoShowPoliciesPathParams struct {
}

func (p *GetNoShowPoliciesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetNoShowPoliciesRequest) PathParams() *GetNoShowPoliciesPathParams {
	return r.pathParams
}

func (r *GetNoShowPoliciesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetNoShowPoliciesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetNoShowPoliciesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetNoShowPoliciesRequestBody() GetNoShowPoliciesRequestBody {
	return GetNoShowPoliciesRequestBody{}
}

type GetNoShowPoliciesRequestBody struct {
}

func (r *GetNoShowPoliciesRequest) RequestBody() *GetNoShowPoliciesRequestBody {
	return nil
}

func (r *GetNoShowPoliciesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetNoShowPoliciesRequest) SetRequestBody(body GetNoShowPoliciesRequestBody) {
	r.requestBody = body
}

func (r *GetNoShowPoliciesRequest) NewResponseBody() *GetNoShowPoliciesResponseBody {
	return &GetNoShowPoliciesResponseBody{}
}

type GetNoShowPoliciesResponseBody struct {
	Count          int            `json:"count"`
	NoShowPolicies NoShowPolicies `json:"noShowPolicies"`
}

func (r *GetNoShowPoliciesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rateplan/v1/no-show-policies", r.PathParams())
	return &u
}

func (r *GetNoShowPoliciesRequest) Do() (GetNoShowPoliciesResponseBody, error) {
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

func (r *GetNoShowPoliciesRequest) All() (NoShowPolicies, error) {
	noShowPolicies := NoShowPolicies{}
	for {
		resp, err := r.Do()
		if err != nil {
			return noShowPolicies, err
		}

		// Break out of loop when no noShowPolicies are found
		if len(resp.NoShowPolicies) == 0 {
			break
		}

		// Add noShowPolicies to list
		noShowPolicies = append(noShowPolicies, resp.NoShowPolicies...)

		if len(noShowPolicies) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return noShowPolicies, nil
}
