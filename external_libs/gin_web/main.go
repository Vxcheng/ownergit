package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 文件上传接口
	r.POST("/upload", func(c *gin.Context) {
		// 获取上传的文件
		header, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 获取文件名
		filename := header.Filename
		uploadPath := "./uploads/" + filename
		c.SaveUploadedFile(header, uploadPath)

		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("File %s uploaded successfully", filename),
			"path":    uploadPath,
		})
	})

	// 文件下载接口
	r.GET("/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filePath := "./uploads/" + filename

		// 检查文件是否存在
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}

		// 设置文件下载的 HTTP 头
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Expires", "0")
		c.Header("Cache-Control", "must-revalidate")
		c.Header("Pragma", "public")

		// 发送文件内容
		c.File(filePath)
	})

	// 启动服务器
	log.Fatal(r.Run(":8080"))
}
