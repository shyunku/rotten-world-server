package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/shyunku-libraries/go-logger"
)

func ping(c *gin.Context) {
	c.String(200, "pong")
}

func UseGlobalRouter() *gin.Engine {
	gin.DefaultWriter = log.GetLogger()
	gin.DefaultErrorWriter = log.GetLogger()

	// setting cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	r := gin.Default()
	r.Use(cors.New(config))
	r.GET("/ping", ping)

	return r
}
