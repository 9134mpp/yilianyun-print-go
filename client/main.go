package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
	"yilianyun-print-go/Lib/pkg/common"
	pd "yilianyun-print-go/Lib/pkg/proto"
	"yilianyun-print-go/Lib/pkg/setting"
)

func main(){
	setting.Setup()
	common.Setup()
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "localhost:9523",nil)
	defer clientConn.Close()

	client := pd.NewPrintServiceClient(clientConn)
	_ = Print(client)
}

func Print(client pd.PrintServiceClient) error {
	timestamp := strconv.FormatInt(time.Now().Unix(),10)
	strTime, err:= strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		log.Println(err)
	}
	p := &pd.GetPrintRequest{
		ClientId: setting.ClientSetting.ClientId,
		AccessToken: "c779a83d842f4b5eb587a7491ea7d155",
		MachineCode: "4004628158",
		Content: "测试",
		OriginId: "123",
		Sign: common.GetSign(timestamp),
		Id:common.GetUUID4(),
		Timestamp: strTime,
	}
	log.Printf("%v\n", p)
	resp, _ := client.Print(context.Background(), p)
	log.Printf("client.SayHello resp: %+v", resp.ErrorDescription)
	return nil
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
