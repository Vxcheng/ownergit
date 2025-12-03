package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// 官方示例
const host = "192.168.10.1"

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{host + ":2379"},
		DialTimeout: 20 * time.Second,
		Username:    "root",
		Password:    "root1234",
	})
	if err != nil {
		// handle error!
		log.Fatal(err)
	}
	defer cli.Close()
	_ = put(cli)

}

func put(cli *clientv3.Client) (err error) {

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	_, err = cli.Put(ctx, "/lmh", "lmh")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	resp, err := cli.Get(ctx, "/lmh")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}

	return
}
