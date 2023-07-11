package util

import (
	"github.com/gin-gonic/gin"
)

type Responses struct {
	Version    string            `json:"version"`
	TraceId    string            `json:"traceId"`
	StatusCode int               `json:"statusCode"`
	Message    string            `json:"message,omitempty"`
	Data       any               `json:"data,omitempty"`
	Meta       map[string]string `json:"meta,omitempty"`
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  int    `json:"errorCode"`
	Message    string `json:"message"`
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Data any, Meta map[string]string) {

	jsonResponse := Responses{
		Version:    "1.0.0",
		TraceId:    "1-1-1-1",
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
		Meta:       Meta,
	}

	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func APIErrorResponse(ctx *gin.Context, StatusCode int, ErrorCode int, Message string) {
	errResponse := ErrorResponse{
		StatusCode: StatusCode,
		ErrorCode:  ErrorCode,
		Message:    Message,
	}

	ctx.JSON(StatusCode, errResponse)
	defer ctx.AbortWithStatus(StatusCode)
}

// func getStatusCode(err error) int {
// 	if err == nil {
// 		return http.StatusOK
// 	}

// 	logrus.Error(err)
// 	switch err {
// 	case domain.ErrInternalServerError:
// 		return http.StatusInternalServerError
// 	case domain.ErrNotFound:
// 		return http.StatusNotFound
// 	case domain.ErrConflict:
// 		return http.StatusConflict
// 	default:
// 		return http.StatusInternalServerError
// 	}
// }
