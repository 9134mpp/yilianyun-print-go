package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pd "yilianyun-print-go/Lib/pkg/proto"
)

func main(){
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "localhost:9523",nil)
	defer clientConn.Close()

	client := pd.NewPrintServiceClient(clientConn)
	_ = Print(client)
}

func Print(client pd.PrintServiceClient) error {
	p := &pd.GetPrintRequest{
		AccessToken: "c779a83d842f4b5eb587a7491ea7d155",
		MachineCode: "4004628158",
		Content: "测试",
		OriginId: "123",
	}
	resp, _ := client.Print(context.Background(), p)
	success := "0"
	if resp.Error != success{
		log.Printf("client.Print resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.Print resp: %+v", resp)
	return nil
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
