package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetFolioByIDRequest() GetFolioByIDRequest {
	return GetFolioByIDRequest{
		client:      c,
		queryParams: c.NewGetFolioByIDQueryParams(),
		pathParams:  c.NewGetFolioByIDPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetFolioByIDRequestBody(),
	}
}

type GetFolioByIDRequest struct {
	client      *Client
	queryParams *GetFolioByIDQueryParams
	pathParams  *GetFolioByIDPathParams
	method      string
	headers     http.Header
	requestBody GetFolioByIDRequestBody
}

func (c *Client) NewGetFolioByIDQueryParams() *GetFolioByIDQueryParams {
	return &GetFolioByIDQueryParams{}
}

type GetFolioByIDQueryParams struct {
	Expand []string `schema:"expand,omitempty"`
}

func (p GetFolioByIDQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetFolioByIDRequest) QueryParams() *GetFolioByIDQueryParams {
	return r.queryParams
}

func (c *Client) NewGetFolioByIDPathParams() *GetFolioByIDPathParams {
	return &GetFolioByIDPathParams{}
}

type GetFolioByIDPathParams struct {
	ID string `schema:"id"`
}

func (p *GetFolioByIDPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *GetFolioByIDRequest) PathParams() *GetFolioByIDPathParams {
	return r.pathParams
}

func (r *GetFolioByIDRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetFolioByIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFolioByIDRequest) Method() string {
	return r.method
}

func (s *Client) NewGetFolioByIDRequestBody() GetFolioByIDRequestBody {
	return GetFolioByIDRequestBody{}
}

type GetFolioByIDRequestBody struct {
}

func (r *GetFolioByIDRequest) RequestBody() *GetFolioByIDRequestBody {
	return nil
}

func (r *GetFolioByIDRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetFolioByIDRequest) SetRequestBody(body GetFolioByIDRequestBody) {
	r.requestBody = body
}

func (r *GetFolioByIDRequest) NewResponseBody() *GetFolioByIDResponseBody {
	return &GetFolioByIDResponseBody{}
}

type GetFolioByIDResponseBody FolioModel

func (r *GetFolioByIDRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v0-nsfw/folios/{{.id}}", r.PathParams())
	return &u
}

func (r *GetFolioByIDRequest) Do(ctx context.Context) (GetFolioByIDResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
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
