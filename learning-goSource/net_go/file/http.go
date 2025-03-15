package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type HttpFServer struct{}

// 处理文件上传
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// 限制上传文件大小为 10MB
	r.ParseMultipartForm(10 << 20)

	// 获取上传文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 创建目标文件
	dst, err := os.Create(handler.Filename + "_upload")
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// 获取文件名
	filename := r.URL.Query().Get("file")
	if filename == "" {
		http.Error(w, "File name is required", http.StatusBadRequest)
		return
	}

	// 构造文件路径
	filePath := filename

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 将文件内容写入响应
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
}

func (HttpFServer) main() (err error) {
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download", downloadHandler)
	fmt.Println("Server started at :10010")
	err = http.ListenAndServe(":10010", nil)
	println(err)
	return
}

type HttpFClient struct{}

func (c HttpFClient) main() {
	c.download()
	c.upload()
}

func (HttpFClient) upload() {
	// 文件路径
	filePath := "./test.txt" // 替换为你的测试文件路径

	// 创建请求体
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加文件字段
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(part, file)

	// 关闭 multipart writer
	writer.Close()

	// 发送 POST 请求
	resp, err := http.Post("http://localhost:10010/upload", writer.FormDataContentType(), body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %s\n", respBody)
}

func downloadFile(url, localFilePath string) error {
	// 发起HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error downloading file: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server responded with status code %d", resp.StatusCode)
	}

	// 创建本地文件
	err = os.MkdirAll(filepath.Dir(localFilePath), 0755)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	file, err := os.Create(localFilePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// 将响应内容写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func (HttpFClient) download() {
	url := "http://localhost:10010/download?file=test.txt"
	localFilePath := "./downloads_test.txt"

	if err := downloadFile(url, localFilePath); err != nil {
		fmt.Println("Error downloading file:", err)
		return
	}

	fmt.Println("File downloaded successfully.")
}
