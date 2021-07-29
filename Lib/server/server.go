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
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintDeleteVoice(ctx, request)
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

func (p *PrintServer) PrintDeletePrinter(ctx context.Context, request *pd.PrintDeletePrinterRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintDeletePrinter(ctx, request)
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

func (p *PrintServer) PrintAddPrintMenu(ctx context.Context, request *pd.PrintAddPrintMenuRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintAddPrintMenu(ctx, request)
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

func (p *PrintServer) PrintShutDownRestart(ctx context.Context, request *pd.PrintShutDownRestartRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintShutDownRestart(ctx, request)
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

func (p *PrintServer) PrintSetSound(ctx context.Context, request *pd.PrintSetSoundRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintSetSound(ctx, request)
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

func (p *PrintServer) PrintInfo(ctx context.Context, request *pd.PrintInfoRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintInfo(ctx, request)
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

func (p *PrintServer) PrintCancelAll(ctx context.Context, request *pd.PrintCancelAllRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintCancelAll(ctx, request)
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

func (p *PrintServer) PrintCaneLone(ctx context.Context, request *pd.PrintCaneLoneRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintCaneLone(ctx, request)
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

func (p *PrintServer) PrintSetIcon(ctx context.Context, request *pd.PrintSetIconRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintSetIcon(ctx, request)
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

func (p *PrintServer) PrintDeleteIcon(ctx context.Context, request *pd.PrintDeleteIconRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintDeleteIcon(ctx, request)
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

func (p *PrintServer) PrintBtnPrint(ctx context.Context, request *pd.PrintBtnPrintRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintBtnPrint(ctx, request)
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

func (p *PrintServer) PrintGetOrder(ctx context.Context, request *pd.PrintGetOrderRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintGetOrder(ctx, request)
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

func (p *PrintServer) PrintSetPushUrl(ctx context.Context, request *pd.PrintSetPushUrlRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintSetPushUrl(ctx, request)
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

func (p *PrintServer) PrintGetOrderStatus(ctx context.Context, request *pd.PrintGetOrderStatusRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintGetOrderStatus(ctx, request)
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

func (p *PrintServer) PrintGetOrderPagingList(ctx context.Context, request *pd.PrintGetOrderPagingListRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintGetOrderPagingList(ctx, request)
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

func (p *PrintServer) PrintGetPrintStatus(ctx context.Context, request *pd.PrintGetPrintStatusRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintGetPrintStatus(ctx, request)
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

func (p *PrintServer) PicturePrint(ctx context.Context, request *pd.PicturePrintRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PicturePrint(ctx, request)
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

type PrintErrorReply struct {
	Error  string `json:"error"`
	ErrorDescription string `json:"error_description"`

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
		return &printReply, nil
	}
	setting.Update("Client", "AccessToken", printReply.Body.AccessToken) //修改配置文件中的access_token
	setting.Setup() // 重新初始化配置
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
		return nil, errcode.TogRPCError(errcode.Fail)
	}

	return &printReply, nil
}
