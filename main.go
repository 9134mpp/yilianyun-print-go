package main

import (
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	pd "yilianyun-print-go/Lib/pkg/proto"
	"yilianyun-print-go/Lib/server"

	"log"
)

var grpcPort string
var httpPort string

func init (){
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
