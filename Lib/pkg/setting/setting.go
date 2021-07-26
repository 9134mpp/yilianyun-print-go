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
}


func Setup(){
	var err error
	Cfg, err = ini.Load("../Lib/Config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'app.ini':%v", err)
	}
	err = Cfg.Section("Client").MapTo(ClientSetting)
	if err!= nil  {
		log.Fatalf("Cfg.MapTo ClientSetting err: %v", err)
	}
}