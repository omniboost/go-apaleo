package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetInvoicesRequest() GetInvoicesRequest {
	return GetInvoicesRequest{
		client:      c,
		queryParams: c.NewGetInvoicesQueryParams(),
		pathParams:  c.NewGetInvoicesPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetInvoicesRequestBody(),
	}
}

type GetInvoicesRequest struct {
	client      *Client
	queryParams *GetInvoicesQueryParams
	pathParams  *GetInvoicesPathParams
	method      string
	headers     http.Header
	requestBody GetInvoicesRequestBody
}

func (c *Client) NewGetInvoicesQueryParams() *GetInvoicesQueryParams {
	return &GetInvoicesQueryParams{}
}

type GetInvoicesQueryParams struct {
	Number                         string   `schema:"number,omitempty"`
	Status                         string   `schema:"status,omitempty"`
	CheckedOutOnAccountsReceivable bool     `schema:"checkedOutOnAccountsReceivable,omitempty"`
	OutstandingPaymentFilter       []string `schema:"outstandingPaymentFilter,omitempty"`
	DateFilter                     []string `schema:"dateFilter,omitempty"`
	PropertyIDs                    []string `schema:"propertyIds,omitempty"`
	ReservationIDs                 []string `schema:"reservationIds,omitempty"`
	BookingIDs                     []string `schema:"bookingIds,omitempty"`
	FolioIDs                       []string `schema:"folioIds,omitempty"`
	NameSearch                     string   `schema:"nameSearch,omitempty"`
	PaymentSettled                 bool     `schema:"paymentSettled,omitempty"`
	CompanyIDs                     []string `schema:"companyIds,omitempty"`
	PageNumber                     int      `schema:"pageNumber,omitempty"`
	PageSize                       int      `schema:"pageSize,omitempty"`
	Expand                         []string `schema:"expand,omitempty"`
}

func (p GetInvoicesQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetInvoicesRequest) QueryParams() *GetInvoicesQueryParams {
	return r.queryParams
}

func (c *Client) NewGetInvoicesPathParams() *GetInvoicesPathParams {
	return &GetInvoicesPathParams{}
}

type GetInvoicesPathParams struct {
}

func (p *GetInvoicesPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetInvoicesRequest) PathParams() *GetInvoicesPathParams {
	return r.pathParams
}

func (r *GetInvoicesRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetInvoicesRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetInvoicesRequest) Method() string {
	return r.method
}

func (s *Client) NewGetInvoicesRequestBody() GetInvoicesRequestBody {
	return GetInvoicesRequestBody{}
}

type GetInvoicesRequestBody struct {
}

func (r *GetInvoicesRequest) RequestBody() *GetInvoicesRequestBody {
	return nil
}

func (r *GetInvoicesRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetInvoicesRequest) SetRequestBody(body GetInvoicesRequestBody) {
	r.requestBody = body
}

func (r *GetInvoicesRequest) NewResponseBody() *GetInvoicesResponseBody {
	return &GetInvoicesResponseBody{}
}

type GetInvoicesResponseBody struct {
	Count    int      `json:"count"`
	Invoices Invoices `json:"invoices"`
}

func (r *GetInvoicesRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("finance/v1/invoices", r.PathParams())
	return &u
}

func (r *GetInvoicesRequest) Do() (GetInvoicesResponseBody, error) {
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

func (r *GetInvoicesRequest) All() (Invoices, error) {
	invoices := Invoices{}
	for {
		resp, err := r.Do()
		if err != nil {
			return invoices, err
		}

		// Break out of loop when no invoices are found
		if len(resp.Invoices) == 0 {
			break
		}

		// Add invoices to list
		invoices = append(invoices, resp.Invoices...)

		if len(invoices) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return invoices, nil
}
