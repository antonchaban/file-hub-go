package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	api := router.Group("/api")
	{
		folders := api.Group("/folders")
		{
			folders.GET("/")
			folders.POST("/")
			folders.GET("/:folder_id")
			folders.PUT("/:folder_id")
			folders.DELETE("/:folder_id")
		}

		files := folders.Group("/:folder_id/files")
		{
			files.GET("/")
			files.POST("/")
			files.GET("/:file_id")
			files.PUT("/:file_id")
			files.DELETE("/:file_id")
		}
	}
	return router
}
