package domain

const (
	// Default config constants:
	DEFAULT_SERVER_ADDRESS       = "127.0.0.1:8080"
	DEFAULT_ORIGINS_ADDRESS_CORS = "http://127.0.0.1:8080"
	DEFAULT_LOG_LEVEL            = "debug"
	DEFAULT_LOG_FILE_PATH        = "logs/serverRestApi/logs.log"

	// For the logger layers strings
	LAYER_MIDDLEWARE_STRING_FOR_LOGGER = "middleware"
	LAYER_DELIVERY_STRING_FOR_LOGGER   = "delivery"
	LAYER_USECASE_STRING_FOR_LOGGER    = "usecase"
	LAYER_REPOSITORY_STRING_FOR_LOGGER = "repository"

	REQUEST_ID_KEY_FOR_CONTEXT = stringTypeAlias("requestID")

	// For the session
	SESSION_COOKIE_HEADER_NAME = "session_id" // Cookie header
	SESSION_ID_STRING_LENGTH   = 32
)

type stringTypeAlias string
