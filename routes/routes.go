package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marcellribeiro/tunity_test/controllers"
)

func Start() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.GET("/collatz/:number", controllers.GetCollatz)
	router.Run(":3000")
}
