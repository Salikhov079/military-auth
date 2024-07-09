package api

import (
	"root/api/handler"
	_ "root/docs"

	_ "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Voting service
// @version 1.0
// @description Voting service
// @host localhost:8081
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization

func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	u := r.Group("/user")
	u.POST("/login", h.LoginUser)
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))
	return r
}
