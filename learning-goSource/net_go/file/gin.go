package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type GinServer struct{}

func (GinServer) uploadFile(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存文件
	savePath := file.Filename + "_upload"
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "filename": file.Filename})
}

func (GinServer) downloadFile(c *gin.Context) {
	// 获取文件名
	filename := c.Param("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Filename is required"})
		return
	}

	// 构造文件路径
	filePath := filename

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// 设置响应头
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Transfer-Encoding", "binary")

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// 将文件内容写入响应
	io.Copy(c.Writer, file)
}

func (s GinServer) main() {
	r := gin.Default()

	// 文件上传路由
	r.POST("/upload", s.uploadFile)

	// 文件下载路由
	r.GET("/download/:filename", s.downloadFile)

	// 启动服务器
	r.Run(":10011")
}

/*
curl -X POST -F "file=@/path/to/your/file.txt" http://localhost:10011/upload
curl -O http://localhost:10011/download/yourfile.txt
*/
