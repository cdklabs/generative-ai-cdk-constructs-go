package piitype


// Types of PII in the domain of Finance.
// Experimental.
type Finance string

const (
	// A three-digit card verification code (CVV) that is present on VISA, MasterCard, and Discover credit and debit cards.
	//
	// For American Express credit or debit cards,
	// the CVV is a four-digit numeric code.
	// Experimental.
	Finance_CREDIT_DEBIT_CARD_CVV Finance = "CREDIT_DEBIT_CARD_CVV"
	// The expiration date for a credit or debit card.
	//
	// This number is usually four digits
	// long and is often formatted as month/year or MM/YY. Guardrails recognizes expiration
	// dates such as 01/21, 01/2021, and Jan 2021.
	// Experimental.
	Finance_CREDIT_DEBIT_CARD_EXPIRY Finance = "CREDIT_DEBIT_CARD_EXPIRY"
	// The number for a credit or debit card.
	//
	// These numbers can vary from 13 to 16 digits
	// in length.
	// Experimental.
	Finance_CREDIT_DEBIT_CARD_NUMBER Finance = "CREDIT_DEBIT_CARD_NUMBER"
	// A four-digit personal identification number (PIN) with which you can access your bank account.
	// Experimental.
	Finance_PIN Finance = "PIN"
	// A SWIFT code is a standard format of Bank Identifier Code (BIC) used to specify a particular bank or branch.
	//
	// Banks use these codes for money transfers such as
	// international wire transfers. SWIFT codes consist of eight or 11 characters.
	// Experimental.
	Finance_SWIFT_CODE Finance = "SWIFT_CODE"
	// An International Bank Account Number (IBAN).
	//
	// It has specific formats in each country.
	// Experimental.
	Finance_INTERNATIONAL_BANK_ACCOUNT_NUMBER Finance = "INTERNATIONAL_BANK_ACCOUNT_NUMBER"
)

