package server

import (
	"context"
	"encoding/json"
	"yilianyun-print-go/Lib/pkg/bapi"
	"yilianyun-print-go/Lib/pkg/errcode"
	pd "yilianyun-print-go/Lib/pkg/proto"
)

type PrintServer struct {}

func NewPrintServer () *PrintServer {
	return &PrintServer{}
}

func (p *PrintServer)GetToken(ctx context.Context, r *pd.GetPrintRequest)(*pd.GetPrintReply,error){
	api := bapi.NewAPI("https://open-api.10ss.net/")
	body, err := api.GetAccessToken()
	if err != nil {
		return nil, err
	}
	printReply := pd.GetPrintReply{}
	err = json.Unmarshal(body, &printReply)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}

	return &printReply, nil
}
func (p *PrintServer)Print(ctx context.Context, r *pd.GetPrintRequest)(*pd.GetPrintReply,error){
	api := bapi.NewAPI("https://open-api.10ss.net/")
	body, err := api.Print(ctx, r)
	if err != nil {
		return nil, err
	}
	printReply := pd.GetPrintReply{}
	err = json.Unmarshal(body, &printReply)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.NewError(10000009,err.Error()))
	}

	return &printReply, nil

}

