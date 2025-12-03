package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

func main() {
	// 配置 etcd 客户端，包含认证信息
	cfg := clientv3.Config{
		Endpoints:   []string{"192.168.10.1:2379"}, // etcd 节点地址
		DialTimeout: 5 * time.Second,               // 连接超时时间
		Username:    "root",                        // etcd 用户名
		Password:    "root1234",                    // etcd 密码
	}

	// 创建 etcd 客户端
	client, err := clientv3.New(cfg)
	if err != nil {
		log.Fatalf("创建 etcd 客户端失败: %v", err)
	}
	defer client.Close()

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Get(ctx, "test")
	if err != nil {
		log.Fatalf("连接测试失败: %v", err)
	}
	fmt.Println("成功连接到 etcd 服务器")

	// 使用分布式锁
	useDistributedLock(client, "my-lock-key")
}

func useDistributedLock(client *clientv3.Client, lockKey string) {
	// 创建会话（Session），会话超时时间设置为 30 秒
	session, err := concurrency.NewSession(client, concurrency.WithTTL(30))
	if err != nil {
		log.Fatalf("创建会话失败: %v", err)
	}
	defer session.Close()

	// 创建互斥锁
	mutex := concurrency.NewMutex(session, lockKey)

	// 获取锁的上下文（带有超时）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 尝试获取锁
	fmt.Printf("尝试获取锁: %s\n", lockKey)
	if err := mutex.Lock(ctx); err != nil {
		log.Fatalf("获取锁失败: %v", err)
	}
	fmt.Println("成功获取锁，开始执行临界区操作")

	// 模拟临界区操作
	fmt.Println("执行关键业务逻辑...")
	time.Sleep(5 * time.Second)

	// 释放锁
	if err := mutex.Unlock(ctx); err != nil {
		log.Fatalf("释放锁失败: %v", err)
	}
	fmt.Println("锁已释放")
}
