package apaleo

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetFoliosRequest() GetFoliosRequest {
	return GetFoliosRequest{
		client:      c,
		queryParams: c.NewGetFoliosQueryParams(),
		pathParams:  c.NewGetFoliosPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetFoliosRequestBody(),
	}
}

type GetFoliosRequest struct {
	client      *Client
	queryParams *GetFoliosQueryParams
	pathParams  *GetFoliosPathParams
	method      string
	headers     http.Header
	requestBody GetFoliosRequestBody
}

func (c *Client) NewGetFoliosQueryParams() *GetFoliosQueryParams {
	return &GetFoliosQueryParams{}
}

type GetFoliosQueryParams struct {
	PropertyIDs                    CommaSeparatedQueryParam `schema:"propertyIds,omitempty"`
	CompanyIDs                     CommaSeparatedQueryParam `schema:"companyIds,omitempty"`
	ReservationIDs                 CommaSeparatedQueryParam `schema:"reservationIds,omitempty"`
	BookingIDs                     CommaSeparatedQueryParam `schema:"bookingIds,omitempty"`
	IsEmpty                        Bool                     `schema:"isEmpty,omitempty"`
	CheckedOutOnAccountsReceivable Bool                     `schema:"checkedOutOnAccountsReceivable,omitempty"`
	ExcludeClosed                  Bool                     `schema:"excludeClosed,omitempty"`
	HasInvoices                    Bool                     `schema:"hasInvoices,omitempty"`
	CreatedFrom                    DateTime                 `schema:"createdFrom,omitempty"`
	CreatedTo                      DateTime                 `schema:"createdTo,omitempty"`
	UpdatedFrom                    DateTime                 `schema:"updatedFrom,omitempty"`
	UpdatedTo                      DateTime                 `schema:"updatedTo,omitempty"`
	OnlyMain                       Bool                     `schema:"onlyMain,omitempty"`
	Type                           string                   `schema:"type,omitempty"`
	ExternalFolioCode              string                   `schema:"externalFolioCode,omitempty"`
	TextSearch                     string                   `schema:"textSearch,omitempty"`
	BalanceFilter                  CommaSeparatedQueryParam `schema:"balanceFilter,omitempty"`
	PageNumber                     int                      `schema:"pageNumber,omitempty"`
	PageSize                       int                      `schema:"pageSize,omitempty"`
	Expand                         []string                 `schema:"expand,omitempty"`
}

func (p GetFoliosQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetFoliosRequest) QueryParams() *GetFoliosQueryParams {
	return r.queryParams
}

func (c *Client) NewGetFoliosPathParams() *GetFoliosPathParams {
	return &GetFoliosPathParams{}
}

type GetFoliosPathParams struct {
}

func (p *GetFoliosPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetFoliosRequest) PathParams() *GetFoliosPathParams {
	return r.pathParams
}

func (r *GetFoliosRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetFoliosRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetFoliosRequest) Method() string {
	return r.method
}

func (s *Client) NewGetFoliosRequestBody() GetFoliosRequestBody {
	return GetFoliosRequestBody{}
}

type GetFoliosRequestBody struct {
}

func (r *GetFoliosRequest) RequestBody() *GetFoliosRequestBody {
	return nil
}

func (r *GetFoliosRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetFoliosRequest) SetRequestBody(body GetFoliosRequestBody) {
	r.requestBody = body
}

func (r *GetFoliosRequest) NewResponseBody() *GetFoliosResponseBody {
	return &GetFoliosResponseBody{}
}

type GetFoliosResponseBody struct {
	Count  int              `json:"count"`
	Folios []FolioItemModel `json:"folios"`
}

func (r *GetFoliosRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/folios", r.PathParams())
	return &u
}

func (r *GetFoliosRequest) Do(ctx context.Context) (GetFoliosResponseBody, error) {
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

func (r *GetFoliosRequest) All(ctx context.Context) ([]FolioItemModel, error) {
	folios := []FolioItemModel{}
	for {
		resp, err := r.Do(ctx)
		if err != nil {
			return folios, err
		}

		// Break out of loop when no folios are found
		if len(resp.Folios) == 0 {
			break
		}

		// Add units to list
		folios = append(folios, resp.Folios...)

		if len(folios) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return folios, nil
}
