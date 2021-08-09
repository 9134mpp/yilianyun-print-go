<h1 align="center">yilianyun-print-sdk</h1>

# 版本

```
Golang >= 1.16
```
# 下载

```shell
git clone https://github.com/9134mpp/yilianyun-print-go.git
```

如果下载过慢
```shell
git clone https://github.com.cnpmjs.org/9134mpp/yilianyun-print-go
```
##目录介绍
```
├─ yilianyun-print-go
│  ├─ Config            ---配置目录
│  ├─├─ app.example.ini  ---配置文件模板
│  ├─ Demo   ---示例调用demo目录
│  ├─├─ client   ---client端调用示例目录
│  ├─├─├─ Lib    ---client的内部封装目录
│  ├─├─├─├─Api   --- Api所有调用接口示例目录
│  ├─├─├─├─├─ printClient.go ---client Api封装文件
│  ├─├─├─ main.go  ---client调用示例
│  ├─ Lib  ---server端的核心文件
│  ├─├─ bapi ---server Api调用目录
│  ├─├─ common ---公共方法目录如生成sign order_id方法
│  ├─├─ errcode ---server端内部错误目录
│  ├─├─ proto ---编译和生成 proto 文件目录
│  ├─├─ server ---server端接口封装目录
│  ├─├─ setting ---读取配置文件目录
│  ├─ vendor 
│  ├─ main.go ---server端启动文件

    
若有什么问题，欢迎来提个issues， 谢谢大家！
```
 # 使用
 先运行
 ```shell script
 go mod vendor
 ```
 根目录下运行 main.go
   ```shell script
  go run main.go grpc_port=9523
  ```
更多详情运行
```shell script
go run main.go -help
````
然后运行Demo/client 开始测试
  ```shell script
 go run main.go -url=localhost -port=9523
 ```
注意 port端口要和配置文件的端口一致
更多详情运行
```shell script
go run main.go -help
````
grpc server端服务 在参考 Demo/Client/main.go代码

## 使用说明
```
client_id、origin_id、sign、id、timestamp参数不用在传递，server端这边会自动生成
自由应用可以不用传assess_token。
开放应用需要传assess_token。

配置文件如果默认填写了access_token默认即自有应用。
对接文档 http://doc2.10ss.net/331992
如果文档失效查看 https://doc3.10ss.net/
```

