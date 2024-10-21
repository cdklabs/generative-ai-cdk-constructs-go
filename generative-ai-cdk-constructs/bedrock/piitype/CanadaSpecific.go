package piitype


// Types of PII specific to Canada.
// Experimental.
type CanadaSpecific string

const (
	// A Canadian Health Service Number is a 10-digit unique identifier, required for individuals to access healthcare benefits.
	// Experimental.
	CanadaSpecific_CA_HEALTH_NUMBER CanadaSpecific = "CA_HEALTH_NUMBER"
	// A Canadian Social Insurance Number (SIN) is a nine-digit unique identifier, required for individuals to access government programs and benefits.
	// Experimental.
	CanadaSpecific_CA_SOCIAL_INSURANCE_NUMBER CanadaSpecific = "CA_SOCIAL_INSURANCE_NUMBER"
)

