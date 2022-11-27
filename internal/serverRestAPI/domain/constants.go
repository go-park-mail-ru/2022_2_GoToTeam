package domain

const (
	// For the logger layers strings
	LAYER_MIDDLEWARE_STRING_FOR_LOGGER = "middleware"
	LAYER_DELIVERY_STRING_FOR_LOGGER   = "delivery"
	LAYER_USECASE_STRING_FOR_LOGGER    = "usecase"
	LAYER_REPOSITORY_STRING_FOR_LOGGER = "repository"

	// For the context keys
	REQUEST_ID_KEY_FOR_CONTEXT = stringTypeAlias("requestID")
	USER_EMAIL_KEY_FOR_CONTEXT = stringTypeAlias("email")

	// For the session
	SESSION_COOKIE_HEADER_NAME = "session_id" // Cookie header
	SESSION_ID_STRING_LENGTH   = 32
)

type stringTypeAlias string
