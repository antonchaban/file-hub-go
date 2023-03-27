package handler

import (
	fhub "github.com/antonchaban/file-hub-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create folder
// @Security ApiKeyAuth
// @Tags folders
// @Description create folder
// @ID create-folder
// @Accept  json
// @Produce  json
// @Param input body fhub.UpdateFolderInput true "folder info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/folders [post]
func (h *Handler) createFolder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input fhub.Folder
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Folder.CreateFolder(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllFoldersResponse struct {
	Data []fhub.Folder `json:"data"`
}

// @Summary Get all folders
// @Security ApiKeyAuth
// @Tags folders
// @Description get all folders
// @ID get-all-folders
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllFoldersResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/folders [get]
func (h *Handler) getAllFolders(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	folders, err := h.services.Folder.GetAllFolders(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllFoldersResponse{
		Data: folders,
	})

}

// @Summary Get folder by id
// @Security ApiKeyAuth
// @Tags folders
// @Description get folder by id
// @ID get-folder-by-id
// @Accept  json
// @Produce  json
// @Param folder_id path int true "folder id"
// @Success 200 {object} fhub.Folder
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/folders/{folder_id} [get]
func (h *Handler) getFolderById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("folder_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid folder id param")
		return
	}

	folder, err := h.services.Folder.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, folder)
}

// @Summary Update folder
// @Security ApiKeyAuth
// @Tags folders
// @Description update folder
// @ID update-folder
// @Accept  json
// @Produce  json
// @Param folder_id path int true "folder id"
// @Param input body fhub.UpdateFolderInput true "folder info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/folders/{folder_id} [put]
func (h *Handler) updateFolder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("folder_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid folder id param")
		return
	}

	var input fhub.UpdateFolderInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete folder
// @Security ApiKeyAuth
// @Tags folders
// @Description delete folder
// @ID delete-folder
// @Accept  json
// @Produce  json
// @Param folder_id path int true "folder id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/folders/{folder_id} [delete]
func (h *Handler) deleteFolder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("folder_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid folder id param")
		return
	}

	err = h.services.Folder.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
