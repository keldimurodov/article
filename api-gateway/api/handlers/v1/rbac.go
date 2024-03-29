package v1

import (
	"encoding/json"
	"net/http"
	"projects/article/api-gateway/api/handlers/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePost ...
// @Summary CreatePost
// @Security ApiKeyAuth
// @Description Api for creating a new user
// @Tags rbac
// @Accept json
// @Produce json
// @Param body body models.CreateUserRoleRequest true "body"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/create_role/ [post]
func (h *handlerV1) CreateRole(ctx *gin.Context) {

	var reqBody models.CreateUserRoleRequest

	err := json.NewDecoder(ctx.Request.Body).Decode(&reqBody)

	if err != nil {
		h.log.Error("invalid request body", zap.Error(err))
		return
	}

	_, err = h.enforcer.AddRoleForUser(reqBody.Role, reqBody.Path, reqBody.Metod)

	if err != nil {
		h.log.Error("error on grantAccess", zap.Error(err))
		return
	}
	h.enforcer.SavePolicy()
	ctx.JSON(http.StatusCreated, reqBody.Role)
}
