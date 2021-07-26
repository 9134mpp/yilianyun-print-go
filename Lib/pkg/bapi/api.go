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

func(a *API)GetAccessToken()([]byte, error){
	data := make(url.Values)
	timestamp := strconv.FormatInt(time.Now().Unix(),10)
	data["client_id"] = []string{setting.ClientSetting.ClientId}
	data["sign"] = []string{common.GetSign(timestamp)}
	data["scope"] = []string{"all"}
	data["timestamp"] = []string{timestamp}
	data["id"] = []string{common.GetUUID4()}

	body, err := a.httpPost("oauth/oauth", data)
	if err != nil {
		return nil , err
	}
	return body, nil
}

func (a *API) Print(ctx context.Context, r *pd.GetPrintRequest)([]byte, error){
	data := make(url.Values)
	timestamp := strconv.FormatInt(r.Timestamp,10)
	data["client_id"] = []string{r.ClientId}
	data["access_token"] = []string{r.AccessToken}
	data["sign"] = []string{r.Sign}
	data["machine_code"] = []string{r.MachineCode}
	data["content"] = []string{r.Content}
	data["origin_id"] = []string{r.OriginId}
	data["timestamp"] = []string{timestamp}
	data["id"] = []string{common.GetUUID4()}
	body, err := a.httpPost("print/index",data)

	if err != nil {
		return nil , err
	}
	return body, nil
}

func(a *API)httpGet(path string)([]byte, error){
	resp, err := http.Get(fmt.Sprintf("%s/%s",a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}

func(a *API)httpPost(path string, data url.Values)([]byte, error){
	resp, err := http.PostForm(fmt.Sprintf("%s/%s",a.URL, path), data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body,_ := ioutil.ReadAll(resp.Body)
	return body, nil
}

