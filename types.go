package apaleo

import (
	"github.com/cydev/zero"
	"github.com/omniboost/go-apaleo/omitempty"
)

var (
	AccountTypeRevenues           AccountType = "Revenues"
	AccountTypePayments           AccountType = "Payments"
	AccountTypeLiabilities        AccountType = "Liabilities"
	AccountTypeReceivables        AccountType = "Receivables"
	AccountTypeVat                AccountType = "Vat"
	AccountTypeHouse              AccountType = "House"
	AccountTypeAccountsReceivable AccountType = "AccountsReceivable"
	AccountTypeCityTaxes          AccountType = "CityTaxes"
	AccountTypeTransitoryItems    AccountType = "TransitoryItems"
	AccountTypeVatOnLiabilities   AccountType = "VatOnLiabilities"

	AccountingSchemaSimple   AccountingSchema = "Simple"
	AccountingSchemaExtended AccountingSchema = "Extended"
)

type AccountType string
type AccountingSchema string

type ActionReasonModel struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ActionModel struct {
	Action    string              `json:"action"`
	IsAllowed bool                `json:"isAllowed"`
	Reasons   []ActionReasonModel `json:"reasons"`
}

type AgeCategoryItemModel struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	PropertyID string `json:"propertyId"`
	Name       string `json:"name"`
	MinAge     int32  `json:"minAge"`
	MaxAge     int32  `json:"maxAge"`
}

type AgeCategories []AgeCategoryItemModel

type CancellationPolicyItemModel struct {
	ID                  string          `json:"id"`
	Name                string          `json:"name"`
	Code                string          `json:"code"`
	Description         string          `json:"description"`
	PropertyID          string          `json:"propertyId"`
	PeriodFromReference PeriodModel     `json:"periodFromReference"`
	Reference           string          `json:"reference"`
	Fee                 FeeDetailsModel `json:"fee"`
}

type CancellationPolicies []CancellationPolicyItemModel

