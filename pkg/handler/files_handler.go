package handler

import (
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createFile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	folderId, err := strconv.Atoi(c.Param("folder_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid folder id param")
		return
	}

	var input fhub.File
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.File.CreateFile(userId, folderId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllFiles(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	folderId, err := strconv.Atoi(c.Param("folder_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid folder id param")
		return
	}

	files, err := h.services.File.GetAllFiles(userId, folderId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, files)
}

func (h *Handler) getFileById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	fileId, err := strconv.Atoi(c.Param("file_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid file id param")
		return
	}

	file, err := h.services.File.GetFileById(userId, fileId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, file)
}

func (h *Handler) updateFile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("file_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid file id param")
		return
	}

	var input fhub.UpdateFileInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.File.UpdateFile(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteFile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	fileId, err := strconv.Atoi(c.Param("file_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid file id param")
		return
	}

	err = h.services.File.DeleteFile(userId, fileId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
