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

type IsBody struct {
	Error string `json:"error"`
}

type ErrBody struct {
	Error string `json:"error"`
	Body string `json:"body"`
	ErrorDescription string `json:"error_description"`
}

type SuccessBody struct {
	Error string `json:"error"`
	Body pd.Body `json:"body"`
	ErrorDescription string `json:"error_description"`
}

type ListBody struct{
	ListBody  []*pd.Body `json:"body"`
}
// 获取Body数据
func getBody(b []byte) (*pd.PrintReply, error) {
	isBody := IsBody{}
	err := json.Unmarshal(b, &isBody)
	// 订单列表接口数据解析
	listBody := ListBody{}
	_ = json.Unmarshal(b, &listBody)
	if err != nil{
		return nil, err
	}
	p := pd.PrintReply{}
	success := "0"
	if isBody.Error != success{
		errBody := ErrBody{}
		_ = json.Unmarshal(b, &errBody)
		p = pd.PrintReply{
			Error: errBody.Error,
			Body: &pd.PrintReply_ErrorBody{ErrorBody: errBody.Body},
			ErrorDescription: errBody.ErrorDescription,
		}
		return &p, nil
	}
	successBody := SuccessBody{}
	_ = json.Unmarshal(b, &successBody)
	p = pd.PrintReply{
		Error: isBody.Error,
		Body: &pd.PrintReply_SuccessBody{SuccessBody: &successBody.Body},
		ErrorDescription: successBody.ErrorDescription,
		ListBody: listBody.ListBody,
	}
	return &p, nil
}

func (p *PrintServer) PrintAddPrinter(ctx context.Context, request *pd.PrintAddPrinterRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintAddPrinter(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintDeleteVoice(ctx context.Context, request *pd.PrintDeleteVoiceRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintDeleteVoice(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintDeletePrinter(ctx context.Context, request *pd.PrintDeletePrinterRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintDeletePrinter(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintAddPrintMenu(ctx context.Context, request *pd.PrintAddPrintMenuRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintAddPrintMenu(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintShutDownRestart(ctx context.Context, request *pd.PrintShutDownRestartRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintShutDownRestart(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintSetSound(ctx context.Context, request *pd.PrintSetSoundRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintSetSound(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintInfo(ctx context.Context, request *pd.PrintInfoRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintInfo(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintCancelAll(ctx context.Context, request *pd.PrintCancelAllRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintCancelAll(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintCaneLone(ctx context.Context, request *pd.PrintCaneLoneRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintCaneLone(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintSetIcon(ctx context.Context, request *pd.PrintSetIconRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintSetIcon(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintDeleteIcon(ctx context.Context, request *pd.PrintDeleteIconRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintDeleteIcon(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintBtnPrint(ctx context.Context, request *pd.PrintBtnPrintRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintBtnPrint(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintGetOrder(ctx context.Context, request *pd.PrintGetOrderRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintGetOrder(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintSetPushUrl(ctx context.Context, request *pd.PrintSetPushUrlRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintSetPushUrl(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintGetOrderStatus(ctx context.Context, request *pd.PrintGetOrderStatusRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintGetOrderStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintGetOrderPagingList(ctx context.Context, request *pd.PrintGetOrderPagingListRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintGetOrderPagingList(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintGetPrintStatus(ctx context.Context, request *pd.PrintGetPrintStatusRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintGetPrintStatus(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PicturePrint(ctx context.Context, request *pd.PicturePrintRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PicturePrint(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) ExpressPrint(ctx context.Context, request *pd.ExpressPrintRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.ExpressPrint(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil
}

func (p *PrintServer) PrintSetVoice(ctx context.Context, request *pd.PrintSetVoiceRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.PrintSetVoice(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)
	if err != nil {
		return nil, err
	}
	return printReply, nil
}

func (p *PrintServer) GetForeignToken(ctx context.Context, request *pd.ForeignOauthRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.GetForeignToken(ctx, request)
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return printReply, nil

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
	printReply, err := getBody(body)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	success := "0"
	if printReply.Error != success{
		return printReply, nil
	}
	setting.Update("Client", "AccessToken", printReply.GetSuccessBody().AccessToken) //修改配置文件中的access_token
	return printReply, nil
}

func (p *PrintServer) Print(ctx context.Context, r *pd.PrintRequest) (*pd.PrintReply, error) {
	api := bapi.NewAPI(ApiUrl)
	body, err := api.Print(ctx, r) // 调用打印
	if err != nil {
		return nil, err
	}
	printReply, err := getBody(body)

	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}

	return printReply, nil
}
