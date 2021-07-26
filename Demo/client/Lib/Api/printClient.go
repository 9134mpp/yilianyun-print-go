package Api

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pd "yilianyun-print-go/Lib/pkg/proto"
)

type PrintParam struct {
	AccessToken string
	MachineCode string
	Content     string
}

// 文本打印
func Print(client pd.PrintServiceClient, p *pd.PrintRequest) error {
	resp, _ := client.Print(context.Background(), p)
	success := "0"
	if resp.Error != success {
		log.Printf("client.Print resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.Print resp: %+v", resp)
	return nil
}

// 获取token
func GetToken(client pd.PrintServiceClient, p *pd.OauthRequest) (string, error) {
	resp, _ := client.GetToken(context.Background(), p)
	success := "0"
	if resp.Error != success {
		log.Printf("client.GetToken resp: %+v", resp.ErrorDescription)
		return "", nil
	}
	log.Printf("client.GetToken resp: %+v", resp)
	return resp.Body.AccessToken, nil

}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
