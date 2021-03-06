package http

// Response body
type Response struct {
	Error error
	Data  interface{}
}

// BaseResponse body
type BaseResponse struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
