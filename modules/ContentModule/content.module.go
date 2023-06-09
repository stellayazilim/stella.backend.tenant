package ContentModule

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitContentModule(router *gin.RouterGroup) {
	router.StaticFS("", http.Dir("public/elitas"))
}
