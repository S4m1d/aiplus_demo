package bl

type BlErrorType int

const (
	BlDefaultError    BlErrorType = 0
	BlValidationError BlErrorType = 1
)

const (
	phoneNumberRegexpEnv = "PHONE_REGEXP"
)
