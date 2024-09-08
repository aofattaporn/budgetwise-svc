package constants

type common struct {
	EMPTY_STRING    string
	SYSYEM          string
	DATETIME_FORMAT string
}

var commonConstantVar = &common{
	EMPTY_STRING:    "",
	SYSYEM:          "System",
	DATETIME_FORMAT: "2006-01-02 15:04:05.000",
}

func COMMON() *common {
	return commonConstantVar
}
