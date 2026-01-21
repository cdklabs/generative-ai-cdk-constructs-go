package piitype


// Types of PII specific to the USA.
// Experimental.
type USASpecific string

const (
	// A US bank account number, which is typically 10 to 12 digits long.
	// Experimental.
	USASpecific_US_BANK_ACCOUNT_NUMBER USASpecific = "US_BANK_ACCOUNT_NUMBER"
	// A US bank account routing number.
	//
	// These are typically nine digits long.
	// Experimental.
	USASpecific_US_BANK_ROUTING_NUMBER USASpecific = "US_BANK_ROUTING_NUMBER"
	// A US Individual Taxpayer Identification Number (ITIN) is a nine-digit number that starts with a "9" and contain a "7" or "8" as the fourth digit.
	// Experimental.
	USASpecific_US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER USASpecific = "US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER"
	// A US passport number.
	//
	// Passport numbers range from six to nine alphanumeric characters.
	// Experimental.
	USASpecific_US_PASSPORT_NUMBER USASpecific = "US_PASSPORT_NUMBER"
	// A US Social Security Number (SSN) is a nine-digit number that is issued to US citizens, permanent residents, and temporary working residents.
	// Experimental.
	USASpecific_US_SOCIAL_SECURITY_NUMBER USASpecific = "US_SOCIAL_SECURITY_NUMBER"
)

