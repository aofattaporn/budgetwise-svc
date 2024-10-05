package customerrors

type errorType struct {
	INVALID_PARAMETER_ERROR string
	BUSINESS_ERROR          string
	Technical_ERROR         string
}

func ERROR_TYPE() *errorType {
	return &errorType{
		INVALID_PARAMETER_ERROR: "INVALID_PARAMETER_ERROR",
		BUSINESS_ERROR:          "BUSINESS_ERROR",
		Technical_ERROR:         "Technical_ERROR",
	}
}
