package handler

// gin-swagger middleware
// swagger embed files
import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/antonchaban/file-hub-go/docs"
)

type Handler struct {
	Authorization
	Folder
	File
}

func NewHandler(authorization Authorization, folder Folder, file File) *Handler {
	return &Handler{Authorization: authorization, Folder: folder, File: file}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

			files := folders.Group("/:folder_id/files")
			{
				files.GET("/", h.getAllFiles)
				files.POST("/", h.createFile)
			}
		}
		files := api.Group("/files")
		{
			files.GET("/:file_id", h.getFileById)
			files.PUT("/:file_id", h.updateFile)
			files.DELETE("/:file_id", h.deleteFile)
		}

	}
	return router
}
