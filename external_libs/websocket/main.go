package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// 定义全局变量：Upgrader 和客户端映射
var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 允许所有跨域请求，生产环境应严格限制
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// 用于广播消息给所有客户端
	clients   = make(map[*websocket.Conn]bool) // 已连接的客户端
	broadcast = make(chan []byte)              // 广播频道
)

func main() {
	// 1. 设置静态文件服务（用于提供 HTML 客户端页面）
	http.Handle("/", http.FileServer(http.Dir("./public")))

	// 2. 配置 WebSocket 路由
	http.HandleFunc("/ws", handleWebSocket)

	// 3. 启动一个单独的 Goroutine 来处理广播消息
	go handleBroadcast()

	// 4. 启动 HTTP 服务器
	log.Println("Server starting on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 1. 升级 HTTP 连接到 WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade failed: %v", err)
		return
	}
	defer conn.Close() // 确保函数退出时关闭连接

	// 2. 注册新客户端
	clients[conn] = true
	log.Printf("Client %s connected. Total clients: %d", conn.RemoteAddr(), len(clients))

	// 3. 设置连接关闭时的处理（清理客户端列表）
	defer func() {
		delete(clients, conn)
		log.Printf("Client %s disconnected. Total clients: %d", conn.RemoteAddr(), len(clients))
	}()

	// 4. 设置读超时（可选，但很重要）
	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	// 5. 循环读取客户端发送的消息
	for {
		// ReadMessage 读取下一个文本或二进制消息
		_, message, err := conn.ReadMessage()
		if err != nil {
			// 检查是否是正常关闭（如客户端断开）
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Read error: %v", err)
			}
			break // 退出循环，关闭连接
		}

		log.Printf("Received: %s", message)

		// 6. 将收到的消息放入广播频道
		broadcast <- message
	}
}

func handleBroadcast() {
	for {
		// 从广播频道中取出消息
		message := <-broadcast

		// 遍历所有已连接的客户端，发送消息
		for client := range clients {
			log.Printf("Sending to client %s: %s", client.RemoteAddr(), message)
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Write error to client %s: %v", client.RemoteAddr(), err)
				client.Close()
				delete(clients, client) // 发送失败，移除客户端
			}
		}
	}
}
