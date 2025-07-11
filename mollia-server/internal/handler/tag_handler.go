package handler

import (
	"mollia/internal/model"
	"mollia/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	svc service.TagService
}

func NewTagHandler(svc service.TagService) *TagHandler {
	return &TagHandler{svc: svc}
}

func (h *TagHandler) CreateTag(c *gin.Context) {
	var req model.TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数无效: " + err.Error()})
		return
	}
	id, err := h.svc.Create(c.Request.Context(), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"id": id}})
}

func (h *TagHandler) GetTags(c *gin.Context) {
	tags, err := h.svc.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": tags})
}

func (h *TagHandler) UpdateTag(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "无效的ID"})
		return
	}
	var req model.TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数无效: " + err.Error()})
		return
	}
	if err := h.svc.Update(c.Request.Context(), uint(id), req.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "更新成功"})
}

func (h *TagHandler) DeleteTag(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "无效的ID"})
		return
	}
	if err := h.svc.Delete(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功"})
}
