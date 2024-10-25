package customerrors

type errorType struct {
	INVALID_PARAMETER_ERROR string
	BUSINESS_ERROR          string
	Technical_ERROR         string
	DATA_NOT_FOUND          string
}

func ERROR_TYPE() *errorType {
	return &errorType{
		INVALID_PARAMETER_ERROR: "INVALID_PARAMETER_ERROR",
		BUSINESS_ERROR:          "BUSINESS_ERROR",
		Technical_ERROR:         "Technical_ERROR",
		DATA_NOT_FOUND:          "DATA_NOT_FOUND",
	}
}
