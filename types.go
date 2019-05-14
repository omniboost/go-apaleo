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
	Location struct {
		AddressLine1 string `json:"addressline1"`
		AddressLine2 string `json:"addressline2"`
		Postalcode   string `json:"postalcode"`
		City         string `json:"city"`
		CountryCode  string `json:"countryCode"`
	}
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
