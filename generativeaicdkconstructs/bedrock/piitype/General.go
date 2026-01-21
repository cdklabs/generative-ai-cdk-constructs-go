package piitype


// Types of PII that are general, and not domain-specific.
// Experimental.
type General string

const (
	// A physical address, such as "100 Main Street, Anytown, USA" or "Suite #12, Building 123".
	//
	// An address can include information such as the street, building,
	// location, city, state, country, county, zip code, precinct, and neighborhood.
	// Experimental.
	General_ADDRESS General = "ADDRESS"
	// An individual's age, including the quantity and unit of time.
	// Experimental.
	General_AGE General = "AGE"
	// The number assigned to a driver's license, which is an official document permitting an individual to operate one or more motorized vehicles on a public road.
	//
	// A driver's license number consists of alphanumeric characters.
	// Experimental.
	General_DRIVER_ID General = "DRIVER_ID"
	// An email address, such as marymajor@email.com.
	// Experimental.
	General_EMAIL General = "EMAIL"
	// A license plate for a vehicle is issued by the state or country where the vehicle is registered.
	//
	// The format for passenger vehicles is typically five
	// to eight digits, consisting of upper-case letters and numbers. The format
	// varies depending on the location of the issuing state or country.
	// Experimental.
	General_LICENSE_PLATE General = "LICENSE_PLATE"
	// An individual's name.
	//
	// This entity type does not include titles, such as Dr.,
	//  Mr., Mrs., or Miss.
	// Experimental.
	General_NAME General = "NAME"
	// An alphanumeric string that is used as a password, such as "*very20special#pass*".
	// Experimental.
	General_PASSWORD General = "PASSWORD"
	// A phone number.
	//
	// This entity type also includes fax and pager numbers.
	// Experimental.
	General_PHONE General = "PHONE"
	// A user name that identifies an account, such as a login name, screen name, nick name, or handle.
	// Experimental.
	General_USERNAME General = "USERNAME"
	// A Vehicle Identification Number (VIN) uniquely identifies a vehicle.
	//
	// VIN
	// content and format are defined in the ISO 3779 specification. Each country
	// has specific codes and formats for VINs.
	// Experimental.
	General_VEHICLE_IDENTIFICATION_NUMBER General = "VEHICLE_IDENTIFICATION_NUMBER"
)

