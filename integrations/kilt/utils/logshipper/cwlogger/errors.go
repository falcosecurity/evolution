package cwlogger

const (
	errCodeDataAlreadyAcceptedException  = "DataAlreadyAcceptedException"
	errCodeInvalidSequenceTokenException = "InvalidSequenceTokenException"
	errCodeThrottlingException           = "ThrottlingException"
	errCodeInternalFailure               = "InternalFailure"
	errCodeServiceUnavailable            = "ServiceUnavailable"
	errCodeServiceUnavailableException   = "ServiceUnavailableException"
)

var retryableErrorCodes = map[string]struct{}{
	errCodeInvalidSequenceTokenException: {},
	errCodeThrottlingException:           {},
	errCodeInternalFailure:               {},
	errCodeServiceUnavailable:            {},
	errCodeServiceUnavailableException:   {},
}

// Error contains the AWS error code and message that caused the PutLogEvents
// action to fail. Errors reported by the LogGroup ErrorReporter function may
// be converted into this type.
type Error struct {
	Code    string
	Message string
}

func (err Error) Error() string {
	if err.Message == "" {
		return err.Code
	}
	return err.Code + ": " + err.Message
}

func shouldRetry(err error) bool {
	if ownErr, ok := err.(Error); ok {
		_, found := retryableErrorCodes[ownErr.Code]
		return found
	}
	return true
}

func isErrorCode(err error, code string) bool {
	if ownErr, ok := err.(Error); ok {
		return ownErr.Code == code
	}
	return false
}

func noopErrorReporter(error) {}
