package bapi

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"yilianyun-print-go/Lib/pkg/common"

	"yilianyun-print-go/Lib/pkg/setting"

	pd "yilianyun-print-go/Lib/pkg/proto"
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
	data["access_token"] = []string{r.AccessToken}
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
