package helper

import "time"

//General Wrapper to wrap response JSON
type GeneralWrapper struct {
	StatusCode string      `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	TimeStamp  time.Time   `json:"timeStamp"`
}

type ErrorMessage struct {
	ErrorMessage string `json:"errorMessage"`
}
