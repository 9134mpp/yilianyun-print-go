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

func (p *PrintServer) PrintDeleteVoice(ctx context.Context, request *pd.PrintDeleteVoiceRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintDeletePrinter(ctx context.Context, request *pd.PrintDeletePrinterRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintAddPrintMenu(ctx context.Context, request *pd.PrintAddPrintMenuRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintShutDownRestart(ctx context.Context, request *pd.PrintShutDownRestartRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintSetSound(ctx context.Context, request *pd.PrintSetSoundRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintInfo(ctx context.Context, request *pd.PrintInfoRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintCancelAll(ctx context.Context, request *pd.PrintCancelAllRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintCaneLone(ctx context.Context, request *pd.PrintCaneLoneRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintSetIcon(ctx context.Context, request *pd.PrintSetIconRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintDeleteIcon(ctx context.Context, request *pd.PrintDeleteIconRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintBtnPrint(ctx context.Context, request *pd.PrintBtnPrintRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintGetOrder(ctx context.Context, request *pd.PrintGetOrderRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintSetPushUrl(ctx context.Context, request *pd.PrintSetPushUrlRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintGetOrderStatus(ctx context.Context, request *pd.PrintGetOrderStatusRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintGetOrderPagingList(ctx context.Context, request *pd.PrintGetOrderPagingListRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PrintGetPrintStatus(ctx context.Context, request *pd.PrintGetPrintStatusRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) PicturePrint(ctx context.Context, request *pd.PicturePrintRequest) (*pd.PrintReply, error) {
	panic("implement me")
}

func (p *PrintServer) ExpressPrint(ctx context.Context, request *pd.ExpressPrintRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.ExpressPrint(ctx, request)
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

func (p *PrintServer) PrintSetVoice(ctx context.Context, request *pd.PrintSetVoiceRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintSetVoice(ctx, request)
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
