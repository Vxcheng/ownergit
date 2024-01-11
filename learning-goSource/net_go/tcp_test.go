package net_go

import (
	"fmt"
	"net"
	"testing"
)

func TestTcpServer(t *testing.T)  {
	t.Run("", func(t *testing.T) {
		go tcpServer()
		select {
		}
	})
}

func TestTcpClient(t *testing.T)  {
	t.Run("", func(t *testing.T) {
		go tcpClient()
		select {
		}
	})
}

func tcpServer() {
	listen_socket, err := net.Listen("tcp", "127.0.0.1:20000") //打开监听接口
	if err != nil {                                            //如果有错误
		fmt.Println("服务器连接失败")
		return
	}
	
	defer listen_socket.Close() //延迟服务器端关闭
	fmt.Println("服务器运行中....")
	
	for {
		conn, err := listen_socket.Accept() //监听客户端的端口
		if err != nil {
			fmt.Println("客户端开始失败")
		}
		fmt.Println("连接服务器成功") //显示服务器端连接成功
		
		var msg string //声明msg为字符串变量
		
		for {
			//开始接收客户端发过来的消息
			msg = ""                         //字符串msg初始为空
			data := make([]byte, 255)        //创建并声明数据变量，为255位
			msg_read, err := conn.Read(data) //接收由客户端发来的消息，字节赋值给msg_read，err为错误
			if msg_read == 0 || err != nil { //如果读取的消息为0字节或者有错误
				fmt.Println("err")
			}
			
			msg_read_str := string(data[0:msg_read]) //将msg_read_str的字节格式转化成字符串形式
			if msg_read_str == "close" {             //如果接收到的客户端消息为close
				conn.Write([]byte("close"))
				break
			}
			//fmt.Println(string(data[0:msg_read]))
			fmt.Println("客户端给服务器发送消息 ", msg_read_str) //接收客户端发来的信息
			
			fmt.Printf("对客户端发送消息: ") //提示向客户端要说的话
			fmt.Scan(&msg)                //输入服务器端要对客户端说的话
			//conn.Write([]byte("hello client\n"))
			//msg_write := []byte(msg)
			conn.Write([]byte(msg)) //把消息发送给客户端
			//此处造成服务器端的端口堵塞
		}
		//fmt.Println("client Close\n")
		conn.Close() //关闭连接
	}
}

func tcpClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	defer conn.Close()
	fmt.Println("连接成功")
	
	var msg string //声明msg为字符串变量
	
	for {
		msg = "" //初始化msg为空值
		fmt.Printf("给服务器发送消息: ")
		fmt.Scan(&msg) //输入客户端向服务器端要发送的消息
		//fmt.Println(msg)
		//msg_write := []byte(msg)
		//conn.Write(msg_write)
		conn.Write([]byte(msg)) //信息转化成字节流形式并向服务器端发送
		//此处造成客户端程序端口堵塞
		//fmt.Println([]byte(msg))
		
		//等待服务器端发送信息回来
		data := make([]byte, 255)
		msg_read, err := conn.Read(data)
		if msg_read == 0 || err != nil {
			fmt.Println("err")
		}
		msg_read_str := string(data[0:msg_read])
		if msg_read_str == "close" {
			conn.Write([]byte("close"))
			break
		}
		
		fmt.Println("服务器回应:", msg_read_str)
	}
	conn.Close()
}
