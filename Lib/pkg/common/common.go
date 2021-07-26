package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"yilianyun-print-go/Lib/pkg/setting"
)

var (
	ClientId string
	ClientServer string
)
func Setup(){
	ClientId = setting.ClientSetting.ClientId
	if ClientId == ""{
		log.Printf("初始化client_id失败！")
	}
	ClientServer = setting.ClientSetting.ClientServer
	if ClientServer == ""{
		log.Printf("初始化client_server失败！")
	}
}
// 获取UUID4
func GetUUID4()string{
	return fmt.Sprintf("%04x%04x-%04x-%04x-%04x-%04x%04x%04x",
		mtRand(0,0xffff), mtRand(0,0xffff),
		mtRand(0,0xffff),
		mtRand(0,0xfff) | 0x4000,
		mtRand(0, 0x3fff) | 0x8000,
		mtRand(0, 0xffff), mtRand(0, 0xffff), mtRand(0, 0xffff),
	)
}

// 获取sign
func GetSign(timestamp string) string {
	clientId := ClientId
	clientServer := ClientServer
	pretreatmentData := fmt.Sprintf("%s%s%s",clientId,timestamp,clientServer)
	m := md5.New()
	m.Write([]byte(pretreatmentData))
	return fmt.Sprintf("%s", hex.EncodeToString(m.Sum(nil)))
}

func mtRand(min, max int) int {
	return rand.Intn(max - min) + min
}



