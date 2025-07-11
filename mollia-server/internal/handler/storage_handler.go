package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StorageHandler struct{}

func NewStorageHandler() *StorageHandler {
	return &StorageHandler{}
}

func (h *StorageHandler) UploadFile(c *gin.Context) {
	// 从 multipart-form 中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败: " + err.Error()})
		return
	}

	// 生成唯一文件名
	ext := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + ext
	dst := filepath.Join("uploads", newFileName)

	// 保存文件到服务器
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "无法保存文件: " + err.Error()})
		return
	}

	// 返回可访问的 URL
	// TODO: URL 应从配置中动态生成
	fileUrl := "http://localhost:8080/static/" + newFileName

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "上传成功",
		"data": gin.H{"url": fileUrl},
	})
}
