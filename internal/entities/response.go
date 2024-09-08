package entities

type Response struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}
