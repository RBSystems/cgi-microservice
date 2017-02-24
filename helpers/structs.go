package helpers

//Body represents a POST payload
type Body struct {
	Req   Request   `json:"req"`
	Param Parameter `json:"param"`
}

//Request represents the request in the body
type Request struct {
	Type     string `json:"type"`
	Category string `json:"category"`
}

//Parameter represents a parameter in the body
type Parameter struct {
	Category string `json:"category"`
	Code     string `json:"code"`
}
