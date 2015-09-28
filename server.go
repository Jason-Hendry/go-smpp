package go_smpp
import (
	"net"
	"fmt"
	"io"
	"code.google.com/p/go-uuid/uuid"
)

type SmppServer struct {
	NodeUuid uuid.UUID
	Node string
	bindAddr string
	OnBind OnPduCallback
	OnSubmit OnPduCallback
	connections map[string]*net.TCPConn
}

type SmppClientConn struct {
	conn *net.TCPConn
	server *SmppServer
	UserId string
}

func (c *SmppClientConn) Write(data []byte) {
	c.conn.Write(data)
}
func (c *SmppClientConn) WritePdu(pdu Pdu) {
	c.conn.Write(pdu.Pack())
}

type OnPduCallback func(Pdu, *SmppClientConn)

func handleClient(conn *net.TCPConn, server SmppServer) {
	buf := make([]byte, 20480)
	read := make([]byte, 10240)
	more := make([]byte, 20480)
	more = more[:0]
	buf = buf[:0]
	client := SmppClientConn{conn, &server, ""}
	var smppPdu Pdu
	for {
		fmt.Printf("Node: %s Waiting to read\n", server.Node)
		readLen,err := conn.Read(read[0:])
		if readLen == 0 {
			fmt.Printf("Node: %s Read %d Bytes and probably quit\n", server.Node, readLen)
			return
		}
		fmt.Printf("Node: %s Add %d bytes to buffer\n", server.Node, readLen)
		smppPdu = readMore(&buf, &more, &read, readLen)
		if smppPdu.complete {
			fmt.Printf("Node: %s Got one\n", server.Node)
			processPdu(&client, smppPdu)

		}
		if err == io.EOF {

		}

	}
}

//      00 00 00 00 00 00 00 00
// buf
// more
// read 01 12 12 32 34 34 34 12 8
// buf  01 12 12 32 34 34 34 12
// more 01 12 12 32 34 34 34 12
// read 02 23 12_00 00 00 00 00 3
// buf  01 12 12 32 34 34 34 12 02 23 12
// more
func readMore(bufP *[]byte, moreP *[]byte, readP *[]byte, readLen int) (smppPdu Pdu) {
	buf := *bufP
	more := *moreP
	read := *readP

	buf = more[:len(more)] // Leftover from previous packet
	buf = append(buf, read[:readLen]...) // Newly ready data from socket

	smppPdu = RawPdu(buf)
	if smppPdu.complete == true {
		if uint32(len(read)) > smppPdu.command_length {
			more = more[:0]
			more = append(more, buf[int(smppPdu.command_length):len(buf)]...)
		}
	} else {
		more = buf[0:len(buf)]
	}
	*moreP = more;
	return smppPdu
}

func processPdu(client *SmppClientConn, pdu Pdu) {

	server := client.server
	fmt.Printf("Node: %s Got %d\n", server.Node, pdu.command_id)

	switch pdu.command_id {
	case PDU_COMMAND_BIND_TX, PDU_COMMAND_BIND_TRX, PDU_COMMAND_BIND_RX:
		fmt.Printf("Node: %s Bind\n", client.server.Node)
		go server.OnBind(pdu, client)
	case PDU_COMMAND_SUBMIT_SM:
		fmt.Printf("Node: %s SUBMIT_SM\n", server.Node)
		go server.OnSubmit(pdu, client)
	case PDU_COMMAND_ENQUIRE:
		fmt.Printf("Node: %s ENQUIRE\n", server.Node)
		resp := EnquireLinkResp(pdu)
		client.Write(resp.Pack())
	}
}

func (server SmppServer) Start() {
	bindAddr,err := net.ResolveTCPAddr("tcp", server.bindAddr)
	if !HandleError(fmt.Sprintf("Failed to resolve %s",server.bindAddr), err) {
		return
	}
	bind,err := net.ListenTCP("tcp", bindAddr)
	if !HandleError(fmt.Sprintf("Failed to bind to %s",server.bindAddr), err) {
		return
	}

	for {
		conn,err := bind.AcceptTCP()
		HandleError("Failed to accept client", err)
		fmt.Printf("Node: %s Connection: %s\n", server.Node, conn.RemoteAddr().String())
		// server.connections[conn.RemoteAddr().String()] = conn
		go handleClient(conn, server)
	}
}


func Server(node string, bindAddr string) (*SmppServer) {
	var server SmppServer
	server.Node = node
	server.bindAddr = bindAddr
	return &server
}