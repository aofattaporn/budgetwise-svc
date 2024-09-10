package constants

type status struct {
	SUCCESS string
}

func STATUS() *status {
	return &status{
		SUCCESS: "success",
	}
}
