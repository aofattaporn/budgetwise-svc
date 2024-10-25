package customerrors

type CustomError struct {
	Code        int
	ErrorType   string
	Description string
}

func (e *CustomError) Error() string {
	return e.Description
}

// **** technical error ***
func DATA_NOT_FOUND(errorMsg string) *CustomError {
	return &CustomError{
		Code:        1699,
		ErrorType:   ERROR_TYPE().DATA_NOT_FOUND,
		Description: errorMsg,
	}
}

// **** invalid parameters error ***
func INVALID_PERAETERS_ERROR(errorMsg string) *CustomError {
	return &CustomError{
		Code:        1799,
		ErrorType:   ERROR_TYPE().INVALID_PARAMETER_ERROR,
		Description: errorMsg,
	}
}

// **** business error ***
func BUSINESS_ERROR(errorMsg string) *CustomError {
	return &CustomError{
		Code:        1899,
		ErrorType:   ERROR_TYPE().BUSINESS_ERROR,
		Description: errorMsg,
	}
}

// **** technical error ***
func TECHNICAL_ERROR(errorMsg string) *CustomError {
	return &CustomError{
		Code:        1999,
		ErrorType:   ERROR_TYPE().Technical_ERROR,
		Description: errorMsg,
	}
}

func FOREIGN_KEY_VIOLATION_ERROR(errorMsg string) *CustomError {
	return &CustomError{
		Code:        1999,
		ErrorType:   ERROR_TYPE().Technical_ERROR,
		Description: errorMsg,
	}
}

// **** transactions error ***
