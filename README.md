<h1 align="center">yilianyun-print-sdk</h1>

# Requirement

```
Golang >= 1.16
```
# Installation

```shell
git clone https://github.com/9134mpp/yilianyun-print-go.git
```


# Usage
  1. 接口类Lib/bapi/api.go，集成了所有的易联云接口
  2. 配置类Config/app.ini
  3. 公共方法Lib/common/common.go 有获取uuid,order_id, sign 等方法
  4. 配置读取类Lib/setting/seeting.go 读取相关配置如client_id, client_server等配置
  5. 错误返回目录Lib/errcode 返回grpc 错误的目录
  6. 目录Lib/Proto 定义Protobuf数据结构的目录
  7. 目录Lib/server grpc Server服务的目录 
  8. Demo/Client client端调用测试demo
  8. 若有什么问题，欢迎来提个issues， 谢谢大家！
 
 首先运行 根目录下的main.go 
```go 
package main

import (
	"context"
	"fmt"
	"strings"
	"yilianyun-print-go/Demo/client/Lib/Api"
	pd "yilianyun-print-go/Lib/proto"
	"yilianyun-print-go/Lib/setting"
)


var Client pd.PrintServiceClient

func main() {
	ctx := context.Background()
	clientConn, _ := Api.GetClientConn(ctx, "localhost:9523", nil)
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
	fmt.Println(token)

	//打印
	p := &pd.PrintRequest{
		AccessToken: setting.ClientSetting.AccessToken,
		MachineCode: "4004628156",
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
	var s  []string
	for i :=0 ;i < multiplier; i ++{
		s = append(s, str)
	}
	return fmt.Sprintf("%s", strings.Join(s, ""))
}

// 开放应用示例demo
func foreign(){
	// 这里示范个人推荐的授权模式 (极速授权模式：http://doc2.10ss.net/343932)
	//授权
	pam := &pd.ForeignOauthRequest{
		MachineCode: "4004628156",
		Msign: "14757845145",
	}
	token, _ := Api.GetForeignToken(Client, pam) // 这里token存在数据库或者缓存中即可
	fmt.Println(token)
	//打印
	p := &pd.PrintRequest{
		AccessToken: token,
		MachineCode: "4004628158",
		Content:     content(),
	}
	_ = Api.Print(Client, p)
}












```

易联云Go SDK