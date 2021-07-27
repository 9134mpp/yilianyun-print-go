package server

import (
	"context"
	"encoding/json"
	"log"
	"yilianyun-print-go/Lib/bapi"
	"yilianyun-print-go/Lib/errcode"
	pd "yilianyun-print-go/Lib/proto"
	"yilianyun-print-go/Lib/setting"
)

var ApiUrl string

func Setup(){
	ApiUrl = setting.ApiSetting.Url
	if ApiUrl == ""{
		log.Printf("初始化ApiUrl失败！")
	}
}
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

func (p *PrintServer) GetForeignToken(ctx context.Context, request *pd.ForeignOauthRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.GetForeignToken(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply := pd.PrintReply{}
	err = json.Unmarshal(body, &printReply)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return &printReply, nil

}

func NewPrintServer() *PrintServer {
	return &PrintServer{}
}

func (p *PrintServer) GetToken(ctx context.Context, r *pd.OauthRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
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
	api := bapi.NewAPI(ApiUrl)
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
