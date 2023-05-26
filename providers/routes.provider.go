package providers

import (
	"main/controllers"
	"main/middlewares"
	"main/models"
	"main/models/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	database.InitDatabase()
	models.Migrate()
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//public routes
	public := r.Group("/api")
	public.POST("/signup", controllers.SignUp)
	public.POST("/login", controllers.Login)

	//protected routes (need token)
	protected := r.Group("/api/user")
	protected.Use(middlewares.JwtAuth())
	protected.POST("/current", controllers.CurrentUser)
	protected.POST("/update", controllers.UpdateUser)
	protected.POST("/update/resume", controllers.UpdateResume)
	return r
}
