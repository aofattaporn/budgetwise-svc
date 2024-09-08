package constants

type route struct {
	HEALTHCHECK string
}

var routeConstantVar = &route{
	HEALTHCHECK: "/health",
}

func ROUTE() *route {
	return routeConstantVar
}
