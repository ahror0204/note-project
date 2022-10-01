package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/note_project/pkg/logger"
	"github.com/note_project/pkg/structures"
)

// @Summary Create user
// @Description This API for creating user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body structures.UserStruct true "user"
// @Success 200 {object} structures.UserStruct
// @Router /v1/create_user/ [post]
func (h handlerV1) CreateUser(c *gin.Context) {
	var req structures.UserStruct

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
	}

	resp, err := h.userStorage.CreateUser(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating user", logger.Error(err))
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary Update user
// @Description This API for updating user
// @Tags Users
// @Accept json
// @Produce json
// @Param update user body structures.UserStruct true "update user"
// @Success 200 {string} Success
// @Router /v1/update_user/ [post]
func (h handlerV1) UpdateUser(c *gin.Context) {
	var req structures.UserStruct

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
	}

	err = h.userStorage.UpdateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating user", logger.Error(err))
	}

	c.JSON(http.StatusOK, "Updated")
}

// @Summary Delete user
// @Description This API for deleting user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {string} Success
// @Router /v1/delete_user/{id} [delete]
func (h handlerV1) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	
	err := h.userStorage.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting user", logger.Error(err))
		c.JSON(http.StatusOK, "Deleted")
	}
}
