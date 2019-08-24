package main

import (
	"fmt"
	"github.com/ghiac/go-smpp"
	"github.com/ghiac/smpp-proxy/internal/authentication"
	"github.com/ghiac/smpp-proxy/internal/util"
)

var bind = func(pdu go_smpp.Pdu, smppClient *go_smpp.SmppClientConn) {
	authentication.AuthCheck(pdu, smppClient)
	if go_smpp.GetSmppClient(smppClient).UserId != "" {
		fmt.Println("request from ========>" + go_smpp.GetSmppClient(smppClient).UserId)
		submitResp := go_smpp.BindResp(go_smpp.Pdu{}, 1, string(pdu.MessageId))
		smppClient.WritePdu(submitResp)
	} else {
		var emptyByte []byte
		smppClient.Write(emptyByte)
	}
}

var submit go_smpp.OnPduCallback = func(pdu go_smpp.Pdu, smppClient *go_smpp.SmppClientConn) {
	fmt.Println("============> Request: User Id  " + smppClient.UserId)
	util.SendMessage(pdu.GetDestination(), pdu.GetSource(), pdu.GetOptionalParameters()[0].GetValue())
	submitResp := go_smpp.SubmitResp(go_smpp.Pdu{}, 1, string(pdu.MessageId))
	go_smpp.GetSmppClient(smppClient).WritePdu(submitResp)
}

func main() {
	server := go_smpp.Server("1", "localhost:7878")
	server.OnSubmit = submit
	server.OnBind = bind
	server.Start()
}
