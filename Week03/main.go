package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出
func main() {
	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	group, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello Go")
	})

	// http server
	group.Go(func() error {
		server := http.Server{
			Addr:    ":8080",
			Handler: mux,
		}
		go func() {
			<-ctx.Done()
			fmt.Println("接受到别的程序出错，准备退出http server程序")
			server.Shutdown(context.Background())
		}()
		return server.ListenAndServe()
	})

	// 监听 linux signal
	group.Go(func() error {
		select {
		case <-q:
			return errors.New("接受到退出信息")
		case <-ctx.Done():
			fmt.Println("接受到别的程序出错，准备退出监听程序")
			return errors.New("退出程序")
		}
	})

	// 测试
	//group.Go(func() error {
	//	time.Sleep(time.Second * 5)
	//	return errors.New("测试")
	//})

	if err := group.Wait(); err != nil {
		fmt.Println("group wait接收到有个程序错误：", err.Error())
		fmt.Println("退出主程序")
	}

}
