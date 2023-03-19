package handler

import (
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create file
// @Security ApiKeyAuth
// @Tags files
// @Description create file
// @ID create-file
// @Accept  json
// @Produce  json
// @Param folder_id path int true "folder id"
// @Param input body fhub.File true "file info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/folders/{folder_id}/files [post]
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

// @Summary Get all files
// @Security ApiKeyAuth
// @Tags files
// @Description get all files
// @ID get-all-files
// @Accept  json
// @Produce  json
// @Param folder_id path int true "folder id"
// @Success 200 {object} []fhub.File
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/folders/{folder_id}/files [get]
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

// @Summary Get file by id
// @Security ApiKeyAuth
// @Tags files
// @Description get file by id
// @ID get-file-by-id
// @Accept  json
// @Produce  json
// @Param file_id path int true "file id"
// @Success 200 {object} fhub.File
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/files/{file_id} [get]
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

// @Summary Update file
// @Security ApiKeyAuth
// @Tags files
// @Description update file
// @ID update-file
// @Accept  json
// @Produce  json
// @Param file_id path int true "file id"
// @Param input body fhub.UpdateFileInput true "file info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/files/{file_id} [put]
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

// @Summary Delete file
// @Security ApiKeyAuth
// @Tags files
// @Description delete file
// @ID delete-file
// @Accept  json
// @Produce  json
// @Param file_id path int true "file id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/files/{file_id} [delete]
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
