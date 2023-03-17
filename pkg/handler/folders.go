package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createFolder(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllFolders(c *gin.Context) {

}

func (h *Handler) getFolderById(c *gin.Context) {

}

func (h *Handler) updateFolder(c *gin.Context) {

}

func (h *Handler) deleteFolder(c *gin.Context) {

}
