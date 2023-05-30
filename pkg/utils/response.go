package utils

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

//In Go, the strings.Split() function is used to split a string into substrings based on a specified delimiter. It returns a slice of substrings.

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ErrorResponse(message string, err string, data interface{}) Response {
	splittedstring := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splittedstring,
		Data:    data,
	}
	return res
}

func SuccessResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   nil,
	}
	return res
}

func ResponseJSON(c gin.Context, data interface{}) {

	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(data)
}
