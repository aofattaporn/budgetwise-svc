package constants

type route struct {
	HEALTHCHECK string

	// user route definind
	SALARY string

	// accounts
	ACCOUNTS           string
	ACCOUNTS_AMOUNT    string
	ACCOUNTS_NAME      string
	ACCOUNTS_ACCOUNTID string

	// transactions
	TRANSACTIONS string

	// plans
	PLANS    string
	PLANS_ID string
}

func ROUTE() *route {
	return &route{
		HEALTHCHECK: "/health",
		SALARY:      "/users/salary",
		ACCOUNTS:    "/accounts",

		// accounts
		ACCOUNTS_ACCOUNTID: "/accounts/:id",

		// transactions
		TRANSACTIONS: "/transactions",
		PLANS:        "/plans",
		PLANS_ID:     "/plans/:id",
	}
}
