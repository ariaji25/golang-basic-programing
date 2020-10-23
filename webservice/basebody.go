package webservice

type BaseBody struct {
	Status  int     `json:"status"`
	Data    float64 `json:"data"`
	Message string  `json:"message"`
	IsError bool    `json:"isError"`
}
