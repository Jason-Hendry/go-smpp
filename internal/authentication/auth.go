package authentication

import (
	"fmt"
	"github.com/ghiac/go-smpp"
)

func AuthCheck(pdu go_smpp.Pdu, smppClient *go_smpp.SmppClientConn) {
	if pdu.GetSystemID() == "admin" && pdu.GetPassword() == "admin" {
		smppClient.UserId = pdu.GetSystemID()
		fmt.Println("sign in : " + smppClient.UserId)
	}
	go_smpp.SetSmppClient(smppClient)
}
