package piitype


// Types of PII specific to the United Kingdom (UK).
// Experimental.
type UKSpecific string

const (
	// A UK National Health Service Number is a 10-17 digit number, such as 485 777 3456.
	// Experimental.
	UKSpecific_UK_NATIONAL_HEALTH_SERVICE_NUMBER UKSpecific = "UK_NATIONAL_HEALTH_SERVICE_NUMBER"
	// A UK National Insurance Number (NINO) provides individuals with access to National Insurance (social security) benefits.
	//
	// It is also used for some purposes in the UK
	// tax system.
	// Experimental.
	UKSpecific_UK_NATIONAL_INSURANCE_NUMBER UKSpecific = "UK_NATIONAL_INSURANCE_NUMBER"
	// A UK Unique Taxpayer Reference (UTR) is a 10-digit number that identifies a taxpayer or a business.
	// Experimental.
	UKSpecific_UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER UKSpecific = "UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER"
)