# 示范代码 
```go 
package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"yilianyun-print-go/Demo/client/Lib/Api"
	pd "yilianyun-print-go/Lib/proto"
	"yilianyun-print-go/Lib/setting"
)

var Client pd.PrintServiceClient

var url string
var port string
func init() {
	flag.StringVar(&url,"url", "localhost","rpcClient 启动地址(默认localhost)")
	flag.StringVar(&port,"port", "9523","rpcClient 启动端口(默认9523端口)")
	flag.Parse()
}
func main() {
	ctx := context.Background()
	clientConn, _ := Api.GetClientConn(ctx, fmt.Sprintf("%s:%s", setting.RpcServerSetting.Url, setting.RpcServerSetting.Port), nil)
	defer clientConn.Close()
	Client = pd.NewPrintServiceClient(clientConn)
	// 自有应用示例子
	have()
	// 开放应用示例
	//foreign()
}

// 自有应用授权打印demo
func have() {
	//授权
	token, _ := Api.GetToken(Client, &pd.OauthRequest{}) // 自有应用access_token有效期时间永久获取后记得保存，不要频繁获取！！！ 不要频繁获取！！！ 不要频繁获取！！！
	fmt.Printf("token:%s\n", token)
	//打印
	p := &pd.PrintRequest{
		AccessToken: token,
		MachineCode: "",
		Content:     content(),
	}
	_ = Api.Print(Client, p)
}

// 示例模板
func content() string {
	content := "<FS2><center>**#1 美团**</center></FS2>"
	content += strRepeat(".", 32)
	content += "<FS2><center>--在线支付--</center></FS2>"
	content += "<FS><center>张周兄弟烧烤</center></FS>"
	content += "订单时间: 2020-12-25 10:15 \n"
	content += "订单编号:40807050607030\n"
	content += strRepeat("*", 14) + "商品" + strRepeat("*", 14)
	content += "<table>"
	content += "<tr><td>烤土豆(超级辣)</td><td>x3</td><td>5.96</td></tr>"
	content += "<tr><td>烤豆干(超级辣)</td><td>x2</td><td>3.88</td></tr>"
	content += "<tr><td>烤鸡翅(超级辣)</td><td>x3</td><td>17.96</td></tr>"
	content += "<tr><td>烤排骨(香辣)</td><td>x3</td><td>12.44</td></tr>"
	content += "<tr><td>烤韭菜(超级辣)</td><td>x3</td><td>8.96</td></tr>"
	content += "</table>"
	content += strRepeat(".", 32)
	content += "<QR>这是二维码内容</QR>"
	content += "小计:￥82\n"
	content += "折扣:￥４ \n"
	content += strRepeat("*", 32)
	content += "订单总价:￥78 \n"
	content += "<FS2><center>**#1 完**</center></FS2>"

	return content
}

func strRepeat(str string, multiplier int) string {
	var s []string
	for i := 0; i < multiplier; i++ {
		s = append(s, str)
	}
	return fmt.Sprintf("%s", strings.Join(s, ""))
}

// 开放应用示例demo
func foreign() {
	// 这里示范个人推荐的授权模式 (极速授权模式：http://doc2.10ss.net/343932)
	//授权
	pam := &pd.ForeignOauthRequest{
		MachineCode: "",
		Msign:       "",
	}
	token, _ := Api.GetForeignToken(Client, pam) // 这里token存在数据库或者缓存中即可
	fmt.Println(token)
	//打印
	p := &pd.PrintRequest{
		AccessToken: token,
		MachineCode: "",
		Content:     content(),
	}
	_ = Api.Print(Client, p)
}
```
# 参考接口
```go
// 文本打印
func Print(client pd.PrintServiceClient, p *pd.PrintRequest) error
// 获取token
func GetToken(client pd.PrintServiceClient, p *pd.OauthRequest) (string, error)
// 开放应用获取token
func GetForeignToken(client pd.PrintServiceClient, p *pd.ForeignOauthRequest)(string, error)
// 图形打印
func PicturePrint(client pd.PrintServiceClient, p *pd.PicturePrintRequest) error
// 面单接口
func ExpressPrint(client pd.PrintServiceClient, p *pd.ExpressPrintRequest) error
// 设置内置语音接口
func PrintSetVoice(client pd.PrintServiceClient, p *pd.PrintSetVoiceRequest) error
// 删除内置语音接口
func PrintDeleteVoice(client pd.PrintServiceClient, p *pd.PrintDeleteVoiceRequest) error
// 删除终端授权
func PrintDeletePrinter(client pd.PrintServiceClient, p *pd.PrintDeletePrinterRequest) error
// 添加应用菜单
func PrintAddPrintMenu(client pd.PrintServiceClient, p *pd.PrintAddPrintMenuRequest) error
// 关机重启接口
func PrintShutDownRestart(client pd.PrintServiceClient, p *pd.PrintShutDownRestartRequest) error
// 声音调节接口
func PrintSetSound(client pd.PrintServiceClient, p *pd.PrintSetSoundRequest) error
// 获取机型打印宽度接口
func PrintInfo(client pd.PrintServiceClient, p *pd.PrintInfoRequest) error
// 取消所有未打印订单
func PrintCancelAll(client pd.PrintServiceClient, p *pd.PrintCancelAllRequest) error
// 取消单条未打印订单
func PrintCaneLone(client pd.PrintServiceClient, p *pd.PrintCaneLoneRequest) error
// 设置logo接口
func PrintSetIcon(client pd.PrintServiceClient, p *pd.PrintSetIconRequest) error 
// 取消logo接口
func PrintDeleteIcon(client pd.PrintServiceClient, p *pd.PrintDeleteIconRequest) error
// 打印方式接口
func PrintBtnPrint(client pd.PrintServiceClient, p *pd.PrintBtnPrintRequest) error
// 接单拒单设置接口
func PrintGetOrder(client pd.PrintServiceClient, p *pd.PrintGetOrderRequest) error
// 设置推送url接口
func PrintSetPushUrl(client pd.PrintServiceClient, p *pd.PrintSetPushUrlRequest) error
// 获取订单状态接口
func PrintGetOrderStatus(client pd.PrintServiceClient, p *pd.PrintGetOrderStatusRequest) error
// 获取订单列表接口
func PrintGetOrderPagingList(client pd.PrintServiceClient, p *pd.PrintGetOrderPagingListRequest) error
//	获取终端状态接口
func PrintGetPrintStatus(client pd.PrintServiceClient, p *pd.PrintGetPrintStatusRequest) error
// 终端授权 (永久授权)
func PrintAddPrinter(client pd.PrintServiceClient, p *pd.PrintAddPrinterRequest) error
```
具体细节参考 Demo/client/Lib/Api/printClient.go文件

说明一下 获取订单列表接口
由于数据结构原因目前我把数据返回的list放在了
`ListBody`里面 所以要获取接口返回的`body`,使用 `GetListBody()`。