package go_smpp
import (
	"net"
	"fmt"
)


func Client(host, username, password, source, destination, message  string) {

	addr,err := net.ResolveTCPAddr("tcp", host)
	HandleError("Failed to resolve", err)

	laddr,err := net.ResolveTCPAddr("tcp", ":0")
	HandleError("Failed to resolve", err)

	conn,err := net.DialTCP("tcp", laddr, addr)
	HandleError("Failed to connect", err)

	bind := Bind(1, PDU_COMMAND_BIND_TX, username, password, "GO-SMPP", 0, 1, 1, "")
	conn.Write(bind.Pack())

	resp := make([]byte, 10240);

	conn.Read(resp)
	bindResp := Pdu(resp)
	fmt.Printf("BIND Response status: %d\n", bindResp.Command_status)

	sms := SubmitSM(2, "GO-SMPP", 1, 1, source, 1, 1, destination, 3, 0, message)
	sms.registered_delivery = 4
	conn.Write(sms.Pack())

	conn.Read(resp)
	submitSMResp := Pdu(resp)
	fmt.Printf("Submit Message ID: %s\n", string(submitSMResp.Message_id))

	conn.Read(resp)
	dlr := Pdu(resp)
	fmt.Printf("DLR Msg: %s\n", string(dlr.short_message))

	conn.Close()
}