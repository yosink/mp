package provider

type ErrorResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
