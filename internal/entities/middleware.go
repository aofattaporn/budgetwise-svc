package entities

type LoggerRequestAndResponse struct {
	Method    string `json:"method"`
	Path      string `json:"path"`
	Req       LogReq `json:"req"`
	Res       LogRes `json:"res"`
	LatencyMS string `json:"latencyMS"`
}
type LogHeaders map[string]string
type LogParams map[string]string
type LogReq struct {
	Headers LogHeaders `json:"headers"`
	Params  LogParams  `json:"params"`
	Body    string     `json:"body"`
	Time    string     `json:"time"`
}
type LogRes struct {
	Status  int        `json:"status"`
	Headers LogHeaders `json:"headers"`
	Body    string     `json:"body"`
	Time    string     `json:"time"`
}