type NoShowPolicyItemModel struct {
	ID          string          `json:"id"`
	Code        string          `json:"code"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	PropertyID  string          `json:"propertyId"`
	Fee         FeeDetailsModel `json:"fee"`
}

type NoShowPolicies []NoShowPolicyItemModel

type FeeDetailsModel struct {
	VatType      string             `json:"vatType"`
	FixedValue   MonetaryValueModel `json:"fixedValue"`
	PercentValue PercentValueModel  `json:"percentValue"`
}

type UnitItemModel struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	Property       EmbeddedPropertyModel  `json:"property"`
	UnitGroup      EmbeddedUnitGroupModel `json:"unitGroup"`
	ConnectingUnit EmbeddedUnitModel      `json:"connectingUnit"`
	Status         UnitItemStatusModel    `json:"status"`
	MaxPersons     int32                  `json:"maxPersons"`
	Created        DateTime               `json:"created"`
	Attributes     []UnitAttributeModel   `json:"attributes"`
	ConnectedUnits ConnectedUnitModel     `json:"connectedUnits"`
}

type Units []UnitItemModel

type UnitItemStatusModel struct {
	IsOccupied  bool                     `json:"isOccupied"`
	Condition   string                   `json:"condition"`
	Maintenance UnitItemMaintenanceModel `json:"maintenance"`
}

type UnitItemMaintenanceModel struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type UnitAttributeModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ConnectedUnitModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UnitGroupID string `json:"unitGroupId"`
	Condition   string `json:"condition"`
	MaxPersons  int32  `json:"maxPersons"`
}

type UnitAttributeDefinitionModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UnitAttributes []UnitAttributeDefinitionModel

type UnitGroupItemModel struct {
	ID                  string                    `json:"id"`
	Code                string                    `json:"code"`
	Name                string                    `json:"name"`
	Description         string                    `json:"description"`
	MemberCount         int32                     `json:"memberCount"`
	MaxPersons          int32                     `json:"maxPersons"`
	Rank                int32                     `json:"rank"`
	Type                string                    `json:"type"`
	Property            EmbeddedPropertyModel     `json:"property"`
	ConnectedUnitGroups []ConnectedUnitGroupModel `json:"connectedUnitGroups"`
}

type UnitGroups []UnitGroupItemModel

type ConnectedUnitGroupModel struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MemberCount int32  `json:"memberCount"`
	MaxPersons  int32  `json:"maxPersons"`
}

type PropertyItemModel struct {
	// The property id
	ID string `json:"id"`

	// The code for the property that can be shown in reports and table views
	Code string `json:"code"`

	// The id of the property used as a template while creating the property
	PropertyTemplateID string `json:"propertyTemplateId"`

	// Whether the property can be used as a template for other properties
	IsTemplate bool `json:"isTemplate"`

	// The name for the property
	Name string `json:"name"`

	// The description for the property
	Description string `json:"description"`

	// The legal name of the company running the property.
	CompanyName string `json:"companyName"`

	// The managing director(s) of the company, as they should appear on invoices
	ManagingDirectors string `json:"managingDirectors"`

	// The entry in the Commercial Register of the company running the property, as it should appear on invoices
	CommercialRegisterEntry string `json:"commercialRegisterEntry"`

	// The Tax-ID of the company running the property, as it should appear on invoices
	TaxID string `json:"taxId"`

	// The location of the property
	Location AddressModel `json:"location"`

	BankAccount BankAccountModel `json:"bankAccount"`

	// 	The payment terms used for all rate plans
	PaymentTerms map[string]string `json:"paymentTerms"`

	// The time zone name of the property from the IANA Time Zone Database.
	// (see: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
	TimeZone string `json:"timeZone"`

	// The currency a property works with.
	CurrencyCode string `json:"currencyCode"`

	// Date of creation
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	Created Date `json:"created"`

	// The status of the property
	Status string `json:"status"`

	// Is the property archived
	IsArchived bool `json:"isArchived"`

	// The list of actions for this property
	Actions []ActionModel `json:"actions"`
}

type Properties []PropertyItemModel

type AddressModel struct {
	AddressLine1 string `json:"addressline1"`
	AddressLine2 string `json:"addressline2"`
	PostalCode   string `json:"postalcode"`
	City         string `json:"city"`
	RegionCode   string `json:"regionCode"`
	CountryCode  string `json:"countryCode"`
}

type PropertyModel struct {
	// The property id
	ID string `json:"id"`

	// The code for the property that can be shown in reports and table views
	Code string `json:"code"`

	// The id of the property used as a template while creating the property
	PropertyTemplateID string `json:"propertyTemplateId"`

	// Whether the property can be used as a template for other properties
	IsTemplate bool `json:"isTemplate"`

	// The name for the property
	Name map[string]string `json:"name"`

	// The description for the property
	Description map[string]string `json:"description"`

	// The legal name of the company running the property.
	CompanyName string `json:"companyName"`

	// The managing director(s) of the company, as they should appear on invoices
	ManagingDirectors string `json:"managingDirectors"`

	// The entry in the Commercial Register of the company running the property, as it should appear on invoices
	CommercialRegisterEntry string `json:"commercialRegisterEntry"`

	// The Tax-ID of the company running the property, as it should appear on invoices
	TaxID string `json:"taxId"`

	// The location of the property
	Location AddressModel `json:"location"`

	BankAccount BankAccountModel `json:"bankAccount"`

	// 	The payment terms used for all rate plans
	PaymentTerms map[string]string `json:"paymentTerms"`

	// The time zone name of the property from the IANA Time Zone Database.
	// (see: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
	TimeZone string `json:"timeZone"`

	// The currency a property works with.
	CurrencyCode string `json:"currencyCode"`

	// Date of creation
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	Created Date `json:"created"`

	// The status of the property
	Status string `json:"status"`

	// Is the property archived
	IsArchived bool `json:"isArchived"`

	// The list of actions for this property
	Actions []ActionModel `json:"actions"`
}

type BankAccountModel struct {
	IBAN string `json:"iban"`
	BIC  string `json:"bic"`
	Bank string `json:"bank"`
}

type EmbeddedCompanyModel struct {
	ID              string `json:"id"`
	Code            string `json:"code"`
	Name            string `json:"name"`
	CanCheckOutOnAR bool   `json:"canCheckOutOnAr"`
}

type InvoiceItemModel struct {
	ID                   string               `json:"id"`
	Number               string               `json:"number"`
	Type                 string               `json:"type"`
	LanguageCode         string               `json:"languageCode"`
	FolioID              string               `json:"folioId"`
	ReservationID        string               `json:"reservationId"`
	BookingID            string               `json:"bookingId"`
	PropertyID           string               `json:"propertyId"`
	RelatedInvoiceNumber string               `json:"relatedInvoiceNumber"`
	WriteOffReason       string               `json:"writeOffReason"`
	SubTotal             MonetaryValueModel   `json:"subTotal"`
	OutstandingPayment   MonetaryValueModel   `json:"outstandingPayment"`
	PaymentSettled       bool                 `json:"paymentSettled"`
	Status               string               `json:"status"`
	Created              string               `json:"created"`
	GuestName            string               `json:"guestName"`
	GuestCompany         string               `json:"guestCompany"`
	AllowedActions       []string             `json:"allowedActions"`
	Company              EmbeddedCompanyModel `json:"company"`
}

type Invoices []InvoiceItemModel

type ExportTransactionItemModel struct {
	// Timestamp with time zone information, when the booking was done
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004 ,
	Timestamp DateTime `json:"timestamp"`
	// The business date when the booking was done
	Date Date `json:"date"`
	// The account being debited (The 'from' in 'book x from account 1 to
	// account 2')
	DebitedAccountNumber string `json:"debitedAccountNumber"`
	// The parent account of the account being debited
	DebitedAccountParentNumber string `json:"debitedAccountParentNumber"`
	// The account being credited (The 'to' in 'book x from account 1 to account
	// 2')
	CreditedAccountNumber string `json:"creditedAccountNumber"`
	// The parent account of the account being credited
	CreditedAccountParentNumber string `json:"creditedAccountParentNumber"`
	// The type of business transaction which triggered the booking =
	// ['PostCharge', 'PostPayment', 'MoveLineItem', 'PostPrepayment',
	// 'PostToAccountsReceivables', 'PostPrepaymentVat', 'System']
	Command string `json:"command"`
	// The amount being booked
	Amount MonetaryValueModel `json:"amount"`
	// The receipt specifying type and number of the receipt for the business
	// transaction behind this entry. The receipt cannot be changed. It can be
	// identified by the combination of type and number
	Receipt ReceiptModel `json:"receipt"`
	// All transactions having the same number form one booking
	EntryNumber string `json:"entryNumber"`
	// The id of the reservation. Can be empty for transactions made on the
	// house account
	ReservationID string `json:"reservationId"`
}

type AmountModel struct {
	GrossAmount float64 `json:"grossAmount"`
	NetAmount   float64 `json:"netAmount"`
	VatType     string  `json:"vatType"`
	VatPercent  float64 `json:"vatPercent"`
	Currency    string  `json:"currency"`
}

type MonetaryValueModel struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type PercentValueModel struct {
	Percent           int32    `json:"percent"`
	Limit             int32    `json:"limit"`
	IncludeServiceIDs []string `json:"includeServiceIds"`
}

func (j MonetaryValueModel) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(j)
}

func (j MonetaryValueModel) IsEmpty() bool {
	return zero.IsZero(j)
}

type ReceiptModel struct {
	// The type of receipt. = ['Custom', 'Reservation', 'Invoice',
	// 'PspReference']
	Type   string `json:"type"`
	Number string `json:"number"`
}

type ReservationItemModel struct {
	// Reservation id
	ID string `json:"id"`

	// Booking id
	BookingID string `json:"bookingId"`

	// Block id
	BlockID string `json:"blockId"`

	// Status of the reservation = ['Confirmed', 'InHouse', 'CheckedOut',
	// 'Canceled', 'NoShow'],
	Status string `json:"status"`

	// Time of check-in
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	CheckInTime DateTime `json:"checkInTime"`

	// Time of check-out
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	CheckOutTime DateTime `json:"checkOutTime"`

	// Time of cancellation, if the reservation was canceled
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	CancellationTime DateTime `json:"cancellationTime"`

	// Time of setting no-show reservation status
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	NoShowTime DateTime `json:"noShowTime"`

	// The property
	Property EmbeddedPropertyModel `json:"property"`

	// The rate plan
	RatePlan EmbeddedRatePlanModel `json:"ratePlan"`

	// The unit group
	UnitGroup EmbeddedUnitGroupModel `json:"unitGroup"`

	// The unit
	Unit EmbeddedUnitModel `json:"unit,omitempty"`

	// The market segment
	MarketSegment EmbeddedMarketSegmentModel `json:"marketSegment,omitempty"`

	// Total amount
	TotalGrossAmount MonetaryValueModel `json:"totalGrossAmount"`

	// Date of arrival
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	Arrival DateTime `json:"arrival"`

	// Date of departure
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	Departure DateTime `json:"departure"`

	// Date of creation
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	Created DateTime `json:"created"`

	// Date of last modification
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	Modified DateTime `json:"modified"`

	// Number of adults
	Adults int `json:"adults"`

	// The ages of the children
	ChildrenAges []int `json:"childrenAges"`

	// Additional information and comments
	Comment string `json:"comment"`

	// Additional information and comment by the guest
	GuestComment string `json:"guestComment"`

	// Code in external system
	ExternalCode string `json:"externalCode"`

	// Channel code = ['Direct', 'BookingCom', 'Ibe', 'ChannelManager']
	ChannelCode string `json:"channelCode"`

	// Source of the reservation (e.g Hotels.com, Orbitz, etc.)
	Source string `json:"source"`

	// The primary guest of the reservation
	PrimaryGuest GuestModel `json:"primaryGuest"`

	// Additional guests of the reservation
	AdditionalGuests []GuestModel `json:"additionalGuests"`

	// The person who made the booking
	Booker BookerModel `json:"booker"`

	// Payment information
	PaymentAccount PaymentAccountModel `json:"paymentAccountModel"`

	// The strongest guarantee for the rate plans booked in this reservation =
	// ['PM6Hold', 'CreditCard', 'Prepayment', 'Company', 'Ota']
	GuaranteeType string `json:"guaranteeType"`

	// Details about the cancellation fee for this reservation<Paste>
	CancellationFee ReservationCancellationFeeModel `json:"cancellationFee"`

	// Details about the no-show fee for this reservation
	NoShowFee ReservationNoShowFeeModel `json:"noShowFee"`

	// The purpose of the trip, leisure or business = ['Business', 'Leisure']
	TravelPurpose string `json:"travelPurpose"`

	// The balance of this reservation
	Balance MonetaryValueModel `json:"balance"`

	// The list of units assigned to this reservation
	AssignedUnits []ReservationAssignedUnitModel `json:"assignedUnits"`
	// The list of time slices with the reserved units / unit groups for the
	// stay
	TimeSlices []TimeSliceModel `json:"timeSlices"`

	// The list of additional services (extras, add-ons) reserved for the stay
	Services []ReservationServiceItemModel `json:"services"`

	// Validation rules are applied to reservations during their lifetime. For
	// example a reservation that was created while the house or unit group is
	// already fully booked. Whenever a rule was or is currently violated, a
	// validation message will be added to this list. They can be deleted
	// whenever the hotel staff worked them off.
	ValidationMessages []ReservationValidationMessageModel `json:"validationMessages"`

	// The list of actions for this reservation
	Actions []ActionModel `json:"actions"`

	Company              EmbeddedCompanyModel `json:"company"`
	CorporateCode        string               `json:"corporateCode"`
	AllFoliosHaveInvoice bool                 `json:"allFoliosHaveInvoice"`
	HasCityTax           bool                 `json:"hasCityTax"`
	Commission           CommissionModel      `json:"commission"`
	PromoCode            string               `json:"promoCode"`
}

type Reservations []ReservationItemModel

type EmbeddedPropertyModel struct {
	// The property id
	ID string `json:"Id"`
	// The code for the property that can be shown in reports and table views
	Code string `json:"code"`
	// The name for the property
	Name string `json:"name"`
	// The description for the property
	Description string `json:"description"`
}

type EmbeddedRatePlanModel struct {
	// The rate plan id
	ID string `json:"id"`
	// The code for the rate plan that can be shown in reports and table views
	Code string `json:"code"`
	// The name for the rate plan
	Name string `json:"name"`
	// The description for the rate plan
	Description string `json:"description"`
	// Whether the rate plan is subject to city tax or not
	IsSubjectToCityTax bool `json:"isSubjectToCityTax"`
}

type EmbeddedUnitGroupModel struct {
	// The unit group id
	ID string `json:"id"`
	// The code for the unit group that can be shown in reports and table views
	Code string `json:"code"`
	// The name for the unit group
	Name string `json:"name"`
	// The description for the unit group
	Description string `json:"description"`
	// The unit group type
	Type string `json:"type"`
}

type EmbeddedUnitModel struct {
	// The unit id
	ID string `json:"id"`
	// The name for the unit
	Name string `json:"name"`
	// The description for the unit
	Description string `json:"description"`
}

type EmbeddedMarketSegmentModel struct {
	// The market segment id
	ID string `json:"id"`

	// The market segment code
	Code string `json:"code"`

	// The market segment name
	Name string `json:"name"`
}

type EmbeddedServiceModel struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EmbeddedCancellationPolicyModel struct {
	ID                   string      `json:"id"`
	Code                 string      `json:"code"`
	Name                 string      `json:"name"`
	Description          string      `json:"description"`
	PeriodPriorToArrival PeriodModel `json:"periodPriorToArrival"`
}

type EmbeddedNoShowPolicyModel struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EmbeddedTimeSliceDefinitionModel struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Template     string `json:"template"`
	CheckInTime  Time   `json:"checkInTime"`
	CheckOutTime Time   `json:"checkOutTime"`
}

type PeriodModel struct {
	Hours  int32 `json:"hours"`
	Days   int32 `json:"days"`
	Months int32 `json:"months"`
}

type GuestModel struct {
	Title                          string             `json:"title"`
	Gender                         string             `json:"gender"`
	FirstName                      string             `json:"firstName"`
	MiddleInitial                  string             `json:"middleInitial"`
	LastName                       string             `json:"lastName"`
	Email                          string             `json:"email"`
	Phone                          string             `json:"phone"`
	Address                        PersonAddressModel `json:"address"`
	NationalityCountryCode         string             `json:"nationalityCountryCode,omitempty"`
	IdentificationNumber           string             `json:"identificationNumber,omitempty"`
	IdentificationAdditionalNumber string             `json:"identificationAdditionalNumber,omitempty"`
	IdentificationIssueDate        string             `json:"identificationIssueDate,omitempty"`
	IdentificationExpiryDate       string             `json:"identificationExpiryDate,omitempty"`
	IdentificationIssuePlace       string             `json:"identificationIssuePlace,omitempty"`
	IdentificationType             string             `json:"identificationType,omitempty"`
	PersonalTaxID                  string             `json:"personalTaxId"`
	Company                        PersonCompanyModel `json:"company"`
	PreferredLanguage              string             `json:"preferredLanguage,omitempty"`
	BirthDate                      Date               `json:"birthDate,omitempty"`
	BirthFirstName                 string             `json:"birthFirstName,omitempty"`
	BirthLastName                  string             `json:"birthLastName,omitempty"`
	MotherFirstName                string             `json:"motherFirstName,omitempty"`
	MotherLastName                 string             `json:"motherLastName,omitempty"`
	BorderCrossingPlace            string             `json:"borderCrossingPlace,omitempty"`
	BorderCrossingDate             string             `json:"borderCrossingDate,omitempty"`
}

type BookingItemModel struct {
	ID             string              `json:"id"`
	GroupID        string              `json:"groupId"`
	Booker         BookerModel         `json:"booker"`
	PaymentAccount PaymentAccountModel `json:"paymentAccount"`
	Comment        string              `json:"comment"`
	BookerComment  string              `json:"bookerComment"`
	Created        string              `json:"created"`
	Modified       string              `json:"modified"`
}

type BookerModel struct {
	Title                    string             `json:"title"`
	Gender                   string             `json:"gender"`
	FirstName                string             `json:"firstName"`
	MiddleInitial            string             `json:"middleInitial"`
	LastName                 string             `json:"lastName"`
	Email                    string             `json:"email"`
	Phone                    string             `json:"phone"`
	Address                  PersonAddressModel `json:"address"`
	NationalityCountryCode   string             `json:"nationalityCountryCode"`
	IdentificationNumber     string             `json:"identificationNumber,omitempty"`
	IdentificationIssueDate  string             `json:"identificationIssueDate,omitempty"`
	IdentificationExpiryDate string             `json:"identificationExpiryDate,omitempty"`
	IdentificationType       string             `json:"identificationType,omitempty"`
	Company                  PersonCompanyModel `json:"company"`
	PreferredLanguage        string             `json:"preferredLanguage,omitempty"`
	BirthDate                Date               `json:"birthDate,omitempty"`
	BirthPlace               string             `json:"birthPlace,omitempty"`
}

type PersonAddressModel struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	PostalCode   string `json:"postalCode"`
	City         string `json:"city"`
	RegionCode   string `json:"regionCode"`
	CountryCode  string `json:"countryCode"`
}

type PersonCompanyModel struct {
	Name  string `json:"name"`
	TaxID string `json:"taxId"`
}

type PaymentAccountModel struct {
	AccountNumber  string `json:"accountNumber"`
	AccountHolder  string `json:"accountHolder"`
	ExpiryMonth    string `json:"expiryMonth"`
	ExpiryYear     string `json:"expiryYear"`
	PaymentMethod  string `json:"paymentMethod"`
	PayerEmail     string `json:"payerEmail"`
	PayerReference string `json:"payerReference"`
	IsVirtual      bool   `json:"isVirtual"`
	InactiveReason string `json:"inactiveReason"`
}

type BookingReservationModel struct {
	ID               string                          `json:"id"`
	Status           string                          `json:"status"`
	ExternalCode     string                          `json:"externalCode"`
	ChannelCode      string                          `json:"channelCode"`
	Source           string                          `json:"source"`
	PaymentAccount   PaymentAccountModel             `json:"paymentAccount"`
	Arrival          string                          `json:"arrival"`
	Departure        string                          `json:"departure"`
	Adults           int32                           `json:"adults"`
	ChildrenAges     []int32                         `json:"childrenAges"`
	TotalGrossAmount MonetaryValueModel              `json:"totalGrossAmount"`
	Property         EmbeddedPropertyModel           `json:"property"`
	RatePlan         EmbeddedRatePlanModel           `json:"ratePlan"`
	UnitGroup        EmbeddedUnitGroupModel          `json:"unitGroup"`
	Services         []ReservationServiceItemModel   `json:"services"`
	GuestComment     string                          `json:"guestComment"`
	CancellationFee  ReservationCancellationFeeModel `json:"cancellationFee"`
	NoShowFee        ReservationNoShowFeeModel       `json:"noShowFee"`
	Company          EmbeddedCompanyModel            `json:"company"`
}

type ReservationCancellationFeeModel struct {
	ID          string             `json:"id"`
	Code        string             `json:"code"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	DueDateTime string             `json:"dueDateTime"`
	Fee         MonetaryValueModel `json:"fee"`
}

