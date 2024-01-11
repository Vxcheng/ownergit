package net_go

import (
	"fmt"
	"net"
	"testing"
)

func TestUdpServer(t *testing.T)  {
	t.Run("", func(t *testing.T) {
		go udpServer()
		select {

		}
	})
}

func TestUdpClient(t *testing.T)  {
	t.Run("", func(t *testing.T) {
		go udpClient()
		select {

		}
	})
}


func Porecss(listener *net.UDPConn) {
	//退出是关闭资源
	defer listener.Close()
	
	var b string
	for {
		//获取连接输出
		var buf [1024]byte
		
		n, addr, err := listener.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("读取udp失败err=", err)
			return
		}
		//信息输出
		fmt.Printf("来自%v的回复,内容是%v\n", addr, string(buf[:n]))
		
		//信息回复
		fmt.Println("给客户端回复")
		fmt.Scan(&b)
		_, err = listener.WriteToUDP([]byte(b), addr)
		if err != nil {
			fmt.Println("回复失败err=", err)
			return
		}
		
	}
}
func udpServer() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.ParseIP("0.0.0.0"),
		
		Port: 38080,
	})
	if err != nil {
		fmt.Println("连接异常，err=", err)
	}
	Porecss(listener)
	
}

func udpClient() {
	conn, err := net.Dial("udp", "127.0.0.1:38081")
	
	if err != nil {
		fmt.Println("连接服务器失败,err:", err)
		return
	}
	
	defer conn.Close()
	//发送消息
	
	var a string
	for {
		fmt.Println("输入要发送的信息")
		fmt.Scan(&a)
		_, err = conn.Write([]byte(a))
		if err != nil {
			fmt.Println("信息发送失败,err", err)
			return
		}
		//接受消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("读取失败,err", err)
			return
		}
		fmt.Println("收到服务端的信息是:", string(buf[:n]))
	}
	
}
