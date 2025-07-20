package apierrors

//noolingo gen

var (
	apiErrors = map[string]*APIError{
		"Forbidden":            {name: "Forbidden", code: 403, message: "permission denied"},
		"InternalServerError":  {name: "InternalServerError", code: 500, message: "internal server error"},
		"InvalidPayload":       {name: "InvalidPayload", code: 400, message: "invalid payload"},
		"NotFound":             {name: "NotFound", code: 404, message: "not found"},
		"NotImplemented":       {name: "NotImplemented", code: 400, message: "route not implemented"},
		"PasswordsNotEqual":    {name: "PasswordsNotEqual", code: 400, message: "passwords not equal"},
		"ServiceDisabled":      {name: "ServiceDisabled", code: 400, message: "service disabled"},
		"ServiceNotConfigured": {name: "ServiceNotConfigured", code: 400, message: "service not configured"},
		"TokenExpired":         {name: "TokenExpired", code: 401, message: "token expired"},
		"Unauthorized":         {name: "Unauthorized", code: 401, message: "unauthorized"},
		"UnknownType":          {name: "UnknownType", code: 400, message: "unknown type"},
		"UserExists":           {name: "UserExists", code: 409, message: "user already exists"},
		"UserIdNotSpecified":   {name: "UserIdNotSpecified", code: 400, message: "user id not specified"},
		"UserNotFoundError":    {name: "UserNotFoundError", code: 404, message: "sendmail service error"},
	}

	// ErrForbidden code: 403 message: permission denied
	ErrForbidden = apiErrors["Forbidden"].Error()
	// ErrInternalServerError code: 500 message: internal server error
	ErrInternalServerError = apiErrors["InternalServerError"].Error()
	// ErrInvalidPayload code: 400 message: invalid payload
	ErrInvalidPayload = apiErrors["InvalidPayload"].Error()
	// ErrNotFound code: 404 message: not found
	ErrNotFound = apiErrors["NotFound"].Error()
	// ErrNotImplemented code: 400 message: route not implemented
	ErrNotImplemented = apiErrors["NotImplemented"].Error()
	// ErrPasswordsNotEqual code: 400 message: passwords not equal
	ErrPasswordsNotEqual = apiErrors["PasswordsNotEqual"].Error()
	// ErrServiceDisabled code: 400 message: service disabled
	ErrServiceDisabled = apiErrors["ServiceDisabled"].Error()
	// ErrServiceNotConfigured code: 400 message: service not configured
	ErrServiceNotConfigured = apiErrors["ServiceNotConfigured"].Error()
	// ErrTokenExpired code: 401 message: token expired
	ErrTokenExpired = apiErrors["TokenExpired"].Error()
	// ErrUnauthorized code: 401 message: unauthorized
	ErrUnauthorized = apiErrors["Unauthorized"].Error()
	// ErrUnknownType code: 400 message: unknown type
	ErrUnknownType = apiErrors["UnknownType"].Error()
	// ErrUserExists code: 409 message: user already exists
	ErrUserExists = apiErrors["UserExists"].Error()
	// ErrUserIdNotSpecified code: 400 message: user id not specified
	ErrUserIdNotSpecified = apiErrors["UserIdNotSpecified"].Error()
	// ErrUserNotFoundError code: 404 message: sendmail service error
	ErrUserNotFoundError = apiErrors["UserNotFoundError"].Error()
)
