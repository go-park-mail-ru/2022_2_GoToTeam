package constants

const (
	// For the grpc metadata keys
	REQUEST_ID_KEY_FOR_METADATA = "requestID"
	USER_EMAIL_KEY_FOR_METADATA = "email"

	// For the context keys
	REQUEST_ID_KEY_FOR_CONTEXT = stringTypeAlias("requestID")
	USER_EMAIL_KEY_FOR_CONTEXT = stringTypeAlias("email")
)

type stringTypeAlias string
