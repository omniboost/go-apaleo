package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetUnitsRequest() GetUnitsRequest {
	return GetUnitsRequest{
		client:      c,
		queryParams: c.NewGetUnitsQueryParams(),
		pathParams:  c.NewGetUnitsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetUnitsRequestBody(),
	}
}

type GetUnitsRequest struct {
	client      *Client
	queryParams *GetUnitsQueryParams
	pathParams  *GetUnitsPathParams
	method      string
	headers     http.Header
	requestBody GetUnitsRequestBody
}

func (c *Client) NewGetUnitsQueryParams() *GetUnitsQueryParams {
	return &GetUnitsQueryParams{}
}

type GetUnitsQueryParams struct {
	PropertyID       string   `schema:"propertyId,omitempty"`
	UnitGroupIDs     []string `schema:"unitGroupIds,omitempty"`
	UnitAttributeIDs []string `schema:"unitAttributeIds,omitempty"`
	IsOccupied       bool     `schema:"isOccupied,omitempty"`
	MaintenaceType   string   `schema:"maintenanceType,omitempty"`
	Condition        string   `schema:"condition,omitempty"`
	TextSearch       string   `schema:"textSearch,omitempty"`
	PageNumber       int      `schema:"pageNumber,omitempty"`
	PageSize         int      `schema:"pageSize,omitempty"`
	Expand           []string `schema:"expand,omitempty"`
}

func (p GetUnitsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetUnitsRequest) QueryParams() *GetUnitsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetUnitsPathParams() *GetUnitsPathParams {
	return &GetUnitsPathParams{}
}

type GetUnitsPathParams struct {
}

func (p *GetUnitsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetUnitsRequest) PathParams() *GetUnitsPathParams {
	return r.pathParams
}

func (r *GetUnitsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetUnitsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetUnitsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetUnitsRequestBody() GetUnitsRequestBody {
	return GetUnitsRequestBody{}
}

type GetUnitsRequestBody struct {
}

func (r *GetUnitsRequest) RequestBody() *GetUnitsRequestBody {
	return nil
}

func (r *GetUnitsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetUnitsRequest) SetRequestBody(body GetUnitsRequestBody) {
	r.requestBody = body
}

func (r *GetUnitsRequest) NewResponseBody() *GetUnitsResponseBody {
	return &GetUnitsResponseBody{}
}

type GetUnitsResponseBody struct {
	Count int   `json:"count"`
	Units Units `json:"units"`
}

func (r *GetUnitsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("inventory/v1/units", r.PathParams())
	return &u
}

func (r *GetUnitsRequest) Do() (GetUnitsResponseBody, error) {
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

func (r *GetUnitsRequest) All() (Units, error) {
	units := Units{}
	for {
		resp, err := r.Do()
		if err != nil {
			return units, err
		}

		// Break out of loop when no units are found
		if len(resp.Units) == 0 {
			break
		}

		// Add units to list
		units = append(units, resp.Units...)

		if len(units) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return units, nil
}
