package webservice

/* Create Base Response for standarize the Response Body */
type BaseResponse struct {
	Status  int     `json:"status"`
	Data    float64 `json:"data"`
	Message string  `json:"message"`
	IsError bool    `json:"isError"`
}
