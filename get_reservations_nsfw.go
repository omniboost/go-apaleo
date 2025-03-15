package apaleo

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-apaleo/utils"
)

func (c *Client) NewGetReservationsNSFWRequest() GetReservationsNSFWRequest {
	return GetReservationsNSFWRequest{
		client:      c,
		queryParams: c.NewGetReservationsNSFWQueryParams(),
		pathParams:  c.NewGetReservationsNSFWPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetReservationsNSFWRequestBody(),
	}
}

type GetReservationsNSFWRequest struct {
	client      *Client
	queryParams *GetReservationsNSFWQueryParams
	pathParams  *GetReservationsNSFWPathParams
	method      string
	headers     http.Header
	requestBody GetReservationsNSFWRequestBody
}

func (c *Client) NewGetReservationsNSFWQueryParams() *GetReservationsNSFWQueryParams {
	return &GetReservationsNSFWQueryParams{}
}

type GetReservationsNSFWQueryParams struct {
	BookingID                 string                   `schema:"booking_id,omitempty"`
	PropertyIDs               []string                 `schema:"propertyIds,omitempty"`
	RatePlanIDs               []string                 `schema:"ratePlanIds,omitempty"`
	CompanyIDs                []string                 `schema:"companyIds,omitempty"`
	UnitIDs                   []string                 `schema:"unitIds,omitempty"`
	UnitGroupIDs              []string                 `schema:"unitGroupIds,omitempty"`
	UnitGroupTypes            []string                 `schema:"unitGroupTypes,omitempty"`
	BlockIDs                  []string                 `schema:"blockIds,omitempty"`
	MarketSegmentIDs          []string                 `schema:"marketSegmentIds,omitempty"`
	Status                    []string                 `schema:"status,omitempty"`
	DateFilter                string                   `schema:"dateFilter,omitempty"`
	From                      DateTime                 `schema:"from,omitempty"`
	To                        DateTime                 `schema:"to,omitempty"`
	ChannelCode               []string                 `schema:"channelCode,omitempty"`
	Sources                   []string                 `schema:"sources,omitempty"`
	ValidationMessageCategory []string                 `schema:"validationMessageCategory,omitempty"`
	ExternalCode              string                   `schema:"externalCode,omitempty"`
	TextSearch                string                   `schema:"textSearch,omitempty"`
	ExternalReferences        CommaSeparatedQueryParam `schema:"externalReferences,omitempty"`
	BalanceFilter             []string                 `schema:"balanceFilter,omitempty"`
	AllFoliosHaveInvoice      bool                     `schema:"allFoliosHaveInvoice,omitempty"`
	IsPreCheckedIn            bool                     `schema:"isPreCheckedIn,omitempty"`
	PageNumber                int                      `schema:"pageNumber,omitempty"`
	PageSize                  int                      `schema:"pageSize,omitempty"`
	Sort                      []string                 `schema:"sort,omitempty"`
	Expand                    []string                 `schema:"expand,omitempty"`
}

func (p GetReservationsNSFWQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetReservationsNSFWRequest) QueryParams() *GetReservationsNSFWQueryParams {
	return r.queryParams
}

func (c *Client) NewGetReservationsNSFWPathParams() *GetReservationsNSFWPathParams {
	return &GetReservationsNSFWPathParams{}
}

type GetReservationsNSFWPathParams struct {
}

func (p *GetReservationsNSFWPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetReservationsNSFWRequest) PathParams() *GetReservationsNSFWPathParams {
	return r.pathParams
}

func (r *GetReservationsNSFWRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetReservationsNSFWRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetReservationsNSFWRequest) Method() string {
	return r.method
}

func (s *Client) NewGetReservationsNSFWRequestBody() GetReservationsNSFWRequestBody {
	return GetReservationsNSFWRequestBody{}
}

type GetReservationsNSFWRequestBody struct {
}

func (r *GetReservationsNSFWRequest) RequestBody() *GetReservationsNSFWRequestBody {
	return nil
}

func (r *GetReservationsNSFWRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetReservationsNSFWRequest) SetRequestBody(body GetReservationsNSFWRequestBody) {
	r.requestBody = body
}

func (r *GetReservationsNSFWRequest) NewResponseBody() *GetReservationsNSFWResponseBody {
	return &GetReservationsNSFWResponseBody{}
}

type GetReservationsNSFWResponseBody struct {
	Count        int              `json:"count"`
	Reservations ReservationsNSFW `json:"reservations"`
}

func (r *GetReservationsNSFWRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("booking/v0-nsfw/reservations", r.PathParams())
	return &u
}

func (r *GetReservationsNSFWRequest) Do() (GetReservationsNSFWResponseBody, error) {
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

func (r *GetReservationsNSFWRequest) All() (ReservationsNSFW, error) {
	reservations := ReservationsNSFW{}
	for {
		resp, err := r.Do()
		if err != nil {
			return reservations, err
		}

		// Break out of loop when no reservations are found
		if len(resp.Reservations) == 0 {
			break
		}

		// Add reservations to list
		reservations = append(reservations, resp.Reservations...)

		if len(reservations) == resp.Count {
			break
		}

		// Increment page number
		r.QueryParams().PageNumber = r.QueryParams().PageNumber + 1
	}

	return reservations, nil
}
