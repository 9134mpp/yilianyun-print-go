package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	Cfg *ini.File
)

var ClientSetting = &Client{}

type Client struct {
	ClientId string
	ClientServer string
	AccessToken string
}


var ApiSetting = &Api{}

type Api struct {
	Url string
}


func Setup(){
	var err error
	Cfg, err = ini.Load("Config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'app.ini':%v", err)
	}
	err = Cfg.Section("Client").MapTo(ClientSetting)
	if err!= nil  {
		log.Fatalf("Cfg.MapTo ClientSetting err: %v", err)
	}
	err = Cfg.Section("Api").MapTo(ApiSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ApiSetting err: %v", err)
	}
}

func Update(section, key, value string){
	Cfg.Section(section).Key(key).SetValue(value)
	_ = Cfg.SaveTo("Config/app.ini")
}