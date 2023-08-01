package Middlewares

import "github.com/gin-gonic/gin"

// ExampleMiddleware
/*
 create middlewares with Closure syntax
*/
func ExampleMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {}
}
