package main

import (
	"context"
	"fmt"
	"yilianyun-print-go/Demo/client/Lib/Api"
	pd "yilianyun-print-go/Lib/pkg/proto"
	"yilianyun-print-go/Lib/pkg/setting"
)

func main() {
	// 自有应用示例子
	have()
	//todo 开放应用示例
}

func have() {
	ctx := context.Background()
	clientConn, _ := Api.GetClientConn(ctx, "localhost:9523", nil)
	defer clientConn.Close()
	client := pd.NewPrintServiceClient(clientConn)
	//授权
	token, _ := Api.GetToken(client, &pd.OauthRequest{}) // 自有应用access_token时间永久
	fmt.Println(token)

	//打印
	p := &pd.PrintRequest{
		AccessToken: setting.ClientSetting.AccessToken,
		MachineCode: "4004628158",
		Content:     "测试",
	}
	_ = Api.Print(client, p)
}
