package apaleo

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

type PropertyListModel struct {
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

type PropertyList []PropertyListModel

type AddressModel struct {
	AddressLine1 string `json:"addressline1"`
	AddressLine2 string `json:"addressline2"`
	PostalCode   string `json:"postalcode"`
	City         string `json:"city"`
	RegionCode   string `json:"regionCode"`
	CountryCode  string `json:"countryCode"`
}

type PropertyModel struct {
	// The property id ,
	ID string `json:"id"`
	// The code for the property that can be shown in reports and table views
	Code string `json:"code"`
	// The name for the property
	Name map[string]string `json:"name"`
	// The description for the property
	Description map[string]string `json:"description"`
	// The legal name of the company running the property.
	CompanyName string `json:"companyName"`
	// The managing director(s) of the company, as they should appear on invoices
	ManagingDirectors string `json:"managingDirectors"`
	// The entry in the Commercial Reigster of the company running the property, as it should appear on invoices ,
	CommercialRegisterEntry string `json:"commercialRegisterEntry"`
	// The Tax-ID of the company running the property, as it should appear on invoices
	TaxID string `json:"taxId"`
	// The location of the property ,
	Location    AddressModel     `json:"location"`
	BankAccount BankAccountModel `json:"bankAccountModel"`
	// The payment terms used for all rate plans
	PaymentTerms map[string]string `json:"paymentTerms"`
	// The time zone name of the property from the IANA Time Zone Database. (see:
	// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
	TimeZone string `json:"timeZone"`
	// The currency a property works with.
	CurrencyCode string `json:"currencyCode"`
	// Date of creation
	// Specify a date and time (without fractional second part) in UTC or with UTC
	// offset as defined in the ISO8601:2004
	Created Date `json:"created"`
}

type BankAccountModel struct {
	IBAN string `json:"iban"`
	BIC  string `json:"bic"`
	Bank string `json:"bank"`
}

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

type MonetaryValueModel struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
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
	Comment string `json:"comment`
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
	// The list of additional services (extras, add-ons) reserved for the stay -
	// DEPRECATED: Please use 'Services'. This field will be removed on
	// 31.03.2019
	ExtraServices []ReservationServiceModel `json:"extraServices"`
	// The list of additional services (extras, add-ons) reserved for the stay
	Services ReservationServiceItemModel `json:"services"`
	// Validation rules are applied to reservations during their lifetime. For
	// example a reservation that was created while the house or unit group is
	// already fully booked. Whenever a rule was or is currently violated, a
	// validation message will be added to this list. They can be deleted
	// whenever the hotel staff worked them off.
	ValidationMessages []ReservationValidationMessageModel `json:"validationMessages"`
	// The list of actions for this reservation
	Actions []ActionModel        `json:"actions"`
	Company EmbeddedCompanyModel `json:"company,omitempty"`
}

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
}

type EmbeddedUnitModel struct {
	// The unit id
	ID string `json:"id"`
	// The name for the unit
	Name string `json:"name"`
	// The description for the unit
	Description string `json:"description"`
}
