package handler

import (
	"github.com/antonchaban/file-hub-go/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		folders := api.Group("/folders")
		{
			folders.GET("/", h.getAllFolders)
			folders.POST("/", h.createFolder)
			folders.GET("/:folder_id", h.getFolderById)
			folders.PUT("/:folder_id", h.updateFolder)
			folders.DELETE("/:folder_id", h.deleteFolder)
		}

		files := folders.Group("/:folder_id/files")
		{
			files.GET("/", h.getAllFiles)
			files.POST("/", h.createFile)
			files.GET("/:file_id", h.getFileById)
			files.PUT("/:file_id", h.updateFile)
			files.DELETE("/:file_id", h.deleteFile)
		}
	}
	return router
}
