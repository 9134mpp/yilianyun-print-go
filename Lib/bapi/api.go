package bapi

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"yilianyun-print-go/Lib/common"

	"yilianyun-print-go/Lib/setting"

	pd "yilianyun-print-go/Lib/proto"
)

type API struct {
	URL string
}

type Response struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Body             Body   `json:"body"`
}
type Body struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	Scope        string `json:"scope"`
}

func NewAPI(url string) *API {
	return &API{URL: url}
}

// 自有应用获取token
func (a *API) GetToken(ctx context.Context, r *pd.OauthRequest) ([]byte, error) {
	data := make(url.Values)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["scope"] = []string{"all"}
	data["timestamp"] = []string{timestamp}
	data["id"] = []string{common.GetUUID4()}
	data["grant_type"] = []string{"client_credentials"}
	body, err := a.httpPost("oauth/oauth", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 文本打印接口
func (a *API) Print(ctx context.Context, r *pd.PrintRequest) ([]byte, error) {
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["machine_code"] = []string{r.MachineCode}
	data["content"] = []string{r.Content}
	data["origin_id"] = []string{common.GetOrderId()}
	data["timestamp"] = []string{timestamp}
	data["id"] = []string{common.GetUUID4()}
	body, err := a.httpPost("print/index", data)

	if err != nil {
		return nil, err
	}
	return body, nil
}

// 授权开放应用token
func (a *API) GetForeignToken(ctx context.Context, r *pd.ForeignOauthRequest) ([]byte, error) {
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["machine_code"] = []string{r.MachineCode}
	data["msign"] = []string{r.Msign}
	data["scope"] = []string{"all"}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("oauth/scancodemodel", data)
	if err != nil {
		return nil, err
	}
	return body, nil

}

// 设置内置语音
func(a *API) PrintSetVoice(ctx context.Context, r *pd.PrintSetVoiceRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["content"] = []string{r.Content}
	data["is_file"] = []string{r.IsFile}
	data["aid"] = []string{r.Aid}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/setvoice", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 面单接口
func(a *API) ExpressPrint(ctx context.Context, r *pd.ExpressPrintRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["origin_id"] = []string{common.GetOrderId()}
	data["id"] = []string{common.GetUUID4()}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["timestamp"] = []string{timestamp}
	data["content"] = []string{r.Content}
	if r.Sandbox != ""{
		data["sandbox"] = []string{r.Sandbox}
	}

	body, err := a.httpPost("expressprint/index", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 图片接口
func (a *API)PicturePrint(ctx context.Context, r *pd.PicturePrintRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["picture_url"] = []string{r.PictureUrl}
	data["origin_id"] = []string{common.GetOrderId()}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("pictureprint/index", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//删除内置语音接口
func (a *API)PrintDeleteVoice(ctx context.Context, r *pd.PrintDeleteVoiceRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["aid"] = []string{r.Aid}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/deletevoice", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 删除终端授权
func (a *API)PrintDeletePrinter(ctx context.Context, r *pd.PrintDeletePrinterRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/deleteprinter", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 添加应用菜单
func (a *API)PrintAddPrintMenu(ctx context.Context, r *pd.PrintAddPrintMenuRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["content"] = []string{r.Content}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printmenu/addprintmenu", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//关机重启接口
func (a *API)PrintShutDownRestart(ctx context.Context, r *pd.PrintShutDownRestartRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["response_type"] = []string{r.ResponseType}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/shutdownrestart", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//声音调节接口
func (a *API)PrintSetSound(ctx context.Context, r *pd.PrintSetSoundRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["response_type"] = []string{r.ResponseType}
	data["voice"] = []string{r.Voice}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/setsound", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 获取机型打印宽度接口
func (a *API)PrintInfo(ctx context.Context, r *pd.PrintInfoRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/printinfo", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}
// 取消所有未打印订单
func (a *API)PrintCancelAll(ctx context.Context, r *pd.PrintCancelAllRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/cancelall", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//取消单条未打印订单
func (a *API)PrintCaneLone(ctx context.Context, r *pd.PrintCaneLoneRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["order_id"] = []string{r.OrderId}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/cancelone", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}
//设置logo接口
func (a *API)PrintSetIcon(ctx context.Context, r *pd.PrintSetIconRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["img_url"] = []string{r.ImgUrl}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/seticon", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}
// 取消logo接口
func (a *API)PrintDeleteIcon(ctx context.Context, r *pd.PrintDeleteIconRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/deleteicon", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 打印方式接口
func (a *API)PrintBtnPrint(ctx context.Context, r *pd.PrintBtnPrintRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["response_type"] = []string{r.ResponseType}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/btnprint", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}
//接单拒单设置接口
func (a *API)PrintGetOrder(ctx context.Context, r *pd.PrintGetOrderRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["response_type"] = []string{r.ResponseType}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/getorder", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//设置推送url接口
func (a *API)PrintSetPushUrl(ctx context.Context, r *pd.PrintSetPushUrlRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["cmd"] = []string{r.Cmd}
	data["url"] = []string{r.Url}
	data["status"] = []string{r.Status}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("oauth/setpushurl", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}
//获取订单状态接口
func (a *API)PrintGetOrderStatus(ctx context.Context, r *pd.PrintGetOrderStatusRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["order_id"] = []string{r.OrderId}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/getorderstatus", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//获取订单列表接口
func (a *API)PrintGetOrderPagingList(ctx context.Context, r *pd.PrintGetOrderPagingListRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["page_index"] = []string{r.PageIndex}
	data["page_size"] = []string{r.PageSize}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/getorderpaginglist", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//获取终端状态接口
func (a *API)PrintGetPrintStatus(ctx context.Context, r *pd.PrintGetPrintStatusRequest)([]byte, error){
	data := make(url.Values)

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["access_token"] = []string{getToken(r.AccessToken)}
	data["machine_code"] = []string{r.MachineCode}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["id"] = []string{common.GetUUID4()}
	data["timestamp"] = []string{timestamp}

	body, err := a.httpPost("printer/getprintstatus", data)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *API) httpGet(path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}

func (a *API) httpPost(path string, data url.Values) ([]byte, error) {
	resp, err := http.PostForm(fmt.Sprintf("%s/%s", a.URL, path), data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

// 获取access_token
func getToken(token string) string {
	if token != ""{
		return token
	}
	return setting.ClientSetting.AccessToken
}
