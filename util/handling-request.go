package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_JSON_UNMARSHAL    = "JSON_UNMARSHAL"
	ERR_CODE_JWT_TOKEN_INVALID = "JWT_TOKEN_INVALID"
)

// ErrorResponse represents an error response in your API.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

var (
	BadRequestResponse       = ErrorResponse{400, "Bad request"}
	BadRequestResponseCustom = func(message string) ErrorResponse {
		return ErrorResponse{400, message}
	}
	InternalServerErrorResponse = ErrorResponse{500, "Internal server error"}
)

func BadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
}

func BadRequestWithMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": msg})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"message": "internal server error"})
}

func InternalServerErrorMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": msg})
}
func Ok(c *gin.Context, msg gin.H) {
	c.JSON(http.StatusOK, msg)
}

func OkMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func UnauthorizedMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": msg})
}

func BadReqWithDetail(c *gin.Context, detail string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status_code": http.StatusBadRequest,
		"error_code":  400,
		"detail":      detail,
	})
}

func Abort(c *gin.Context, httpStatusCode int, errorCode, message interface{}) {
	resp := gin.H{
		"status_code": httpStatusCode,
		"error_code":  errorCode,
		"message":     message,
	}

	c.AbortWithStatusJSON(httpStatusCode, resp)
}

func AbortUnauthorized(c *gin.Context, errorCode, message interface{}) {
	Abort(c, http.StatusUnauthorized, errorCode, message)
}

func AbortInternalServerError(c *gin.Context, errorCode, message interface{}) {
	Abort(c, http.StatusInternalServerError, errorCode, message)
}

func AbortJSONBadRequest(c *gin.Context) {
	Abort(c, http.StatusBadRequest, ERR_CODE_JSON_UNMARSHAL, "Can't parse data")
}

func ResponseOk(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
