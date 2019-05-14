package apaleo

type PropertyItemModel struct {
	// The property id
	ID string `json:"id"`
	// The code for the property that can be shown in reports and table views
	Code string `json:"code"`
	// The name for the property
	Name string `json:"name"`
	// The description for the property
	Description string `json:"description"`
	// The location of the property
	Location AddressModel `json:"location"`
	// The time zone name of the property from the IANA Time Zone Database.
	// (see: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
	TimeZone string `json:"timeZone"`
	// The currency a property works with.
	CurrencyCode string `json:"currencyCode"`
	// Date of creation
	// Specify a date and time (without fractional second part) in UTC or with
	// UTC offset as defined in the ISO8601:2004
	Created Date `json:"created"`
}

type AddressModel struct {
	AddressLine1 string `json:"addressline1"`
	AddressLine2 string `json:"addressline2"`
	Postalcode   string `json:"postalcode"`
	City         string `json:"city"`
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
