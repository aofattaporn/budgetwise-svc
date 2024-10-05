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
	TRANSACTIONS    string
	TRANSACTIONS_ID string

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
		TRANSACTIONS:    "/transactions",
		TRANSACTIONS_ID: "/transactions/:id",

		// plans
		PLANS:    "/plans",
		PLANS_ID: "/plans/:id",
	}
}
