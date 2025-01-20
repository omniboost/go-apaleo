package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetCancellationPoliciesRequest() GetCancellationPoliciesRequest {
	return GetCancellationPoliciesRequest{
		client:      c,
		queryParams: c.NewGetCancellationPoliciesQueryParams(),
		pathParams:  c.NewGetCancellationPoliciesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetCancellationPoliciesRequestBody(),
	}
}

type GetCancellationPoliciesRequest struct {
	client      *Client
	queryParams *GetCancellationPoliciesQueryParams
	pathParams  *GetCancellationPoliciesPathParams
	method      string
	headers     http.Header
	requestBody GetCancellationPoliciesRequestBody
}

func (c *Client) NewGetCancellationPoliciesQueryParams() *GetCancellationPoliciesQueryParams {
	return &GetCancellationPoliciesQueryParams{}
}

type GetCancellationPoliciesQueryParams struct {
	PropertyID string `schema:"propertyId,omitempty"`
	PageNumber int    `schema:"pageNumber,omitempty"`
	PageSize   int    `schema:"pageSize,omitempty"`
}

func (p GetCancellationPoliciesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetCancellationPoliciesRequest) QueryParams() *GetCancellationPoliciesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetCancellationPoliciesPathParams() *GetCancellationPoliciesPathParams {
	return &GetCancellationPoliciesPathParams{}
}

type GetCancellationPoliciesPathParams struct {
}

func (p *GetCancellationPoliciesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetCancellationPoliciesRequest) PathParams() *GetCancellationPoliciesPathParams {
	return r.pathParams
}

func (r *GetCancellationPoliciesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCancellationPoliciesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCancellationPoliciesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetCancellationPoliciesRequestBody() GetCancellationPoliciesRequestBody {
	return GetCancellationPoliciesRequestBody{}
}

type GetCancellationPoliciesRequestBody struct {
}

func (r *GetCancellationPoliciesRequest) RequestBody() *GetCancellationPoliciesRequestBody {
	return nil
}

func (r *GetCancellationPoliciesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetCancellationPoliciesRequest) SetRequestBody(body GetCancellationPoliciesRequestBody) {
	r.requestBody = body
}

func (r *GetCancellationPoliciesRequest) NewResponseBody() *GetCancellationPoliciesResponseBody {
	return &GetCancellationPoliciesResponseBody{}
}

type GetCancellationPoliciesResponseBody struct {
	Count                int                  `json:"count"`
	CancellationPolicies CancellationPolicies `json:"cancellationPolicies"`
}

func (r *GetCancellationPoliciesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rateplan/v1/cancellation-policies", r.PathParams())
	return &u
}

func (r *GetCancellationPoliciesRequest) Do() (GetCancellationPoliciesResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *GetCancellationPoliciesRequest) All() (CancellationPolicies, error) {
	cancellationPolicies := CancellationPolicies{}
	for {
		resp, err := r.Do()
		if err != nil {
			return cancellationPolicies, err
		}

		// Break out of loop when no cancellationPolicies are found
		if len(resp.CancellationPolicies) == 0 {
			break
		}

		// Add cancellationPolicies to list
		cancellationPolicies = append(cancellationPolicies, resp.CancellationPolicies...)

		if len(cancellationPolicies) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return cancellationPolicies, nil
}