type ReservationNoShowFeeModel struct {
	ID          string             `json:"id"`
	Code        string             `json:"code"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Fee         MonetaryValueModel `json:"fee"`
}

type ReservationAssignedUnitModel struct {
	Unit      EmbeddedUnitModel                       `json:"unit"`
	TimeRange []ReservationAssignedUnitTimeRangeModel `json:"timeRanges"`
}

type ReservationAssignedUnitTimeRangeModel struct {
	From DateTime `json:"from"`
	To   DateTime `json:"to"`
}

type ReservationServiceModel struct {
	Service       EmbeddedServiceModel `json:"service"`
	ServiceDate   string               `json:"serviceDate"`
	Count         int32                `json:"count"`
	Amount        AmountModel          `json:"amount"`
	BookedAsExtra bool                 `json:"bookedAsExtra"`
}

type ReservationServiceItemModel struct {
	Service     ServiceModel           `json:"service"`
	TotalAmount AmountModel            `json:"totalAmount"`
	Dates       []ServiceDateItemModel `json:"dates"`
}

type ServiceModel struct {
	ID                string             `json:"id"`
	Code              string             `json:"code"`
	Name              string             `json:"name"`
	Description       string             `json:"description"`
	PricingUnit       string             `json:"pricingUnit"`
	DefaultGrossPrice MonetaryValueModel `json:"defaultGrossPrice"`
}

type ServiceDateItemModel struct {
	ServiceDate string      `json:"serviceDate"`
	Count       int32       `json:"count"`
	Amount      AmountModel `json:"amount"`
	IsMandatory bool        `json:"isMandatory"`
}

type ReservationValidationMessageModel struct {
	Service     ServiceModel           `json:"service"`
	TotalAmount AmountModel            `json:"totalAmount"`
	Dates       []ServiceDateItemModel `json:"dates"`
}

type TimeSliceModel struct {
	From             DateTime                  `json:"from"`
	To               DateTime                  `json:"to"`
	ServiceDate      string                    `json:"serviceDate"`
	RatePlan         EmbeddedRatePlanModel     `json:"ratePlan"`
	UnitGroup        EmbeddedUnitGroupModel    `json:"unitGroup"`
	Unit             EmbeddedUnitModel         `json:"unit"`
	BaseAmount       AmountModel               `json:"baseAmount"`
	TotalGrossAmount MonetaryValueModel        `json:"totalGrossAmount"`
	IncludedServices []ReservationServiceModel `json:"includedServices"`
	Actions          ActionModel               `json:"actions"`
}

type ServiceItemModel struct {
	ID                string                `json:"id"`
	Name              string                `json:"name"`
	Code              string                `json:"code"`
	Description       string                `json:"description"`
	DefaultGrossPrice MonetaryValueModel    `json:"defaultGrossPrice"`
	PricingUnit       string                `json:"pricingUnit"`
	PostNextDay       bool                  `json:"postNextDay"`
	ServiceType       string                `json:"serviceType"`
	VatType           string                `json:"vatType"`
	Availability      AvailabilityModel     `json:"availability"`
	Property          EmbeddedPropertyModel `json:"property"`
	SubAccountID      string                `json:"subAccountId"`
	ChannelCodes      []string              `json:"channelCodes"`
	AgeCategoryID     string                `json:"ageCategoryId"`
}

type Services []ServiceItemModel

type AvailabilityModel struct {
	Mode       string   `json:"mode"`
	Quantity   int32    `json:"quantity"`
	DaysOfWeek []string `json:"daysOfWeek"`
}

type RatePlanItemModel struct {
	ID                   string                           `json:"id"`
	Code                 string                           `json:"code"`
	Name                 string                           `json:"name"`
	Description          string                           `json:"description"`
	MinGuaranteeType     string                           `json:"minGuaranteeType"`
	PriceCalculationMode string                           `json:"priceCalculationMode"`
	Property             EmbeddedPropertyModel            `json:"property"`
	UnitGroup            EmbeddedUnitGroupModel           `json:"unitGroup"`
	CancellationPolicy   EmbeddedCancellationPolicyModel  `json:"cancellationPolicy"`
	NoShowPolicy         EmbeddedNoShowPolicyModel        `json:"noShowPolicy"`
	ChannelCodes         []string                         `json:"channelCodes"`
	PromoCodes           []string                         `json:"promoCodes"`
	TimeSliceDefinition  EmbeddedTimeSliceDefinitionModel `json:"timeSliceDefinition"`
	Restrictions         BookingRestrictionsModel         `json:"restrictions"`
	BookingPeriods       []BookingPeriodModel             `json:"bookingPeriods"`
	IsBookable           bool                             `json:"isBookable"`
	IsSubjectToCityTax   bool                             `json:"isSubjectToCityTax"`
	PricingRule          PricingRuleModel                 `json:"pricingRule"`
	IsDerived            bool                             `json:"isDerived"`
	DerivationLevel      int32                            `json:"derivationLevel"`
	Surcharges           []SurchargeModel                 `json:"surcharges"`
	AgeCategories        []RatePlanAgeCategoryModel       `json:"ageCategories"`
	IncludedServices     []RatePlanServiceItemModel       `json:"includedServices"`
	Companies            []CompanyRatePlanModel           `json:"companies"`
	RatesRange           RatesRangeModel                  `json:"ratesRange"`
	AccountingConfigs    []AccountingConfigModel          `json:"accountingConfigs"`
}

type RatePlans []RatePlanItemModel

type CompanyModel struct {
	ID              string                 `json:"id"`
	Code            string                 `json:"code"`
	PropertyID      string                 `json:"propertyId"`
	Name            string                 `json:"name"`
	TaxID           string                 `json:"taxId"`
	AdditionalTaxID string                 `json:"additionalTaxId"`
	Address         CompanyAddressModel    `json:"address"`
	CanCheckOutOnAr bool                   `json:"canCheckOutOnAr"`
	RatePlans       []RatePlanCompanyModel `json:"ratePlans"`
}

type Companies []CompanyModel

type CompanyAddressModel struct {
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	PostalCode   string `json:"postalCode"`
	City         string `json:"city"`
	RegionCode   string `json:"regionCode"`
	CountryCode  string `json:"countryCode"`
}

type RatePlanCompanyModel struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	CorporateCode string `json:"corporateCode"`
	Name          string `json:"name"`
}

type CreateCompanyModel struct {
	Code            string                 `json:"code"`
	PropertyID      string                 `json:"propertyId"`
	Name            string                 `json:"name"`
	TaxID           string                 `json:"taxId"`
	AdditionalTaxID string                 `json:"additionalTaxId"`
	Address         CompanyAddressModel    `json:"address"`
	CanCheckOutOnAr bool                   `json:"canCheckOutOnAr"`
	RatePlans       []RatePlanCompanyModel `json:"ratePlans"`
}

func (j CreateCompanyModel) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(j)
}

func (j CreateCompanyModel) IsEmpty() bool {
	return zero.IsZero(j)
}

type CreateRatePlanCompanyModel struct {
	ID            string `json:"id"`
	CorporateCode string `json:"corporateCode,omitempty"`
}

func (j CreateRatePlanCompanyModel) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(j)
}

func (j CreateRatePlanCompanyModel) IsEmpty() bool {
	return zero.IsZero(j)
}

type BookingRestrictionsModel struct {
	MinAdvance       PeriodModel `json:"minAdvance"`
	MaxAdvance       PeriodModel `json:"maxAdvance"`
	LateBookingUntil Time        `json:"lateBookingUntil"`
}

type BookingPeriodModel struct {
	From DateTime `json:"from"`
	To   DateTime `json:"to"`
}

type PricingRuleModel struct {
	BaseRatePlan EmbeddedRatePlanModel `json:"baseRatePlan"`
	Type         string                `json:"type"`
	Value        float64               `json:"value"`
}

type SurchargeModel struct {
	Adults int32   `json:"adults"`
	Type   string  `json:"type"`
	Value  float64 `json:"value"`
}

type RatePlanAgeCategoryModel struct {
	ID         string `json:"id"`
	Surcharges []AgeCategorySurchageModel
}

type AgeCategorySurchageModel struct {
	Adults int32   `json:"adults"`
	Value  float64 `json:"value"`
}

type RatePlanServiceItemModel struct {
	Service     EmbeddedServiceModel `json:"service"`
	GrossPrice  MonetaryValueModel   `json:"grossPrice"`
	PricingMode string               `json:"pricingMode"`
}

type CompanyRatePlanModel struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	CorporateCode string `json:"corporateCode"`
	Name          string `json:"name"`
}

type RatesRangeModel struct {
	From Date `json:"from"`
	To   Date `json:"to"`
}

type AccountingConfigModel struct {
	VatType      string `json:"vatType"`
	ServiceType  string `json:"serviceType"`
	SubAccountID string `json:"subAccountId"`
	ValidFrom    Date   `json:"validFrom"`
}

type CommissionModel struct {
	ComissionAmount        MonetaryValueModel `json:"comissionAmount"`
	BeforeCommissionAmount MonetaryValueModel `json:"beforeCommissionAmount"`
}

func (j CommissionModel) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(j)
}

func (j CommissionModel) IsEmpty() bool {
	return zero.IsZero(j)
}

type CreatePaymentAccountModel struct {
	AccountNumber  string `json:"accountNumber"`
	AccountHolder  string `json:"accountHolder"`
	ExpiryMonth    string `json:"expiryMonth"`
	ExpiryYear     string `json:"expiryYear"`
	PaymentMethod  string `json:"paymentMethod"`
	PayerEmail     string `json:"payerEmail"`
	PayerReference string `json:"payerReference"`
	IsVirtual      bool   `json:"isVirtual"`
	InactiveReason string `json:"inactiveReason,omitempty"`
}

func (j CreatePaymentAccountModel) IsEmpty() bool {
	return zero.IsZero(j)
}

type CreateReservationModel struct {
	Arrival          Date                              `json:"arrival"`
	Departure        Date                              `json:"departure"`
	Adults           int32                             `json:"adults"`
	ChildrenAges     []int32                           `json:"childrenAges,omitempty"`
	Comment          string                            `json:"comment,omitempty"`
	GuestComment     string                            `json:"guestComment,omitempty"`
	ExternalCode     string                            `json:"externalCode,omitempty"`
	ChannelCode      string                            `json:"channelCode"`
	Source           string                            `json:"source,omitempty"`
	PrimaryGuest     GuestModel                        `json:"primaryGuest,omitempty"`
	AdditionalGuest  []GuestModel                      `json:"additionalGuest,omitempty"`
	GuaranteeType    string                            `json:"guaranteeType,omitempty"`
	TravelPurpose    string                            `json:"travelPurpose,omitempty"`
	TimeSlices       []CreateReservationTimeSliceModel `json:"timeSlices"`
	Services         []BookReservationServiceModel     `json:"services,omitempty"`
	CompanyID        string                            `json:"companyId,omitempty"`
	CorporateCode    string                            `json:"corporateCode,omitempty"`
	PrePaymentAmount MonetaryValueModel                `json:"prePaymentAmount,omitempty"`
	Commission       CommissionModel                   `json:"commission,omitempty"`
	PromoCode        string                            `json:"promoCode,omitempty"`
}

func (j CreateReservationModel) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(j)
}

func (j CreateReservationModel) IsEmpty() bool {
	return zero.IsZero(j)
}

type CreateReservationTimeSliceModel struct {
	RatePlanID  string             `json:"ratePlanId"`
	TotalAmount MonetaryValueModel `json:"totalAmount"`
}

type BookReservationServiceModel struct {
	ServiceID string             `json:"serviceId"`
	Count     int32              `json:"count,omitempty"`
	Amount    MonetaryValueModel `json:"amount,omitempty"`
	Dates     []struct {
		ServiceDate Date               `json:"serviceDate"`
		Count       int32              `json:"count,omitempty"`
		Amount      MonetaryValueModel `json:"amount,omitempty"`
	} `json:"dates,omitempty"`
}
