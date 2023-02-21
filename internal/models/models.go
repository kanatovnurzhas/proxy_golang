package models

type RequestProxy struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type ResponseProxy struct {
	Id      int                 `json:"id"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int                 `json:"length"`
}

type KeyRequest struct {
	Method  string
	Url     string
	Headers string
}
