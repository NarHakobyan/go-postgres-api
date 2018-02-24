package response

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func HandleFunc(handler func(*Context)) func(*gin.Context) {
	return func(c *gin.Context) {
		customContext := Context{c}
		handler(&customContext)
	}
}

func (context *Context) Ok(message string, body interface{}) {
	context.Send(http.StatusOK, message, body)
}

func (context *Context) Created(message string, body interface{}) {
	context.Send(http.StatusCreated, message, body)
}

func (context *Context) NotFound(message string, body interface{}) {
	context.Send(http.StatusNotFound, message, body)
}

func (context *Context) BadRequest(message string, body interface{}) {
	context.Send(http.StatusBadRequest, message, body)
}

func (context *Context) InternalServerError(message string, body interface{}) {
	context.Send(http.StatusInternalServerError, message, body)
}

func (context *Context) UnprocessableEntity(message string, body interface{}) {
	bodyType := reflect.TypeOf(body)
	fmt.Println(bodyType.Kind())

	switch body.(type) {
	case []error:
		var errs []string
		errors := body.([]error)
		for _, e := range errors {
			errs = append(errs, e.Error())
		}
		body = errs
	}
	context.Send(http.StatusUnprocessableEntity, message, body)
}

func (context *Context) Send(status int, message string, body interface{}) {
	data := gin.H{
		"statusText": http.StatusText(status),
		"statusCode": status,
		"body":       body,
		"message":    message,
	}

	context.JSON(status, data)
}
