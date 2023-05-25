package providers

import (
	"main/controllers"
	"main/middlewares"
	"main/models"
	"main/models/database"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	database.InitDatabase()
	models.Migrate()
	r := gin.Default()
	r.Use(middlewares.Cors())
	public := r.Group("/api")

	public.POST("/signup", controllers.SignUp)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/user")
	protected.Use(middlewares.JwtAuth())
	protected.POST("/current", controllers.CurrentUser)
	protected.POST("/update", controllers.UpdateUser)
	return r
}
