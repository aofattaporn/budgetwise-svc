package constants

type route struct {
	HEALTHCHECK string

	// user route definind
	SALARY string

	// accounts
	ACCOUNTS string

	// transactions
	TRANSACTIONS string

	// plans
	PLANS string
}

func ROUTE() *route {
	return &route{
		HEALTHCHECK:  "/health",
		SALARY:       "/users/salary",
		ACCOUNTS:     "/accounts",
		TRANSACTIONS: "/transactions",
		PLANS:        "/plans",
	}
}
