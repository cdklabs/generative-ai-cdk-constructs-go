package piitype

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PIIType.CanadaSpecific",
		reflect.TypeOf((*CanadaSpecific)(nil)).Elem(),
		map[string]interface{}{
			"CA_HEALTH_NUMBER": CanadaSpecific_CA_HEALTH_NUMBER,
			"CA_SOCIAL_INSURANCE_NUMBER": CanadaSpecific_CA_SOCIAL_INSURANCE_NUMBER,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PIIType.Finance",
		reflect.TypeOf((*Finance)(nil)).Elem(),
		map[string]interface{}{
			"CREDIT_DEBIT_CARD_CVV": Finance_CREDIT_DEBIT_CARD_CVV,
			"CREDIT_DEBIT_CARD_EXPIRY": Finance_CREDIT_DEBIT_CARD_EXPIRY,
			"CREDIT_DEBIT_CARD_NUMBER": Finance_CREDIT_DEBIT_CARD_NUMBER,
			"PIN": Finance_PIN,
			"SWIFT_CODE": Finance_SWIFT_CODE,
			"INTERNATIONAL_BANK_ACCOUNT_NUMBER": Finance_INTERNATIONAL_BANK_ACCOUNT_NUMBER,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PIIType.General",
		reflect.TypeOf((*General)(nil)).Elem(),
		map[string]interface{}{
			"ADDRESS": General_ADDRESS,
			"AGE": General_AGE,
			"DRIVER_ID": General_DRIVER_ID,
			"EMAIL": General_EMAIL,
			"LICENSE_PLATE": General_LICENSE_PLATE,
			"NAME": General_NAME,
			"PASSWORD": General_PASSWORD,
			"PHONE": General_PHONE,
			"USERNAME": General_USERNAME,
			"VEHICLE_IDENTIFICATION_NUMBER": General_VEHICLE_IDENTIFICATION_NUMBER,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PIIType.InformationTechnology",
		reflect.TypeOf((*InformationTechnology)(nil)).Elem(),
		map[string]interface{}{
			"URL": InformationTechnology_URL,
			"IP_ADDRESS": InformationTechnology_IP_ADDRESS,
			"MAC_ADDRESS": InformationTechnology_MAC_ADDRESS,
			"AWS_ACCESS_KEY": InformationTechnology_AWS_ACCESS_KEY,
			"AWS_SECRET_KEY": InformationTechnology_AWS_SECRET_KEY,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PIIType.UKSpecific",
		reflect.TypeOf((*UKSpecific)(nil)).Elem(),
		map[string]interface{}{
			"UK_NATIONAL_HEALTH_SERVICE_NUMBER": UKSpecific_UK_NATIONAL_HEALTH_SERVICE_NUMBER,
			"UK_NATIONAL_INSURANCE_NUMBER": UKSpecific_UK_NATIONAL_INSURANCE_NUMBER,
			"UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER": UKSpecific_UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PIIType.USASpecific",
		reflect.TypeOf((*USASpecific)(nil)).Elem(),
		map[string]interface{}{
			"US_BANK_ACCOUNT_NUMBER": USASpecific_US_BANK_ACCOUNT_NUMBER,
			"US_BANK_ROUTING_NUMBER": USASpecific_US_BANK_ROUTING_NUMBER,
			"US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER": USASpecific_US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER,
			"US_PASSPORT_NUMBER": USASpecific_US_PASSPORT_NUMBER,
			"US_SOCIAL_SECURITY_NUMBER": USASpecific_US_SOCIAL_SECURITY_NUMBER,
		},
	)
}
