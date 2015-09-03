package go_smpp
import (
	"net"
	"fmt"
	"os/signal"
	"os"
	"time"
)


func listen(conn *net.TCPConn) {
	resp := make([]byte, 10240);
	for {
		fmt.Printf("Waiting to read\n")
		readLen,err := conn.Read(resp[0:])
		if err != nil {
			fmt.Printf("Something broke\n")
			return
		}
		if readLen == 0 {
			fmt.Printf("Server stopped\n")
			return
		}

		fmt.Printf("Read %d Bytes\n", readLen)
		pduResp := RawPdu(resp)

		switch pduResp.command_id {
		case PDU_COMMAND_SUBMIT_SM+PDU_COMMAND_RESP:
			fmt.Printf("Listener: Submit Message ID: %s\n", string(pduResp.Message_id))
		case PDU_COMMAND_BIND_RX+PDU_COMMAND_RESP,PDU_COMMAND_BIND_TX+PDU_COMMAND_RESP,PDU_COMMAND_BIND_TRX+PDU_COMMAND_RESP:
			fmt.Printf("Listener: BIND Response status: %d\n", pduResp.Command_status)
		case PDU_COMMAND_DELIVER_SM:
			fmt.Printf("Listener: DLR status: %d\n", pduResp.Command_status)
		}
	}
}

func keepAlive(conn *net.TCPConn) {
	for {
		time.Sleep(60 * time.Second)
		enq := EnquireLink(1)
		conn.Write(enq.Pack())
	}
}

func Client(host, username, password, source, destination, message  string) {

	addr,err := net.ResolveTCPAddr("tcp", host)
	if !HandleError("Failed to resolve", err) {
		return
	}

	laddr,err := net.ResolveTCPAddr("tcp", ":0")
	if !HandleError("Failed to resolve", err) {
		return
	}

	conn,err := net.DialTCP("tcp", laddr, addr)
	if !HandleError("Failed to connect", err) {
		return
	}
	go listen(conn)
	go keepAlive(conn)

	bind := Bind(1, PDU_COMMAND_BIND_TX, username, password, "GO-SMPP", 0, 1, 1, "")
	data := bind.Pack()
	fmt.Printf("Wrote %d bytes\n", len(data))
	conn.Write(data)


	sms := SubmitSM(2, "GO-SMPP", 1, 1, source, 1, 1, destination, 3, 0, message)
	sms.registered_delivery = 4
	data = sms.Pack()
	fmt.Printf("Wrote %d bytes\n", len(data))
	conn.Write(data)



	fmt.Println("Wait to die")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("Got signal:", s)
	conn.Close()

}