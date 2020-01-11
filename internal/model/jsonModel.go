package model

type JsonResult struct {
	Status int 			`json:"status"`
	Data   interface{} 	`json:"data"`
}
type JsonResultError struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error"`
}