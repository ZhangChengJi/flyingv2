package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR        = 7
	SUCCESS      = 0
	UNAUTHORIZED = 401
)

type re interface {
	Result(code int, data interface{}, msg string, c *gin.Context)
}

func Result(StatusCode, code int, data interface{}, msg string, c *gin.Context) {

	// 开始时间
	c.JSON(StatusCode, Resp{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(http.StatusOK, SUCCESS, map[string]interface{}{}, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, data, "success", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(http.StatusOK, ERROR, map[string]interface{}{}, "error", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, ERROR, data, message, c)
}

func Unauthorized(message string, c *gin.Context) {
	Result(http.StatusUnauthorized, ERROR, map[string]interface{}{}, message, c)
}
