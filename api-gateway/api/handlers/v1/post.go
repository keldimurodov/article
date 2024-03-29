package v1

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	models "projects/article/api-gateway/api/handlers/models"
	p "projects/article/api-gateway/genproto/post"
	l "projects/article/api-gateway/pkg/logger"
	"projects/article/api-gateway/pkg/utils"
)

// CreatePost ...
// @Summary CreatePost
// @Security ApiKeyAuth
// @Description Api for creating a new user
// @Tags post
// @Accept json
// @Produce json
// @Param Post body models.Post true "createPostModel"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/ [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body        models.Post
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Create(ctx, &p.Post{
		Id:      body.Id,
		Picture: body.Picture,
		Title:   body.Title,
		Article: body.Article,
		OwnerId: body.OwnerId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetPost gets post by id
// @Summary GetPost by id
// @Security ApiKeyAuth
// @Description Api for getting post by id
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/{id} [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	id64, err := strconv.Atoi(id)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Get(
		ctx, &p.GetPostRequest{
			Id: int64(id64),
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllPosts returns list of posts
// @Summary GetAllPosts
// @Security ApiKeyAuth
// @Description Api for getting post by page and limit
// @Tags post
// @Accept json
// @Produce json
// @Param page path string true "PAGE"
// @Param limit path string true "LIMIT"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/ [get]
func (h *handlerV1) GetAllPosts(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetAll(
		ctx, &p.GetAllRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdatePost updates user by id
// @Summary UpdatePost
// @Security ApiKeyAuth
// @Description Api for updating post by id
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/{id} [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        p.Post
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}


	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Update(ctx, &p.Post{
		Id:      body.Id,
		Title:   body.Title,
		Picture: body.Picture,
		OwnerId: body.OwnerId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeletePost deletes post by id
// @Summary DeletePost
// @Security ApiKeyAuth
// @Description Api for deleting post by id
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/posts/{id} [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	id64, err := strconv.Atoi(id)

	if err != nil{
		log.Fatal("Error converting")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Delete(ctx, &p.GetPostRequest{
		Id: int64(id64),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
