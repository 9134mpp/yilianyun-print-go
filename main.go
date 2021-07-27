package main

import (
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"yilianyun-print-go/Lib/common"
	pd "yilianyun-print-go/Lib/proto"
	"yilianyun-print-go/Lib/server"
	"yilianyun-print-go/Lib/setting"

	"log"
)

var grpcPort string
var httpPort string

func init (){
	setting.Setup() //初始化配置
	common.Setup()  //初始化常用方法配置
	server.Setup()  //初始化Api配置
	flag.StringVar(&grpcPort, "grpc_port", "9523", "gRPC 启动端口号")
	flag.StringVar(&httpPort, "http_port", "9001", "HTTP 启动端口号")
	flag.Parse()
}
func main(){
	s := grpc.NewServer()
	pd.RegisterPrintServiceServer(s, server.NewPrintServer())
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":"+grpcPort)

	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}

func RunHttpServer(port string) error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return http.ListenAndServe(":"+port, serveMux)
}

func RunGrpcServer(port string) error {
	s := grpc.NewServer()
	pd.RegisterPrintServiceServer(s, server.NewPrintServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	return s.Serve(lis)
}
