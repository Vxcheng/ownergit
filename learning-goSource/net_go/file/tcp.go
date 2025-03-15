package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
)

type TcpServer struct {
}

type TcpClient struct {
}

func (s TcpServer) main() (err error) {
	ln, err := net.Listen("tcp", ":10012")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())

		go s.handleConn(conn)
	}
}

func (TcpServer) handleConn(conn net.Conn) (err error) {
	defer conn.Close()
	buf := make([]byte, 1024)

	// 读取文件名
	var destFileName string
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading filename:", err)
		return
	}
	if n >= len(buf) {
		fmt.Println("Filename too long")
		return
	}
	destFileName = string(buf[:n])

	var fileSize int64
	if err = binary.Read(conn, binary.LittleEndian, &fileSize); err != nil {
		return
	}

	fmt.Printf("Receiving file: %s, Size: %d\n", destFileName, fileSize)
	file, err := os.Create(destFileName)
	if err != nil {
		return
	}
	defer file.Close()

	total := int64(0)
	for total < fileSize {
		n, err = conn.Read(buf)
		if err != nil {
			return err
		}

		total += int64(n)
		if _, err = file.Write(buf[:n]); err != nil {
			return err
		}
	}
	conn.Write([]byte("ack"))
	return
}

func (TcpClient) main() (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:10012")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	filePath := "test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fileSize := fileInfo.Size()

	if _, err = conn.Write([]byte(filePath)); err != nil {
		return
	}
	if err = binary.Write(conn, binary.LittleEndian, fileSize); err != nil {
		return
	}
	fmt.Printf("filePath: %s, File size: %d\n", filePath, fileSize)

	buf := make([]byte, 1024)
	for {
		var n int
		n, err = file.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			var ack [3]byte
			if _, err = conn.Read(ack[:]); err != nil {
				return
			}
			fmt.Printf("File upload complete. %v\n", string(ack[:]))
			return
		}

		if _, err = conn.Write(buf[:n]); err != nil {
			return
		}
	}
}
