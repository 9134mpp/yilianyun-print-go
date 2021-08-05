package Api

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pd "yilianyun-print-go/Lib/proto"
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
	if resp.GetError() != success {
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
	if resp.GetError() != success {
		log.Printf("client.GetToken resp: %+v", resp.ErrorDescription)
		return "", nil
	}
	log.Printf("client.GetToken resp: %+v", resp)
	return resp.GetSuccessBody().AccessToken, nil

}
// 开放应用获取token
func GetForeignToken(client pd.PrintServiceClient, p *pd.ForeignOauthRequest)(string, error){
	resp, _ := client.GetForeignToken(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.GetForeignToken resp: %+v", resp.ErrorDescription)
		return "", nil
	}
	log.Printf("client.GetForeignToken resp: %+v", resp)
	return resp.GetSuccessBody().AccessToken, nil
}

// 图形打印
func PicturePrint(client pd.PrintServiceClient, p *pd.PicturePrintRequest) error {
	resp, _ := client.PicturePrint(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PicturePrint resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PicturePrint resp: %+v", resp)
	return nil
}

// 面单接口
func ExpressPrint(client pd.PrintServiceClient, p *pd.ExpressPrintRequest) error {
	resp, _ := client.ExpressPrint(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.ExpressPrint resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.ExpressPrint resp: %+v", resp)
	return nil
}

// 设置内置语音接口
func PrintSetVoice(client pd.PrintServiceClient, p *pd.PrintSetVoiceRequest) error {
	resp, _ := client.PrintSetVoice(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintSetVoice resp: %+v", resp.GetErrorBody())
		return nil
	}
	log.Printf("client.PrintSetVoice resp: %+v", resp)
	return nil
}

// 删除内置语音接口
func PrintDeleteVoice(client pd.PrintServiceClient, p *pd.PrintDeleteVoiceRequest) error {
	resp, _ := client.PrintDeleteVoice(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintDeleteVoice resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintDeleteVoice resp: %+v", resp)
	return nil
}

// 删除终端授权
func PrintDeletePrinter(client pd.PrintServiceClient, p *pd.PrintDeletePrinterRequest) error {
	resp, _ := client.PrintDeletePrinter(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintDeletePrinter resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintDeletePrinter resp: %+v", resp)
	return nil
}

// 添加应用菜单
func PrintAddPrintMenu(client pd.PrintServiceClient, p *pd.PrintAddPrintMenuRequest) error {
	resp, _ := client.PrintAddPrintMenu(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintAddPrintMenu resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintAddPrintMenu resp: %+v", resp)
	return nil
}

// 关机重启接口
func PrintShutDownRestart(client pd.PrintServiceClient, p *pd.PrintShutDownRestartRequest) error {
	resp, _ := client.PrintShutDownRestart(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintShutDownRestart resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintShutDownRestart resp: %+v", resp)
	return nil
}
// 声音调节接口
func PrintSetSound(client pd.PrintServiceClient, p *pd.PrintSetSoundRequest) error {
	resp, _ := client.PrintSetSound(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintSetSound resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintSetSound resp: %+v", resp)
	return nil
}

// 获取机型打印宽度接口
func PrintInfo(client pd.PrintServiceClient, p *pd.PrintInfoRequest) error {
	resp, _ := client.PrintInfo(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintInfo resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintInfo resp: %+v", resp)
	return nil
}

// 取消所有未打印订单
func PrintCancelAll(client pd.PrintServiceClient, p *pd.PrintCancelAllRequest) error {
	resp, _ := client.PrintCancelAll(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintCancelAll resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintCancelAll resp: %+v", resp)
	return nil
}

// 取消单条未打印订单
func PrintCaneLone(client pd.PrintServiceClient, p *pd.PrintCaneLoneRequest) error {
	resp, _ := client.PrintCaneLone(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintCaneLone resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintCaneLone resp: %+v", resp)
	return nil
}

// 设置logo接口
func PrintSetIcon(client pd.PrintServiceClient, p *pd.PrintSetIconRequest) error {
	resp, _ := client.PrintSetIcon(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintSetIcon resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintSetIcon resp: %+v", resp)
	return nil
}

// 取消logo接口
func PrintDeleteIcon(client pd.PrintServiceClient, p *pd.PrintDeleteIconRequest) error {
	resp, _ := client.PrintDeleteIcon(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintDeleteIcon resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintDeleteIcon resp: %+v", resp)
	return nil
}

// 打印方式接口

func PrintBtnPrint(client pd.PrintServiceClient, p *pd.PrintBtnPrintRequest) error {
	resp, _ := client.PrintBtnPrint(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintBtnPrint resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintBtnPrint resp: %+v", resp)
	return nil
}

// 接单拒单设置接口
func PrintGetOrder(client pd.PrintServiceClient, p *pd.PrintGetOrderRequest) error {
	resp, _ := client.PrintGetOrder(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintGetOrder resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintGetOrder resp: %+v", resp)
	return nil
}

// 设置推送url接口
func PrintSetPushUrl(client pd.PrintServiceClient, p *pd.PrintSetPushUrlRequest) error {
	resp, _ := client.PrintSetPushUrl(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintSetPushUrl resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintSetPushUrl resp: %+v", resp)
	return nil
}

// 获取订单状态接口
func PrintGetOrderStatus(client pd.PrintServiceClient, p *pd.PrintGetOrderStatusRequest) error {
	resp, _ := client.PrintGetOrderStatus(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintGetOrderStatus resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintGetOrderStatus resp: %+v", resp)
	return nil
}

// 获取订单列表接口
func PrintGetOrderPagingList(client pd.PrintServiceClient, p *pd.PrintGetOrderPagingListRequest) error {
	resp, _ := client.PrintGetOrderPagingList(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintGetOrderPagingList resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintGetOrderPagingList resp: %+v", resp)
	return nil
}

//	获取终端状态接口
func PrintGetPrintStatus(client pd.PrintServiceClient, p *pd.PrintGetPrintStatusRequest) error {
	resp, _ := client.PrintGetPrintStatus(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintGetPrintStatus resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintGetPrintStatus resp: %+v", resp)
	return nil
}

// 终端授权 (永久授权)
func PrintAddPrinter(client pd.PrintServiceClient, p *pd.PrintAddPrinterRequest) error {
	resp, _ := client.PrintAddPrinter(context.Background(), p)
	success := "0"
	if resp.GetError() != success {
		log.Printf("client.PrintAddPrinter resp: %+v", resp.ErrorDescription)
		return nil
	}
	log.Printf("client.PrintAddPrinter resp: %+v", resp)
	return nil
}

















// 获取连接
func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
