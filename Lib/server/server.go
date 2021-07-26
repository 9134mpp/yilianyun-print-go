package server

import (
	"context"
	"encoding/json"
	"yilianyun-print-go/Lib/pkg/bapi"
	"yilianyun-print-go/Lib/pkg/errcode"
	pd "yilianyun-print-go/Lib/pkg/proto"
	"yilianyun-print-go/Lib/pkg/setting"

)

type PrintServer struct{}

func (p *PrintServer) PicturePrint(ctx context.Context, request *pd.PicturePrintRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) ExpressPrint(ctx context.Context, request *pd.ExpressPrintRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintSetVoice(ctx context.Context, request *pd.PrintSetVoiceRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) GetOpenToken(ctx context.Context, request *pd.OpenOauthRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func NewPrintServer() *PrintServer {
	return &PrintServer{}
}

func (p *PrintServer) GetToken(ctx context.Context, r *pd.OauthRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI("https://open-api.10ss.net/")
	body, err := api.GetToken(ctx, r)
	if err != nil {
		return nil, err
	}
	printReply := pd.PrintReply{}
	err = json.Unmarshal(body, &printReply)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	success := "0"
	if printReply.Error != success{
		setting.Update("Client", "AccessToken", printReply.Body.AccessToken) //修改配置文件中的access_token
	}
	return &printReply, nil
}

func (p *PrintServer) Print(ctx context.Context, r *pd.PrintRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI("https://open-api.10ss.net/")
	body, err := api.Print(ctx, r) // 调用打印
	if err != nil {
		return nil, err
	}
	printReply := pd.PrintReply{}
	err = json.Unmarshal(body, &printReply)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.NewError(10000009, err.Error()))
	}

	return &printReply, nil
}
