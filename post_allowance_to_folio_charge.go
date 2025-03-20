package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewPostAllowanceToFolioChargeRequest() PostAllowanceToFolioChargeRequest {
	return PostAllowanceToFolioChargeRequest{
		client:      c,
		queryParams: c.NewPostAllowanceToFolioChargeQueryParams(),
		pathParams:  c.NewPostAllowanceToFolioChargePathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostAllowanceToFolioChargeRequestBody(),
	}
}

type PostAllowanceToFolioChargeRequest struct {
	client      *Client
	queryParams *PostAllowanceToFolioChargeQueryParams
	pathParams  *PostAllowanceToFolioChargePathParams
	method      string
	headers     http.Header
	requestBody PostAllowanceToFolioChargeRequestBody
}

func (c *Client) NewPostAllowanceToFolioChargeQueryParams() *PostAllowanceToFolioChargeQueryParams {
	return &PostAllowanceToFolioChargeQueryParams{}
}

type PostAllowanceToFolioChargeQueryParams struct{}

func (p PostAllowanceToFolioChargeQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(CommaSeparatedQueryParam{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostAllowanceToFolioChargeRequest) QueryParams() *PostAllowanceToFolioChargeQueryParams {
	return r.queryParams
}

func (c *Client) NewPostAllowanceToFolioChargePathParams() *PostAllowanceToFolioChargePathParams {
	return &PostAllowanceToFolioChargePathParams{}
}

type PostAllowanceToFolioChargePathParams struct {
	FolioID  string `schema:"id"`
	ChargeID string `schema:"chargeId"`
}

func (p *PostAllowanceToFolioChargePathParams) Params() map[string]string {
	return map[string]string{
		"id":       p.FolioID,
		"chargeId": p.ChargeID,
	}
}

func (r *PostAllowanceToFolioChargeRequest) PathParams() *PostAllowanceToFolioChargePathParams {
	return r.pathParams
}

func (r *PostAllowanceToFolioChargeRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostAllowanceToFolioChargeRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostAllowanceToFolioChargeRequest) Method() string {
	return r.method
}

func (s *Client) NewPostAllowanceToFolioChargeRequestBody() PostAllowanceToFolioChargeRequestBody {
	return PostAllowanceToFolioChargeRequestBody{}
}

type PostAllowanceToFolioChargeRequestBody struct {
	Reason       string             `json:"reason,omitempty"`
	Amount       MonetaryValueModel `json:"amount"`
	BusinessDate *Date              `json:"businessDate,omitempty"`
}

func (r PostAllowanceToFolioChargeRequest) NewRequestBody() PostAllowanceToFolioChargeRequestBody {
	return PostAllowanceToFolioChargeRequestBody{}
}

func (r *PostAllowanceToFolioChargeRequest) RequestBody() *PostAllowanceToFolioChargeRequestBody {
	return &r.requestBody
}

func (r *PostAllowanceToFolioChargeRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *PostAllowanceToFolioChargeRequest) SetRequestBody(body PostAllowanceToFolioChargeRequestBody) {
	r.requestBody = body
}

func (r *PostAllowanceToFolioChargeRequest) NewResponseBody() *PostAllowanceToFolioChargeResponseBody {
	return &PostAllowanceToFolioChargeResponseBody{}
}

type PostAllowanceToFolioChargeResponseBody struct {
	ID string `json:"id"`
}

func (r *PostAllowanceToFolioChargeRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v0-nsfw/folio-actions/{{.id}}/charges/{{.chargeId}}/allowances", r.PathParams())
	return &u
}

func (r *PostAllowanceToFolioChargeRequest) Do() (PostAllowanceToFolioChargeResponseBody, error) {
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
