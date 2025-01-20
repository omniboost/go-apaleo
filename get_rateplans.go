package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetRatePlansRequest() GetRatePlansRequest {
	return GetRatePlansRequest{
		client:      c,
		queryParams: c.NewGetRatePlansQueryParams(),
		pathParams:  c.NewGetRatePlansPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetRatePlansRequestBody(),
	}
}

type GetRatePlansRequest struct {
	client      *Client
	queryParams *GetRatePlansQueryParams
	pathParams  *GetRatePlansPathParams
	method      string
	headers     http.Header
	requestBody GetRatePlansRequestBody
}

func (c *Client) NewGetRatePlansQueryParams() *GetRatePlansQueryParams {
	return &GetRatePlansQueryParams{}
}

type GetRatePlansQueryParams struct {
	PropertyID             string   `schema:"propertyId,omitempty"`
	RatePlanCodes          []string `schema:"ratePlanCodes,omitempty"`
	IncludedServiceIDs     []string `schema:"includedServiceIds,omitempty"`
	ChannelCodes           []string `schema:"channelCodes,omitempty"`
	PromoCodes             []string `schema:"promoCodes,omitempty"`
	CompanyIDs             []string `schema:"companyIds,omitempty"`
	BaseRatePlanIDs        []string `schema:"baseRatePlanIds,omitempty"`
	UnitGroupIDs           []string `schema:"unitGroupIds,omitempty"`
	TimeSliceDefinitionIDs []string `schema:"timeSliceDefinitionIds,omitempty"`
	UnitGroupTypes         []string `schema:"unitGroupTypes,omitempty"`
	TimeSliceTemplate      string   `schema:"timeSliceTemplate,omitempty"`
	MinGuaranteeTypes      []string `schema:"minGuaranteeTypes,omitempty"`
	CancellationPolicyIDs  []string `schema:"cancellationPolicyIds,omitempty"`
	NoShowPolicyIDs        []string `schema:"noShowPolicyIds,omitempty"`
	IsDerived              bool     `schema:"isDerived,omitempty"`
	DerivationLevelFilter  []string `schema:"derivationLevelFilter,omitempty"`
	PageNumber             int      `schema:"pageNumber,omitempty"`
	PageSize               int      `schema:"pageSize,omitempty"`
	Expand                 []string `schema:"expand,omitempty"`
}

func (p GetRatePlansQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetRatePlansRequest) QueryParams() *GetRatePlansQueryParams {
	return r.queryParams
}

func (c *Client) NewGetRatePlansPathParams() *GetRatePlansPathParams {
	return &GetRatePlansPathParams{}
}

type GetRatePlansPathParams struct {
}

func (p *GetRatePlansPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetRatePlansRequest) PathParams() *GetRatePlansPathParams {
	return r.pathParams
}

func (r *GetRatePlansRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetRatePlansRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetRatePlansRequest) Method() string {
	return r.method
}

func (s *Client) NewGetRatePlansRequestBody() GetRatePlansRequestBody {
	return GetRatePlansRequestBody{}
}

type GetRatePlansRequestBody struct {
}

func (r *GetRatePlansRequest) RequestBody() *GetRatePlansRequestBody {
	return nil
}

func (r *GetRatePlansRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetRatePlansRequest) SetRequestBody(body GetRatePlansRequestBody) {
	r.requestBody = body
}

func (r *GetRatePlansRequest) NewResponseBody() *GetRatePlansResponseBody {
	return &GetRatePlansResponseBody{}
}

type GetRatePlansResponseBody struct {
	Count     int       `json:"count"`
	RatePlans RatePlans `json:"ratePlans"`
}

func (r *GetRatePlansRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("rateplan/v1/rate-plans", r.PathParams())
	return &u
}

func (r *GetRatePlansRequest) Do() (GetRatePlansResponseBody, error) {
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

func (r *GetRatePlansRequest) All() (RatePlans, error) {
	rateplans := RatePlans{}
	for {
		resp, err := r.Do()
		if err != nil {
			return rateplans, err
		}

		// Break out of loop when no rateplans are found
		if len(resp.RatePlans) == 0 {
			break
		}

		// Add rateplans to list
		rateplans = append(rateplans, resp.RatePlans...)

		if len(rateplans) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return rateplans, nil
}
