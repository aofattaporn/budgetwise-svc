package constants

type route struct {
	HEALTHCHECK string
	SALARY      string
}

func ROUTE() *route {
	return &route{
		HEALTHCHECK: "/health",
		SALARY:      "/users/salary",
	}

}
