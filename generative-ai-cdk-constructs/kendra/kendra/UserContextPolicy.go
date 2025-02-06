package kendra


// The different policies available to filter search results based on user context.
// Experimental.
type UserContextPolicy string

const (
	// All indexed content is searchable and displayable for all users.
	//
	// If you want to filter search results on user context, you can use
	// the attribute filters of _user_id and _group_ids or you can provide
	// user and group information in UserContext .
	// Experimental.
	UserContextPolicy_ATTRIBUTE_FILTER UserContextPolicy = "ATTRIBUTE_FILTER"
	// Enables token-based user access control to filter search results on user context.
	//
	// All documents with no access control and all documents
	// accessible to the user will be searchable and displayable.
	// Experimental.
	UserContextPolicy_USER_TOKEN UserContextPolicy = "USER_TOKEN"
)

